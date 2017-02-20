package app

import "time"

// ErrorResponse standard error response model.
type ErrorResponse struct {
	ID     string                   `json:"id,omitempty" xml:"id" form:"id"`                           // ID is the unique error instance identifier.
	Code   string                   `json:"code,omitempty" xml:"code" form:"code"`                     // Code identifies the class of errors.
	Status int                      `json:"status" xml:"status" form:"status"`                         // Status is the HTTP status code used by responses that cary the error.
	Detail string                   `json:"detail" xml:"detail" form:"detail"`                         // Detail describes the specific error occurrence.
	Meta   []map[string]interface{} `json:"meta,omitempty" xml:"meta,omitempty" form:"meta,omitempty"` // Meta contains additional key/value pairs useful to clients.
}

// Confirmation generic request confirmation.
type Confirmation struct {
	Message string `json:"message"` // confirmation message
}

// BuildResp returns information about the executables build.
type BuildResp struct {
	Build string `json:"build"`
}

// EventResp returns information about an event
type EventResp struct {
	Data EventData `json:"data"`
}

type EventData struct {
	ID         int             `json:"id"`
	Attributes EventAttributes `json:"attributes"`
}

type EventAttributes struct {
	Name string     `json:"name"`
	City string     `json:"city"`
	Date EventDates `json:"date"`
}

type EventDates struct {
	InitDate time.Time `json:"init-date"`
	EndDate  time.Time `json:"end-date"`
}

//SEARCH TICKET DATA
type SearchTicketResponse struct {
	Data SearchTicketData `json:"data"`
}

type SearchTicketData struct {
	ID string `json:"id"`
}

//LANGUAGE STRUCTURES
//LanguagesResp returns information about languages
type LanguagesResp struct {
	Data Language `json:"data"`
}

type Language struct {
	Languages []LanguagesData `json:"languages"`
}

type LanguagesData struct {
	Name    string `json:"name"`
	Locale  string `json:"locale"`
	Default bool   `json:"default"`
	Flag    string `json:"flag"`
}

//TRANSLATIONS STRUCTS
type TranslationsResp struct {
	Translations  map[string]string `json:"-"`
	LastCacheTime time.Time         `json:"-"`
}

//Declare a type to get name from POST body
type RegistrantName struct {
	Name         string `json:"name"`
	PrintedBadge string `json:"status"`
}

// BadgePrintResponse printable badge the kiosks will print.
type BadgePrintResponse struct {
	Name   string `json:"name"`
	Markup string `json:"html"`
}
