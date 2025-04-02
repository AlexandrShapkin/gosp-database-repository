package repositories

import (
	"context"

	"github.com/AlexandrShapkin/gosp-schema/models"
	"github.com/AlexandrShapkin/gosp-schema/repositories"
	"gorm.io/gorm"
)

type GormChatParticipantsRepository struct {
	db *gorm.DB
}

// AddUserToChat implements repositories.ChatParticipantsRepository.
func (g *GormChatParticipantsRepository) AddUserToChat(ctx context.Context, chatID string, userID string) error {
	participant := &models.ChatParticipants{
		ChatID: chatID,
		UserID: userID,
	}
	return g.db.WithContext(ctx).Create(participant).Error
}

// RemoveUserFromChat implements repositories.ChatParticipantsRepository.
func (g *GormChatParticipantsRepository) RemoveUserFromChat(ctx context.Context, chatID string, userID string) error {
	return g.db.WithContext(ctx).
		Where("chat_id = ? AND user_id = ?", chatID, userID).
		Delete(&models.ChatParticipants{}).Error
}

func NewChatParticipantsRepository(db *gorm.DB) repositories.ChatParticipantsRepository {
	return &GormChatParticipantsRepository{
		db: db,
	}
}
