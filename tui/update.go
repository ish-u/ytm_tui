package tui

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/faiface/beep/speaker"
)

// Update Method
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	var cmd tea.Cmd
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		// If we set a width on the help menu it can it can gracefully truncate
		// its view as needed.
		m.help.Width = msg.Width
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.Exit):
			go removeMP3()
			cmds = append(cmds, tea.Quit)
		case key.Matches(msg, m.keys.PlayPause):
			if len(m.queue.Streamers) > 0 {
				speaker.Lock()
				m.queue.Paused = !m.queue.Paused
				speaker.Unlock()
			}
		case key.Matches(msg, m.keys.AddToQueue):
			ID := m.textInput.Value()

			if len(ID) == 11 {
				m.queue.Loading = true
				video, err := m.client.GetVideo(ID)
				if err != nil {
					log.Println(err)
					break
				}
				m.enteredText = "https://www.youtube.com/watch?v=" + video.ID
				m.queue.Start = 0
				m.queue.End = 1
				go playVideo(ID, &m)
			} else if len(ID) == 34 {
				playlist, err := m.client.GetPlaylist(ID)
				if err != nil {
					log.Println(err)
					break
				}
				header := fmt.Sprintf("%s by %s", playlist.Title, playlist.Author)
				m.enteredText = header
				m.queue.Start = 0
				m.queue.End = len(playlist.Videos)
				go playPlaylist(*playlist, &m)
			}
			m.textInput.Reset()
			m.textInput.Blur()
		case key.Matches(msg, m.keys.FocusToInput):
			m.textInput.Focus()
		case key.Matches(msg, m.keys.SeekForward):
			speaker.Lock()
			m.queue.Streamers[m.queue.Current].Seek(m.queue.Streamers[m.queue.Current].Position() + int(m.sampleRate*5))
			speaker.Unlock()
		case key.Matches(msg, m.keys.SeekBackward):
			speaker.Lock()
			m.queue.Streamers[m.queue.Current].Seek(m.queue.Streamers[m.queue.Current].Position() - int(m.sampleRate*5))
			speaker.Unlock()
		case key.Matches(msg, m.keys.NextSong):
			m.queue.Next()
		case key.Matches(msg, m.keys.PrevSong):
			m.queue.Prev()
		case key.Matches(msg, m.keys.Help):
			m.help.ShowAll = !m.help.ShowAll
		}

	}

	if len(Queue.Streamers) > 0 {
		m.position = m.sampleRate.D((Queue.Streamers[Queue.Current].Position())).Round(time.Second)
	}
	m.streamers = len(Queue.Streamers)
	m.textInput, cmd = m.textInput.Update(msg)
	cmds = append(cmds, cmd)
	m.spinner, cmd = m.spinner.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func removeMP3() string {
	dirname := "." + string(filepath.Separator) + "audio" + string(filepath.Separator)

	d, err := os.Open(dirname)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer d.Close()

	files, err := d.Readdir(-1)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Reading " + dirname)

	for _, file := range files {
		if file.Mode().IsRegular() {
			if filepath.Ext(file.Name()) == ".mp3" {
				os.Remove("file.Name()")
				fmt.Println("Deleted - ", file.Name())
			}
		}
	}

	return "success"
}
