package main

import (
	"log"
	"github.com/joho/godotenv"
	"lore-fetcher/internal/adapters/repository/database/postgres"
	"lore-fetcher/internal/adapters/repository/patchArchive/lore"
	gitlabCIRepo "lore-fetcher/internal/adapters/repository/gitlabCI"
	"lore-fetcher/internal/core/services/database"
	"lore-fetcher/internal/core/services/patchArchive"
	"lore-fetcher/internal/core/services/gitlabCI"
	"lore-fetcher/internal/fetcher"
	"lore-fetcher/cmd/ui/tui"
)

var (
	loreService     *patchArchive.PatchArchiveService
	postgresService *database.DatabaseService
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}
	postgresRepository := postgres.NewPostgresRepository()
	postgresService = database.NewDatabaseService(postgresRepository)
	loreRepository := lore.NewLoreRepository()
	loreService = patchArchive.NewPatchArchiveService(loreRepository)
	gitlabCIRepository := gitlabCIRepo.NewGitlabCIRepository()
	gitlabCIService := gitlabCI.NewGitlabCIService(gitlabCIRepository)
	fetcher := fetcher.NewFetcher(*loreService, *postgresService, gitlabCIService)
	go fetcher.FetchDaemon()
	tui.RenderTuiMenu(*postgresService)
}
