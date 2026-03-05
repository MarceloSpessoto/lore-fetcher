package fetcher

import (
	"log"
	"regexp"
	"time"
	"lore-fetcher/internal/core/services/patchArchive"
	"lore-fetcher/internal/core/services/database"
	"lore-fetcher/internal/core/services/gitlabCI"
	"lore-fetcher/internal/core/domain"
	"lore-fetcher/internal/jobManager"
)

type Fetcher struct {
	lastHref            string
	patchArchiveService patchArchive.PatchArchiveService
	databaseService     database.DatabaseService
	gitlabCIService     *gitlabCI.GitlabCIService
}

func NewFetcher(pasvc patchArchive.PatchArchiveService, dbsvc database.DatabaseService, glsvc *gitlabCI.GitlabCIService) *Fetcher {
	var fetcher Fetcher
	fetcher.patchArchiveService = pasvc
	fetcher.databaseService = dbsvc
	fetcher.gitlabCIService = glsvc
	return &fetcher
}

func (fetcher *Fetcher) FetchDaemon() {
	log.Println("Starting fetcher daemon...")
	for {
		patches := fetcher.patchArchiveService.GetRecentPatches()
		if len(patches) == 0 {
			time.Sleep(20 * time.Second)
			continue
		}
		patch := patches[0]
		log.Println("Most recent patch:", patch.Title)
		fetcher.databaseService.SavePatch(&patch)
		jobManager.MockedJobPipeline(fetcher.databaseService, patch)
		if err := fetcher.gitlabCIService.TriggerPipeline(patch.PatchHref); err != nil {
			log.Println("Failed to trigger GitLab CI pipeline:", err)
		}
		log.Println("New patch found:", patch.Title)
		fetcher.lastHref = patch.PatchHref
		break
	}

	for {
		time.Sleep(30 * time.Second)
		log.Println("Searching for new patches...")
		patches := fetcher.patchArchiveService.GetRecentPatches()
		fetcher.processPatches(patches)
	}
}

func (fetcher *Fetcher) processPatches(patches []domain.Patch) {
	for i := 0; i < len(patches); i++ {
		patch := patches[i]
		if patch.PatchHref != fetcher.lastHref {
			if isPatch(patch.Title) {
				fetcher.databaseService.SavePatch(&patch)
				if err := fetcher.gitlabCIService.TriggerPipeline(patch.PatchHref); err != nil {
					log.Println("Failed to trigger GitLab CI pipeline:", err)
				}
				log.Println("New patch found:", patch.Title)
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
	log.Println("Checking patch title:", patchTitle)
	hasPattern, _ := regexp.Match(`.*\[.*PATCH.*\].*`, []byte(patchTitle))
	return hasPattern
}
