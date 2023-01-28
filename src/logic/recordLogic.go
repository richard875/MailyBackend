package logic

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/teris-io/shortid"
	"golang.org/x/exp/slices"
	"io"
	"maily/go-backend/src/database"
	"maily/go-backend/src/dtos"
	"maily/go-backend/src/enums"
	"maily/go-backend/src/functions"
	"maily/go-backend/src/global"
	"maily/go-backend/src/models"
	"maily/go-backend/src/structure"
	"maily/go-backend/src/telegramBot"
	mailyWebsocket "maily/go-backend/src/websocket"
	"os"
	"strconv"
	"strings"
)

func LogEmailOpen(c *gin.Context) error {
	db := database.DB

	// Check if the tracking number is in the database
	rawTrackingNumber := c.Param("trackingId")
	trackingNumber := rawTrackingNumber[:strings.Index(rawTrackingNumber, ".")]
	var currentTracker models.Tracker
	if err := db.First(&currentTracker, "id = ?", trackingNumber).Error; err != nil {
		return err
	}

	// Load user agent files
	userAgents := openJsonFile() // List of browser user agents

	// Gather data for tracker
	ipAddress := c.ClientIP()
	if ipAddress == "::1" {
		ipAddress = "99.232.45.12" // Hardcoded IP address for localhost (Whitby, Ontario)
	}
	data := functions.GetIpData(ipAddress) // Get IP address data

	// Check if user agent is in the list of email clients (Subject to change)
	userAgent := c.Request.Header.Get("User-Agent")
	confidentWithEmailClient := slices.IndexFunc(userAgents, func(agent string) bool { return agent == userAgent }) != -1

	// Get and store Header data
	headers := c.Request.Header
	headersString := ""
	for key, value := range headers {
		headersString += key + global.Delimiter + value[0] + global.Delimiter
	}

	// Create tracker record
	record := createTrackerRecord(data, trackingNumber, ipAddress, confidentWithEmailClient, headersString)
	db.Create(&record)

	// Update tracker
	db.Model(&currentTracker).Update("TimesOpened", currentTracker.TimesOpened+1)
	db.Model(&currentTracker).Update("Updated", true)

	// Send WebSockets message update
	websocketStringList := []string{mailyWebsocket.UpdateSignal, currentTracker.Subject, record.IpCity, record.IpCountry, record.EmojiFlag}

	// Send APN message update if websocket is established successfully
	if mailyWebsocket.Websocket != nil {
		mailyWebsocket.Websocket.WriteMessage(1, []byte(strings.Join(websocketStringList, global.Delimiter)))
	}

	// Update user total clicks
	var user models.User
	db.First(&user, "id = ?", currentTracker.UserID)
	db.Model(&user).Update("TotalClicks", user.TotalClicks+1)

	// Send Telegram message if user has linked their Telegram ID
	if user.TelegramID != 0 {
		telegramBot.NotifyUser(user.TelegramID, currentTracker.Subject, record)
	}

	return nil
}

func openJsonFile() []string {
	jsonFile, _ := os.Open("static/data/user_agents_email_client.json")
	var userAgents []string // List of browser user agents
	byteValue, _ := io.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &userAgents)
	jsonFile.Close()

	return userAgents
}

func createTrackerRecord(ipData structure.IpLocationInfo, trackingNumber string, ipAddress string, confidentWithEmailClient bool, headersString string) models.Record {
	var record models.Record

	record.ID = uuid.New()
	record.PublicTrackingNumber = trackingNumber
	record.IpAddress = ipAddress
	record.IpCity = ipData.City
	record.IpCountry = ipData.CountryName
	record.EmojiFlag = ipData.EmojiFlag
	record.IsEU = ipData.IsEU
	record.IsTor = ipData.Threat.IsTor
	record.IsIcloudRelay = ipData.Threat.IsICloudRelay
	record.IsProxy = ipData.Threat.IsProxy
	record.IsDataCenter = ipData.Threat.IsDatacenter
	record.IsAnonymous = ipData.Threat.IsAnonymous
	record.IsKnownAttacker = ipData.Threat.IsKnownAttacker
	record.IsKnownAbuser = ipData.Threat.IsKnownAbuser
	record.IsThreat = ipData.Threat.IsThreat
	record.IsBogon = ipData.Threat.IsBogon
	record.Latitude = ipData.Latitude
	record.Longitude = ipData.Longitude
	record.ConfidentWithEmailClient = confidentWithEmailClient
	record.Headers = headersString

	return record
}

func GeneratePublicTrackingNumber() string {
	trackingNumber, _ := shortid.Generate()
	return trackingNumber
}

func AssignTrackingNumber(c *gin.Context, userId string) error {
	// Bind request body to TrackingNumber DTO
	db := database.DB
	var trackingNumber dtos.TrackingNumber
	_ = c.ShouldBindJSON(&trackingNumber)

	// Construct new Tracker model and create in database
	var tracker models.Tracker
	tracker.ID = trackingNumber.TrackingNumber
	tracker.UserID = userId
	tracker.ComposeAction = trackingNumber.ComposeAction
	tracker.Subject = trackingNumber.Subject
	tracker.FromAddress = trackingNumber.FromAddress
	tracker.ToAddresses = trackingNumber.ToAddresses
	tracker.CcAddresses = trackingNumber.CcAddresses
	tracker.BccAddresses = trackingNumber.BccAddresses
	tracker.ReplyToAddresses = trackingNumber.ReplyToAddresses
	tracker.InternalMessageID = trackingNumber.InternalMessageID
	tracker.Updated = true // Set default to true so that the tracker will appear as notification in frontend

	result := db.Create(&tracker)

	// Send WebSockets message update
	mailyWebsocket.Websocket.WriteMessage(1, []byte(mailyWebsocket.UpdateSignal))

	// Update User
	var user models.User
	db.First(&user, "id = ?", userId)
	db.Model(&user).Update("EmailsSent", user.EmailsSent+1)
	return result.Error
}

func GetUserTrackers(c *gin.Context, userId string) ([]models.Tracker, error) {
	db := database.DB
	limit := 10
	indexEmail := c.Param("indexEmail")
	pageNumber, err := strconv.Atoi(c.Param("page")) // Default to page 1 if not provided
	if err != nil || pageNumber < 1 {
		pageNumber = 1
	}

	var trackers []models.Tracker
	result := db.Where("user_id = ?", userId).Order("updated_at desc").Limit(limit).Offset((pageNumber - 1) * limit)
	if indexEmail == string(enums.Opened) {
		result = result.Where("times_opened > ?", 0)
	} else if indexEmail == string(enums.Unopened) {
		result = result.Where("times_opened = ?", 0)
	}
	result.Find(&trackers)
	return trackers, result.Error
}

func SearchTrackers(c *gin.Context, userId string) ([]models.Tracker, error) {
	db := database.DB
	limit := 10
	searchQuery := c.Param("searchQuery")
	pageNumber, conversionError := strconv.Atoi(c.Param("page")) // Default to page 1 if not provided
	if conversionError != nil || pageNumber < 1 {
		pageNumber = 1
	}

	var trackers []models.Tracker
	err := db.Order("updated_at desc").Where("MATCH(id, subject, from_address, to_addresses, cc_addresses, bcc_addresses, reply_to_addresses, internal_message_id) AGAINST (?) AND user_id = ?", searchQuery, userId).Limit(limit).Offset((pageNumber - 1) * limit).Find(&trackers).Error
	if err != nil {
		return nil, err
	}

	return trackers, nil
}

func GetTrackerClicks(c *gin.Context) ([]models.Record, error) {
	db := database.DB
	limit := 10
	trackingNumber := c.Param("trackingNumber")
	emailViewSort := c.Param("emailViewSort")
	pageNumber, convertErr := strconv.Atoi(c.Param("page")) // Default to page 1 if not provided
	if convertErr != nil || pageNumber < 1 {
		pageNumber = 1
	}

	sortDirection := "desc"
	if emailViewSort == string(enums.OldestToLatest) {
		sortDirection = "asc"
	}

	// Get tracker clicks database record
	var currentTracker models.Tracker
	if err := db.First(&currentTracker, "id = ?", trackingNumber).Error; err != nil {
		return []models.Record{}, err
	}

	// Update tracker
	db.Model(&currentTracker).UpdateColumn("Updated", false)

	var records []models.Record
	err := db.Order(fmt.Sprintf("created_at %s", sortDirection)).Where("public_tracking_number = ?", trackingNumber).Limit(limit).Offset((pageNumber - 1) * limit).Find(&records).Error
	if err != nil {
		return nil, err
	}

	return records, nil
}

func ReGenerateTelegramToken(userId string) string {
	db := database.DB
	newTelegramToken, _ := shortid.Generate()

	// Update user Telegram token
	var user models.User
	db.First(&user, "id = ?", userId)
	db.Model(&user).Update("TelegramToken", newTelegramToken)

	return newTelegramToken
}
