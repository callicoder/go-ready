package service

type ChatService interface {
}

type chatService struct {
}

func NewChatService() ChatService {
	return &chatService{}
}
