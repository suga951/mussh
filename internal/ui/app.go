package ui

import (
	"fmt"
	"mussh/internal/player"
	"mussh/internal/ytmusic"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	song      string
	artist    string
	duration  string
	paused    bool
	player    player.Player
	searching bool
	textInput textinput.Model
}

func New() *Model {
	ti := textinput.New()
	ti.Placeholder = "Search a song..."
	ti.Focus()
	ti.CharLimit = 100
	ti.Width = 40

	return &Model{
		song:      "Through Struggle",
		artist:    "As I Lay Dying",
		duration:  "0:23 / 3:23",
		paused:    false,
		player:    player.Player{},
		searching: false,
		textInput: ti,
	}
}

func (m *Model) Init() tea.Cmd { return nil }
func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

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
		m.duration = msg.Duration

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

	return fmt.Sprintf(
		"\n %s  %s\n %s\n\n [%s]  %s\n\n %s\n",
		status,
		TitleStyle.Render(m.song),
		ArtistStyle.Render(m.artist),
		BarStyle.Render("████──────"),
		TimeStyle.Render(m.duration),
		HelpStyle.Render("[Space]Play/Pause [S]Search [Q]Quit"),
	)
}
