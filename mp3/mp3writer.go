package mp3

import (
	"io"
	"os"
	"os/exec"
)

// Writer reads the video flux read in `in`, writes it out to `out`
type Writer interface {
	Write(in io.Reader, out io.Writer) error
}

type mp3Writer struct {
}

// NewWriter creates a Writer with an underlying implementation using ffmpeg
func NewWriter() Writer {
	return &mp3Writer{}
}

// Write writes mp3 from reader using ffmpeg and writes to provided output
func (w *mp3Writer) Write(in io.Reader, out io.Writer) error {
	ffmpeg, err := exec.LookPath("ffmpeg")
	if err != nil {
		return err
	}

	cmd := exec.Command(ffmpeg, "-y", "-loglevel", "quiet", "-i", "-", "-f", "mp3", "-vn", "-")
	cmd.Stdin = in
	cmd.Stdout = out
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
