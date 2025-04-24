







build:
	go build -o ./bin/watch_linuxarm64 . && \
	GOOS=windows GOARCH=arm64 go build -o ./bin/watch_arm64.exe . && \
	GOOS=windows GOARCH=amd64 go build -o ./bin/watch_amd64.exe . && \
	GOOS=darwin GOARCH=arm64 go build -o ./bin/watch_darwin .
