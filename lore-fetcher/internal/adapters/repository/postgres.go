package repository

import (
  "lore-fetcher/internal/core/domain"
  "fmt"

  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/postgres"
)

type LFPostgresRepository struct {
  db *gorm.DB
}

func NewLFPostgresRepository() *LFPostgresRepository {
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

  db.AutoMigrate(&domain.Patch{})
  return &LFPostgresRepository{db}
}

func (lf *LFPostgresRepository) SavePatch(patch domain.Patch) error {
  if err := lf.db.Create(&patch).Error; err != nil {
    return err
  }
  return nil
}

func (lf *LFPostgresRepository) ReadPatches() ([]*domain.Patch, error) {
  var patches []*domain.Patch
  if err := lf.db.Find(&patches).Error; err != nil {
    return nil, err
  }
  return patches, nil
}
