package controllers

import (
	"github.com/joho/godotenv"
	"github.com/teris-io/shortid"
	"maily/go-backend/src/utils"
	"maily/go-backend/src/utils/token"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	ipdata "github.com/ipdata/go"
	"gorm.io/gorm"

	"maily/go-backend/src/models"

	"github.com/aidarkhanov/nanoid"
	"github.com/google/uuid"
)

// Beep Example 1 godoc
// @Summary  1234
// @Schemes
// @Description  do ping
// @Tags         example
// @Accept       json
// @Produce      json
// @Success      200  {string}  Helloworld
// @Router       /example/12345 [get]
func Beep(c *gin.Context) {
	db := c.MustGet("DB").(*gorm.DB)
	_ = godotenv.Load(".env")
	ipd, _ := ipdata.NewClient(os.Getenv("IP_ADDRESS_API_KEY"))
	jsonFile, err := os.Open("static/data/user_agents_email_client.json")
	if err != nil {
		utils.HandleError(c, err)
		return
	}
	var userAgents []string // List of browser user agents
	byteValue, _ := io.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &userAgents)
	jsonFile.Close()

	var tracker models.Record
	// ipAddress := c.ClientIP()
	ipAddress := "118.102.80.22"
	userAgent := c.Request.Header.Get("User-Agent")
	data, _ := ipd.Lookup(ipAddress) // Get IP address data

	tracker.ID = uuid.New()
	tracker.PublicTrackingNumber = c.Param("trackingId")
	tracker.IpAddress = ipAddress
	tracker.IpCity = data.City
	tracker.IpCountry = data.CountryName
	tracker.IsEU = data.IsEU
	tracker.IsTor = data.Threat.IsTOR
	tracker.IsProxy = data.Threat.IsProxy
	tracker.IsAnonymous = data.Threat.IsAnonymous
	tracker.IsKnownAttacker = data.Threat.IsKnownAttacker
	tracker.IsKnownAbuser = data.Threat.IsKnownAbuser
	tracker.IsThreat = data.Threat.IsThreat
	tracker.IsBogon = data.Threat.IsBogon
	tracker.ConfidentWithEmailClient = slices.IndexFunc(userAgents, func(agent string) bool { return agent == userAgent }) != -1
	db.Create(&tracker)

	c.File("static/images/beep.gif")
}

func IpAddress(c *gin.Context) {
	_ = godotenv.Load(".env")
	ipd, _ := ipdata.NewClient(os.Getenv("IP_ADDRESS_API_KEY"))
	data, _ := ipd.Lookup("118.102.80.22")
	//data, _ := ipd.Lookup(c.ClientIP())

	fmt.Println(c.ClientIP())
	fmt.Println(c.Request.Header.Get("User-Agent"))

	c.IndentedJSON(http.StatusOK, data)
}

func BrowserTest(c *gin.Context) {
	jsonFile, err := os.Open("static/data/user_agents_email_client.json")
	if err != nil {
		fmt.Println(err)
	}

	var userAgents []string // List of browser user agents
	byteValue, _ := io.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &userAgents)
	jsonFile.Close()

	idx := slices.IndexFunc(userAgents, func(c string) bool {
		return c == "1"
	})
	fmt.Println(idx)

	//for _, s := range userAgents {
	//	fmt.Println(s)
	//}
	c.IndentedJSON(http.StatusOK, userAgents)
}

func Generate(c *gin.Context) {
	userId, tokenError := token.ExtractUserID(c)
	if tokenError != nil {
		utils.HandleError(c, tokenError)
		return
	}

	db := c.MustGet("DB").(*gorm.DB)
	publicTrackingNumber, _ := shortid.Generate()
	privateTrackingNumber, _ := shortid.Generate()

	result := db.Create(&models.Tracker{ID: publicTrackingNumber, PrivateTrackingNumber: privateTrackingNumber, UserID: userId})
	if result.Error != nil {
		utils.HandleError(c, result.Error)
		return
	}

	c.JSON(http.StatusOK, gin.H{"beaconUrl": fmt.Sprintf("http://%s/api/beep/%s", c.Request.Host, publicTrackingNumber)})
}
