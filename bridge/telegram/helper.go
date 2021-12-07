package btelegram

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

func getUsername(user *tgbotapi.User) string {
	if user.LastName != "" {
		return user.FirstName + " " + user.LastName
	}

	return user.FirstName
}
