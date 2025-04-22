package patchArchive
import "lore-fetcher/internal/core/domain"

type PatchArchiveService interface {
  GetRecentPatches() []domain.Patch
}

type PatchArchiveRepository interface {
  GetRecentPatches() []domain.Patch
}
