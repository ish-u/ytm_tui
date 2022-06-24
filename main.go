package main

import (
	"fmt"
	"log"
	"os"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/faiface/beep"
	"github.com/faiface/beep/speaker"

	tui "github.com/ish-u/ytm_tui/tui"
)

func main() {

	path := "./audio"
	// Make audio dirtectory if doesn't exists
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.Mkdir(path, 0700)
		// TODO: handle error
		if err != nil {
			log.Fatal("./audio does not exists")
			return
		}
	}

	// Initial Sample Rate
	sr := beep.SampleRate(44100)
	speaker.Init(sr, sr.N(time.Second/10))

	tui.Queue.Current = 0
	tui.Queue.Loading = false
	// A zero Queue is an empty Queue.
	speaker.Play(&tui.Queue)

	m := tui.InitialModel()
	p := tea.NewProgram(m, tea.WithAltScreen())

	if err := p.Start(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
