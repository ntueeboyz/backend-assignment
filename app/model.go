package main

import "time"

// Advertisement represents the structure for an advertisement
type Advertisement struct {
	Title      string     `json:"title"`
	StartAt    time.Time  `json:"startAt"`
	EndAt      time.Time  `json:"endAt"`
	Conditions Conditions `json:"conditions"`
}

// Conditions represents the conditions under which an advertisement is shown
type Conditions struct {
	AgeStart int            `json:"ageStart,omitempty"`
	AgeEnd   int            `json:"ageEnd,omitempty"`
	Country  []ISO3166      `json:"country,omitempty"`
	Platform []PlatformType `json:"platform,omitempty"`
	Gender   GenderType     `json:"gender,omitempty"`
}

// PlatformType defines a type for platform enums
type PlatformType string

// GenderType defines a type for gender enums
type GenderType string

// Enum values for PlatformType
const (
    PlatformAndroid PlatformType = "android"
    PlatformIOS     PlatformType = "ios"
    PlatformWeb     PlatformType = "web"
)

// Enum values for GenderType
const (
    GenderMale   GenderType = "M"
    GenderFemale GenderType = "F"
)

// ISO3166Country type for country codes based on ISO 3166-1 alpha-2
type ISO3166 string

// Define constants for ISO 3166-1 alpha-2 country codes
const (
    CountryUSA ISO3166 = "US"
    CountryCanada ISO3166 = "CA"
    CountryJapan ISO3166 = "JP"
    CountrySouthKorea ISO3166 = "KR"
	CountryTaiwan ISO3166 = "TW"
    // Add more countries as needed
)