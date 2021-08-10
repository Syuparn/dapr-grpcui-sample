package main

import (
	"math/rand"
	"time"

	"github.com/oklog/ulid"
)

type Order struct {
	ID    string
	Drink string
	Time  time.Time
}

var Orders = []*Order{}

func newID() string {
	t := time.Now()
	entropy := ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)
	return ulid.MustNew(ulid.Timestamp(t), entropy).String()
}
