/*

Exercise 10.4: Construct a tool that reports the set of all packages in the
workspace that transitively depend on the packages specified by the arguments.
Hint: you will need to run go list twice, once for the initial packages and
once for all packages. You may want to parse its JSON output using the
encoding/json package (§4.5).

*/

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"sort"
)

type Package struct {
	ImportPath string   // import path of package in dir
	Deps       []string // all (recursively) imported dependencies
}

func resolveImportPaths(pkgNames []string) map[string]bool {

	result := make(map[string]bool)
	// go list -json names
	cmdArg := []string{
		"list", "-json",
	}
	cmdArg = append(cmdArg, pkgNames...)

	out, err := exec.Command("go", cmdArg...).Output()
	if err != nil {
		log.Fatal(err)
	}

	dec := json.NewDecoder(bytes.NewReader(out))
	for {
		var p Package
		if err := dec.Decode(&p); err != nil {
			if err == io.EOF {
				break
			}
			log.Fatalf("reading go list output: %v", err)
		}
		result[p.ImportPath] = true
	}

	return result
}

func getAncestors(pkgs map[string]bool) []string {
	cmdArg := []string{
		"list", "-json", "all",
	}
	out, err := exec.Command("go", cmdArg...).Output()
	if err != nil {
		log.Fatal(err)
	}

	set := make(map[string]bool)
	dec := json.NewDecoder(bytes.NewReader(out))
	for {
		var p Package
		if err := dec.Decode(&p); err != nil {
			if err == io.EOF {
				break
			}
			log.Fatalf("reading go list output: %v", err)
		}
		for _, d := range p.Deps {
			if pkgs[d] {
				set[p.ImportPath] = true
			}
		}
	}

	result := make([]string, len(set))

	for k := range set {
		result = append(result, k)
	}

	sort.Strings(result)

	return result
}

// usage: go run . fmt net/url
func main() {
	if len(os.Args) == 1 {
		log.Fatal("Please specify package name(s)")
	}

	pkgs := resolveImportPaths(os.Args[1:])
	for _, n := range getAncestors(pkgs) {
		fmt.Println(n)
	}
}
