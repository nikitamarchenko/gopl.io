/*

Exercise 7.18: Using the token-based decoder API, write a program that will read
an arbitrary XML document and construct a tree of generic nodes that represents
it. Nodes are of two kinds: CharData nodes represent text strings, and Element
nodes represent named elements and their attributes. Each element node has a
slice of child nodes.

You may find the following declarations helpful.

Click here to view code image

import "encoding/xml"

type Node interface{} // CharData or *Element

type CharData string

type Element struct {
    Type     xml.Name
    Attr     []xml.Attr
    Children []Node
}

*/

package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
)

type Node interface{} // CharData or *Element

type CharData string

type Element struct {
	Type     xml.Name
	Attr     []xml.Attr
	Children []Node
}

func (el *Element) appendChildren(n Node) {
	el.Children = append(el.Children, n)
}

func main() {
	n := decode(xml.NewDecoder(os.Stdin))
	print(n)
	fmt.Println()
}

func decode(d *xml.Decoder) Node {
	var root Node
	var stack []*Element // stack of element names
	for {
		tok, err := d.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "xmlselect: %v\n", err)
			os.Exit(1)
		}
		switch tok := tok.(type) {
		case xml.StartElement:
			el := &Element{Type: tok.Name, Attr: tok.Attr, Children: []Node{}}
			if root == nil {
				root = el
			} else {
				stack[len(stack)-1].appendChildren(el)
			}
			stack = append(stack, el) // push
		case xml.EndElement:
			stack = stack[:len(stack)-1] // pop
		case xml.CharData:
			stack[len(stack)-1].appendChildren(CharData(string(tok)))
		}
	}
	return root
}

func print(root Node) {
	switch tok := root.(type) {
	case *Element:
		fmt.Printf("<%s>", tok.Type.Local)
		for _, c := range tok.Children {
			print(c)
		}
		fmt.Printf("</%s>", tok.Type.Local)
	case CharData:
		fmt.Printf("TEXT{%s}", string(tok))
	}
}
