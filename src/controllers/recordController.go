package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	ipdata "github.com/ipdata/go"
	"gorm.io/gorm"

	"maily/go-backend/src/models"
	"maily/go-backend/src/secrets"

	"github.com/aidarkhanov/nanoid"
	"github.com/google/uuid"
)

// 123 godoc
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
	ipd, _ := ipdata.NewClient(secrets.IpAddressApiKey)

	data, err := ipd.Lookup("118.102.80.22")
	if err != nil {
		// handle error
	}

	c.IndentedJSON(http.StatusOK, data)
}
