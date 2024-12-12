package membership

import "time"

type Membership struct {
	Id      int
	UserId  int
	StartAt time.Time
	EndAt   time.Time
}
