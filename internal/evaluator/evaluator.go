package evaluator

import(
  "github.com/MarceloSpessoto/lore-fetcher/internal/types"
  "fmt"
  "time"
)

type Evaluator struct {  
}

func (evaluator Evaluator) ReceivePatches(fetchBuffer chan types.Patch, resultBuffer chan string) {
  for {
    patch := <- fetchBuffer
    fmt.Println(time.Now(), ": Testing patch ", patch.Title)
    result, _ := ParseResult(patch)
    fmt.Println(time.Now(), ": Patch ", patch.Title, " got the following result: ", result, ". Preparing to send the result")
    resultBuffer <- result
  }
  
}

func ParseResult(patch types.Patch) (string, bool) {
  return "Success", true
}
