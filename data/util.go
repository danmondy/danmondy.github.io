package data

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"log"
	"time"
)

const timestampFmt = "20060102150405"

// NewUniqueID Returns a unique id that fills 12 bytes
func NewUniqueID() string {
	timestamp := time.Now().Format(timestampFmt)
	b := make([]byte, 12)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		log.Println(err)
		return ""
	}
	return fmt.Sprintf("%s%s%s", base64.URLEncoding.EncodeToString(b), "&tt&", timestamp)
}

// TimeToString converts a time.Time to RFC3339 string
func TimeToString(t time.Time) string {
	return t.Format(time.RFC3339)
}

// StringToTime used to convert a RFC3339 to time.Time
func StringToTime(s string) (time.Time, error) {
	return time.Parse(time.RFC3339, s) //returns both a time and an error so it can be returned directly.
}
func DateOnly(s string) string {
	if len(s) >= 10 {
		return s[:10]
	}
	return ""
}

func Contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
