package storage

import (
	"auth-service/internal/entity"
	"context"
	"database/sql"
)

func (s *Storage) CreateUser(user entity.CreateUserEntry) (string, error) {
	db := s.DB
	tx, err := db.BeginTx(context.TODO(), nil)
	if err != nil {
		return "", err
	}
	defer tx.Commit()

	query := `
		WITH u AS (
			INSERT INTO users (username, email, icon) 
			VALUES ($1, $2, $3) 
			RETURNING user_id
		)
		INSERT INTO user_passwords (user_id, password, algorithm) 
		SELECT user_id, $4, 'plain' FROM u
		RETURNING user_id;
	`

	var userID string
	row := tx.QueryRow(query, user.Username, user.Email, user.Icon, user.Password)
	err = row.Scan(&userID)
	if err != nil {
		tx.Rollback()
		return "", err
	}

	return userID, nil
}

func (s *Storage) UpdateUser(user entity.User) (string, error) {
	db := s.DB
	tx, err := db.BeginTx(context.TODO(), nil)
	if err != nil {
		return "", err
	}
	defer tx.Commit()

	query := `
		Update users SET
		username = $1,
		email = $2,
		icon = $3
		WHERE user_id = $4;
	`

	_, err = tx.Exec(query, user.Username, user.Email, user.Icon, user.UserID)
	if err != nil {
		tx.Rollback()
		return "", err
	}

	return user.UserID, nil
}

func (s *Storage) GetUserById(userID string) (*entity.User, error) {
	db := s.DB
	query := `
		SELECT user_id, username, email, icon, created_at, updated_at 
		FROM users WHERE user_id = $1;
	`

	var user entity.User
	row := db.QueryRow(query, userID)
	err := row.Scan(
		&user.UserID,
		&user.Username,
		&user.Email,
		&user.Icon,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err == nil {
		return &user, nil
	}

	if err == sql.ErrNoRows {
		return nil, nil
	}

	return nil, err
}

func (s *Storage) GetUserByEmail(email string) (*entity.User, error) {
	db := s.DB
	query := `
		SELECT user_id, username, email, icon, created_at, updated_at 
		FROM users WHERE email = $1;
	`

	var user entity.User
	row := db.QueryRow(query, email)
	err := row.Scan(
		&user.UserID,
		&user.Username,
		&user.Email,
		&user.Icon,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err == nil {
		return &user, nil
	}

	if err == sql.ErrNoRows {
		return nil, nil
	}

	return nil, err
}

func (s *Storage) GetUserByUsername(username string) (*entity.User, error) {
	db := s.DB
	query := `
		SELECT user_id, username, email, icon, created_at, updated_at 
		FROM users WHERE username = $1;
	`

	var user entity.User
	row := db.QueryRow(query, username)
	err := row.Scan(
		&user.UserID,
		&user.Username,
		&user.Email,
		&user.Icon,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err == nil {
		return &user, nil
	}

	if err == sql.ErrNoRows {
		return nil, nil
	}

	return nil, err
}

func (s *Storage) DeleteUser(userID string) error {
	db := s.DB
	tx, err := db.BeginTx(context.TODO(), nil)
	if err != nil {
		return err
	}
	defer tx.Commit()

	query := `DELETE FROM users WHERE user_id = $1;`
	_, err = tx.Exec(query, userID)
	if err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

func (s *Storage) UpdateUserPassword(userPassword entity.UpdatePasswordEntry) (string, error) {
	db := s.DB
	tx, err := db.BeginTx(context.TODO(), nil)
	if err != nil {
		return "", err
	}
	defer tx.Commit()

	updateQuery := `
		UPDATE user_passwords
		SET enabled = false
		WHERE user_id = $1 AND password = $2 AND enabled = true;
	`
	_, err = tx.Exec(updateQuery, userPassword.UserID, userPassword.OldPassword)
	if err != nil {
		tx.Rollback()
		return "", err
	}

	insertQuery := `
		INSERT INTO user_passwords (user_id, password, algorithm)
		VALUES ($1, $2, 'plain')
	`
	_, err = tx.Exec(insertQuery, userPassword.UserID, userPassword.Password)
	if err != nil {
		tx.Rollback()
		return "", err
	}

	return userPassword.UserID, nil
}

func (s *Storage) VerifyUserPassword(userID, password string) bool {
	db := s.DB
	query := `
		SELECT COUNT(*) FROM user_passwords
		WHERE user_id = $1 AND password = $2 AND enabled = true;
	`

	var countRow int
	row := db.QueryRow(query, userID, password)
	err := row.Scan(&countRow)
	if err == sql.ErrNoRows || countRow < 1 {
		return false
	}

	return true
}
