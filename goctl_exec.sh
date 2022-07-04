1.安装
git rm -r --cached .
git add .
git commit -m 'update .gitignore'
框架安装 go get -u github.com/zeromicro/go-zero
框架代码生成工具安装 go get -u github.com/zeromicro/go-zero/tools/goctl
grpcui
grpcui -plaintext localhost:12345
sql2pb
sql2pb -go_package ./pb -host 114.67.110.204 -package pb -password password -port 32033 -schema hzjt-oam -table sys_dept,sys_role,sys_role_dept,sys_user -service_name usersrv -user root > sys.proto
goctl model mysql datasource -url="root:password@tcp(114.67.110.204:32033)/hzjt-oam" -table="sys*"  -dir="model" --style=go_zero
goctl model mysql datasource -url="root:password@tcp(114.67.110.204:32033)/hzjt-oam" -table="sys*"  -dir="model/user" --style=go_zero
goctl api plugin -p goctl-swagger="swagger -filename hzjt.json" --api admin.api --dir .
goctl api doc --dir ./

2.创建api
进到api/doc/目录执行
goctl api -o admin.api
goctl api go -api admin.api -dir ../ --style go_zero

front-pai
front-api/doc
goctl api -o front.api
goctl api go -api front.api -dir ../

3.创建rpc
进到rpc/sys/目录操作
goctl rpc protoc sys.proto --go_out=./ --go-grpc_out=./ --zrpc_out=. --style=go_zero

进到rpc/ums/目录操作
goctl rpc template -o ums.proto
goctl rpc protoc ums.proto --go_out=./ --go-grpc_out=./ --zrpc_out=.

进到rpc/pms/目录操作
goctl rpc template -o pms.proto
goctl rpc protoc pms.proto --go_out=./ --go-grpc_out=./ --zrpc_out=.

进到rpc/oms/目录操作
goctl rpc template -o oms.proto
goctl rpc protoc oms.proto --go_out=./ --go-grpc_out=./ --zrpc_out=.

进到rpc/sms/目录操作
goctl rpc template -o sms.proto
goctl rpc protoc sms.proto --go_out=./ --go-grpc_out=./ --zrpc_out=.

进到rpc/pay/目录操作
goctl rpc template -o pay.proto
goctl rpc protoc pay.proto --go_out=./ --go-grpc_out=./ --zrpc_out=.

进到rpc/cms/目录操作
goctl rpc template -o cms.proto
goctl rpc protoc cms.proto --go_out=./ --go-grpc_out=./ --zrpc_out=.

4.创建model
进到rpc/目录操作
goctl model mysql ddl -c -src book.sql -dir .
goctl model mysql datasource -url="haozheyu_1:123123Aa@@@tcp(qc84hv2knijar7uxb5tgq36woesy1m9z.mysql.qingcloud.link:3306)/oam_system" -table="oam_user*" -dir ./model/user --style=go_zero
goctl model mysql datasource -url="root:123456@tcp(127.0.0.1:3306)/gozero" -table="sys*" -dir ./model/sysmodel

goctl model mysql datasource -url="root:123456@tcp(127.0.0.1:3306)/gozero" -table="ums*" -dir ./model/umsmodel
goctl model mysql datasource -url="root:123456@tcp(127.0.0.1:3306)/gozero" -table="pms*" -dir ./model/pmsmodel
goctl model mysql datasource -url="root:123456@tcp(127.0.0.1:3306)/gozero" -table="oms*" -dir ./model/omsmodel
goctl model mysql datasource -url="root:123456@tcp(127.0.0.1:3306)/gozero" -table="sms*" -dir ./model/smsmodel
goctl model mysql datasource -url="root:123456@tcp(127.0.0.1:3306)/gozero" -table="pay*" -dir ./model/paymodel
