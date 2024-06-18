windows:
	export GOOS=windows
	go build -ldflags="-s -w" -trimpath -o ./bin/windows/auto-merge.exe ./main.go

linux:
	export GOOS=linux
	go build -ldflags="-s -w" -trimpath -o ./bin/linux/auto-merge ./main.go
