package gitlabCI

type GitlabCIService interface {
	TriggerPipeline(patchURL string) error
}

type GitlabCIRepository interface {
	TriggerPipeline(patchURL string) error
}
