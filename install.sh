#!/bin/bash

go build lore-fetcher/main.go
mv main /usr/bin/lore-fetcher

mkdir -p /etc/lore-fetcher
cp config.toml /etc/lore-fetcher/lore-fetcher.toml
