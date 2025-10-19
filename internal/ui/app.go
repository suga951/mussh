package ui

import (
	"fmt"
	"mussh/internal/player"
	"mussh/internal/ytmusic"
	"time"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type tickMsg struct{}

type Model struct {
	song       string
	artist     string
	current    time.Duration
	total      time.Duration
	pulseFrame int
	ticker     *time.Ticker
	paused     bool
	player     player.Player
	searching  bool
	textInput  textinput.Model
}

func New() *Model {
	ti := textinput.New()
	ti.Placeholder = "Search a song..."
	ti.Focus()
	ti.CharLimit = 100
	ti.Width = 40

	return &Model{
		song:      "",
		artist:    "",
		current:   0,
		total:     0,
		paused:    false,
		player:    player.Player{},
		searching: false,
		textInput: ti,
	}
}

func (m *Model) Init() tea.Cmd {
	m.ticker = time.NewTicker(100 * time.Millisecond)
	return tick()
}

func tick() tea.Cmd {
	return tea.Tick(100*time.Millisecond, func(t time.Time) tea.Msg {
		return t
	})
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case time.Time:
		if !m.paused {
			m.current += 100 * time.Millisecond
			if m.current > m.total {
				m.current = m.total
			}
		}
		m.pulseFrame = (m.pulseFrame + 1) % 6
		return m, tick()

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			m.player.Stop()
			return m, tea.Quit
		}

		if m.searching {
			var cmd tea.Cmd
			m.textInput, cmd = m.textInput.Update(msg)

			switch msg.String() {
			case "enter":
				query := m.textInput.Value()
				m.textInput.SetValue("")
				m.searching = false
				return m, fetchSong(query)
			case "esc":
				m.textInput.SetValue("")
				m.searching = false
			}

			return m, cmd
		}

		switch msg.String() {
		case " ":
			m.paused = !m.paused
		case "s":
			m.searching = true
			m.textInput.Focus()
		}

	case *ytmusic.Song:
		m.song = msg.Title
		m.artist = msg.Artist
		dur, err := ParseDuration(msg.Duration)
		if err != nil {
			fmt.Println("error parsing duration:", err)
			dur = 0
		}
		m.total = dur
		m.current = 0
		m.pulseFrame = 0
		go m.player.Play(msg.URL)
	}

	return m, nil
}

func fetchSong(query string) tea.Cmd {
	return func() tea.Msg {
		song, err := ytmusic.FetchSong(query)
		if err != nil {
			return fmt.Sprintf("Error: %v", err)

		}
		return song
	}
}
func (m *Model) View() string {
	if m.searching {
		return "\nSearch: " + m.textInput.View() + "\n\n"
	}

	status := "Now playing..."
	if m.paused {
		status = "Paused"
	}

	bar := ProgressBar(m.current, m.total, 30, m.pulseFrame)

	return fmt.Sprintf(
		"\n %s  %s\n %s\n\n %s\n\n %s\n",
		status,
		TitleStyle.Render(m.song),
		ArtistStyle.Render(m.artist),
		bar,
		HelpStyle.Render("[Space]Play/Pause [S]Search [Q]Quit"),
	)

}
