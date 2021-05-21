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
// As soon as one of the writers returns an error, AsyncMultiWriter returns the error.
// The error returned is not necessarily from the first writer that would result in error.
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
