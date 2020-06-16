package db

import (
	"net/http"
	"net/url"
	"time"
)

//RequestDetails - Request Details Struct.
// Only used for making JSON. Not a GORM Model
type RequestDetails struct {
	Method     string
	Parameters url.Values
	Origin     string
	URL        string
	Headers    http.Header
}

//ShowList output send to html template
type ShowList struct {
	Data string
	Date time.Time
}
