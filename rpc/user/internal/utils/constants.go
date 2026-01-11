package utils

const (
	// Redis key constants
	LoginCodeKey = "login:code:"
	LoginCodeTTL = 2 * 60 // 2 minutes in seconds
	LoginUserKey = "login:token:"
	LoginUserTTL = 36000 * 60 // 36000 minutes in seconds
	UserSignKey  = "sign:"
	FollowsKey   = "follows:"

	// System constants
	UserNickNamePrefix = "user_"
	DefaultPageSize    = 5
	MaxPageSize        = 10
)
