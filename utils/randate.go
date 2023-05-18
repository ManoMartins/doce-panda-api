package utils

import (
	"math/rand"
	"time"
)

func Randate() time.Time {
	min := time.Date(2022, 1, 0, 0, 0, 0, 0, time.Local).Unix()
	max := time.Date(2023, 3, 20, 0, 0, 0, 0, time.Local).Unix()
	delta := max - min

	sec := rand.Int63n(delta) + min
	return time.Unix(sec, 0)
}
