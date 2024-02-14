package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	eval "gopl.io/ch7/ex7.14"
)

type TemplateData struct {
	Expr, Result, Error string
}

func main() {
	indexHtml, err := os.ReadFile("./index.html")
	if err != nil {
		log.Fatalf("error while read template: %s", err)
	}
	t := template.Must(template.New("index").Parse(string(indexHtml)))
	http.HandleFunc("/", calcHandler(t))
	log.Fatal(http.ListenAndServe("127.0.0.1:8000", nil))
}

func calcHandler(t *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			r.ParseForm()
			exprText := r.PostForm.Get("expr")
			_, _, result, err := calc(exprText)
			if err != nil {
				err_txt := fmt.Sprintf(`Error: %s , usage "var1=1,var2=2;expr"`, err)
				t.Execute(w, TemplateData{
					exprText, "", err_txt,
				})
			} else {
				t.Execute(w, TemplateData{
					exprText, fmt.Sprintf("%g", result), "",
				})
			}
		} else {
			t.Execute(w, nil)
		}
	}
}

func calc(command string) (eval.Expr, eval.Env, float64, error) {
	var parseArgs bool
	switch strings.Count(command, ";") {
	case 0:
		parseArgs = false
	case 1:
		parseArgs = true
	default:
		return nil, nil, 0, fmt.Errorf(`too many ";" in expr, allowed only one`)
	}
	exprCommand := command
	env := eval.Env{}
	vars := make(map[eval.Var]bool)
	if parseArgs {
		tokens := strings.Split(command, ";")
		exprCommand = tokens[1]
		argCommand := tokens[0]
		for _, pairs := range strings.Split(argCommand, ",") {
			if strings.Count(pairs, "=") != 1 {
				return nil, nil, 0, fmt.Errorf("invalid var expr: must contain =")
			}
			parts := strings.Split(pairs, "=")
			val, err := strconv.ParseFloat(parts[1], 64)
			if err != nil {
				return nil, nil, 0, fmt.Errorf("error: can't parse float %s", err)
			}
			varName := eval.Var(strings.TrimSpace(parts[0]))
			env[varName] = val
			vars[varName] = true
		}
	}

	expr, err := eval.Parse(exprCommand)
	if err != nil {
		return nil, nil, 0, fmt.Errorf("can't parse with error %s", err)
	}
	err = expr.Check(vars)
	if err != nil {
		return nil, nil, 0, fmt.Errorf("expr check failed with error %s", err)
	}
	r := expr.Eval(env)

	return expr, env, r, nil
}
