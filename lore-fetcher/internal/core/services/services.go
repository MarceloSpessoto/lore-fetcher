package services
import (
  "lore-fetcher/internal/core/domain"
  "lore-fetcher/internal/core/ports"
  "github.com/google/uuid"
)

type LFService struct {
  repo ports.LFRepository
}

func NewLFService(repo ports.LFRepository) *LFService {
  return &LFService{
    repo: repo,
  }
}

func (lf *LFService) SavePatch(patch domain.Patch) error {
  patch.ID = uuid.New().String()
  return lf.repo.SavePatch(patch)
}

func (lf *LFService) ReadPatches() ([]*domain.Patch, error) {
  return lf.repo.ReadPatches()
}
