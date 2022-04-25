/* For license and copyright information please see LEGAL file in repository */

package org

import (
	"../libgo/protocol"
)

// CommentGroup indicate the comment-group domain record data fields.
type CommentGroup interface {
	GroupID() [16]byte                 // or ChannelID. Use to store and retrieve comments in time order. It can be UUID, URL hash or any random id.
	OwnerUserID() [16]byte             // user-status domain
	BlockedWords() []string            //
	GetPolicy() CommentGroupPolicy     // Visibility level
	SetPolicy() CommentGroupPolicy     //
	ForwardPolicy() CommentGroupPolicy //
	Time() protocol.Time               // Save time
	RequestID() [16]byte               // user-request domain
}

type CommentGroupPolicy uint8

const (
	CommentGroupPolicy_Unset CommentGroupPolicy = iota
	CommentGroupPolicy_JustOwner
	CommentGroupPolicy_GroupManagers
	CommentGroupPolicy_GroupMembers
	CommentGroupPolicy_Mentioned
	CommentGroupPolicy_AnyOne
	// Below policies are person type authorization
	CommentGroupPolicy_1stFriendshipCycle
	CommentGroupPolicy_2stFriendshipCycle
	CommentGroupPolicy_3stFriendshipCycle
	CommentGroupPolicy_4stFriendshipCycle
	CommentGroupPolicy_5stFriendshipCycle
	CommentGroupPolicy_6stFriendshipCycle
)

type CommentGroupStorageServices interface {
	Save(c CommentGroup) protocol.Error

	Count(groupID [16]byte) (numbers uint64, err protocol.Error)
	Get(groupID [16]byte, versionOffset uint64) (c CommentGroup, err protocol.Error)

	FindByUserID(userID [16]byte, offset, limit uint64) (ids [][16]byte, numbers uint64, err protocol.Error)
}
