package logic

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	ipdata "github.com/ipdata/go"
	"github.com/joho/godotenv"
	"github.com/teris-io/shortid"
	"golang.org/x/exp/slices"
	"gorm.io/gorm"
	"io"
	"maily/go-backend/src/dtos"
	"maily/go-backend/src/enums"
	"maily/go-backend/src/models"
	"os"
	"strings"
)

func LogEmailOpen(c *gin.Context) error {
	db := c.MustGet("DB").(*gorm.DB)

	// Check if the tracking number is in the database
	rawTrackingNumber := c.Param("trackingId")
	trackingNumber := rawTrackingNumber[:strings.Index(rawTrackingNumber, ".")]
	var currentTracker models.Tracker
	if err := db.First(&currentTracker, "id = ?", trackingNumber).Error; err != nil {
		return nil
	}

	// Setup IP client, load .env and user agent files
	_ = godotenv.Load(".env")
	ipd, _ := ipdata.NewClient(os.Getenv("IP_ADDRESS_API_KEY"))
	userAgents := openJsonFile() // List of browser user agents

	// Gather data for tracker
	ipAddress := c.ClientIP()
	data, _ := ipd.Lookup(ipAddress) // Get IP address data
	userAgent := c.Request.Header.Get("User-Agent")
	confidentWithEmailClient := slices.IndexFunc(userAgents, func(agent string) bool { return agent == userAgent }) != -1

	// Create tracker record
	tracker := createTrackerRecord(data, trackingNumber, ipAddress, confidentWithEmailClient)
	db.Create(&tracker)

	// Update tracker
	db.Model(&currentTracker).Update("TimesOpened", currentTracker.TimesOpened+1)

	// Update user total clicks
	var user models.User
	db.First(&user, "id = ?", currentTracker.UserID)
	db.Model(&user).Update("TotalClicks", user.TotalClicks+1)

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

func createTrackerRecord(ipData ipdata.IP, trackingNumber string, ipAddress string, confidentWithEmailClient bool) models.Record {
	var record models.Record

	record.ID = uuid.New()
	record.PublicTrackingNumber = trackingNumber
	record.IpAddress = ipAddress
	record.IpCity = ipData.City
	record.IpCountry = ipData.CountryName
	record.EmojiFlag = ipData.EmojiFlag
	record.IsEU = ipData.IsEU
	record.IsTor = ipData.Threat.IsTOR
	record.IsProxy = ipData.Threat.IsProxy
	record.IsAnonymous = ipData.Threat.IsAnonymous
	record.IsKnownAttacker = ipData.Threat.IsKnownAttacker
	record.IsKnownAbuser = ipData.Threat.IsKnownAbuser
	record.IsThreat = ipData.Threat.IsThreat
	record.IsBogon = ipData.Threat.IsBogon
	record.Latitude = ipData.Latitude
	record.Longitude = ipData.Longitude
	record.ConfidentWithEmailClient = confidentWithEmailClient

	return record
}

func GeneratePublicTrackingNumber() string {
	trackingNumber, _ := shortid.Generate()
	return trackingNumber
}

func AssignTrackingNumber(c *gin.Context, userId string) error {
	// Bind request body to TrackingNumber DTO
	db := c.MustGet("DB").(*gorm.DB)
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

	result := db.Create(&tracker)

	// Update User
	var user models.User
	db.First(&user, "id = ?", userId)
	db.Model(&user).Update("EmailsSent", user.EmailsSent+1)
	return result.Error
}

func GetUserTrackers(c *gin.Context, userId string) ([]models.Tracker, error) {
	db := c.MustGet("DB").(*gorm.DB)
	indexEmail := c.Param("indexEmail")

	var trackers []models.Tracker
	result := db.Where("user_id = ?", userId).Order("updated_at desc")
	if indexEmail == string(enums.Opened) {
		result = result.Where("times_opened > ?", 0)
	} else if indexEmail == string(enums.Unopened) {
		result = result.Where("times_opened = ?", 0)
	}
	result.Find(&trackers)
	return trackers, result.Error
}

func SearchTrackers(c *gin.Context, userId string) ([]models.Tracker, error) {
	db := c.MustGet("DB").(*gorm.DB)
	searchQuery := c.Param("searchQuery")

	var trackers []models.Tracker
	err := db.Order("updated_at desc").Where("MATCH(id, subject, from_address, to_addresses, cc_addresses, bcc_addresses, reply_to_addresses, internal_message_id) AGAINST (?) AND user_id = ?", searchQuery, userId).Find(&trackers).Error
	if err != nil {
		return nil, err
	}

	return trackers, nil
}

func GetTrackerClicks(c *gin.Context) ([]models.Record, error) {
	db := c.MustGet("DB").(*gorm.DB)
	trackingNumber := c.Param("trackingNumber")
	emailViewSort := c.Param("emailViewSort")

	sortDirection := "desc"
	if emailViewSort == string(enums.OldestToLatest) {
		sortDirection = "asc"
	}

	var records []models.Record
	err := db.Order(fmt.Sprintf("created_at %s", sortDirection)).Where("public_tracking_number = ?", trackingNumber).Find(&records).Error
	if err != nil {
		return nil, err
	}

	return records, nil
}
