package tar

import (
	"archive/tar"
	"fmt"
	"io"
	"os"

	"gopl.io/ch10/ex10.2/archive"
)

func init() {
	archive.SetReader(NewReader)
}

type reader struct {
	//r *tar.ReadCloser
	r []string
	i int
}

func (r *reader) Get() string {
	if r.i >= 0 && r.i < len(r.r) {
		return r.r[r.i]
	}
	return ""
}

func (r *reader) Next() bool {
	r.i++
	return r.i < len(r.r)
}

func (r *reader) Close() {
}

func NewReader(p string) (archive.Reader, error) {

	buf, err := os.Open(p)
	if err != nil {
		return nil, err
	}
	defer buf.Close()

	var r []string
	tr := tar.NewReader(buf)
	for {
		hdr, err := tr.Next()
		if err == io.EOF {
			break // End of archive
		}
		if err != nil {
			return nil, fmt.Errorf("tar: %s", err)
		}
		fmt.Println("-", hdr.Name)
		r = append(r, hdr.Name)
	}

	return &reader{r: r, i: -1}, nil
}
