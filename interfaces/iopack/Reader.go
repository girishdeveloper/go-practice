package iopack

type Reader interface {
	Read(p []byte) (n int, err error)
}
