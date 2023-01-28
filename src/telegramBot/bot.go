package telegramBot

import (
	TelegramBotAPI "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"gorm.io/gorm"
	"maily/go-backend/src/database"
	"maily/go-backend/src/models"
	"os"
	"strings"
)

func StartTelegramBot() {
	// Load Telegram bot
	var err error
	bot, err = TelegramBotAPI.NewBotAPI(os.Getenv("TELEGRAM_BOT_TOKEN"))
	if err != nil {
		return
	}

	// Development
	bot.Debug = false
	//log.Printf("Authorized on account %s", bot.Self.UserName)

	// Config
	registerCommands() // Register commands
	u := TelegramBotAPI.NewUpdate(0)
	u.Timeout = 60

	// Start listening for messages
	updates := bot.GetUpdatesChan(u)

	go func() {
		// Listen for messages
		for update := range updates {
			// If the received update has a message
			if update.CallbackQuery != nil {
				// Handle callback query for user about to enter their Maily Telegram token
				if update.CallbackQuery.Data == tokenCallbackQuery {
					msg := TelegramBotAPI.NewMessage(update.CallbackQuery.Message.Chat.ID, enterTokenMessage)
					msg.ReplyMarkup = TelegramBotAPI.ForceReply{ForceReply: true, Selective: true}
					bot.Send(msg)
				}
			} else if update.Message != nil {
				if update.Message.ReplyToMessage != nil { // If received token message from the reply
					handleTokenMessage(update)
				} else if update.Message.Command() == startCommand { // If received /start command
					handleStartCommand(update)
				} else if update.Message.Command() == stopCommand { // If received /stop command
					handleStopCommand(update)
				}
			}
		}
	}()
}

func registerCommands() *TelegramBotAPI.APIResponse {
	// Register commands
	botStartCommand := TelegramBotAPI.BotCommand{Command: startCommand, Description: startCommandDescription}
	botStopCommand := TelegramBotAPI.BotCommand{Command: stopCommand, Description: stopCommandDescription}

	config := TelegramBotAPI.NewSetMyCommands(botStartCommand, botStopCommand)
	request, _ := bot.Request(config)

	return request
}

func handleTokenMessage(update TelegramBotAPI.Update) {
	// User token
	token := strings.TrimSpace(update.Message.Text)

	// Get Database connection
	var db = database.DB

	// Get user and check if user exists
	var user models.User
	err := db.Where("telegram_token = ?", token).First(&user).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			// Handle case where user does not exist
			msg := TelegramBotAPI.NewMessage(update.Message.Chat.ID, userNotFoundMessage)

			// Add inline reply option, after user click on this button, they should be able to enter their Maily Telegram token
			msg.ReplyMarkup = TelegramBotAPI.NewInlineKeyboardMarkup(
				TelegramBotAPI.NewInlineKeyboardRow(
					TelegramBotAPI.NewInlineKeyboardButtonData(enterTokenButton2Message, tokenCallbackQuery),
				),
			)

			bot.Send(msg)
		} else {
			// Handle other errors
		}
	} else {
		notLinked := user.TelegramID == 0

		// User exists, update their telegram_chat_id
		db.Model(&user).Update("telegram_id", update.Message.Chat.ID)

		// Send message to user
		if notLinked {
			msg := TelegramBotAPI.NewMessage(update.Message.Chat.ID, initialSetupCompletedMessage)
			bot.Send(msg)
		} else {
			msg := TelegramBotAPI.NewMessage(update.Message.Chat.ID, relinkSetupCompletedMessage)
			bot.Send(msg)
		}
	}
}

func handleStartCommand(update TelegramBotAPI.Update) {
	// Get Database connection
	var db = database.DB

	// Get user by their Telegram ID and check if user exists
	var user models.User
	err := db.Where("telegram_id = ?", update.Message.Chat.ID).First(&user).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			msg := TelegramBotAPI.NewMessage(update.Message.Chat.ID, welcomeMessage)
			msg.ReplyToMessageID = update.Message.MessageID

			// Add inline reply option, after user click on this button, they should be able to enter their Maily Telegram token
			msg.ReplyMarkup = TelegramBotAPI.NewInlineKeyboardMarkup(
				TelegramBotAPI.NewInlineKeyboardRow(
					TelegramBotAPI.NewInlineKeyboardButtonData(enterTokenButton1Message, tokenCallbackQuery),
				),
			)

			bot.Send(msg)
		} else {
			// Handle other errors
		}
	} else {
		// User exists, send message to user
		msg := TelegramBotAPI.NewMessage(update.Message.Chat.ID, alreadySetupMessage)
		bot.Send(msg)
	}

}

func handleStopCommand(update TelegramBotAPI.Update) {
	// Get Database connection
	var db = database.DB

	// Get user by their Telegram ID and check if user exists
	var user models.User
	err := db.Where("telegram_id = ?", update.Message.Chat.ID).First(&user).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			// Handle case where user does not exist
			msg := TelegramBotAPI.NewMessage(update.Message.Chat.ID, stopCommandUserNotFoundMessage)
			bot.Send(msg)
		} else {
			// Handle other errors
		}
	} else {
		// User exists, remove their Telegram ID
		db.Model(&user).Update("telegram_id", nil)

		// Send message to user
		msg := TelegramBotAPI.NewMessage(update.Message.Chat.ID, stopCommandCompletedMessage)
		bot.Send(msg)
	}
}
