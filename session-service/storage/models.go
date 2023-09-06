package storage

// TODO::add a new session
func WriteSession(session Session) error {
	return nil
}

// TODO::retreive an active session of user if exists
func GetActiveSession(id string) (Session, error) {
	return Session{}, nil
}

// TODO::verify token is belonged to user and is active
func VerifyToken(id string, token string) bool {
	return false
}

// TODO::expire all sessions of a user
func ExpireUserSessions(id string) error {
	return nil
}
