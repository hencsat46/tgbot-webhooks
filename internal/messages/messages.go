package messages

import "fmt"

func Profile(chatId int64, picCount int) string {
	message := fmt.Sprintf("Ваш профиль\nChatId: %d\nКоличество фото: %d", chatId, picCount)
	return message
}
