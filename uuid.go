package main

import (
	"fmt"
	"math/rand"

	"github.com/gofrs/uuid"
)

func generateSessionId() string {
	u2, err := uuid.NewV4()
	if err != nil {
		milli := getCurrentMilli()
		rand.Seed(milli)
		sessionId := fmt.Sprintf("%d-%d", milli, rand.Intn(10000000))
		return sessionId
	}
	return u2.String()
}
