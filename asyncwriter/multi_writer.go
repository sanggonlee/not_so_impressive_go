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
// When at least one of the writers returns an error, AsyncMultiWriter returns the error.
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
	var wg sync.WaitGroup
	wg.Add(len(aw.writers))

	type nAndErrType struct {
		n   int
		err error
	}

	nAndErrs := make([]nAndErrType, len(aw.writers))

	for i, w := range aw.writers {
		go func(i int, w io.Writer) {
			nAndErrs[i].n, nAndErrs[i].err = w.Write(p)
			wg.Done()
		}(i, w)
	}

	var total int
	var err error

	wg.Wait()

	for _, elt := range nAndErrs {
		total += elt.n
		if elt.err != nil {
			err = elt.err
		}
	}

	return total, err
}
