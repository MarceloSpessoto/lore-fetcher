package fetcher

import (
	"log"
	"regexp"
	"strconv"
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
	// Collect new patches (up to lastHref)
	var newPatches []domain.Patch
	for i := 0; i < len(patches); i++ {
		if patches[i].PatchHref == fetcher.lastHref {
			break
		}
		if isPatch(patches[i].Title) {
			newPatches = append(newPatches, patches[i])
		}
	}
	if len(newPatches) == 0 {
		if len(patches) > 0 {
			fetcher.lastHref = patches[0].PatchHref
		}
		return
	}

	// Find which series totals have a cover letter (index 0) present
	seriesWithCoverLetter := map[int]bool{}
	for _, p := range newPatches {
		idx, total := parsePatchSeriesIndex(p.Title)
		if idx == 0 && total > 0 {
			seriesWithCoverLetter[total] = true
		}
	}

	for _, patch := range newPatches {
		idx, total := parsePatchSeriesIndex(patch.Title)
		if total > 1 && idx > 0 {
			// Part of a multi-patch series; skip unless it's the first patch
			// and no cover letter exists for this series size.
			if !(idx == 1 && !seriesWithCoverLetter[total]) {
				continue
			}
		}
		fetcher.databaseService.SavePatch(&patch)
		if err := fetcher.gitlabCIService.TriggerPipeline(patch.PatchHref); err != nil {
			log.Println("Failed to trigger GitLab CI pipeline:", err)
		}
		log.Println("New patch found:", patch.Title)
	}

	fetcher.lastHref = patches[0].PatchHref
}

// isPatch checks if the title contains the [PATCH] tag.
func isPatch(patchTitle string) bool {
	hasPattern, _ := regexp.MatchString(`\[.*PATCH.*\]`, patchTitle)
	return hasPattern
}

// parsePatchSeriesIndex extracts (index, total) from a [PATCH X/N] subject.
// Returns (-1, -1) if not a numbered series.
func parsePatchSeriesIndex(title string) (int, int) {
	re := regexp.MustCompile(`\[.*?PATCH\s+(\d+)/(\d+).*?\]`)
	m := re.FindStringSubmatch(title)
	if m == nil {
		return -1, -1
	}
	idx, _ := strconv.Atoi(m[1])
	total, _ := strconv.Atoi(m[2])
	return idx, total
}
