/*
Exercise 7.11: Add additional handlers so that clients can create, read,
update, and delete database entries. For example, a request of the
form /update?item=socks&price=6 will update the price of an item in the
inventory and report an error if the item does not exist or if the price is
invalid. (Warning: this change introduces concurrent variable updates.)
*/

package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
)

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
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
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
