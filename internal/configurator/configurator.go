package configurator

import (
	"fmt"
	"strconv"

  "github.com/MarceloSpessoto/lore-fetcher/internal/fetcher"
	"github.com/spf13/viper"
)

type Configurator struct {
  configuration map[string]string
}

func NewConfigurator() *Configurator {
  var configurator Configurator
  configurator.configuration = make(map[string]string)
  return &configurator
}

func (configurator *Configurator) ParseConfiguration(){
  viper.SetConfigType("toml")
  viper.SetConfigName("lore-fetcher")
  viper.AddConfigPath("/etc/lore-fetcher/")
  if err := viper.ReadInConfig(); err != nil {
    panic(fmt.Errorf("error reading config file: %w", err))
  }

  configurator.configuration["mailing-list"] = viper.GetString("fetcher.mailing-list")
  configurator.configuration["fetch-interval"] = strconv.Itoa(viper.GetInt("fetcher.fetch-interval"))

  configurator.configuration["from-mail"] = viper.GetString("mailer.from_mail")
  configurator.configuration["to-mail"] = viper.GetString("mailer.to_mail")
  configurator.configuration["password"] = viper.GetString("mailer.password")
}

func (configurator *Configurator) ConfigureFetch(fetcher *fetcher.Fetcher){
  fetcher.MailingList = configurator.configuration["mailing-list"]
  fetcher.FetchInterval, _ = strconv.Atoi(configurator.configuration["mailing-list"])
}
