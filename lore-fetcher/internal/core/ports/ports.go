package ports
import "lore-fetcher/internal/core/domain"

type LFService interface {
  SavePatch(patch domain.Patch) error
  ReadPatches() ([]*domain.Patch, error)
}

type LFRepository interface {
  SavePatch(patch domain.Patch) error
  ReadPatches() ([]*domain.Patch, error)
}
