/*
Exercise 7.9: Use the html/template package (§4.6) to replace printTracks with
a function that displays the tracks as an HTML table. Use the solution to the
previous exercise to arrange that each click on a column head makes an HTTP
request to sort the table.
*/

package main

import (
	"html/template"
	"log"
	"net/http"
	"sort"
	"time"
)

// !+main
type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

type TrackTableSort struct {
	t   []*Track
	comp []func(i, j *Track) int
}

func (tts TrackTableSort) Len() int      { return len(tts.t) }
func (tts TrackTableSort) Swap(i, j int) { tts.t[i], tts.t[j] = tts.t[j], tts.t[i] }
func (tts TrackTableSort) Less(i, j int) bool {
	if len(tts.comp) > 0 {
		for _, comp := range tts.comp {
			r := comp(tts.t[i], tts.t[j])
			if r == 0 {
				continue
			}
			return r < 0
		}
	}
	return tts.t[i].Title < tts.t[j].Title
}

func (tts *TrackTableSort) Add(f func(i, j *Track) int) {
	tts.comp = append([]func(i, j *Track) int{f}, tts.comp...)
	if len(tts.comp) > 3 {
		tts.comp = tts.comp[:3]
	}
}

var tracks = []*Track{
	{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
	{"Go", "Moby", "Moby", 1992, length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
}

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

var trackList = template.Must(template.New("trackList").Parse(`
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Document</title>
</head>
<body>
<h1>Track List</h1>
<table>
<tr style='text-align: left'>
  <th><a href="?sort=title">Title</a></th>
  <th><a href="?sort=artist">Artist</a></th>
  <th><a href="?sort=album">Album</a></th>
  <th><a href="?sort=year">Year</a></th>
  <th><a href="?sort=length">Length</a></th>
</tr>
{{range .}}
<tr>
  <td>{{.Title}}</td>
  <td>{{.Artist}}</td>
  <td>{{.Album}}</td>
  <td>{{.Year}}</td>
  <td>{{.Length}}</td>
</tr>
{{end}}
</table>
</body>
</html>
`))

func TrackTableSortByArtist(i, j *Track) int {
	if i.Artist == j.Artist {
		return 0
	} else if i.Artist < j.Artist {
		return -1
	}
	return 1
}

func TrackTableSortByAlbum(i, j *Track) int {
	if i.Album == j.Album {
		return 0
	} else if i.Album < j.Album {
		return -1
	}
	return 1
}
func TrackTableSortByYear(i, j *Track) int {
	if i.Year == j.Year {
		return 0
	} else if i.Year < j.Year {
		return -1
	}
	return 1
}
func TrackTableSortByTitle(i, j *Track) int {
	if i.Title == j.Title {
		return 0
	} else if i.Title < j.Title {
		return -1
	}
	return 1
}
func TrackTableSortByLength(i, j *Track) int {
	if i.Length == j.Length {
		return 0
	} else if i.Length < j.Length {
		return -1
	}
	return 1
}

func main() {

	tts := TrackTableSort{t: tracks}
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {

		s := req.URL.Query().Get("sort")
		if s != "" {
			switch s {
			case "artist":
				tts.Add(TrackTableSortByArtist)
			case "album":
				tts.Add(TrackTableSortByAlbum)
			case "year":
				tts.Add(TrackTableSortByYear)
			case "title":
				tts.Add(TrackTableSortByTitle)
			case "length":
				tts.Add(TrackTableSortByLength)
			}
			sort.Sort(tts)
		}
		err := trackList.Execute(w, tracks)
		if err != nil {
			log.Printf("error on rendering template %s", err)
		}
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
