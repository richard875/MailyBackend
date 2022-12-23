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
