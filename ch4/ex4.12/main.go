/*
Exercise 4.12: The popular web comic xkcd has a JSON interface. For example,
a request to https://xkcd.com/571/info.0.json produces a detailed description
of comic 571, one of many favorites. Download each URL (once!) and build an
offline index. Write a tool xkcd that, using this index, prints the URL and
transcript of each comic that matches a search term provided on the command
line.
*/

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"unicode"
)

func printHelp() {
	fmt.Println("Usage: search <term> - find comix by term word")
	fmt.Println("       download - get data from xkcd.com")
	fmt.Println("       build - create/update search index")
}

func main() {

	if len(os.Args) == 1 {
		printHelp()
		os.Exit(1)
	}

	switch os.Args[1] {
	case "search":
		search(os.Args[2:])
	case "download":
		download()
	case "build":
		build()
	default:
		printHelp()
	}
}

type LatestComix struct {
	Num int `json:"num"`
}

type Comix struct {
	Alt        string `json:"alt"`
	Day        string `json:"day"`
	Img        string `json:"img"`
	Link       string `json:"link"`
	Month      string `json:"month"`
	News       string `json:"news"`
	Num        int    `json:"num"`
	SafeTitle  string `json:"safe_title"`
	Title      string `json:"title"`
	Transcript string `json:"transcript"`
	Year       string `json:"year"`
}

func getLatestComix() (int, error) {

	resp, err := http.Get("https://xkcd.com/info.0.json")
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("query failed: %s", resp.Status)
	}

	var result LatestComix
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return 0, err
	}
	return result.Num, nil
}

func getComix(id int) ([]byte, error) {

	resp, err := http.Get(fmt.Sprintf("https://xkcd.com/%d/info.0.json", id))
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

	return data, nil
}

func downloadAndSaveComixById(workDir string, i int) {
	fmt.Printf("\nprocess %d", i)
	filename := filepath.Join(workDir, fmt.Sprintf("%d.json", i))
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		fmt.Print(".")
		data, err := getComix(i)
		fmt.Print(".")
		if err != nil {
			fmt.Printf("Can't get data for id %d with error %s", i, err)
			return
		}
		fo, err := os.Create(filename)
		if err != nil {
			fmt.Printf("Can't create "+
				"file %s for id %d with error %s", filename, i, err)
			return
		}
		fmt.Print(".")
		defer func() {
			err = fo.Close()
			if err != nil {
				fmt.Printf("Can't close file %s for id %d with error %s",
					filename, i, err)
			}
		}()
		_, err = fo.Write(data)
		if err != nil {
			fmt.Printf("Can't write data to "+
				"file %s for id %d with error %s", filename, i, err)
		}
		fmt.Print(".")
	} else {
		fmt.Print(" already downloaded skip.")
	}
}

func getDataDir() string {
	homePath := os.Getenv("HOME")
	return filepath.Join(homePath, ".config", "xkcd")
}

func getIndexFilename() string {
	dd := getDataDir()
	return filepath.Join(dd, "index.json")
}

func SpaceMap(str string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return ' '
		}
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			return r
		}
		return -1
	}, str)
}

func download() {
	workDir := getDataDir()
	err := os.Mkdir(workDir, 0775)
	if os.IsNotExist(err) {
		fmt.Printf("Can't create dir %s with error: %v", workDir, err)
		os.Exit(1)
	}

	latest, err := getLatestComix()

	if err != nil {
		fmt.Printf("Can't get latest comix id : %s", err)
		os.Exit(1)
	}

	fmt.Printf("Latest comix id is %d.\n", latest)

	for i := 1; i <= latest; i++ {
		downloadAndSaveComixById(workDir, i)
	}

	fmt.Print("\n")
}

type index map[string][]int

func appendIndex(data index, text string, num int) {
	text = SpaceMap(text)
	for _, s := range strings.Split(text, " ") {
		s = strings.ToLower(s)
		data[s] = append(data[s], num)
	}
}

func build() {
	data := make(index)
	workDir := getDataDir()
	files, err := os.ReadDir(workDir)
	if err != nil {
		fmt.Printf("Can't read directory %s. Error: %s", workDir, err)
		os.Exit(1)
	}

	for _, file := range files {
		if file.Name() == "index.json" {
			continue
		}
		c := readComix(filepath.Join(workDir, file.Name()))
		appendIndex(data, c.Title, c.Num)
		appendIndex(data, c.Transcript, c.Num)
	}

	marshaledDate, err := json.Marshal(data)
	if err != nil {
		fmt.Printf("Can't marshal index file. Error: %s", err)
		os.Exit(1)
	}

	err = os.WriteFile(getIndexFilename(), marshaledDate, 0644)
	if err != nil {
		fmt.Printf("Can't write index file %s. Error: %s",
			getIndexFilename(), err)
		os.Exit(1)
	}
}

func readComix(filename string) *Comix {
	fileData, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("Can't open file %s. Error: %s", filename, err)
		os.Exit(1)
	}

	var c Comix
	err = json.Unmarshal(fileData, &c)
	if err != nil {
		fmt.Printf("Can't unmarshal file %s. Error: %s", filename, err)
		os.Exit(1)
	}

	return &c
}

func search(tokens []string) {
	indexText, err := os.ReadFile(getIndexFilename())

	if err != nil {
		fmt.Printf("Can't open index file. Error: %s", err)
		os.Exit(1)
	}

	var data index
	err = json.Unmarshal(indexText, &data)

	if err != nil {
		fmt.Printf("Can't unmarshal index file. Error: %s", err)
		os.Exit(1)
	}

	for i, t := range tokens {
		tokens[i] = strings.ToLower(t)
	}

	fmt.Printf("Search word(s): %s\n", strings.Join(tokens, " "))

	result := map[int]map[string]bool{}
	for k, v := range data {
		for _, s := range tokens {
			if strings.Contains(k, s) {
				for _, i := range v {
					if _, exists := result[i]; !exists {
						result[i] = map[string]bool{}
					}
					result[i][s] = true
				}
			}
		}
	}

	var total int
	tokensCount := len(tokens)
	for k, v := range result {
		if len(v) == tokensCount {
			total++
			c := readComix(filepath.Join(getDataDir(), fmt.Sprintf("%d.json", k)))
			fmt.Printf("Id(%d) %s\n %s\n", k, c.Title, c.Transcript)
		}
	}
	fmt.Printf("Found: %d\n", total)
}
