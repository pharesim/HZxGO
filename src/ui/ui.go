package ui

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func loadPage(file string) ([]byte, error) {
	if file == "" {
		file = "index.html"
	}

	filename := "ui/"+file
	body, err := ioutil.ReadFile(filename)

	if err != nil {
		return nil, err
	}

	return body, nil
}

func uiHandler(w http.ResponseWriter, r *http.Request) {
	p, _ := loadPage(r.URL.Path[1:])
	fmt.Fprintf(w, "%s", p)
}

func ServeUI(listen string, port int64) {
	serverMuxUI := http.NewServeMux()
	serverMuxUI.HandleFunc("/", uiHandler)
	http.ListenAndServe(fmt.Sprintf("%s:%d",listen,port), serverMuxUI)
}