ENTRY_FILE := "./cmd/gget/main.go"

build:
	GOARCH=amd64 GOOS=linux go build -o build/gget_linux ${ENTRY_FILE}
	GOARCH=amd64 GOOS=windows go build -o build/gget_window.exe ${ENTRY_FILE}
	GOARCH=amd64 GOOS=darwin go build -o build/gget_macos ${ENTRY_FILE}
	GOARCH=arm64 GOOS=darwin go build -o build/gget_macos_arm ${ENTRY_FILE}

clean:
	rm -r build
