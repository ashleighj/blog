package model

import "time"

type Contributor struct {
	ID         int64
	FirstNames string
	LastName   string
	Email      string
	DOB        time.Time
	Website    string
	City       string
	Country    string
	ZipCode    string
	IP         string
	SignUpDate time.Time
	PwHash     string
}
