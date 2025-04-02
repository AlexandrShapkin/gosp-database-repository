package repositories

import (
	"context"
	"github.com/AlexandrShapkin/gosp-schema/models"
	"github.com/AlexandrShapkin/gosp-schema/repositories"
	"gorm.io/gorm"
)

type GormMessageRepository struct {
	db *gorm.DB
}

// Create implements repositories.MessageRepository.
func (g *GormMessageRepository) Create(ctx context.Context, message *models.Message) error {
	return g.db.WithContext(ctx).Create(message).Error
}

// DeleteByChat implements repositories.MessageRepository.
func (g *GormMessageRepository) DeleteByChat(ctx context.Context, chatID string) error {
	return g.db.WithContext(ctx).Where("chat_id = ?", chatID).Delete(&models.Message{}).Error
}

// DeleteByID implements repositories.MessageRepository.
func (g *GormMessageRepository) DeleteByID(ctx context.Context, id string) error {
	return g.db.WithContext(ctx).Where("id = ?", id).Delete(&models.Message{}).Error
}

// GetAll implements repositories.MessageRepository.
func (g *GormMessageRepository) GetAll(ctx context.Context, offset int, limit int) ([]models.Message, error) {
	var messages []models.Message
	query := g.db.WithContext(ctx)
	if limit > 0 {
		query = query.Offset(offset).Limit(limit)
	}
	err := query.Order("timestamp DESC").Find(&messages).Error
	return messages, err
}

// GetByChat implements repositories.MessageRepository.
func (g *GormMessageRepository) GetByChat(ctx context.Context, chatID string, offset int, limit int, order string) ([]models.Message, error) {
	var messages []models.Message
	query := g.db.WithContext(ctx).Where("chat_id = ?", chatID)
	if limit > 0 {
		query = query.Offset(offset).Limit(limit)
	}
	if order != "ASC" && order != "DESC" {
		order = "DESC"
	}
	err := query.Order("timestamp" + order).Find(&messages).Error
	return messages, err
}

// GetByID implements repositories.MessageRepository.
func (g *GormMessageRepository) GetByID(ctx context.Context, id string) (*models.Message, error) {
	var message models.Message
	err := g.db.WithContext(ctx).First(&message, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &message, nil
}

// GetFromID implements repositories.MessageRepository.
func (g *GormMessageRepository) GetFromID(ctx context.Context, chatID string, startFrom string, limit int, order string) ([]models.Message, error) {
	var messages []models.Message
	query := g.db.WithContext(ctx).Where("chat_id = ? AND id >= ?", chatID, startFrom)
	if limit > 0 {
		query = query.Limit(limit)
	}
	if order != "ASC" && order != "DESC" {
		order = "DESC"
	}
	err := query.Order("timestamp" + order).Find(&messages).Error
	return messages, err
}

// GetLastInChat implements repositories.MessageRepository.
func (g *GormMessageRepository) GetLastInChat(ctx context.Context, chatID string) (*models.Message, error) {
	var message models.Message
	err := g.db.WithContext(ctx).
		Where("chat_id = ?", chatID).
		Order("timestamp DESC").
		First(&message).Error
	
	if err != nil {
		return nil, err
	}
	return &message, nil
}

// Update implements repositories.MessageRepository.
func (g *GormMessageRepository) Update(ctx context.Context, message *models.Message) error {
	return g.db.WithContext(ctx).Save(message).Error
}

func NewMessageRepository(db *gorm.DB) repositories.MessageRepository {
	return &GormMessageRepository{
		db: db,
	}
}
