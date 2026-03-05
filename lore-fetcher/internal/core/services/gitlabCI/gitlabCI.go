package gitlabCI

import "lore-fetcher/internal/core/ports/gitlabCI"

type GitlabCIService struct {
	repo gitlabCI.GitlabCIRepository
}

func NewGitlabCIService(repo gitlabCI.GitlabCIRepository) *GitlabCIService {
	return &GitlabCIService{repo: repo}
}

func (svc *GitlabCIService) TriggerPipeline(patchURL string) error {
	return svc.repo.TriggerPipeline(patchURL)
}
