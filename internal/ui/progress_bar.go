package ui

import (
	"fmt"
	"strings"
	"time"
)

func ProgressBar(current, total time.Duration, width int, pulseFrame int) string {
	progress := float64(current) / float64(total)
	filled := int(progress * float64(width))

	if filled >= width {
		filled = width - 1
	}

	full := "█"
	pulseChars := []string{"▒", "▓"}
	empty := "░"

	pulse := pulseChars[pulseFrame%len(pulseChars)]

	bar := fmt.Sprintf("[%s%s%s]",
		strings.Repeat(full, filled),
		pulse,
		strings.Repeat(empty, width-filled-1),
	)

	return fmt.Sprintf("%s  %s / %s",
		bar,
		formatTime(current),
		formatTime(total),
	)
}

func formatTime(d time.Duration) string {
	min := int(d.Minutes())
	sec := int(d.Seconds()) % 60
	return fmt.Sprintf("%02d:%02d", min, sec)
}
