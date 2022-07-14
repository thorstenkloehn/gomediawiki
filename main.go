package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"text/template"
)

var vorlagen, _ = template.ParseGlob("view/*")

type jsonausgabe struct {
	Parse struct {
		Title  string `json:"title"`
		Pageid int    `json:"pageid"`
		Revid  int    `json:"revid"`
		Text   struct {
			Field1 string `json:"*"`
		} `json:"text"`
		Langlinks  []interface{} `json:"langlinks"`
		Categories []struct {
			Sortkey string `json:"sortkey"`
			Field2  string `json:"*"`
		} `json:"categories"`
		Links []struct {
			Ns     int    `json:"ns"`
			Exists string `json:"exists"`
			Field3 string `json:"*"`
		} `json:"links"`
		Templates []struct {
			Ns     int    `json:"ns"`
			Exists string `json:"exists"`
			Field3 string `json:"*"`
		} `json:"templates"`
		Images        []interface{} `json:"images"`
		Externallinks []string      `json:"externallinks"`
		Sections      []struct {
			Toclevel   int    `json:"toclevel"`
			Level      string `json:"level"`
			Line       string `json:"line"`
			Number     string `json:"number"`
			Index      string `json:"index"`
			Fromtitle  string `json:"fromtitle"`
			Byteoffset *int   `json:"byteoffset"`
			Anchor     string `json:"anchor"`
		} `json:"sections"`
		Parsewarnings []interface{} `json:"parsewarnings"`
		Displaytitle  string        `json:"displaytitle"`
		Iwlinks       []interface{} `json:"iwlinks"`
		Properties    []struct {
			Name   string `json:"name"`
			Field2 string `json:"*"`
		} `json:"properties"`
	} `json:"parse"`
}

func main() {
	server := http.NewServeMux()
	server.HandleFunc("/", testausgabe)
	http.ListenAndServe(":8080", server)
}
func testausgabe(w http.ResponseWriter, r *http.Request) {
	var start jsonausgabe
	resp, err := http.Get("https://ahrensburg.wiki/api.php?action=parse&format=json&page=Hauptseite")
	if err != nil {
		panic(err)
	}
	body1, _ := ioutil.ReadAll(resp.Body)
	error := json.Unmarshal((body1), &start)
	if error != nil {
		fmt.Printf("Fehört")
	}

	vorlagen.ExecuteTemplate(w, "index.html", start)

}
