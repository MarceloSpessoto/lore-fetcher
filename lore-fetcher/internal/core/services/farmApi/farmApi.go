package farmApi
import (
  "lore-fetcher/internal/core/domain"
  "lore-fetcher/internal/core/ports/farmApi"
)

type FarmApiService struct {
  repo farmApi.FarmApiRepository
}

func NewFarmApiService(repo farmApi.FarmApiRepository) *FarmApiService {
  return &FarmApiService{
    repo: repo,
  }
}

func (svc *FarmApiService) GetCitronInstances() ([]*domain.CitronInstance, error) {
  return svc.repo.GetCitronInstances(patch)
}
