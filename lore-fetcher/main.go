package main

import (
	"lore-fetcher/internal/configurator"
	"lore-fetcher/internal/fetcher"
	"flag"
	"fmt"
)

func main(){

  var configurator *configurator.Configurator = configurator.NewConfigurator()
  configurator.ParseConfiguration()

  var options = make(map[string]*bool)
  var params = make(map[string]*string)

  parseOptions(options)
  parseParameters(params)

  flag.Parse()

  for key, value := range params {
    if *params[key] != "" {
      configurator.SetConfiguration(key, *value)
    }
  }

  var chosen_option string = ""
  for key, value := range options {
    if *value {
      if chosen_option != "" {
        fmt.Println("ERROR: Multiple options have been chosen")
        return
      } else {
        chosen_option = key
      }
    }
  }

  if chosen_option == "" {
    fmt.Println("Choose a single option below to start using lore fetcher:")
    displayHelp()
    return
  }

  switch chosen_option {
    case "fetch":
    fetcher := fetcher.NewFetcher()
    configurator.ConfigureFetch(fetcher)
    fetcher.FetchDaemon()
    case "apply":
    fmt.Println("APPLYING")
    case "send":
    fmt.Println("SENDING")
    case "help":
    displayHelp()
  }

}

func displayHelp(){
  fmt.Println("* fetch: Enable fetch mode - keep listening to new patches")
  fmt.Println("* apply: Use apply operation - prepare a kernel repo with a given patch applied to it")
  fmt.Println("* send: Send test results to a given mail address or mail thread")
  fmt.Println("* help: Display this text")
  fmt.Println("The main operations require the specification of additional parameters, through /etc/lorefetcher config file or CLI flags.")
  fmt.Println("To display flags or more information about an option, use lore-fetcher <option> --help")
}

func parseOptions(options map[string]*bool){
  options["fetch"] = flag.Bool("fetch", false, "Enable fetch mode - keep listening to new patches")
  options["apply"] = flag.Bool("apply", false, "Use apply operation - prepare a kernel repo with a given patch applied to it")
  options["send"] = flag.Bool("send", false, "Send test results to a given mail address or mail thread")
  options["help"] = flag.Bool("help", false, "Display basic information about lore fetcher and its possible command options")
}

func parseParameters(params map[string]*string){
  params["mailing-list"] = flag.String("mailing-list", "", "[Requires --fetch] Set the mailing list to be tracked")
  params["fetch-interval"] = flag.String("fetch-interval", "", "[Requires --fetch] Interval in seconds between each attempt to find new patches")
  params["jenkins-server"] = flag.String("jenkins-server", "", "[Requires --fetch] Server URL (<jenkins-address>:<port>)")
  params["jenkins-pipeline"] = flag.String("jenkins-pipeline", "", "[Requires --fetch] The name of the Jenkins pipeline to trigger")
  params["jenkins-token"] = flag.String("jenkins-token", "", "[Requires --fetch] The name of the Jenkins auth token to trigger pipeline")

  params["from-mail"] = flag.String("from-mail", "", "[Requires --send] Mail address that will be used to send test reports")
  params["to-mail"] = flag.String("to-mail", "", "[Requires --send] Mail address where report will be sent")
  params["password"] = flag.String("auth", "", "[Requires --send] authentication string to use 'from_mail'. Required for Gmail addresses, for example")
}
