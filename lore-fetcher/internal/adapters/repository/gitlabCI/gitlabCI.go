package gitlabCI

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
)

type GitlabCIRepository struct {
	instanceURL string
	projectID   string
	triggerToken string
	ref         string
}

func NewGitlabCIRepository() *GitlabCIRepository {
	return &GitlabCIRepository{
		instanceURL:  os.Getenv("GITLAB_URL"),
		projectID:    os.Getenv("GITLAB_PROJECT_ID"),
		triggerToken: os.Getenv("GITLAB_TRIGGER_TOKEN"),
		ref:          os.Getenv("GITLAB_REF"),
	}
}

func (r *GitlabCIRepository) TriggerPipeline(patch string) error {
	endpoint := fmt.Sprintf("%s/api/v4/projects/%s/trigger/pipeline", r.instanceURL, r.projectID)

	formData := url.Values{}
	formData.Set("token", r.triggerToken)
	formData.Set("ref", r.ref)
	formData.Set("variables[PATCH]", patch)

	resp, err := http.PostForm(endpoint, formData)
	if err != nil {
		return fmt.Errorf("gitlab CI trigger request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("gitlab CI trigger returned status %d", resp.StatusCode)
	}

	return nil
}
