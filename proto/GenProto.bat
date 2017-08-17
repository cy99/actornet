set CURR=%cd%
cd ..\..\..\..\..
set GOPATH=%cd%
cd %CURR%

go build -o objprotogen.exe github.com/davyxu/cellnet/objprotogen
@IF %ERRORLEVEL% NEQ 0 pause

objprotogen.exe --out objproto_gen.go test.go sys.go
@IF %ERRORLEVEL% NEQ 0 pause

ping -n 1 127.1>nul

del /q objprotogen.exe