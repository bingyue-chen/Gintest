all: app

app:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -tags=jsoniter -o app

clean:
	rm -f app

test:
	go test ./tests/... -v

testnc:
	go test ./tests/... -v -count=1

genmock:
	./mocks/mock.sh
