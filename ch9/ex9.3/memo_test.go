package memo

import (
	"fmt"
	"reflect"
	"sync"
	"testing"
)

func TestMemo_Get(t *testing.T) {

	tests := []struct {
		name      string
		arg       string
		want      interface{}
		wantErr   bool
		cancel    bool
		cacheSize uint
	}{
		{"normal", "key", nil, false, false, 1},
		{"canceled", "key", nil, true, true, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			f := func(k string, done chan struct{}) (interface{}, error) {
				if done != nil {
					<-done
					return nil, fmt.Errorf("canceled")
				}
				return nil, nil
			}

			memo := New(f)

			var c chan struct{}
			if tt.cancel {
				c = make(chan struct{})
			}

			var wg sync.WaitGroup
			var got interface{}
			var err error
			wg.Add(1)
			go func() {
				defer wg.Done()
				got, err = memo.Get(tt.arg, c)
			}()
			if c != nil {
				close(c)
			}
			wg.Wait()

			if (err != nil) != tt.wantErr {
				t.Errorf("Memo.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Memo.Get() = %v, want %v", got, tt.want)
			}
			if memo.cacheSize != tt.cacheSize {
				t.Errorf("memo.cacheSize = %v, want %v", memo.cacheSize, tt.cacheSize)
			}
		})
	}
}
