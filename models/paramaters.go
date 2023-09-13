package models

import (
	"time"
)

var (
	// StatusSuccess to store a status response success
	StatusSuccess = "Success"

	// MessageDataSuccess to store a success message response of data
	MessageDataSuccess = "Data Berhasil Dikirim"

	// MessageUnprocessableEntity to store a message response of unproccessable entity
	MessageUnprocessableEntity = "Entitas Tidak Dapat Diproses"

	MessageDataProcessing = "Data Sedang di proses"

	// DateFormat to store a date format of timestamp
	DateFormat = "2006-01-02"

	// DateTimeFormat to store a date format of timestamp
	DateTimeFormat = "2006-01-02 15:04:05"

	// DateTimeFormatMikro to store a date format of timestamp micro time
	DateTimeFormatMikro = "2006-01-02 15:04:05.000"

	// DateFormatNoSpace to store a date format of timestamp
	DateFormatNoSpace = "20060102"
)

// Parameter struct is represent a data for parameters model
type Parameter struct {
	ID          int64     `json:"id"`
	Key         string    `json:"key"`
	Value       string    `json:"value"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

// NowDbpg this function only because db pg when insert or update conver time format with UTC
// so when you INSERT/UPDATE using DBPG then you need this for to get time now
func NowDbpg() time.Time {
	return time.Now()
}

// NowUTC to get real current datetime but UTC format
func NowUTC() time.Time {
	return time.Now().UTC().Add(7 * time.Hour)
}
