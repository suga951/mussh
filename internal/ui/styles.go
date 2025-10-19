package ui

import "github.com/charmbracelet/lipgloss"

type Theme struct {
	TitleColor   string
	ArtistColor  string
	TimeColor    string
	PlayBarColor string
	HelpBarColor string
}

var neonCyberTheme = Theme{
	TitleColor:   "#005B80",
	ArtistColor:  "#005B80",
	TimeColor:    "#005B80",
	PlayBarColor: "#00FFAA",
	HelpBarColor: "#00B3A1",
}

var (
	TitleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color(neonCyberTheme.TitleColor))

	ArtistStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color(neonCyberTheme.ArtistColor))

	TimeStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color(neonCyberTheme.TimeColor))

	BarStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color(neonCyberTheme.PlayBarColor))

	HelpStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color(neonCyberTheme.HelpBarColor)).
			Italic(true)
)
