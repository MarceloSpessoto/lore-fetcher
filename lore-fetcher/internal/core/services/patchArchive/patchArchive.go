package patchArchive
import (
  "lore-fetcher/internal/core/domain"
  "lore-fetcher/internal/core/ports/patchArchive"
)

type PatchArchiveService struct {
  repo patchArchive.PatchArchiveRepository
}

func NewPatchArchiveService(repo patchArchive.PatchArchiveRepository) *PatchArchiveService {
  return &PatchArchiveService{
    repo: repo,
  }
}

func (svc *PatchArchiveService) GetRecentPatches() []domain.Patch {
  return svc.repo.GetRecentPatches()
}
