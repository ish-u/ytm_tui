package tui

import (
	"time"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/faiface/beep"
	queuestreamer "github.com/ish-u/ytm_tui/queue_streamer"
	"github.com/kkdai/youtube/v2"
)

var Queue queuestreamer.QueueStreamer

type Model struct {
	textInput   textinput.Model
	err         error
	enteredText string
	client      youtube.Client
	position    time.Duration
	streamers   int
	sampleRate  beep.SampleRate
	loading     bool
	queue       *queuestreamer.QueueStreamer
	spinner     spinner.Model
	help        help.Model
	keys        KeyMap
}
