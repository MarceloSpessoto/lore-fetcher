package configurator

import (
	"fmt"
	"github.com/spf13/viper"
  "github.com/MarceloSpessoto/lore-fetcher/internal/fetcher"
  "github.com/MarceloSpessoto/lore-fetcher/internal/evaluator"
  "github.com/MarceloSpessoto/lore-fetcher/internal/mailer"
)

type Configurator struct {
}

func (configurator Configurator) ParseConfiguration(fetcher *fetcher.Fetcher, evaluator *evaluator.Evaluator, mailer *mailer.Mailer){
  viper.SetConfigType("toml")
  viper.SetConfigName("config")
  viper.AddConfigPath("./lore-fetcher")
  if err := viper.ReadInConfig(); err != nil {
    panic(fmt.Errorf("error reading config file: %w", err))
  }

  fetcher.MailingList = viper.GetString("fetcher.mailing_list")
  fetcher.FetchInterval = viper.GetInt("fetcher.fetch_interval")

  mailer.FromMail = viper.GetString("mailer.from_mail")
  mailer.ToMail = viper.GetString("mailer.to_mail")
  mailer.Password = viper.GetString("password")
}
