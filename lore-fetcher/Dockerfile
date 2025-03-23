FROM golang

COPY . /lore-fetcher
WORKDIR /lore-fetcher/
RUN go build cmd/main.go
ENTRYPOINT ./main
