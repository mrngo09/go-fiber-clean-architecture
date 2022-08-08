server-start:
	nodemon --exec go run .\src\app\main.go --ext go
build-prod:
	go build .\src\app\main.go  
run-prod:
	.\main.exe