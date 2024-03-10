package zip

import (
	"archive/zip"
	"fmt"

	"gopl.io/ch10/ex10.2/archive"
)

func init() {
	archive.SetReader(NewReader)
}

type reader struct {
	r *zip.ReadCloser
	i int
}

func (r *reader) Get() string {
	if r.i >= 0 && r.i < len(r.r.File) {
		return r.r.File[r.i].Name
	}
	return ""
}

func (r *reader) Next() bool {
	r.i++
	return r.i < len(r.r.File)
}

func (r *reader) Close() {
	r.r.Close()
}

func NewReader(p string) (archive.Reader, error) {
	r, err := zip.OpenReader(p)
	if err != nil {
		return nil, fmt.Errorf("zip: %s", err)
	}

	return &reader{r: r, i: -1}, nil
}
