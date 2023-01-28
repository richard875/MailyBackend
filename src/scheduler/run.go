package scheduler

import (
	"maily/go-backend/src/database"
	"time"
)

func Run() {
	// Run the createFulltextIndex function every 24 hours
	ticker := time.NewTicker(time.Hour * 24)
	go func() {
		for range ticker.C {
			//log.Println("Running create fulltext index")
			FulltextIndex(database.DB)
		}
	}()
}
