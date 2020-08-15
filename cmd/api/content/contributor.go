package content

import "time"

type Contributor struct {
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
}
