package chat

import (
	"core_chat/application/chat/dto"
	"core_chat/application/chat/entity"
	"core_chat/application/chat/repository"
	"database/sql"
	"errors"
	"strconv"
	"time"
)

type chatRepositoryImpl struct {
	db *sql.DB
}

func NewChatRepository(db *sql.DB) repository.ChatRepository {
	return &chatRepositoryImpl{db: db}
}

func (r *chatRepositoryImpl) SaveImageMetadata(metadata dto.ChatImageMetadata) error {
	query := `
		INSERT INTO chat_image (sender, receiver, image_path)
		VALUES (?, ?, ?)
	`

	_, err := r.db.Exec(query, metadata.Sender, metadata.Receiver, metadata.ImagePath)
	return err
}

func (r *chatRepositoryImpl) ExistsByIdentifier(identifier string) bool {
	var existingIdentifier string
	err := r.db.QueryRow("SELECT identifier FROM person WHERE identifier = ?", identifier).Scan(&existingIdentifier)

	if err != nil {
		if err == sql.ErrNoRows {
			return false // Identifier not found — doesn't exist
		}
		return false // For safety, treat errors as "not existing"
	}

	return true // Identifier exists
}

func (r *chatRepositoryImpl) IsImageCanBeRetrieved(imagePath, identifier string) bool {
	var path string
	err := r.db.QueryRow(
		`SELECT image_path FROM chat_image 
		 WHERE image_path = ? 
		 AND (sender = ? OR receiver = ?) 
		 LIMIT 1`,
		imagePath, identifier, identifier,
	).Scan(&path)

	if err != nil {
		if err == sql.ErrNoRows {
			return false // No access to image or image doesn't exist
		}
		return false // DB error, treat as unauthorized for safety
	}

	return true // Access allowed
}

func (r *chatRepositoryImpl) SaveMessage(input dto.SendMessageInput) (string, error) {
	query := `
		INSERT INTO chat_message (sender, receiver, type, title, body, payload)
		VALUES (?, ?, ?, ?, ?, ?)
	`

	result, err := r.db.Exec(query,
		input.Sender,
		input.Receiver,
		input.Type,
		input.Title,
		input.Body,
		input.Payload,
	)

	messageID, err := result.LastInsertId()

	if err != nil {
		return "", err
	}

	msgID := strconv.FormatInt(messageID, 10)

	return msgID, err
}

func (r *chatRepositoryImpl) FindChatMessage(msgID, identifier string) (*entity.Message, error) {
	var message entity.Message
	var id int64
	var createdAt time.Time
	var readAt sql.NullTime

	msgIDInt, err := strconv.ParseInt(msgID, 10, 64)
	if err != nil {
		return nil, err
	}

	err = r.db.QueryRow("SELECT id, sender, receiver, type, title, body, payload, created_at, read_at FROM chat_message WHERE id = ? AND (sender = ? OR receiver = ?)", msgIDInt, identifier, identifier).
		Scan(id, &message.Sender, &message.Receiver, &message.Type, &message.Title, &message.Body, &message.Payload, &createdAt, &readAt)

	if err != nil {
		return nil, err
	}

	message.ID = strconv.FormatInt(id, 10)
	message.CreatedAt = createdAt.Format("2006-01-02 15:04:05")
	if readAt.Valid {
		message.ReadAt = readAt.Time.Format("2006-01-02 15:04:05")
	} else {
		message.ReadAt = ""
	}

	return &message, nil
}

func (r *chatRepositoryImpl) MarkMessageAsRead(msgID string, receiver string) error {
	msgIDInt, err := strconv.ParseInt(msgID, 10, 64)
	if err != nil {
		return err
	}

	query := `
		UPDATE chat_message
		SET read_at = CURRENT_TIMESTAMP
		WHERE id = ? AND receiver = ? AND read_at IS NULL
	`

	result, err := r.db.Exec(query, msgIDInt, receiver)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("no message updated: either message doesn't exist, not the receiver, or already read")
	}

	return nil
}
