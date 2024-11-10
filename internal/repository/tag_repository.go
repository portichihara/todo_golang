package repository

import (
    "todo-api/internal/domain"
    "gorm.io/gorm"
)

type TagRepository interface {
    FindOrCreate(name string) (*domain.Tag, error)
    FindByNames(names []string) ([]domain.Tag, error)
}

type tagRepository struct {
    db *gorm.DB
}

func NewTagRepository(db *gorm.DB) TagRepository {
    return &tagRepository{db: db}
}

func (r *tagRepository) FindOrCreate(name string) (*domain.Tag, error) {
    var tag domain.Tag
    err := r.db.Where(domain.Tag{Name: name}).FirstOrCreate(&tag).Error
    if err != nil {
        return nil, err
    }
    return &tag, nil
}

func (r *tagRepository) FindByNames(names []string) ([]domain.Tag, error) {
    var tags []domain.Tag
    err := r.db.Where("name IN ?", names).Find(&tags).Error
    if err != nil {
        return nil, err
    }
    return tags, nil
}