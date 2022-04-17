/* For license and copyright information please see LEGAL file in repository */

package pa

// PersonOTP indicate the person-otp domain record data fields.
type PersonOTP interface {
	PersonID() [16]byte   // user-status domain
	OTPPattern() []byte   // https://tools.ietf.org/html/rfc6238
	OTPAdditional() int32 // easy to remember and must be 2 to 7 digit. https://en.wikipedia.org/wiki/Personal_identification_number
	RequestID() [16]byte  // user-request domain
}
