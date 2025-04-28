package database
import "lore-fetcher/internal/core/domain"

type DatabaseService interface {
  SavePatch(patch domain.Patch) error
  ReadPatches() ([]*domain.Patch, error)
	ReadPatch(id string) (*domain.Patch, error)
	SaveJob(job domain.Job) error
	ReadJobs() ([]*domain.Job, error)
	ReadJob(id string) (domain.Job, error)
}

type DatabaseRepository interface {
  SavePatch(patch domain.Patch) error
  ReadPatches() ([]*domain.Patch, error)
	ReadPatch(id string) (*domain.Patch, error)
	SaveJob(job domain.Job) error
	ReadJobs() ([]*domain.Job, error)
	ReadJob(id string) (*domain.Job, error)
}
