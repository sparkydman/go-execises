package main

import (
	"net/http"

	"gopkg.in/yaml.v3"
)

func MapUrlHandler(data map[string]string, fallback http.Handler)(http.HandlerFunc){
	return func(w http.ResponseWriter, r *http.Request){
		urlPath := r.URL.Path
		if dest, ok := data[urlPath]; ok {
			http.Redirect(w,r, dest, http.StatusFound)
			return
		}
		fallback.ServeHTTP(w, r)
	}
}

func YamlUrlHandler(data []byte, fallback http.Handler)(http.HandlerFunc, error){
	var pathUrls []YamlRequest
	if err := yaml.Unmarshal(data, &pathUrls); err != nil {
		return nil, err
	}

	pathUrlMap := buildMap(pathUrls)
	return MapUrlHandler(pathUrlMap, fallback), nil
}

func buildMap(pathUrls []YamlRequest) map[string]string{
	pathUrlMap := make(map[string]string)
	for _, p := range pathUrls {
		pathUrlMap[p.Path] = p.Url
	}
	return pathUrlMap
}

type YamlRequest struct {
	Path string `yaml:"path"`
	Url  string `yaml:"url"`
}
