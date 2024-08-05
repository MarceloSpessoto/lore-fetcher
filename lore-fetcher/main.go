package main

import (
	//"fmt"
	//"sync"

	//"github.com/MarceloSpessoto/lore-fetcher/internal/configurator"
	//"github.com/MarceloSpessoto/lore-fetcher/internal/evaluator"
	"github.com/MarceloSpessoto/lore-fetcher/internal/fetcher"
	//"github.com/MarceloSpessoto/lore-fetcher/internal/mailer"
	//"github.com/MarceloSpessoto/lore-fetcher/internal/types"
	"flag"
	"fmt"
)

//func main(){
//  var wg sync.WaitGroup
//  fetcher := fetcher.NewFetcher()
//  evaluator := evaluator.Evaluator{}
//  mailer := mailer.Mailer{}
//  configurator := configurator.Configurator{}
//  configurator.ParseConfiguration(fetcher, &evaluator, &mailer, "./lore-fetcher")
//  fmt.Println(fetcher)
//  fetchBuffer := make(chan types.Patch, 100)
//  resultBuffer := make(chan types.Patch, 100)
//  wg.Add(3)
//  go fetcher.FetchDaemon(fetchBuffer)
//  go evaluator.ReceivePatches(fetchBuffer, resultBuffer)
//  go mailer.SendResults(resultBuffer)
//  wg.Wait()
//}

func main(){

  var options = make(map[string]*bool)
  var params = make(map[string]*string)

  parseOptions(options)
  parseParameters(params)

  flag.Parse()

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
    fetcher := fetcher.NewFetcher(*params["mailing_list"], *params["fetch_interval"])
    fetcher.FetchDaemon()
    case "apply":
    fmt.Println("APPLYING")
    case "send":
    fmt.Println("SENDING")
    case "help":
    fmt.Println("HELPING")
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
  params["mailing_list"] = flag.String("mailing_list", "all", "Set the mailing list to be tracked")
  params["fetch_interval"] = flag.String("fetch_interval", "60", "Interval in seconds between each attempt to find new patches")

  params["from_mail"] = flag.String("from_mail", "", "Mail address that will be used to send test reports")
  params["to_mail"] = flag.String("to_mail", "", "Mail address where report will be sent")
  params["password"] = flag.String("auth", "", "authentication string to use 'from_mail'. Required for Gmail addresses, for example")
}
