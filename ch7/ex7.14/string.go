package eval

import (
	"fmt"
	"strings"
)

func (b binary) String() string {
	return fmt.Sprintf("%s %c %s", b.x, b.op, b.y)
}

func (u unary) String() string {
	return fmt.Sprintf("%c%s", u.op, u.x)
}

func (v Var) String() string {
	return string(v)
}

func (c call) String() string {
	b := make([]string, len(c.args))
	for i, v := range c.args {
		b[i] = v.String()
	}
	return fmt.Sprintf("%s(%s)", c.fn, strings.Join(b, ", "))
}

func (l literal) String() string {
	return fmt.Sprintf("%g", float64(l))
}

func (c min) String() string {
	b := make([]string, len(c.args))
	for i, v := range c.args {
		b[i] = v.String()
	}
	return fmt.Sprintf("min(%s)", strings.Join(b, ", "))
}
