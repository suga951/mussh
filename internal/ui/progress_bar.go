package ui

import (
	"fmt"
	"math"
	"strings"
	"time"
)

var gradientFrames = []string{"█", "▓", "▒", "░"} // from solid to empty

func ProgressBar(current, total time.Duration, width int, pulseFrame int) string {
	if total <= 0 {
		total = 1
	}

	progress := float64(current) / float64(total)
	if progress > 1 {
		progress = 1
	}

	filled := int(progress * float64(width))
	filled = int(math.Max(0, math.Min(float64(filled), float64(width))))

	bar := ""
	for i := 0; i < filled; i++ {
		if i >= filled-3 {
			idx := (i + pulseFrame) % len(gradientFrames)
			bar += gradientFrames[idx]
		} else {
			bar += gradientFrames[0]
		}
	}

	emptyCount := int(math.Max(0, float64(width-filled)))
	bar += strings.Repeat("─", emptyCount)

	return fmt.Sprintf("[%s]  %s / %s", bar, formatDuration(current), formatDuration(total))
}
