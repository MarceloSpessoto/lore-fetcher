package main

import (
	"fmt"
  "sync"

	"github.com/MarceloSpessoto/lore-fetcher/internal/configurator"
	"github.com/MarceloSpessoto/lore-fetcher/internal/evaluator"
	"github.com/MarceloSpessoto/lore-fetcher/internal/fetcher"
	"github.com/MarceloSpessoto/lore-fetcher/internal/mailer"
	"github.com/MarceloSpessoto/lore-fetcher/internal/types"
)

func main(){
  var wg sync.WaitGroup
  fetcher := fetcher.NewFetcher()
  evaluator := evaluator.Evaluator{}
  mailer := mailer.Mailer{}
  configurator := configurator.Configurator{}
  configurator.ParseConfiguration(fetcher, &evaluator, &mailer, "./lore-fetcher")
  fmt.Println(fetcher)
  fetchBuffer := make(chan types.Patch, 100)
  resultBuffer := make(chan types.Patch, 100)
  wg.Add(3)
  go fetcher.FetchDaemon(fetchBuffer)
  go evaluator.ReceivePatches(fetchBuffer, resultBuffer)
  go mailer.SendResults(resultBuffer)
  wg.Wait()
}
