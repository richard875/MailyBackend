package telegramBot

import (
	"fmt"
	TelegramBotAPI "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"maily/go-backend/src/models"
)

func NotifyUser(userID int64, emailSubject string, record models.Record) {
	//timeZone, _ := time.LoadLocation("America/New_York")
	// Create message
	location := record.IpCountry
	if record.IpCity != "" {
		location = fmt.Sprintf("%s, %s", record.IpCity, record.IpCountry)
	}
	message := fmt.Sprintf(notifyMessage, location, record.EmojiFlag, emailSubject, record.IpAddress, record.CreatedAt.Format("02 Jan 2006, 15:04:05 PM"))

	msg := TelegramBotAPI.NewMessage(userID, message)
	msg.ParseMode = "Markdown"
	bot.Send(msg)
}
