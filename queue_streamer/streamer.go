package queuestreamer

import (
	"github.com/faiface/beep"
	"github.com/faiface/beep/speaker"
	"github.com/kkdai/youtube/v2"
)

type QueueStreamer struct {
	Streamers []beep.StreamSeekCloser
	Current   int
	Paused    bool
	Videos    []youtube.Video
	Loading   bool
	Start     int
	End       int
}

func (q *QueueStreamer) AddVideo(Videos ...youtube.Video) {
	q.Videos = append(q.Videos, Videos...)
}

func (q *QueueStreamer) Add(Streamers ...beep.StreamSeekCloser) {
	q.Streamers = append(q.Streamers, Streamers...)
}

func (q *QueueStreamer) Stream(samples [][2]float64) (n int, ok bool) {
	// We use the filled variable to track how many samples we've
	// successfully filled already. We loop until all samples are filled.
	filled := 0
	for filled < len(samples) {
		// There are no Streamers in the QueueStreamer, so we stream silence.
		if len(q.Streamers) == 0 || q.Paused || q.Current < 0 {
			for i := range samples[filled:] {
				samples[i][0] = 0
				samples[i][1] = 0
			}
			break
		}

		// We stream from the first streamer in the QueueStreamer.
		n, ok := q.Streamers[q.Current].Stream(samples[filled:])
		// If it's drained, we pop it from the QueueStreamer, thus continuing with
		// the next streamer.
		if !ok {
			streamer := q.Streamers[q.Current]
			streamer.Seek(0)
			q.Current++
			if q.Current == len(q.Streamers) {
				q.Current = 0
			}
			// q.Streamers = q.Streamers[1:]
			// q.Streamers = append(q.Streamers, streamer)

			// video := q.Videos[0]
			// q.Videos = q.Videos[1:]
			// q.Videos = append(q.Videos, video)

		}
		// We update the number of filled samples.
		filled += n
	}
	return len(samples), true
}

func (q *QueueStreamer) Err() error {
	return nil
}

func (q *QueueStreamer) Next() {
	speaker.Lock()
	if len(q.Streamers) > 1 {
		streamer := q.Streamers[q.Current]
		streamer.Seek(0)

		q.Current++
		if q.Current == len(q.Streamers) {
			q.Current = 0
		}

	}
	speaker.Unlock()
}

func (q *QueueStreamer) Prev() {
	speaker.Lock()
	if len(q.Streamers) > 1 {

		streamer := q.Streamers[q.Current]
		streamer.Seek(0)
		q.Current--
		if q.Current < 0 {
			q.Current = len(q.Streamers) - 1
		}

	}
	speaker.Unlock()
}
