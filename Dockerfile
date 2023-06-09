############# STAGE 1 #############

FROM golang:1.20.5-alpine3.18 as builder

WORKDIR /app

COPY . ./

RUN go mod download

RUN go build -ldflags "-w -s -extldflags '-static'" -a -o main

############# STAGE 2 #############

FROM gcr.io/distroless/static

WORKDIR /app

COPY --from=builder /app/main ./

ENTRYPOINT ["/app/main"]