package repositories

import (
	"context"

	"gorm.io/gorm"
)

type BaseRepository[T any] interface {
	Create(ctx context.Context, entity *T) error
	GetByID(ctx context.Context, id string) (*T, error)
	Update(ctx context.Context, entity *T) error
	Delete(ctx context.Context, id string) error
	List(ctx context.Context, offset, limit int) ([]T, error)
}

type baseRepository[T any] struct {
	db *gorm.DB
}

func NewBaseRepository[T any](db *gorm.DB) BaseRepository[T] {
	return &baseRepository[T]{db: db}
}

func (r *baseRepository[T]) Create(ctx context.Context, entity *T) error {
	return r.db.Create(entity).Error
}

func (r *baseRepository[T]) GetByID(ctx context.Context, id string) (*T, error) {
	var entity T
	if err := r.db.First(&entity, "id=?", id).Error; err != nil {
		return nil, err
	}
	return &entity, nil
}

func (r *baseRepository[T]) Update(ctx context.Context, entity *T) error {
	return r.db.Save(entity).Error
}

func (r *baseRepository[T]) Delete(ctx context.Context, id string) error {
	var entity T
	return r.db.Delete(&entity, "id=?", id).Error
}

func (r *baseRepository[T]) List(ctx context.Context, offset, limit int) ([]T, error) {
	var entities []T
	if err := r.db.Offset(offset).Limit(limit).Find(&entities).Error; err != nil {
		return nil, err
	}
	return entities, nil
}
