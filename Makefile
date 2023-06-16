.PHONY:

linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -installsuffix cgo -ldflags '-w -s' -o x12_parser

windows:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -installsuffix cgo -ldflags '-w -s' -o x12_parser.exe
