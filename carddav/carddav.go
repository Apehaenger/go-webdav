package carddav

import (
	"encoding/xml"

	"github.com/emersion/go-vcard"
)

type AddressBook struct {
	Href        string
	Description string
}

type AddressBookQuery struct {
	Props []string
}

type AddressBookMultiGet struct {
	Hrefs []string
	Props []string
}

type AddressObject struct {
	Href string
	Card vcard.Card
}
