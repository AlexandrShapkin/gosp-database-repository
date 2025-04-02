package gospdatabaserepository

import (
	"github.com/AlexandrShapkin/gosp-database-repository/repositories"
	gospschema "github.com/AlexandrShapkin/gosp-schema"
	sr "github.com/AlexandrShapkin/gosp-schema/repositories"
	"gorm.io/gorm"
)

type GormRepository struct {
	db *gorm.DB
}

// Chat implements gospschema.Repository.
func (g *GormRepository) Chat() sr.ChatRepository {
	return repositories.NewChatRepository(g.db)
}

// ChatParticipants implements gospschema.Repository.
func (g *GormRepository) ChatParticipants() sr.ChatParticipantsRepository {
	return repositories.NewChatParticipantsRepository(g.db)
}

// Message implements gospschema.Repository.
func (g *GormRepository) Message() sr.MessageRepository {
	return repositories.NewMessageRepository(g.db)
}

// Token implements gospschema.Repository.
func (g *GormRepository) Token() sr.TokenRepository {
	return repositories.NewTokenRepository(g.db)
}

// User implements gospschema.Repository.
func (g *GormRepository) User() sr.UserRepository {
	return repositories.NewUserRepository(g.db)
}

func NewRepository(db *gorm.DB) gospschema.Repository {
	return &GormRepository{
		db: db,
	}
}
