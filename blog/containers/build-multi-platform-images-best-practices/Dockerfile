FROM xcbeyond/go-builder:1.16

LABEL maintainer xcbeyond

WORKDIR /app

COPY multi-arch-test.go /app
RUN go build -o multi-arch-test /app/multi-arch-test.go

CMD ["./multi-arch-test"]