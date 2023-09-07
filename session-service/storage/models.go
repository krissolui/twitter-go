package storage

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

// TODO::add a new session
func (s *Storage) WriteSession(session Session) error {
	_, err := s.Collection.InsertOne(context.TODO(), session)

	return err
}

// TODO::retreive an active session of user if exists
func (s *Storage) GetActiveSession(userID string) (*Session, error) {
	result := s.Collection.FindOne(context.TODO(), bson.D{{"user_id", userID}})

	if result == nil {
		return nil, nil
	}

	session := Session{}
	err := result.Decode(&session)
	if err != nil {
		return &Session{}, err
	}

	return &session, nil
}

// TODO::verify token is belonged to user and is active
func (s *Storage) VerifyToken(userID string, token string) bool {
	result := s.Collection.FindOne(context.TODO(), bson.D{
		{"user_id", userID},
		{"token", token},
	})

	if result == nil {
		return false
	}

	session := Session{}
	err := result.Decode(&session)

	return err == nil
}

// TODO::expire all sessions of a user
func (s *Storage) ExpireUserSessions(userID string) error {
	return nil
}
