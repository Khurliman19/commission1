package utils

import (
	"math/rand"
	"strings"
	"time"
)

func MaskCard(card string) string {
	if len(card) < 4 {
		return "INVALID"
	}
	last4 := card[len(card)-4:]
	return "UZCARD**" + last4
}

func GenerateTransactionID() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(900000000) + 100000000
}

func CurrentDateTime() string {
	return time.Now().Format("02.01.2006 15:04")
}

func ToUpperFullName(first, last string) string {
	return strings.ToUpper(last + " " + first)
}
