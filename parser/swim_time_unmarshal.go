package parser

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

type SwimTime struct {
	time.Duration
}

// implement encoding.TextUnmarshaler
func (st *SwimTime) UnmarshalText(text []byte) error {
	s := strings.TrimSpace(string(text))

	// Example: "01:02:03,45"
	parts := strings.SplitN(s, ".", 2)
	if len(parts) != 2 {
		return fmt.Errorf("invalid duration format: %q", s)
	}

	hms := parts[0]        // "01:02:03"
	hundredths := parts[1] // "45"

	hmsParts := strings.Split(hms, ":")
	if len(hmsParts) != 3 {
		return fmt.Errorf("invalid HH:MM:SS part: %q", hms)
	}

	hours, err := strconv.Atoi(hmsParts[0])
	if err != nil {
		return fmt.Errorf("invalid hours: %w", err)
	}
	minutes, err := strconv.Atoi(hmsParts[1])
	if err != nil {
		return fmt.Errorf("invalid minutes: %w", err)
	}
	seconds, err := strconv.Atoi(hmsParts[2])
	if err != nil {
		return fmt.Errorf("invalid seconds: %w", err)
	}
	hundredthsInt, err := strconv.Atoi(hundredths)
	if err != nil {
		return fmt.Errorf("invalid hundredths: %w", err)
	}

	// compute total duration
	total := time.Duration(hours)*time.Hour +
		time.Duration(minutes)*time.Minute +
		time.Duration(seconds)*time.Second +
		time.Duration(hundredthsInt*10)*time.Millisecond

	st.Duration = total
	return nil
}
