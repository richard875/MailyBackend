package controllers

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
	"maily/go-backend/src/utils"
	"maily/go-backend/src/utils/token"
	"net/http"
	"os"

	"maily/go-backend/src/models"
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
	err := logic.LogEmailOpen(c)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	c.File("static/images/beep.gif")
}


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

	result := db.Create(&models.Tracker{ID: publicTrackingNumber, UserID: userId})
	if result.Error != nil {
		utils.HandleError(c, result.Error)
		return
	}

	c.JSON(http.StatusOK, gin.H{"beaconUrl": fmt.Sprintf("http://%s/api/beep/%s", c.Request.Host, publicTrackingNumber)})
}
