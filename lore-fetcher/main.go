package main

import (
	"lore-fetcher/internal/fetcher"
)

func main(){
  fetcher := fetcher.NewFetcher()
  fetcher.FetchDaemon()
}
