/* For license and copyright information please see LEGAL file in repository */

package org

// HomeTimeline indicate the timeline-home domain record data fields.
type HomeTimeline interface {
	UserID() [16]byte       // user-status domain
	CommentID() [16]byte    // comment domain
	ReasonUserID() [16]byte // it can be the owner of comment or any other user base on reason
	Reason() HomeTimelineReason
}

type HomeTimelineReason uint8

const (
	HomeTimelineReason_Unset HomeTimeline = iota
	HomeTimelineReason_UserMentioned

	HomeTimelineReason_1stFriendshipCycle // new comment from friend of user
	HomeTimelineReason_2stFriendshipCycle
	HomeTimelineReason_3stFriendshipCycle
	HomeTimelineReason_4stFriendshipCycle
	HomeTimelineReason_5stFriendshipCycle
	HomeTimelineReason_6stFriendshipCycle
	HomeTimelineReason_UserFollow // just user follow other, not any friend relationship

	HomeTimelineReason_FriendForward        // retweet or forward other comment to the user timeline
	HomeTimelineReason_FriendReaction       // interest domain:  Liked, ...
	HomeTimelineReason_FriendReplied        // new comment to a user comment
	HomeTimelineReason_FriendReceivedAReply // a friend of user received a reply
	HomeTimelineReason_FollowByFriend       // usually 1stFriendshipCycle

	HomeTimelineReason_BaseOnUserReaction
	HomeTimelineReason_BaseOnUserFollow
	HomeTimelineReason_BaseOnUserTopics
)
