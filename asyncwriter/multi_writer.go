package asyncwriter

import (
	"io"
	"sync"
)

// AsyncMultiWriter behaves similarly as io.MultiWriter, but the write on each writer
// is run concurrently.
// Writing on every writer is attempted, regardless of any errors returned from
// any of the writers.
// The n returned is the sum of number of bytes for all the writes performed.
// The err returned includes all the error messages from any writers which returned
// error and formats them by delimiting with commas.
// The err could also result from attempting to write on a closed AsyncMultiWriter.
func AsyncMultiWriter(writers ...io.Writer) io.Writer {
	return &asyncMultiWriter{
		writers: writers,
	}
}

type asyncMultiWriter struct {
	writers []io.Writer
}

func (aw *asyncMultiWriter) Write(p []byte) (int, error) {
	nchan := make(chan int, len(aw.writers))
	errchan := make(chan error, len(aw.writers))

	var wg sync.WaitGroup
	wg.Add(len(aw.writers))

	for _, w := range aw.writers {
		go func(w io.Writer, nchan chan<- int, errchan chan<- error) {
			n, err := w.Write(p)
			go func() { nchan <- n }()
			go func() { errchan <- err }()
			wg.Done()
		}(w, nchan, errchan)
	}

	wg.Wait()

	var total int
	var err error
	for i := 0; i < 2*len(aw.writers); i++ {
		select {
		case n := <-nchan:
			total += n
		case writeErr := <-errchan:
			if writeErr != nil {
				err = writeErr
			}
		}
	}

	return total, err
}
