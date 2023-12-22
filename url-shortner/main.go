package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	mapPathUrl := map[string]string{
		"/gg": "https://google.com",
		"/tb":  "https://www.youtube.com",
	}

	mapHdl := MapUrlHandler(mapPathUrl, defaultHandler())

	f, err := os.Open("urls.yaml")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	b, err := io.ReadAll(f)
	if err != nil {
		panic(err)
	}

	ymlHdl, err := YamlUrlHandler(b, mapHdl)
	if err != nil {
		panic(err)
	}

	http.ListenAndServe(":8999", ymlHdl)
}

func defaultHandler()*http.ServeMux{
	mux := http.NewServeMux()
	mux.HandleFunc("/", defaultEndpoint)
	return mux
}

func defaultEndpoint(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, "This is default endpoint")
}