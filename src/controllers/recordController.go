package controllers

import (
	"github.com/joho/godotenv"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	ipdata "github.com/ipdata/go"
	"gorm.io/gorm"

	"maily/go-backend/src/models"

	"github.com/aidarkhanov/nanoid"
	"github.com/google/uuid"
)

// Example 1 godoc
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
	db.Create(&models.Record{UserID: uuid.New().String(), LogNumber: nanoid.New()})

	c.File("static/images/1.jpg")
}

// abc godoc
// @Summary  abcd
// @Schemes
// @Description  do ping
// @Tags         example
// @Accept       json
// @Produce      json
// @Success      200  {string}  Helloworld
// @Router       /example/abcde [get]
func IpAddress(c *gin.Context) {
	_ = godotenv.Load(".env")
	ipd, _ := ipdata.NewClient(os.Getenv("IP_ADDRESS_API_KEY"))
	data, _ := ipd.Lookup("118.102.80.22")

	c.IndentedJSON(http.StatusOK, data)
}
