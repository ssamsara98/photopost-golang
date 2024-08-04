FROM golang:1.21-alpine

# Required because go requires gcc to build
RUN apk update && apk upgrade && apk add --no-cache build-base bash git inotify-tools make
RUN go install github.com/rubenv/sql-migrate/...@latest

WORKDIR /app
COPY ["../go.mod", "../go.sum", "./"]
RUN go mod tidy
COPY . .

RUN cp template.production.env .env
RUN go build -buildvcs=false -o photopost

ENV PORT=8080
ENV SERVER_PORT=8080
ENV ENVIRONMENT=production
ENV ENV=production
ENV LOG_LEVEL=info
ENV LOG_OUTPUT=server.log
ENV DB_TYPE=postgres
ENV MAX_MULTIPART_MEMORY=10485760
ENV JWT_ACCESS_SECRET=9f4b7b1e0c2c166aa1733cdcf1f1c1a3f2ded287b5d28c2194378273f5530bd8
ENV JWT_REFRESH_SECRET=9bc0195061abea1f9461ba84412c9d8819594a1b166b0e1996ca097414224ced
ENV ACCESS_TOKEN_DURATION=24h
ENV REFRESH_TOKEN_DURATION=720h

CMD ["./photopost", "app:serve"]
