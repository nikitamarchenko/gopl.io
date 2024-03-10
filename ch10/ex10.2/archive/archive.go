package archive

import "fmt"


type Reader interface {
	Get() string
	Next() bool 
	Close()
}

type NewReaderFunc func (string) (Reader, error)
var readers []NewReaderFunc

func SetReader(f NewReaderFunc) {
	readers = append(readers, f)
}

func NewReader(path string) (Reader, error) {
	for _, r := range readers {
		if reader, err := r(path); err == nil {
			return reader, nil
		}
	}
	return nil, fmt.Errorf("archive: zero readers can read %s file", path)
}