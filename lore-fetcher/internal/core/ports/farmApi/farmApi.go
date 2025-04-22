package farmApi
import "lore-fetcher/internal/core/domain"

type FarmApiService interface {
  GetCitronInstances() ([]*domain.CitronInstance, error)
}

type FarmApiRepository interface {
  GetCitronInstances() ([]*domain.CitronInstance, error)
}
