#!/bin/bash

go build lore-fetcher/main.go
cp lore-fetcher/lore-fetcher /usr/bin/

mkdir /etc/lore-fetcher
cp config.toml /etc/lore-fetcher/lore-fetcher.toml
