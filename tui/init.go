package tui

import (
	"time"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/faiface/beep"
	"github.com/faiface/beep/speaker"
	"github.com/kkdai/youtube/v2"
)

func (m Model) Init() tea.Cmd {
	return m.spinner.Tick
}

func InitialModel() Model {

	// Initial Sample Rate
	sr := beep.SampleRate(44100)
	speaker.Init(sr, sr.N(time.Second/10))

	Queue.Current = 0
	// A zero Queue is an empty Queue.
	speaker.Play(&Queue)

	ti := textinput.New()
	ti.Placeholder = "Enter Video/Playlist ID\n"
	ti.Focus()
	ti.CharLimit = 34
	ti.Width = 34

	s := spinner.New()
	s.Spinner = spinner.Dot

	return Model{
		textInput:   ti,
		err:         nil,
		enteredText: "",
		client:      youtube.Client{},
		streamers:   0,
		sampleRate:  beep.SampleRate(44100),
		loading:     false,
		queue:       &Queue,
		spinner:     s,
		help:        help.New(),
		keys:        Keys,
	}
}
