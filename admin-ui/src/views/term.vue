<template>
  <div ref="terminal" id="terminal">

    <el-button class="termButton" type="text" icon="el-icon-link" @click="gobanckfun()">返回</el-button>
  </div>

</template>

<script>
import {defineComponent, onMounted, ref, onUnmounted} from 'vue'
import { useRouter } from 'vue-router'
import { Terminal } from 'xterm';
import { FitAddon  } from 'xterm-addon-fit';
import Base64 from "crypto-js/enc-base64";
import Utf8 from "crypto-js/enc-utf8";
import 'xterm/css/xterm.css'
import { ElMessage } from 'element-plus'
import {hostsStore} from "../store/hosts";
import ZmodemBrowser from 'nora-zmodemjs/src/zmodem_browser';


export default defineComponent({
  name: 'Xterm',
  setup(){
    //实例化路由
    const msgData = '1'
    const msgResize = '2'

    const shellWs = ref(null)
    const rows = ref(null)
    const cols = ref(null)
    const get_term_size = () => {
      let init_width = 1;
      let init_height = 21;
      let windows_width = document.body.clientWidth;
      let windows_height = document.body.clientHeight;
      return {
        cols: Math.floor(windows_width / init_width),
        rows: Math.floor(windows_height / init_height),
      }
    }
    const clos_rows = get_term_size();

    const term = new Terminal({
      rendererType: 'canvas',
      cursorBlink: true,
      convertEol: true,
      cols: clos_rows.cols,
      rows: clos_rows.rows,
      // scrollback: 1200,
      // row: 200,
      theme: {
        foreground: 'white',
        background: '#181E29'
      }
    })
    const fitAddon = new FitAddon();
    // canvas背景全屏
    term.loadAddon(fitAddon);
    fitAddon.fit();
    const router = useRouter();
    const sshStore = hostsStore();

    let curl = `ws://127.0.0.1:8888/api/host/ws?user=` + sshStore.HostItem.user + `&&addr=` + sshStore.HostItem.addr + `&&port=` + sshStore.HostItem.port
    //let curl = `ws://127.0.0.1:8888/api/host/ws?user=`+ router.currentRoute.value.query.user `&&addr=` + router.currentRoute.value.query.addr + `&&port=` + router.currentRoute.value.query.port
    const ws = new WebSocket(curl)

    const gobanckfun = () => {
      router.go(-1)
    }

    onMounted(()=>{
      term.open(document.getElementById('terminal'));  //绑定dom节点
      term.focus() // 取得输入焦点
      term.writeln('Connecting...');  // 写一行测试

      ws.onclose = function (e) {
        ElMessage.warning({
          message: '链接已关闭',
          type: 'warning',
          center: true,
        });
        router.go(-1)
      }
      ws.onmessage = function (e) { // 服务端ssh输出, 写到web shell展示
        term.write(e.data)
      }
      ws.onerror = function (e) {
        ElMessage.error({
          message: '连接断开',
          type: 'error',
          center: true,
        });
        router.go(-1)
      }
      // 当向web终端敲入字符时候的回调
      // term.onKey(e => {  //给后端发送数据   不支持粘贴
      //   // 写给服务端, 由服务端发给container
      //   console.log(e.key)
      //   ws.send(msgData + Base64.stringify(Utf8.parse(e.key)))
      // })
      // 支持输入与粘贴方法
      term.onData(input => {
        ws.send(msgData + Base64.stringify(Utf8.parse(input.valueOf())))
      })
    })
    onUnmounted(()=>{
      ws.close()
    })
    return {
      shellWs, term, rows, cols,gobanckfun
    }
  },
})

</script>

<!--<script>-->
<!--import 'xterm/css/xterm.css'-->
<!--import { Terminal } from 'xterm'-->
<!--import { FitAddon } from 'xterm-addon-fit'-->
<!--import Base64 from "crypto-js/enc-base64"-->
<!--import Utf8 from "crypto-js/enc-utf8"-->
<!--import { hostsStore } from "../store/hosts";-->
<!--import { useRouter } from "vue-router";-->
<!--import {ElMessage} from "element-plus";-->
<!--import "../utils/webssh";-->

<!--const sshStore = hostsStore();-->
<!--const msgData = '1'-->
<!--const msgResize = '2'-->
<!--export default {-->
<!--  name:"App",-->
<!--  mounted() {-->
<!--    const router = useRouter();-->
<!--    const terminal = new Terminal({})-->
<!--    const fitAddon = new FitAddon()-->
<!--    terminal.loadAddon(fitAddon)-->
<!--    fitAddon.fit()-->

<!--    let terminalContainer = document.getElementById("app")-->
<!--    // console.log(sshStore.HostItem, router.currentRoute.value.query.addr)-->
<!--    if (sshStore.HostItem.user === "") {-->
<!--      ElMessage.warning("链接未建立")-->
<!--      router.go(-1)-->
<!--    }-->
<!--    let curl = `ws://127.0.0.1:8888/api/host/ws?user=` + sshStore.HostItem.user + `&&addr=` + sshStore.HostItem.addr + `&&port=` + sshStore.HostItem.port-->
<!--    //let curl = `ws://127.0.0.1:8888/api/host/ws?user=`+ router.currentRoute.value.query.user `&&addr=` + router.currentRoute.value.query.addr + `&&port=` + router.currentRoute.value.query.port-->
<!--    const webSocket = new WebSocket(curl)-->
<!--    /*-->
<!--    *  new add-->
<!--    */-->
<!--    // webSocket.binaryType = "arraybuffer";-->

<!--    webSocket.onopen = () => {-->
<!--      terminal.open(terminalContainer)-->
<!--      fitAddon.fit()-->
<!--      terminal.write("welcome to WebSSH ☺\r\n")-->

<!--      terminal.focus()-->
<!--    }-->

<!--    webSocket.onclose = () => {-->
<!--      terminal.write("\r\nWebSSH quit!")-->
<!--      webSocket.close()-->
<!--      location.reload()-->

<!--    }-->

<!--    webSocket.onerror = (event) => {-->
<!--      console.error(event)-->
<!--      webSocket.close()-->
<!--      location.reload()-->
<!--    }-->

<!--    terminal.onKey((event) => {-->
<!--      webSocket.send(msgData + Base64.stringify(Utf8.parse(event.key)))-->
<!--    })-->

<!--    terminal.onResize(({ cols, rows }) => {-->
<!--      console.log(cols,rows)-->
<!--      webSocket.send(msgResize +-->
<!--          Base64.stringify(-->
<!--              Utf8.parse(-->
<!--                  JSON.stringify({-->
<!--                    columns: cols,-->
<!--                    rows: rows-->
<!--                  })-->
<!--              )-->
<!--          )-->
<!--      )-->
<!--    })-->
<!--    // 内容全屏显示-窗口大小发生改变时-->
<!--    // resizeScreen-->
<!--    window.addEventListener("resize", () =>{-->
<!--      fitAddon.fit()-->
<!--    },true)-->

<!--  },-->
<!--  methods: {-->
<!--    gobanckfun(){-->
<!--      location.reload()-->
<!--    }-->

<!--  },-->
<!--}-->
<!--</script>-->

<style scoped>
.active {
  color: #000;
  background: rgb(211, 24, 24);
}
.termButton {
  position: fixed;
  bottom: 10%;
  right: 8%;
  z-index: 888;
  border-radius: 40px;
}
#terminal {
  width: 100%;
  height: 100%;
}

.xterm {
  height: 100%;
  cursor: text;
  position: relative;
  user-select: none;
  -ms-user-select: none;
  -webkit-user-select: none;
  z-index: 99;
}

div {
  background-color: red;
  height: 100%;
  width: 100%;
}
* {
  margin: 0;
  padding: 0;
}

</style>
