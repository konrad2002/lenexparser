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

	// Example: "01:23,45"
	parts := strings.SplitN(s, ",", 2)
	if len(parts) != 2 {
		return fmt.Errorf("invalid duration format: %q", s)
	}

	minSec := parts[0]     // "01:23"
	hundredths := parts[1] // "45"

	msParts := strings.SplitN(minSec, ":", 2)
	if len(msParts) != 2 {
		return fmt.Errorf("invalid minutes:seconds part: %q", minSec)
	}

	minutes, err := strconv.Atoi(msParts[0])
	if err != nil {
		return fmt.Errorf("invalid minutes: %w", err)
	}
	seconds, err := strconv.Atoi(msParts[1])
	if err != nil {
		return fmt.Errorf("invalid seconds: %w", err)
	}
	hundredthsInt, err := strconv.Atoi(hundredths)
	if err != nil {
		return fmt.Errorf("invalid hundredths: %w", err)
	}

	// calculate total duration
	total := time.Duration(minutes)*time.Minute +
		time.Duration(seconds)*time.Second +
		time.Duration(hundredthsInt*10)*time.Millisecond

	st.Duration = total
	return nil
}
