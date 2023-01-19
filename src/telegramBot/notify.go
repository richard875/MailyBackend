package telegramBot

import (
	"fmt"
	TelegramBotAPI "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"maily/go-backend/src/models"
	"time"
)

func NotifyUser(userID int64, emailSubject string, record models.Record) {
	timeZone, _ := time.LoadLocation("America/New_York")
	// Create message
	var location string
	if record.IpCity != "" {
		location = fmt.Sprintf("%s, %s", record.IpCity, record.IpCountry)
	} else {
		location = record.IpCountry
	}
	message := fmt.Sprintf(`
✉️ Someone opened your email in %s %s!

*Subject:* %s
*IP Address:* %s
*Opened At:* %s Eastern Time

🌐 You can view more details in your *Maily dashboard*.
`, location, record.EmojiFlag, emailSubject, record.IpAddress, record.CreatedAt.In(timeZone).Format("02 Jan 2006, 15:04:05 PM"))

	msg := TelegramBotAPI.NewMessage(userID, message)
	msg.ParseMode = "Markdown"
	bot.Send(msg)
}
