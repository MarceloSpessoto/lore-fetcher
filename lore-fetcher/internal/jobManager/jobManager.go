package jobManager

import (
	"lore-fetcher/internal/core/services/database"
	"lore-fetcher/internal/core/domain"
)

func MockedJobPipeline(dbsvc database.DatabaseService, patch domain.Patch) {
	// TODO: Implement the pipeline
	// 1) Compile the kernel image
	// 2) Use it to trigger a CI-Tron Job
	// 3) Track the job
	// For now, let's just mark the received job as completed
	
	newJob := domain.Job{}
	newJob.PatchID = patch.ID
	newJob.Status = "COMPLETED"
	dbsvc.SaveJob(&newJob)
}

