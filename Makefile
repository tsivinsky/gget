build:
	GOARCH=amd64 GOOS=linux go build -o build/gget_linux
	GOARCH=amd64 GOOS=windows go build -o build/gget_window.exe
	GOARCH=amd64 GOOS=darwin go build -o build/gget_macos
	GOARCH=arm64 GOOS=darwin go build -o build/gget_macos_arm

clean:
	rm -r build
