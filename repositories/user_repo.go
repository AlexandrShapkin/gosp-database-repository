package repositories

import (
	"context"
	"github.com/AlexandrShapkin/gosp-schema/models"
	"github.com/AlexandrShapkin/gosp-schema/repositories"
	"gorm.io/gorm"
)

type GormUserRepository struct {
	db *gorm.DB
}

// Create implements repositories.UserRepository.
func (g *GormUserRepository) Create(ctx context.Context, user *models.User) error {
	return g.db.WithContext(ctx).Model(models.User{}).Create(user).Error
}

// DeleteByID implements repositories.UserRepository.
func (g *GormUserRepository) DeleteByID(ctx context.Context, id string) error {
	return g.db.WithContext(ctx).Model(models.User{}).Delete(&models.User{}, "id = ?", id).Error
}

// GetAll implements repositories.UserRepository.
func (g *GormUserRepository) GetAll(ctx context.Context) ([]models.User, error) {
	var users []models.User
	err := g.db.WithContext(ctx).Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

// GetByID implements repositories.UserRepository.
func (g *GormUserRepository) GetByID(ctx context.Context, id string) (*models.User, error) {
	var user models.User
	err := g.db.WithContext(ctx).First(&user, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// GetByUsername implements repositories.UserRepository.
func (g *GormUserRepository) GetByUsername(ctx context.Context, username string) (*models.User, error) {
	var user models.User
	err := g.db.WithContext(ctx).First(&user, "username = ?", username).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// Update implements repositories.UserRepository.
func (g *GormUserRepository) Update(ctx context.Context, user *models.User) error {
	return g.db.WithContext(ctx).Save(user).Error
}

// UpdateStatus implements repositories.UserRepository.
func (g *GormUserRepository) UpdateStatus(ctx context.Context, id string, status string) error {
	return g.db.WithContext(ctx).
		Model(&models.User{}).
		Where("id = ?", id).
		Update("status", status).
		Error
}

func NewUserRepository(db *gorm.DB) repositories.UserRepository {
	return &GormUserRepository{
		db: db,
	}
}
