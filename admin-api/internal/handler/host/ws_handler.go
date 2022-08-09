package host

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/websocket"
	"golang.org/x/crypto/ssh"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"sync"
	"time"

	"github.com/haozheyu/oam_system/admin-api/internal/logic/host"
	"github.com/haozheyu/oam_system/admin-api/internal/svc"
	"github.com/haozheyu/oam_system/admin-api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024 * 10,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var bufPool = sync.Pool{
	New: func() interface{} {
		return new(bytes.Buffer)
	},
}

func WsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetWSReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := host.NewWsLogic(r.Context(), svcCtx)
		resp, err := l.Ws(&req)

		if err != nil {
			httpx.Error(w, err)
			return
		} else {
			// 进行协议升级
			wsConn, err := upgrader.Upgrade(w, r, nil)
			if err != nil {
				httpx.Error(w, errors.New(fmt.Sprintf("协议升级失败, err: %s", err)))
			}
			defer wsConn.Close()
			var config *SSHClientConfig
			config = SSHClientConfigPassword(
				fmt.Sprintf("%s:%s", resp.Addr, resp.Port),
				resp.User,
				resp.Password,
			)
			client, err := NewSSHClient(config)
			if err != nil {
				wsConn.WriteControl(websocket.CloseMessage,
					[]byte(err.Error()), time.Now().Add(time.Second))
				return
			}
			defer client.Close()
			var recorder *Recorder

			//mask := syscall.Umask(0)
			//defer syscall.Umask(mask)
			os.MkdirAll("./rec", 0766)
			fileName := path.Join("./rec", fmt.Sprintf("%s_%s.cast", resp.Addr, time.Now().Format("20060102_150405")))
			f, err := os.OpenFile(fileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0766)
			if err != nil {
				httpx.Error(w, errors.New(fmt.Sprintf("rec目录生产失败, err: %s", err)))
			}
			defer f.Close()
			recorder = NewRecorder(f)

			turn, err := NewTurn(wsConn, client, recorder)

			if err != nil {
				wsConn.WriteControl(websocket.CloseMessage,
					[]byte(err.Error()), time.Now().Add(time.Second))
				return
			}
			defer turn.Close()

			var logBuff = bufPool.Get().(*bytes.Buffer)
			logBuff.Reset()
			defer bufPool.Put(logBuff)

			ctx, cancel := context.WithCancel(context.Background())
			wg := sync.WaitGroup{}
			wg.Add(2)
			go func() {
				defer wg.Done()
				err := turn.LoopRead(logBuff, ctx)
				if err != nil {
					log.Printf("%#v", err)
				}
			}()
			go func() {
				defer wg.Done()
				err := turn.SessionWait()
				if err != nil {
					log.Printf("%#v", err)
				}
				cancel()
			}()
			wg.Wait()
		}
	}
}

// Turn
const (
	MsgData   = '1'
	MsgResize = '2'
)

type Turn struct {
	StdinPipe io.WriteCloser
	Session   *ssh.Session
	WsConn    *websocket.Conn
	Recorder  *Recorder
}

func NewTurn(wsConn *websocket.Conn, sshClient *ssh.Client, rec *Recorder) (*Turn, error) {
	sess, err := sshClient.NewSession()
	if err != nil {
		return nil, err
	}

	stdinPipe, err := sess.StdinPipe()
	if err != nil {
		return nil, err
	}

	turn := &Turn{StdinPipe: stdinPipe, Session: sess, WsConn: wsConn}
	sess.Stdout = turn
	sess.Stderr = turn

	modes := ssh.TerminalModes{
		ssh.ECHO:          1,     // disable echo
		ssh.TTY_OP_ISPEED: 14400, // input speed = 14.4kbaud
		ssh.TTY_OP_OSPEED: 14400, // output speed = 14.4kbaud
	}
	if err := sess.RequestPty("xterm", 150, 30, modes); err != nil {
		return nil, err
	}
	if err := sess.Shell(); err != nil {
		return nil, err
	}

	if rec != nil {
		turn.Recorder = rec
		turn.Recorder.Lock()
		turn.Recorder.WriteHeader(30, 150)
		turn.Recorder.Unlock()
	}

	return turn, nil
}
func (t *Turn) Write(p []byte) (n int, err error) {
	writer, err := t.WsConn.NextWriter(websocket.TextMessage)
	if err != nil {
		return 0, err
	}
	defer writer.Close()
	if t.Recorder != nil {
		t.Recorder.Lock()
		t.Recorder.WriteData(OutPutType, string(p))
		t.Recorder.Unlock()
	}
	return writer.Write(p)
}
func (t *Turn) Close() error {
	if t.Session != nil {
		t.Session.Close()
	}
	return t.WsConn.Close()
}
func (t *Turn) Read(p []byte) (n int, err error) {
	for {
		msgType, reader, err := t.WsConn.NextReader()
		if err != nil {
			return 0, err
		}
		if msgType != websocket.TextMessage {
			continue
		}
		return reader.Read(p)
	}
}
func (t *Turn) LoopRead(logBuff *bytes.Buffer, context context.Context) error {
	for {
		select {
		case <-context.Done():
			return errors.New("LoopRead exit")
		default:
			_, wsData, err := t.WsConn.ReadMessage()
			if err != nil {
				return fmt.Errorf("reading webSocket message err:%s", err)
			}
			body := decode(wsData[1:])
			switch wsData[0] {
			case MsgResize:
				var args Resize
				err := json.Unmarshal(body, &args)
				if err != nil {
					return fmt.Errorf("ssh pty resize windows err:%s", err)
				}
				if args.Columns > 0 && args.Rows > 0 {
					if err := t.Session.WindowChange(args.Rows, args.Columns); err != nil {
						return fmt.Errorf("ssh pty resize windows err:%s", err)
					}
				}
			case MsgData:
				if _, err := t.StdinPipe.Write(body); err != nil {
					return fmt.Errorf("StdinPipe write err:%s", err)
				}
				if _, err := logBuff.Write(body); err != nil {
					return fmt.Errorf("logBuff write err:%s", err)
				}
			}
		}
	}
}
func (t *Turn) SessionWait() error {
	if err := t.Session.Wait(); err != nil {
		return err
	}
	return nil
}
func decode(p []byte) []byte {
	decodeString, _ := base64.StdEncoding.DecodeString(string(p))
	return decodeString
}
func encode(p []byte) []byte {
	encodeToString := base64.StdEncoding.EncodeToString(p)
	return []byte(encodeToString)
}

type Resize struct {
	Columns int
	Rows    int
}

// recorder
type RecType string

const (
	InputType  RecType = "i"
	OutPutType RecType = "o"
)

type RecHeader struct {
	Version   int   `json:"version"`
	Width     int   `json:"width"`
	Height    int   `json:"height"`
	Timestamp int64 `json:"timestamp"`
	Env       struct {
		Shell string `json:"SHELL"`
		Term  string `json:"TERM"`
	} `json:"env"`
}

func defaultRecHeader() *RecHeader {
	recHeader := new(RecHeader)
	recHeader.Version = 2
	recHeader.Env.Shell = "/bin/bash"
	recHeader.Env.Term = "xterm-256color"
	return recHeader
}

type Recorder struct {
	StartTime time.Time
	Writer    io.Writer
	sync.Mutex
}

func NewRecorder(writer io.Writer) *Recorder {
	return &Recorder{
		StartTime: time.Now(),
		Writer:    writer,
	}
}
func (rec *Recorder) WriteHeader(height, width int) {
	header := defaultRecHeader()
	header.Timestamp = rec.StartTime.Unix()
	header.Height = height
	header.Width = width
	b, _ := json.Marshal(header)
	rec.Writer.Write(b)
	rec.Writer.Write([]byte("\r\n"))
}
func (rec *Recorder) WriteData(rectype RecType, data string) {
	recData := make([]interface{}, 3)
	recData[0] = float64(time.Since(rec.StartTime).Microseconds()) / float64(1000000)
	recData[1] = rectype
	recData[2] = data
	b, _ := json.Marshal(recData)
	rec.Writer.Write(b)
	rec.Writer.Write([]byte("\r\n"))
}

// ssh config
type AuthModel int8

const (
	PASSWORD AuthModel = iota + 1
	PUBLICKEY
)

type SSHClientConfig struct {
	AuthModel AuthModel
	HostAddr  string
	User      string
	Password  string
	KeyPath   string
	Timeout   time.Duration
}

func SSHClientConfigPassword(hostAddr, user, Password string) *SSHClientConfig {
	return &SSHClientConfig{
		Timeout:   time.Second * 5,
		AuthModel: PASSWORD,
		HostAddr:  hostAddr,
		User:      user,
		Password:  Password,
	}
}

func SSHClientConfigPulicKey(hostAddr, user, keyPath string) *SSHClientConfig {
	return &SSHClientConfig{
		Timeout:   time.Second * 5,
		AuthModel: PUBLICKEY,
		HostAddr:  hostAddr,
		User:      user,
		KeyPath:   keyPath,
	}
}

func NewSSHClient(conf *SSHClientConfig) (*ssh.Client, error) {
	config := &ssh.ClientConfig{
		Timeout:         conf.Timeout,
		User:            conf.User,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(), //忽略know_hosts检查
	}
	switch conf.AuthModel {
	case PASSWORD:
		config.Auth = []ssh.AuthMethod{ssh.Password(conf.Password)}
	case PUBLICKEY:
		signer, err := getKey(conf.KeyPath)
		if err != nil {
			return nil, err
		}
		config.Auth = []ssh.AuthMethod{ssh.PublicKeys(signer)}
	}
	c, err := ssh.Dial("tcp", conf.HostAddr, config)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func getKey(keyPath string) (ssh.Signer, error) {
	key, err := ioutil.ReadFile(keyPath)
	if err != nil {
		return nil, err
	}
	return ssh.ParsePrivateKey(key)
}
