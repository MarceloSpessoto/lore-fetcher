package configurator

import (
	"fmt"
	"strconv"

	"lore-fetcher/internal/fetcher"

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
  configurator.configuration["jenkins-server"] = viper.GetString("fetcher.jenkins-server")
  configurator.configuration["jenkins-token"] = viper.GetString("fetcher.jenkins-token")
  configurator.configuration["jenkins-pipeline"] = viper.GetString("fetcher.jenkins-pipeline")

  configurator.configuration["from-mail"] = viper.GetString("mailer.from-mail")
  configurator.configuration["to-mail"] = viper.GetString("mailer.to-mail")
  configurator.configuration["password"] = viper.GetString("mailer.password")
}

func (configurator *Configurator) IsConfigurated(key string) bool{
  return configurator.configuration[key] != ""
}

func (configurator *Configurator) GetConfiguration(key string) string{
  return configurator.configuration[key]
}

func (configurator *Configurator) SetConfiguration(key string, value string) {
  configurator.configuration[key] = value
}

func (configurator *Configurator) ConfigureFetch(fetcher *fetcher.Fetcher){
  fetcher.MailingList = configurator.configuration["mailing-list"]
  fetcher.FetchInterval, _ = strconv.Atoi(configurator.configuration["fetch-interval"])
  fetcher.JenkinsServer = configurator.configuration["jenkins-server"]
  fetcher.JenkinsPipeline = configurator.configuration["jenkins-pipeline"]
  fetcher.JenkinsToken = configurator.configuration["jenkins-token"]
}
