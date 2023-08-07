package chatstorage

import (
	"errors"
	"time"

	"github.com/SergeyCherepiuk/chat-app/models"
)

type ChatStorageMock struct{}

func NewMock() *ChatStorageMock {
	return &ChatStorageMock{}
}

var messages []models.ChatMessage

func (storage ChatStorageMock) reset() {
	messages = []models.ChatMessage{
		{ID: 1, Message: "First message", From: 1, To: 2, CreatedAt: time.Now()},
		{ID: 2, Message: "Second message", From: 2, To: 1, CreatedAt: time.Now()},
	}
}

func (storage ChatStorageMock) GetHistory(userId, companionId uint) ([]models.ChatMessage, error) {
	storage.reset()
	history := []models.ChatMessage{}
	for _, message := range messages {
		if (message.From == userId && message.To == companionId) || (message.From == companionId && message.To == userId) {
			history = append(history, message)
		}
	}
	return history, nil
}

func (storage ChatStorageMock) Create(message *models.ChatMessage) error {
	storage.reset()
	messages = append(messages, *message)
	return nil
}

func (storage ChatStorageMock) Update(messageId uint, updatedMessage string) error {
	storage.reset()
	for _, message := range messages {
		if message.ID == messageId {
			message.Message = updatedMessage
			message.IsEdited = true
			return nil
		}
	}
	return errors.New("message not found")
}

func (storage ChatStorageMock) Delete(messageId uint) error {
	storage.reset()
	for i, message := range messages {
		if message.ID == messageId {
			messages = append(messages[:i], messages[i+1:]...)
			return nil
		}
	}
	return errors.New("message not found")
}

func (storage ChatStorageMock) DeleteAll(userId, companionId uint) error {
	storage.reset()
	for i := 0; i < len(messages); i++ {
		message := messages[i]
		if (message.From == userId && message.To == companionId) || (message.From == companionId && message.To == userId) {
			messages = append(messages[:i], messages[i+1:]...)
			i--
		}
	}
	return nil
}

func (storage ChatStorageMock) IsBelongsToChat(messageId, userId, companionId uint) (bool, error) {
	storage.reset()
	for _, message := range messages {
		if message.ID == messageId && ((message.From == userId && message.To == companionId) || (message.From == companionId && message.To == userId)) {
			return true, nil
		}
	}
	return false, errors.New("message not found in chat")
}

func (storage ChatStorageMock) IsAuthor(messageId, userId uint) (bool, error) {
	storage.reset()
	for _, message := range messages {
		if message.ID == messageId {
			return message.From == userId, nil
		}
	}
	return false, errors.New("message not found")
}