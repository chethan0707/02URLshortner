package urlshortner

import (
	"fmt"
	"net/http"

	"gopkg.in/yaml.v2"
)

func MapHandler(pathToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// if we can match path
		//we have to redirect to path
		//else
		//fallback
		path := r.URL.Path
		fmt.Println(pathToUrls[path])
		if dest, ok := pathToUrls[path]; ok {

			http.Redirect(w, r, dest, http.StatusFound)
		}
		fallback.ServeHTTP(w, r)
	}
}

func YAMLHandler(yamlBytes []byte, fallback http.Handler) (http.HandlerFunc, error) {
	//1. parse the YAML somehow
	pathUrls := parseYAML(yamlBytes)
	//2. Convert YAML into map array
	pathToUrls := intoMapArray(pathUrls)
	// 3. return a map handler using the map
	return MapHandler(pathToUrls, fallback), nil
}

func parseYAML(data []byte) []pathUrl {
	var pathUrls []pathUrl
	err := yaml.Unmarshal(data, &pathUrls)
	if err != nil {
		fmt.Println("error is ", err)
	}
	return pathUrls
}
func intoMapArray(paths []pathUrl) map[string]string {
	pathToUrls := make(map[string]string)
	for _, pu := range paths {
		pathToUrls[pu.Path] = pu.URL
	}
	return pathToUrls
}

type pathUrl struct {
	Path string `yaml:"path"`
	URL  string `yaml:"url"`
}

// func JsonHandler(jsonD map[string]string) (http.Handler, error) {

// 	// byte, err := parseJsonData(jsonD)
// 	// if err != nil {
// 	// 	fmt.Println("Error")
// 	// }

// 	// pathsToUrl := intoMapArray()
// 	// var handler http.Handler
// 	// return MapHandler(pathsToUrl, handler), nil
// }

// func parseJsonData(jsonD map[string]string) ([]byte, error) {
// 	byte, err := json.Marshal(jsonD)
// 	if err != nil {
// 		return nil, err
// 	}
// 	var jsonData []jsonPaths
// 	er := json.Unmarshal(byte, &jsonData)
// 	if er != nil {
// 		fmt.Println(er)
// 	}
// 	return byte, nil
// }

// type jsonPaths struct {
// 	Path string `json:"path"`
// 	URL  string `json:"url"`
// }
