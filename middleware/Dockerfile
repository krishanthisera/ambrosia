FROM golang:1-alpine3.10 AS build

# Change working directory to go path
WORKDIR $GOPATH/github.com/krishanthisera/ambrosia.git

# Copy go source
COPY . .

#Download the deps and build go
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o /go/bin/middle .


# Stage 02
FROM scratch AS bin

ENV AMB_DB_NAME=topics
ENV AMB_DB_PASS=123
ENV AMB_RED_CHAN=chan-1
ENV AMB_DB_TBL=stories
ENV AMB_DB_USER=user
ENV AMB_RED_HOST=localhost:6379
ENV AMB_DB_HOST=localhost
ENV AMB_DB_PORT=5432
COPY --from=build /go/bin/middle /go/bin/middle

ENTRYPOINT [ "/go/bin/middle" ]
