package evaluator

import(
  "github.com/MarceloSpessoto/lore-fetcher/internal/types"
  "fmt"
  "time"
)

type Evaluator struct {  
}

func (evaluator Evaluator) ReceivePatches(fetchBuffer chan types.Patch, resultBuffer chan types.Patch) {
  for {
    patch := <- fetchBuffer
    fmt.Println("[", time.Now(), "]: Testing patch '", patch.Title, "'")
    patch.ResultString, patch.Result = ParseResult(patch)
    
    fmt.Println(time.Now(), ": Patch '", patch.Title, "' got the following result: '", patch.ResultString, "'. Preparing to send the result.")
    resultBuffer <- patch
  }
  
}

func ParseResult(patch types.Patch) (string, bool) {
  return "Success", true
}
