/*
Exercise 4.13: The JSON-based web service of the Open Movie Database lets you
search https://omdbapi.com/ for a movie by name and download its poster image.
Write a tool poster that downloads the poster image for the movie named on the
command line.
*/

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

type Movie struct {
	Title    string `json:"Title"`
	Year     string `json:"Year"`
	Rated    string `json:"Rated"`
	Released string `json:"Released"`
	Runtime  string `json:"Runtime"`
	Genre    string `json:"Genre"`
	Director string `json:"Director"`
	Writer   string `json:"Writer"`
	Actors   string `json:"Actors"`
	Plot     string `json:"Plot"`
	Language string `json:"Language"`
	Country  string `json:"Country"`
	Awards   string `json:"Awards"`
	Poster   string `json:"Poster"`
	Ratings  []struct {
		Source string `json:"Source"`
		Value  string `json:"Value"`
	} `json:"Ratings"`
	Metascore  string `json:"Metascore"`
	ImdbRating string `json:"imdbRating"`
	ImdbVotes  string `json:"imdbVotes"`
	ImdbID     string `json:"imdbID"`
	Type       string `json:"Type"`
	Dvd        string `json:"DVD"`
	BoxOffice  string `json:"BoxOffice"`
	Production string `json:"Production"`
	Website    string `json:"Website"`
	Response   string `json:"Response"`
}

func searchOpenMovieDb(search []string, apiKey string) (*Movie, error) {

	t := strings.Join(search, "+")
	url := fmt.Sprintf("http://www.omdbapi.com/?t=%s&apikey=%s", t, apiKey)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("query failed: %s", resp.Status)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if len(data) == 0 {
		return nil, fmt.Errorf("zero len buffer")
	}

	var movie Movie

	err = json.Unmarshal(data, &movie)

	if err != nil {
		return nil, err
	}

	return &movie, nil
}

func downloadFile(url, filename string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("query failed: %s", resp.Status)
	}
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	fo, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer func() {
		err = fo.Close()
		if err != nil {
			fmt.Printf("Can't close file %s with error %s",
				filename, err)
		}
	}()
	_, err = fo.Write(data)
	if err != nil {
		return fmt.Errorf("can't write data to "+
			"file %s with error %s", filename, err)
	}
	return nil
}

func printHelp(message string) {
	if len(message) > 0 {
		fmt.Println(message)
	}
	fmt.Printf("Usage of %s <movie name>\n", os.Args[0])
	flag.PrintDefaults()
	os.Exit(1)
}

func main() {
	apiKey := flag.String("api-key", os.Getenv("OMDBAPI_KEY"),
		"Api key from omdbapi.com")

	filename := flag.String("filename", "",
		"New filename for poster")

	flag.Parse()

	if len(*apiKey) == 0 {
		printHelp("Please specify the Api key as " +
			"command line or env var OMDBAPI_KEY")
	}

	if len(*filename) == 0 {
		printHelp("Please specify poster filename")
	}

	if flag.NArg() == 0 {
		printHelp("Please specify movie name")
	}

	movie, err := searchOpenMovieDb(flag.Args(), *apiKey)

	if err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}

	err = downloadFile(movie.Poster, *filename)

	if err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}
}
