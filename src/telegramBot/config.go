package telegramBot

import TelegramBotAPI "github.com/go-telegram-bot-api/telegram-bot-api/v5"

// Bot object
var bot *TelegramBotAPI.BotAPI

// Message
// Parameters: location, record.EmojiFlag, emailSubject, record.IpAddress, record.CreatedAt
var notifyMessage = `
âī¸ Someone opened your email in %s %s

*Subject:* %s
*IP Address:* %s
*Opened At:* %s Eastern Time

đ You can view more details in your *Maily dashboard*.
`

var welcomeMessage = `
âī¸ Welcome to Maily Telegram bot!
 
đ Maily Telegram bot sends you notifications when your email is opened.

đ To start using Maily Telegram bot, reply with your Maily Telegram token.

âšī¸ You can find your Maily Telegram token in your Maily account settings.
After logging in, click on the settings icon in the top right corner of the page and find the "My Telegram token" option.

đ We hope you enjoy using Maily!

đĻđē Make in Australia with đ
`

var userNotFoundMessage = `
đĢ User not found, please try again.

âšī¸ You can find your Maily Telegram token in your Maily account settings.
`

var setupCompletedMessage = `

đ If you wish to stop receiving notifications, just reply with /stop.

đ Please do not share your Maily Telegram token with anyone else.

đ Thanks again for choosing Maily!
`

var initialSetupCompletedMessage = "â Setup completed! You will now receive notifications through Telegram when your email is clicked." + setupCompletedMessage
var relinkSetupCompletedMessage = "â Relink completed! You will now receive notifications on this device when your email is clicked." + setupCompletedMessage
var alreadySetupMessage = "You have already setup the Maily Telegram bot, if you wish to stop receiving notifications or to relink your device, reply with /stop."
var enterTokenMessage = "Please enter your Maily Telegram token:"
var enterTokenButton1Message = "đ Enter Maily Telegram token"
var enterTokenButton2Message = "đ Try again"
var stopCommandUserNotFoundMessage = "It seems like you haven't setup Maily Telegram bot yet. You can setup Maily Telegram bot by replying /start."
var stopCommandCompletedMessage = "You have successfully stopped receiving notifications from Maily Telegram bot. You can start receiving notifications again by replying /start."

// Callback queries
var tokenCallbackQuery = "token"

// Commands
var startCommand = "start"
var stopCommand = "stop"
var startCommandDescription = "Setup Maily Telegram bot"
var stopCommandDescription = "Stop receiving notifications"
