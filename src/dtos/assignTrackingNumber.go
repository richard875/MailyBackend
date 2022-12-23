package dtos

type TrackingNumber struct {
	TrackingNumber string `json:"trackingNumber" binding:"required"`
}
