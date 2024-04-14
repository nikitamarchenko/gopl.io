package bzip

/*
#cgo CFLAGS: -I/usr/include
#cgo LDFLAGS: -L/usr/lib -lbz2
#include <bzlib.h>
#include <stdlib.h>
int bz2compress(bz_stream *s, int action,
                char *in, unsigned *inlen, char *out, unsigned *outlen);
bz_stream* bz2alloc() { return calloc(1, sizeof(bz_stream)); }
void bz2free(bz_stream* s) { free(s); }
*/
import "C"

import (
	"io"
	"sync"
	"unsafe"
)

type writer struct {
	w      io.Writer // underlying output stream
	stream *C.bz_stream
	outbuf [64 * 1024]byte
	sync.Mutex
}

func NewWriter(out io.Writer) io.WriteCloser {
	const blockSize = 9
	const verbosity = 0
	const workFactor = 30
	w := &writer{w: out, stream: C.bz2alloc()}
	C.BZ2_bzCompressInit(w.stream, blockSize, verbosity, workFactor)
	return w
}

func (w *writer) Write(data []byte) (int, error) {
	if w.stream == nil {
		panic("closed")
	}
	var total int // uncompressed bytes written

	w.Lock()
	defer w.Unlock()

	for len(data) > 0 {
		inlen, outlen := C.uint(len(data)), C.uint(cap(w.outbuf))
		C.bz2compress(w.stream, C.BZ_RUN,
			(*C.char)(unsafe.Pointer(&data[0])), &inlen,
			(*C.char)(unsafe.Pointer(&w.outbuf)), &outlen)
		total += int(inlen)
		data = data[inlen:]
		if _, err := w.w.Write(w.outbuf[:outlen]); err != nil {
			return total, err
		}
	}
	return total, nil
}

func (w *writer) Close() error {
	if w.stream == nil {
		panic("closed")
	}
	w.Lock()
	defer func() {
		C.BZ2_bzCompressEnd(w.stream)
		C.bz2free(w.stream)
		w.stream = nil
		w.Unlock()
	}()
	for {
		inlen, outlen := C.uint(0), C.uint(cap(w.outbuf))
		r := C.bz2compress(w.stream, C.BZ_FINISH, nil, &inlen,
			(*C.char)(unsafe.Pointer(&w.outbuf)), &outlen)
		if _, err := w.w.Write(w.outbuf[:outlen]); err != nil {
			return err
		}
		if r == C.BZ_STREAM_END {
			return nil
		}
	}
}
