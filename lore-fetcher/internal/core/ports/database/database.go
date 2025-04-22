package database
import "lore-fetcher/internal/core/domain"

type DatabaseService interface {
  SavePatch(patch domain.Patch) error
  ReadPatches() ([]*domain.Patch, error)
}

type DatabaseRepository interface {
  SavePatch(patch domain.Patch) error
  ReadPatches() ([]*domain.Patch, error)
}
