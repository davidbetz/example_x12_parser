.PHONY:

# FILENAME=m837.84146.20180725.00996large
FILENAME=837.x12

json:
	go run . $(FILENAME) json

xml:
	go run . $(FILENAME) xml

defaults:
	go run . $(FILENAME)

linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -installsuffix cgo -ldflags '-w -s' -o x12decode

windows:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -installsuffix cgo -ldflags '-w -s' -o x12decode.exe

