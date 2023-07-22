package tests

import (
	"time"

	"github.com/google/uuid"
)

func GenUUID() uuid.UUID {
	return uuid.MustParse("a1c790af-48bd-4081-9086-604f6564303e")
}

func GenTime() time.Time {
	tt, _ := time.Parse("2006-01-02T15:04:05", "2020-01-02T03:04:05")

	return tt
}
