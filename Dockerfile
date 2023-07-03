FROM golang:1.20-alpine

# Required because go requires gcc to build
RUN apk update && apk upgrade && apk add --no-cache build-base bash git inotify-tools make
RUN echo $GOPATH
RUN go install github.com/rubenv/sql-migrate/...@latest
RUN go install github.com/go-delve/delve/cmd/dlv@latest

WORKDIR /app
COPY ["../go.mod", "../go.sum", "./"]
RUN go mod tidy
COPY . .

RUN cp template.prod.env .env
RUN go build -buildvcs=false -o binary

ENV PORT=8080
ENV ENVIRONMENT=production
ENV LOG_LEVEL=info
ENV DB_TYPE=postgres
ENV MAX_MULTIPART_MEMORY=10485760
ENV JWT_SECRET=71781f7e9db3b3d1e36235a4cc059896fed0eb0668ebcc70deb2ecb430ade39f

CMD ["./binary", "app:serve"]
