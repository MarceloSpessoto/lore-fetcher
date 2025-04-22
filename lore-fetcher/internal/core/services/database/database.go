package database
import (
  "lore-fetcher/internal/core/domain"
  "lore-fetcher/internal/core/ports/database"
  "github.com/google/uuid"
)

type DatabaseService struct {
  repo database.DatabaseRepository
}

func NewDatabaseService(repo database.DatabaseRepository) *DatabaseService {
  return &DatabaseService{
    repo: repo,
  }
}

func (svc *DatabaseService) SavePatch(patch domain.Patch) error {
  patch.ID = uuid.New().String()
  return svc.repo.SavePatch(patch)
}

func (svc *DatabaseService) ReadPatches() ([]*domain.Patch, error) {
  return svc.repo.ReadPatches()
}
