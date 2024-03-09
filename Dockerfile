FROM golang:latest as builder
WORKDIR /app

COPY ["main.go", "go.mod", "go.sum", "go.work", "/app/"]
COPY assets /app/assets
COPY telegram /app/telegram
COPY utils /app/utils

ENV CGO_ENABLED=0
ENV GOOS=linux
RUN ["go", "build", "-o", "main"]

FROM alpine as runner
WORKDIR /app
COPY assets /app/assets
COPY --from=builder "/app/main" "/app/"
ENV DOCKER=1

CMD [ "/app/main" ]