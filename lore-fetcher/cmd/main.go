package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	gitlabCIRepo "lore-fetcher/internal/adapters/repository/gitlabCI"
	"lore-fetcher/internal/adapters/repository/database/postgres"
	"lore-fetcher/internal/adapters/repository/patchArchive/mailingList"
	"lore-fetcher/internal/adapters/repository/patchArchive/lore"
	"lore-fetcher/internal/core/ports/patchArchive"
	"lore-fetcher/internal/core/services/database"
	"lore-fetcher/internal/core/services/gitlabCI"
	patchArchiveSvc "lore-fetcher/internal/core/services/patchArchive"
	"lore-fetcher/internal/fetcher"
	"lore-fetcher/internal/logger"
	"lore-fetcher/cmd/ui/tui"
)

var (
	patchService    *patchArchiveSvc.PatchArchiveService
	postgresService *database.DatabaseService
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	lg := logger.New()
	log.SetOutput(lg)
	log.SetFlags(log.Ldate | log.Ltime)

	postgresRepository := postgres.NewPostgresRepository()
	postgresService = database.NewDatabaseService(postgresRepository)

	var patchRepo patchArchive.PatchArchiveRepository
	switch os.Getenv("PATCH_SOURCE") {
	case "lore":
		log.Println("Patch source: lore.kernel.org")
		patchRepo = lore.NewLoreRepository()
	default:
		log.Println("Patch source: IMAP mailing list")
		patchRepo = mailingList.NewMailingListRepository()
	}
	patchService = patchArchiveSvc.NewPatchArchiveService(patchRepo)

	gitlabCIRepository := gitlabCIRepo.NewGitlabCIRepository()
	gitlabCIService := gitlabCI.NewGitlabCIService(gitlabCIRepository)
	fetcher := fetcher.NewFetcher(*patchService, *postgresService, gitlabCIService)
	go fetcher.FetchDaemon()
	tui.RenderTuiMenu(*postgresService, lg)
}
