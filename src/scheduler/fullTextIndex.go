package scheduler

import (
	"gorm.io/gorm"
	"log"
	"strings"
)

func FulltextIndex(db *gorm.DB) {
	// Create the FULLTEXT index
	columns := []string{"id", "subject", "from_address", "to_addresses", "cc_addresses", "bcc_addresses", "reply_to_addresses", "internal_message_id"}
	query := "CREATE FULLTEXT INDEX idx_trackers_fulltext ON trackers (" + strings.Join(columns, ", ") + ")"

	err := db.Exec(query).Error
	if err != nil {
		log.Println(err)
	}
}
