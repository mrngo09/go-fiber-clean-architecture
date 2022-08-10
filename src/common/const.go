package common

import "time"

const (
	KeyUserHeader                      = "userId"
	KeyTokenCache                      = "tokenCache"
	TimeExpireTokenCache time.Duration = 60 * 60 * 60 * 7
)
