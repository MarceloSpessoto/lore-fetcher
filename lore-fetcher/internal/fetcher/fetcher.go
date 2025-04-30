package fetcher

import (
	"fmt"
	"regexp"
	"time"
  "lore-fetcher/internal/core/services/patchArchive"
	"lore-fetcher/internal/core/services/database"
  "lore-fetcher/internal/core/domain"
	"lore-fetcher/internal/jobManager"
)

type Fetcher struct {
	lastHref string
  patchArchiveService patchArchive.PatchArchiveService
	databaseService database.DatabaseService
}

func NewFetcher(pasvc patchArchive.PatchArchiveService, dbsvc database.DatabaseService) *Fetcher {
  var fetcher Fetcher
	fetcher.patchArchiveService = pasvc
	fetcher.databaseService = dbsvc
  return &fetcher
}

func (fetcher *Fetcher) FetchDaemon(){
	fmt.Println("Starting fetcher daemon...")
  for {
		patches := fetcher.patchArchiveService.GetRecentPatches()
    if len(patches) == 0 {
      time.Sleep(20 * time.Second)
      continue
    }
		patch := patches[0]
    fmt.Println("Most recent patch from all:\n", patch.Title)
		fetcher.databaseService.SavePatch(&patch)
		jobManager.MockedJobPipeline(fetcher.databaseService, patch)
    fmt.Println("New patch found: ", patch.Title)
		fetcher.lastHref = patch.PatchHref
    break
  }

  for {
    time.Sleep(30 * time.Second)

    fmt.Println("[", time.Now(), "]: Searching for new patches in all")
		patches := fetcher.patchArchiveService.GetRecentPatches()
    fetcher.processPatches(patches)
  }
}

func (fetcher *Fetcher) processPatches(patches []domain.Patch) {
  for i := 0; i < len(patches); i++ {
		patch := patches[i]

		if patch.PatchHref != fetcher.lastHref {
      if isPatch(patch.Title){
        fetcher.databaseService.SavePatch(&patch)
        fmt.Println("New patch found: ", patch.Title)
      }
    } else {
			fetcher.lastHref = patches[0].PatchHref
      break
    }
  }
}

// A temporary method to assert we're fetching a Patch message:
// checking if the title string contains the [PATCH] tag
func isPatch(patchTitle string) bool {
  fmt.Println(patchTitle)
  hasPattern, _ := regexp.Match(`.*\[.*PATCH.*\].*`, []byte(patchTitle))
  return hasPattern
}
