package repository

import (
	"a21hc3NpZ25tZW50/model"

	"gorm.io/gorm"
)

type SessionsRepository interface {
	AddSessions(session model.Session) error
	DeleteSession(token string) error
	UpdateSessions(session model.Session) error
	SessionAvailName(name string) error
	SessionAvailToken(token string) (model.Session, error)
}

type sessionsRepoImpl struct {
	db *gorm.DB
}

func NewSessionRepo(db *gorm.DB) *sessionsRepoImpl {
	return &sessionsRepoImpl{db}
}

func (s *sessionsRepoImpl) AddSessions(session model.Session) error {
	// TODO: replace this
	if err := s.db.Create(&session).Error; err != nil {
		return err
	}
	return nil
}

func (s *sessionsRepoImpl) DeleteSession(token string) error {
	// TODO: replace this
	if err := s.db.Where("token = ?", token).Delete(&model.Session{}).Error; err != nil {
		return err
	}
	return nil
}

func (s *sessionsRepoImpl) UpdateSessions(session model.Session) error {
	// TODO: replace this
	var existingSession model.Session

	if err := s.db.First(&existingSession).Error; err != nil {
		return err
	}

	existingSession.Token = session.Token
	existingSession.Username = session.Username
	existingSession.Expiry = session.Expiry

	if err := s.db.Save(&existingSession).Error; err != nil {
		return err
	}

	return nil
}

func (s *sessionsRepoImpl) SessionAvailName(username string) error {
	// TODO: replace this
	var session model.Session
	if err := s.db.Where("username = ?", username).First(&session).Error; err != nil {
		return err
	}
	return nil
}

func (s *sessionsRepoImpl) SessionAvailToken(token string) (model.Session, error) {
	// TODO: replace this
	var session model.Session
	if err := s.db.Where("token = ?", token).First(&session).Error; err != nil {
		return model.Session{}, err
	}
	return session, nil
}
