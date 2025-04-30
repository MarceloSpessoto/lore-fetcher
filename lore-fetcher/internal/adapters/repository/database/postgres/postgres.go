package postgres
import (
  "lore-fetcher/internal/core/domain"
  "fmt"

  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/postgres"
)

type PostgresRepository struct {
  db *gorm.DB
}

func NewPostgresRepository() *PostgresRepository {
  host := "0.0.0.0"
  port := "5432"
  user := "postgres"
  password := "1234"
  dbname := "lore-fetcher"

  conn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", host, port, user, dbname, password)

  db, err := gorm.Open("postgres", conn)
  if err != nil {
    panic(err)
  }

	db.AutoMigrate(&domain.Job{})
  db.AutoMigrate(&domain.Patch{})
  return &PostgresRepository{db}
}

func (lf *PostgresRepository) SavePatch(patch *domain.Patch) error {
  if err := lf.db.Create(&patch).Error; err != nil {
    return err
  }
  return nil
}

func (lf *PostgresRepository) ReadPatches() ([]*domain.Patch, error) {
  var patches []*domain.Patch
  if err := lf.db.Find(&patches).Error; err != nil {
    return nil, err
  }
  return patches, nil
}

func (lf *PostgresRepository) ReadPatch(id string) (*domain.Patch, error) {
	var patch domain.Patch
	if err := lf.db.Where("id = ?", id).First(&patch).Error; err != nil {
		return nil, err
	}
	return &patch, nil
}

func (lf *PostgresRepository) SaveJob(job *domain.Job) error {
	if err := lf.db.Create(&job).Error; err != nil {
		return err
	}
	return nil
}

func (lf *PostgresRepository) ReadJobs() ([]*domain.Job, error) {
	var jobs []*domain.Job
	if err := lf.db.Find(&jobs).Error; err != nil {
		return nil, err
	}
	return jobs, nil
}

func (lf *PostgresRepository) ReadJob(id string) (*domain.Job, error) {
	var job domain.Job
	if err := lf.db.Where("id = ?", id).First(&job).Error; err != nil {
		return nil, err
	}
	return &job, nil
}
