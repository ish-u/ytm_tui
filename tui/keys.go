package tui

import "github.com/charmbracelet/bubbles/key"

type KeyMap struct {
	Exit         key.Binding
	PlayPause    key.Binding
	AddToQueue   key.Binding
	FocusToInput key.Binding
	SeekForward  key.Binding
	SeekBackward key.Binding
	NextSong     key.Binding
	PrevSong     key.Binding
	Help         key.Binding
}

func (k KeyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.PlayPause, k.NextSong, k.PrevSong}
}

func (k KeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.PlayPause, k.NextSong, k.PrevSong}, // second column
		{k.SeekForward, k.SeekBackward},
		{k.FocusToInput, k.Exit}, // first column
	}
}

// DefaultKeyMap returns a set of default keybindings.
var Keys = KeyMap{
	Exit: key.NewBinding(
		key.WithKeys("ctrl+c"),
		key.WithHelp("ctrl+c", "exit"),
	),
	PlayPause: key.NewBinding(
		key.WithKeys(" "),
		key.WithHelp("space", "Play/Pause"),
	),
	AddToQueue: key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("enter", "Add to Queue"),
	),
	FocusToInput: key.NewBinding(
		key.WithKeys("ctrl+a"),
		key.WithHelp("ctrl+a", "Focus on Input"),
	),
	SeekForward: key.NewBinding(
		key.WithKeys("right"),
		key.WithHelp("right", "Seek Forward"),
	),
	SeekBackward: key.NewBinding(
		key.WithKeys("left"),
		key.WithHelp("left", "Seek Backward"),
	),
	NextSong: key.NewBinding(
		key.WithKeys("ctrl+right"),
		key.WithHelp("ctrl+right", "Next Song"),
	),
	PrevSong: key.NewBinding(
		key.WithKeys("ctrl+left"),
		key.WithHelp("ctrl+left", "Next Song"),
	),
	Help: key.NewBinding(
		key.WithKeys("?"),
		key.WithHelp("?", "toggle help"),
	),
}
