package configurator

import (
	"testing"

	"lore-fetcher/internal/evaluator"
	"lore-fetcher/internal/fetcher"
	"lore-fetcher/internal/mailer"
)

func TestConfigurator(t *testing.T){
  fetcher := fetcher.NewFetcher()
  evaluator := evaluator.Evaluator{}
  mailer := mailer.Mailer{}
  configurator := Configurator{}
  configurator.ParseConfiguration(fetcher, &evaluator, &mailer, "../../testdata/")
  if(mailer.ToMail != "mock_receiver"){
    t.Errorf("got %s, expected %s", mailer.ToMail, "mock_receiver")
  }
}
