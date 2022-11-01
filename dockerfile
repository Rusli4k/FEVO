FROM golang:1.19

ENV GOPATH=/

COPY . .

RUN go mod download 



RUN go get -d -v ./...

RUN go install -v ./... 

RUN go build  -o fevo ./cmd/rest/main.go

CMD ["./fevo"]
