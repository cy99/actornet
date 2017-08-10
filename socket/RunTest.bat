set CURR=%cd%
cd ..\..\..\..\..
set GOPATH=%cd%
cd %CURR%
go test -c -o sockettest.exe github.com/davyxu/actornet/socket
start sockettest.exe -test.v -test.run TestServer
sockettest.exe -test.v -test.run TestClient
ping -n 1 127.1>nul
del sockettest.exe