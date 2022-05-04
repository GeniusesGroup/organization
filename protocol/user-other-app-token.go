/* For license and copyright information please see LEGAL file in repository */

package org

import (
	"../libgo/protocol"
)

type UserOtherAppToken interface {
	UserOtherAppID() [16]byte // user-other-app domain
	Token() string            // DelegateToken , AccessToken
	Time() protocol.Time      // Save time
	RequestID() [16]byte      // user-request domain
}
