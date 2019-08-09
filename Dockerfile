FROM golang:1.12.7-alpine3.10 as builder

RUN adduser -D -g '' appuser

WORKDIR $GOPATH/src/mypackage/myapp/
COPY . .

RUN go get -d -v
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -a -installsuffix cgo -o /go/bin/app .


FROM scratch

COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /go/bin/app /go/bin/app

USER appuser

ENTRYPOINT ["/go/bin/app"]
