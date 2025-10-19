package ui

import (
	"fmt"
	"time"
)

func formatDuration(d time.Duration) string {
	min := int(d.Minutes())
	sec := int(d.Seconds()) % 60
	return fmt.Sprintf("%02d:%02d", min, sec)
}

func ParseDuration(s string) (time.Duration, error) {
	var m, sec int

	//MM:SS first
	n, err := fmt.Sscanf(s, "%d:%d", &m, &sec)
	if err == nil && n == 2 {
		return time.Duration(m)*time.Minute + time.Duration(sec)*time.Second, nil
	}

	// HH:MM:SS fallback
	var h int
	n, err = fmt.Sscanf(s, "%d:%d:%d", &h, &m, &sec)
	if err == nil && n == 3 {
		return time.Duration(h)*time.Hour + time.Duration(m)*time.Minute + time.Duration(sec)*time.Second, nil
	}

	return 0, fmt.Errorf("invalid duration format: %s", s)
}
