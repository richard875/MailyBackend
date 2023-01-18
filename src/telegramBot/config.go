package telegramBot

// Message
var welcomeMessage = `
✉️ Welcome to Maily Telegram bot!
 
📌 Maily Telegram bot sends you notifications when your email is opened.

🔑 To start using Maily Telegram bot, reply with your Maily Telegram token.

ℹ️ You can find your Maily Telegram token in your Maily account settings.
After logging in, click on the settings icon in the top right corner of the page and find the "My Telegram token" option.

🚀 We hope you enjoy using Maily!

🇦🇺 Make in Australia with 💜
`
var userNotFoundMessage = `
🚫 User not found, please try again.

ℹ️ You can find your Maily Telegram token in your Maily account settings.
`

var setupCompletedMessage = `
✅ Setup completed! You will now receive notifications through Telegram when your email is clicked.

📌 If you wish to stop receiving notifications, just reply with /stop.

💜 Thanks again for choosing Maily!
`

var alreadySetupMessage = "You have already setup the Maily Telegram bot, if you wish to stop receiving notifications or to relink your device, reply with /stop."
var enterTokenMessage = "Please enter your Maily Telegram token:"
var enterTokenButton1Message = "🔑 Enter Maily Telegram token"
var enterTokenButton2Message = "🔑 Try again"
var stopCommandUserNotFoundMessage = "It seems like you haven't setup Maily Telegram bot yet. You can setup Maily Telegram bot by replying /start."
var stopCommandCompletedMessage = "You have successfully stopped receiving notifications from Maily Telegram bot. You can start receiving notifications again by replying /start."

// Callback queries
var tokenCallbackQuery = "token"

// Commands
var startCommand = "start"
var stopCommand = "stop"
var startCommandDescription = "Setup Maily Telegram bot"
var stopCommandDescription = "Stop receiving notifications"
