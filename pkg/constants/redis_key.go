package constants

import "fmt"

func GetUserMsgKey(userID int64) string {
	return fmt.Sprintf("user_message_%d", userID)
}

func GetUserLikeListKey(userID int64) string {
	return fmt.Sprintf("user_like_list_%d", userID)
}

func GetUserFollowListKey(userID int64) string {
	return fmt.Sprintf("user_follow_list_%d", userID)
}

func GetVideoMsgKey(videoID int64) string {
	return fmt.Sprintf("video_message_%d", videoID)
}