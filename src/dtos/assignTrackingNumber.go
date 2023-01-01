package dtos

type TrackingNumber struct {
	TrackingNumber    string `json:"trackingNumber" binding:"required"`
	ComposeAction     int    `json:"composeAction" binding:"required"`
	Subject           string `json:"subject" binding:"required"`
	FromAddress       string `json:"fromAddress" binding:"required"`
	ToAddresses       string `json:"toAddresses" binding:"required"`
	CcAddresses       string `json:"ccAddresses" binding:"required"`
	BccAddresses      string `json:"bccAddresses" binding:"required"`
	ReplyToAddresses  string `json:"replyToAddresses" binding:"required"`
	InternalMessageID string `json:"internalMessageID" binding:"required"`
}
