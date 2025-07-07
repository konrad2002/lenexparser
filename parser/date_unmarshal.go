package parser

import (
	"fmt"
	"strings"
	"time"
)

type CustomTime struct {
	time.Time
}

// implement encoding.TextUnmarshaler
func (ct *CustomTime) UnmarshalText(text []byte) error {
	s := strings.TrimSpace(string(text))

	// try formats in order
	layouts := []string{
		"02.01.2006", // 15.10.2002
		"15:04",      // 15:34
		time.RFC3339, // fallback
		"2006-01-02", // fallback
	}

	var err error
	for _, layout := range layouts {
		var t time.Time
		t, err = time.Parse(layout, s)
		if err == nil {
			ct.Time = t
			return nil
		}
	}
	return fmt.Errorf("cannot parse date: %q", s)
}
