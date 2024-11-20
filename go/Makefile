all: app

app: *.go go.mod go.sum
	GOOS=linux GOARCH=arm64 go build -o app
