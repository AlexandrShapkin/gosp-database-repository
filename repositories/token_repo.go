package repositories

import (
	"context"
	"time"

	"github.com/AlexandrShapkin/gosp-schema/models"
	"github.com/AlexandrShapkin/gosp-schema/repositories"
	"gorm.io/gorm"
)

type GormTokenRepository struct {
	db *gorm.DB
}

// Create implements repositories.TokenRepository.
func (g *GormTokenRepository) Create(ctx context.Context, token *models.Token) error {
	return g.db.WithContext(ctx).Create(token).Error
}

// DeleteExpiredTokens implements repositories.TokenRepository.
func (g *GormTokenRepository) DeleteExpiredTokens(ctx context.Context) error {
	return g.db.WithContext(ctx).Where("expires_at < ?", time.Now()).Delete(&models.Token{}).Error
}

// DeleteToken implements repositories.TokenRepository.
func (g *GormTokenRepository) DeleteToken(ctx context.Context, tokenString string) error {
	return g.db.WithContext(ctx).Where("refresh_token = ?", tokenString).Delete(&models.Token{}).Error
}

// DeleteTokensByUser implements repositories.TokenRepository.
func (g *GormTokenRepository) DeleteTokensByUser(ctx context.Context, userID string) error {
	return g.db.WithContext(ctx).Where("user_id = ?", userID).Delete(&models.Token{}).Error
}

// GetToken implements repositories.TokenRepository.
func (g *GormTokenRepository) GetToken(ctx context.Context, tokenString string) (*models.Token, error) {
	var token models.Token
	err := g.db.WithContext(ctx).Where("refresh_token = ?", tokenString).First(&token).Error
	if err != nil {
		return nil, err
	}
	return &token, nil
}

func NewTokenRepository(db *gorm.DB) repositories.TokenRepository {
	return &GormTokenRepository{
		db: db,
	}
}
