package logic

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	ipdata "github.com/ipdata/go"
	"github.com/joho/godotenv"
	"github.com/teris-io/shortid"
	"golang.org/x/exp/slices"
	"gorm.io/gorm"
	"io"
	"maily/go-backend/src/dtos"
	"maily/go-backend/src/models"
	"os"
)

func LogEmailOpen(c *gin.Context) error {
	// Setup IP client, load .env and user agent files
	db := c.MustGet("DB").(*gorm.DB)
	_ = godotenv.Load(".env")
	ipd, _ := ipdata.NewClient(os.Getenv("IP_ADDRESS_API_KEY"))
	jsonFile, openFileError := os.Open("static/data/user_agents_email_client.json")
	if openFileError != nil {
		return openFileError
	}
	var userAgents []string // List of browser user agents
	byteValue, _ := io.ReadAll(jsonFile)
	jsonUnmarshalError := json.Unmarshal(byteValue, &userAgents)
	if jsonUnmarshalError != nil {
		return jsonUnmarshalError
	}
	closeFileError := jsonFile.Close()
	if closeFileError != nil {
		return closeFileError
	}

	// Gather data for tracker
	trackingNumber := c.Param("trackingId")
	// ipAddress := c.ClientIP()
	ipAddress := "118.102.80.22"     // Mock IP data
	data, _ := ipd.Lookup(ipAddress) // Get IP address data
	userAgent := c.Request.Header.Get("User-Agent")
	confidentWithEmailClient := slices.IndexFunc(userAgents, func(agent string) bool { return agent == userAgent }) != -1

	tracker := createTracker(data, trackingNumber, ipAddress, confidentWithEmailClient)
	db.Create(&tracker)

	return nil
}

func createTracker(ipData ipdata.IP, trackingNumber string, ipAddress string, confidentWithEmailClient bool) models.Record {
	var tracker models.Record

	tracker.ID = uuid.New()
	tracker.PublicTrackingNumber = trackingNumber
	tracker.IpAddress = ipAddress
	tracker.IpCity = ipData.City
	tracker.IpCountry = ipData.CountryName
	tracker.IsEU = ipData.IsEU
	tracker.IsTor = ipData.Threat.IsTOR
	tracker.IsProxy = ipData.Threat.IsProxy
	tracker.IsAnonymous = ipData.Threat.IsAnonymous
	tracker.IsKnownAttacker = ipData.Threat.IsKnownAttacker
	tracker.IsKnownAbuser = ipData.Threat.IsKnownAbuser
	tracker.IsThreat = ipData.Threat.IsThreat
	tracker.IsBogon = ipData.Threat.IsBogon
	tracker.ConfidentWithEmailClient = confidentWithEmailClient

	return tracker
}
