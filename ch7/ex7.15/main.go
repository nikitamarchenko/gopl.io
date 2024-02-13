package main

import (
	"fmt"
	"os"
	"path"
	"strconv"
	"strings"

	eval "gopl.io/ch7/ex7.14"
)

func printErrorAndExit(e error) {

	app := os.Args[0]
	if strings.Contains(app, "/") {
		app = path.Base(app)
	}

	fmt.Printf("\n%s: Error: %s\n\n", app, e)
	fmt.Printf(`%s: usage %s "var1=1,var2=2;expr"`, app, app)
	fmt.Printf("\n\n")
	os.Exit(1)
}

func main() {
	env := eval.Env{}
	vars := make(map[eval.Var]bool)

	if len(os.Args) != 2 {
		printErrorAndExit(fmt.Errorf("error: no expr"))
	}

	command := os.Args[1]
	var parseArgs bool
	switch strings.Count(command, ";") {
	case 0:
		parseArgs = false
	case 1:
		parseArgs = true
	default:
		printErrorAndExit(fmt.Errorf(`too many ";" in expr, allowed only one`))
	}

	exprCommand := command
	if parseArgs {
		tokens := strings.Split(command, ";")
		exprCommand = tokens[1]
		argCommand := tokens[0]
		for _, pairs := range strings.Split(argCommand, ",") {
			if strings.Count(pairs, "=") != 1 {
				printErrorAndExit(fmt.Errorf("invalid var expr: must contain ="))
			}
			parts := strings.Split(pairs, "=")
			val, err := strconv.ParseFloat(parts[1], 64)
			if err != nil {
				printErrorAndExit(fmt.Errorf("error: can't parse float %s", err))
			}
			varName := eval.Var(strings.TrimSpace(parts[0]))
			env[varName] = val
			vars[varName] = true
		}
	}

	expr, err := eval.Parse(exprCommand)
	if err != nil {
		printErrorAndExit(fmt.Errorf("can't parse with error %s", err))
	}
	err = expr.Check(vars)
	if err != nil {
		printErrorAndExit(fmt.Errorf("expr check failed with error %s", err))
		os.Exit(1)
	}
	r := expr.Eval(env)
	fmt.Printf("expr: %s env: %s res: %g\n", expr, env, r)
}
