/*
Exercise 7.12: Change the handler for /list to print its output as an HTML
table, not text. You may find the html/template package (§4.6) useful.
*/

package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
	"text/template"
)

var priceList = template.Must(template.New("priceList").Parse(`
<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8" />
    <title>title</title>
</head>
<body>
<h1>Price list</h1>
<table>
<tr style='text-align: left'>
  <th>Item</th>
  <th>Price</th>
</tr>
{{range $item, $price := .}}
<tr>
  <td>{{ $item }}</td>
  <td>{{ $price }}</td>
</tr>
{{end}}
</table>
</body>
</html>
`))

//!+main

func main() {
	db := database{"shoes": 50, "socks": 5}
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	http.HandleFunc("/create", db.create)
	http.HandleFunc("/read", db.read)
	http.HandleFunc("/update", db.update)
	http.HandleFunc("/delete", db.delete)
	log.Fatal(http.ListenAndServe("127.0.0.1:8000", nil))
}

//!-main

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type database map[string]dollars

var databaseMu sync.Mutex

const (
	ITEM_NAME_LENGTH = 255
)

func (db database) list(w http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	databaseMu.Lock()
	defer databaseMu.Unlock()
	var b bytes.Buffer
	if err := priceList.Execute(&b, db); err != nil {
		log.Print("err on list template execute", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(b.Bytes())
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	item := req.URL.Query().Get("item")
	databaseMu.Lock()
	defer databaseMu.Unlock()
	if price, ok := db[item]; ok {
		fmt.Fprintf(w, "%s\n", price)
	} else {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
	}
}

func (db *database) create(w http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	item := req.URL.Query().Get("item")
	price := req.URL.Query().Get("price")
	if (len(item) > 0 && len(item) < ITEM_NAME_LENGTH) && len(price) > 0 {
		databaseMu.Lock()
		defer databaseMu.Unlock()
		if _, ok := (*db)[item]; ok {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, "item with that name already in db")
			return
		}
		p, err := strconv.ParseFloat(price, 32)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, "invalid price format")
			return
		}
		(*db)[item] = dollars(p)
		w.WriteHeader(http.StatusCreated)
	} else {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "invalid data")
	}
}

func (db *database) read(w http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	item := req.URL.Query().Get("item")
	if len(item) > 0 && len(item) < ITEM_NAME_LENGTH {
		databaseMu.Lock()
		defer databaseMu.Unlock()
		if price, ok := (*db)[item]; ok {
			w.WriteHeader(http.StatusOK)
			fmt.Fprintf(w, "%s: %s\n", item, price)
		} else {
			w.WriteHeader(http.StatusNotFound) // 404
			fmt.Fprintf(w, "no such item: %q\n", item)
		}
	} else {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "invalid query")
	}
}

func (db *database) update(w http.ResponseWriter, req *http.Request) {
	if req.Method != "PUT" {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	item := req.URL.Query().Get("item")
	price := req.URL.Query().Get("price")
	if (len(item) > 0 && len(item) < ITEM_NAME_LENGTH) && len(price) > 0 {
		databaseMu.Lock()
		defer databaseMu.Unlock()
		if _, ok := (*db)[item]; ok {
			p, err := strconv.ParseFloat(price, 32)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				fmt.Fprint(w, "invalid price format")
				return
			}
			(*db)[item] = dollars(p)
			w.WriteHeader(http.StatusOK)
		} else {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, "item with that name already in db")
		}
	} else {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "invalid data")
	}
}

func (db *database) delete(w http.ResponseWriter, req *http.Request) {
	if req.Method != "DELETE" {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	item := req.URL.Query().Get("item")
	if len(item) > 0 && len(item) < ITEM_NAME_LENGTH {
		databaseMu.Lock()
		defer databaseMu.Unlock()
		if _, ok := (*db)[item]; ok {
			w.WriteHeader(http.StatusOK)
			delete(*db, item)
		}
	} else {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "invalid data")
	}
}
