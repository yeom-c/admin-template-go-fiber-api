package database

import "time"

type User struct {
	Id             int32 `xorm:"pk autoincr"`
	Email          string
	Name           string
	HashedPassword string
	CreatedAt      time.Time
}
