/*

Exercise 13.4: Depending on C libraries has its drawbacks. Provide an
alternative pure-Go implementation of bzip.NewWriter that uses the os/exec
package to run /bin/bzip2 as a subprocess.

*/

package bzip

import (
	"io"
	"os/exec"
	"sync"
)

type writer struct {
	w     io.Writer
	wg    *sync.WaitGroup
	cmd   *exec.Cmd
	stdin io.WriteCloser
}

func NewWriter(out io.Writer) io.WriteCloser {
	return &writer{w: out}
}

func (w *writer) Write(data []byte) (int, error) {
	if w.cmd == nil {
		w.wg = &sync.WaitGroup{}
		w.cmd = exec.Command("/bin/bzip2")
		var err error
		w.stdin, err = w.cmd.StdinPipe()
		if err != nil {
			return 0, err
		}
		stdout, err := w.cmd.StdoutPipe()
		if err != nil {
			return 0, err
		}
		w.wg.Add(1)
		go func() {
			io.Copy(w.w, stdout)
			w.wg.Done()
		}()
		err = w.cmd.Start()
		if err != nil {
			return 0, err
		}
	}
	return w.stdin.Write(data)
}

func (w *writer) Close() error {
	err := w.stdin.Close()
	if err != nil {
		return err
	}
	err = w.cmd.Wait()
	if err != nil {
		return err
	}
	w.wg.Wait()
	return nil
}
