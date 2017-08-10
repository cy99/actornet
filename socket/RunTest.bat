set CURR=%cd%
cd ..\..\..\..\..
set GOPATH=%cd%
cd %CURR%
go test -c -o socket_test.exe github.com/davyxu/actornet/socket
start socket_test.exe -test.v -test.run TestServer
socket_test.exe -test.v -test.run TestClient
ping -n 1 127.1>nul
del socket_test.exe