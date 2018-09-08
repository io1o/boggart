package players

import (
	"errors"
	"io"
	"net/http"
	"os"
	"sync"
	"sync/atomic"
	"time"

	"github.com/hajimehoshi/go-mp3"
	"github.com/hajimehoshi/oto"
)

/*
	TODO: wav, flac
*/

const (
	AudioFormatMP3 = "mp3"
)

var (
	ErrorAlreadyPlaying     = errors.New("Already playing")
	ErrorUnknownAudioFormat = errors.New("Unknown audio format")
	ErrorStreamEmpty        = errors.New("Stream isn't set")
)

type AudioPlayer struct {
	status        int64
	volumePercent int64

	mutex   sync.RWMutex
	speaker *SpeakerWrapper
	stream  *StreamWrapper
}

func NewAudio() *AudioPlayer {
	p := &AudioPlayer{}
	p.setStatus(StatusStopped)
	p.SetVolume(50)

	return p
}

func (p *AudioPlayer) PlayFromURL(url string) error {
	if p.Status() == StatusPlaying {
		return ErrorAlreadyPlaying
	}

	client := &http.Client{
		Transport: http.DefaultTransport,
		// что бы не рвались стримы
		Timeout: 0,
	}

	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return err
	}

	response, err := client.Do(request)
	if err != nil {
		return err
	}

	mimeType, err := mimeFromHeader(response.Header)
	if err != nil {
		return err
	}

	if mimeType == "" {
		mimeType, err = mimeFromData(response.Body)
	}

	var format string

	switch mimeType {
	case "audio/mpeg":
		format = AudioFormatMP3

		// TODO: wav, flac

	default:
		return ErrorUnknownAudioFormat
	}

	if err := p.initStream(response.Body, format); err != nil {
		return err
	}

	return p.Play()
}

func (p *AudioPlayer) PlayFromReader(reader io.ReadCloser) error {
	if p.Status() == StatusPlaying {
		return ErrorAlreadyPlaying
	}

	mimeType, err := mimeFromData(reader)
	if err != nil {
		return err
	}

	var format string

	switch mimeType {
	case "audio/mpeg":
		format = AudioFormatMP3

		// TODO: wav, flac

	default:
		return ErrorUnknownAudioFormat
	}

	if err := p.initStream(reader, format); err != nil {
		return err
	}

	return p.Play()
}

func (p *AudioPlayer) PlayFromFile(file string) error {
	if p.Status() == StatusPlaying {
		return ErrorAlreadyPlaying
	}

	f, err := os.Open(file)
	if err != nil {
		return err
	}

	return p.PlayFromReader(f)
}

func (p *AudioPlayer) Play() error {
	if p.Status() == StatusPlaying {
		return ErrorAlreadyPlaying
	}

	if p.getStream().IsClosed() {
		return ErrorStreamEmpty
	}

	if p.getSpeaker().IsClosed() {
		if err := p.initSpeaker(); err != nil {
			return err
		}
	}

	go p.play()

	return nil
}

func (p *AudioPlayer) Stop() error {
	p.setStatus(StatusStopped)
	p.getSpeaker().Close()
	p.getStream().Close()

	return nil
}

func (p *AudioPlayer) Pause() error {
	if p.Status() == StatusPlaying {
		p.setStatus(StatusPause)
		p.getSpeaker().Close()
	}

	return nil
}

func (p *AudioPlayer) Volume() int64 {
	return atomic.LoadInt64(&p.volumePercent)
}

func (p *AudioPlayer) VolumeUp() error {
	vol := p.Volume()
	if vol == 100 {
		return nil
	}

	return p.SetVolume(vol + 1)
}

func (p *AudioPlayer) VolumeDown() error {
	vol := p.Volume()
	if vol == 0 {
		return nil
	}

	return p.SetVolume(vol - 1)
}

func (p *AudioPlayer) SetVolume(percent int64) error {
	if percent > 100 {
		percent = 100
	} else if percent < 0 {
		percent = 0
	}

	atomic.StoreInt64(&p.volumePercent, percent)

	return nil
}

func (p *AudioPlayer) Close() {
	p.Stop()
}

func (p *AudioPlayer) initSpeaker() error {
	p.getSpeaker().Close()

	stream := p.getStream()
	if stream.IsClosed() {
		return ErrorStreamEmpty
	}

	// FIXME: magic
	numBytes := int(time.Second/10*time.Duration(stream.SampleRate())/time.Second) * 4

	speaker, err := oto.NewPlayer(stream.SampleRate(), 2, 2, numBytes)
	if err != nil {
		return err
	}

	p.mutex.Lock()
	p.speaker = NewSpeakerWrapper(speaker)
	p.mutex.Unlock()

	return nil
}

func (p *AudioPlayer) initStream(s io.ReadCloser, format string) error {
	p.getStream().Close()

	switch format {
	case AudioFormatMP3:
		stream, err := mp3.NewDecoder(s)
		if err != nil {
			return err
		}

		p.mutex.Lock()
		p.stream = NewStreamWrapper(stream)
		p.mutex.Unlock()

	default:
		return ErrorUnknownAudioFormat
	}

	return nil
}

func (p *AudioPlayer) getSpeaker() *SpeakerWrapper {
	p.mutex.RLock()
	defer p.mutex.RUnlock()

	return p.speaker
}

func (p *AudioPlayer) getStream() *StreamWrapper {
	p.mutex.RLock()
	defer p.mutex.RUnlock()

	return p.stream
}

func (p *AudioPlayer) setStatus(status Status) {
	atomic.StoreInt64(&p.status, status.Int64())
}

func (p *AudioPlayer) Status() Status {
	return Status(atomic.LoadInt64(&p.status))
}

func (p *AudioPlayer) play() {
	p.setStatus(StatusPlaying)

	defer func() {
		p.getSpeaker().Close()

		if p.Status() != StatusPause {
			p.getStream().Close()
		}
	}()

	_, _ = io.Copy(p.getSpeaker(), p.getStream())
}
