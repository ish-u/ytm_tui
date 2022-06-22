package tui

import (
	"io"
	"log"
	"os"
	"time"

	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
	"github.com/kkdai/youtube/v2"
	fluentffmpeg "github.com/modfy/fluent-ffmpeg"
)

func getData(videoID string, stream io.ReadCloser, file *os.File) {
	err := fluentffmpeg.NewCommand("").
		PipeInput(stream).
		OutputFormat("mp3").
		PipeOutput(file).
		Run()
	if err != nil {
		log.Fatal(err)
	}
}

func playVideo(videoID string, m *Model) {

	file, err := os.OpenFile("./audio/"+videoID+".mp3", os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	video, err := m.client.GetVideo(videoID)
	if err != nil {
		panic(err)
	}
	stream, _, err := m.client.GetStream(video, video.Formats.FindByItag(140))
	if err != nil {
		panic(err)
	}

	go getData(videoID, stream, file)
	time.Sleep(time.Second * 1)

	f, err := os.Open("./audio/" + videoID + ".mp3")
	if err != nil {
		// log.Fatal(err)
		return
	}

	streamer, _, err := mp3.Decode(f)
	if err != nil {
		return
		// fmt.Println(err)
		// log.Fatal(err)
	}

	// The speaker's sample rate is fixed at 44100. Therefore, we need to
	// resample the file in case it's in a different sample rate.
	// resampled := beep.Resample(4, format.SampleRate, sr, streamer)

	// And finally, we add the song to the Queue.
	speaker.Lock()
	Queue.Add(streamer)
	Queue.AddVideo(*video)
	speaker.Unlock()

	m.queue.Loading = false
	m.queue.Start++

}

func playPlaylist(playlist youtube.Playlist, m *Model) {

	for _, v := range playlist.Videos {
		// if k == 10 {
		// 	break
		// }
		m.queue.Loading = true
		playVideo(v.ID, m)

	}
	m.queue.Loading = false

}
