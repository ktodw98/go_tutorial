package main

import (
	"io"
)

type progressReader struct {
	Reader  io.Reader
	Total   int64
	Current int64
	pChan   chan int64
}

func (pr *progressReader) Read(p []byte) (n int, err error) {
	n, err = pr.Reader.Read(p)

	if n > 0 {
		pr.Current += int64(n)

		select {
		case pr.pChan <- pr.Current:
		default:
		}
	}
	return n, err
}
