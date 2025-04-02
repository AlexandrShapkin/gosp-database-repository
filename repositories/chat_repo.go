package repositories

import (
	"context"
	"errors"

	"github.com/AlexandrShapkin/gosp-schema/models"
	"github.com/AlexandrShapkin/gosp-schema/repositories"
	"gorm.io/gorm"
)

type GormChatRepository struct {
	db *gorm.DB
}

// Create implements repositories.ChatRepository.
func (g *GormChatRepository) Create(ctx context.Context, chat *models.Chat) error {
	return g.db.WithContext(ctx).Create(chat).Error
}

// DeleteByID implements repositories.ChatRepository.
func (g *GormChatRepository) DeleteByID(ctx context.Context, id string) error {
	return g.db.WithContext(ctx).Where("id = ?", id).Delete(&models.Chat{}).Error
}

// GetAll implements repositories.ChatRepository.
func (g *GormChatRepository) GetAll(ctx context.Context, offset int, limit int) ([]models.Chat, error) {
	var chats []models.Chat
	query := g.db.WithContext(ctx).Preload("Messages").Preload("Participants")

	if limit > 0 {
		query = query.Offset(offset).Limit(limit)
	}

	err := query.Find(&chats).Error
	return chats, err
}

// GetByID implements repositories.ChatRepository.
func (g *GormChatRepository) GetByID(ctx context.Context, id string) (*models.Chat, error) {
	var chat models.Chat
	err := g.db.WithContext(ctx).Preload("Messages").Preload("Participants").First(&chat, "id = ?", id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &chat, err
}

// GetByUser implements repositories.ChatRepository.
func (g *GormChatRepository) GetByUser(ctx context.Context, userID string, offset int, limit int) ([]models.Chat, error) {
	var chats []models.Chat
	query := g.db.WithContext(ctx).
		Preload("Messages").
		Preload("Participants").
		Joins("JOIN chat_participants ON chat_participants.chat_id = chats.id").
		Where("chat_participants.user_id = ?", userID)

	if limit > 0 {
		query = query.Offset(offset).Limit(limit)
	}

	err := query.Find(&chats).Error
	return chats, err
}

// Update implements repositories.ChatRepository.
func (g *GormChatRepository) Update(ctx context.Context, chat *models.Chat) error {
	return g.db.WithContext(ctx).Save(chat).Error
}

// UpdateName implements repositories.ChatRepository.
func (g *GormChatRepository) UpdateName(ctx context.Context, id string, name string) error {
	return g.db.WithContext(ctx).
		Model(&models.Chat{}).
		Where("id = ?", id).
		Update("name", name).
		Error
}

func NewChatRepository(db *gorm.DB) repositories.ChatRepository {
	return &GormChatRepository{
		db: db,
	}
}
