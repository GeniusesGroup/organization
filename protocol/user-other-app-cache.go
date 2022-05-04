/* For license and copyright information please see LEGAL file in repository */

package org

import (
	"../libgo/protocol"
)

// UserOtherAppCache cache all foreign fundamental user data
type UserOtherAppCache interface {
	UserOtherAppID() [16]byte    // user-other-app domain
	UserID() [16]byte            // user-status domain
	Gender() string              // In ISO format
	PassportID() string          // If exist mean user have that country Nationality
	Surename() string            // Firstname
	Middlename() string          //
	Givenname() string           // Lastname
	Nickname() string            //
	Username() string            // it is also be user email address
	Phone() string               //
	TaxID() string               //
	Picture() [16]byte           // Object UUID
	Bio() string                 //
	LocationID() string          //
	PostalCode() string          //
	DateOfBirth() protocol.Time  //
	PlaceOfBirth() string        //
	VerifiedDate() protocol.Time // If !=0 user got famous blue badge(tick) like instagram, twitter,...
	Time() protocol.Time         // Save time
}
