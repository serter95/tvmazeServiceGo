package tvmazeservicego

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
	"strings"
)

// Result struct ...
type Result struct {
	URL  string `json:"url"`
	Name string `json:"name"`
}

// TvmazeResponse struct ...
type TvmazeResponse struct {
	Score float32 `json:"score"`
	Show  Result  `json:"show"`
}

// StandardResponse struct ...
type StandardResponse struct {
	Category   string `json:"category"`
	Name       string `json:"name"`
	Author     string `json:"author"`
	PreviewURL string `json:"previewUrl"`
	Origin     string `json:"origin"`
}

// FormatText Return ...
func FormatText(text string) string {
	// se quitan los espacios del principio y el final
	// se quitan los dobles espacios y se deja un solo espacio
	// luego de eso se reemplaza el espacio por un '+'
	// todas las letras se colocan en minusculas
	filteredText := strings.ToLower(regexp.MustCompile(`\s\s*`).ReplaceAllString(strings.TrimSpace(text), " "))
	// fmt.Println(filteredText)
	return filteredText
}

// FindResults Return ...
func FindResults(textToFind string) ([]StandardResponse, error) {
	// creo el array que tendra la respuesta unificada
	responseArray := make([]StandardResponse, 0)

	// url de la api
	apiURL := "http://api.tvmaze.com"
	// recurso o ruta
	resource := "/search/shows/"
	// variables de la peticion
	data := url.Values{}
	data.Set("q", FormatText(textToFind))
	// se parsea la url
	u, _ := url.ParseRequestURI(apiURL)
	// se iguala el path al recurso solicitado
	u.Path = resource
	// se codifican las variables en la query
	u.RawQuery = data.Encode()
	// se asigna la url completa a la variable final
	urlStr := u.String() // "https://itunes.apple.com/search/?limit=25&media=all&term=foo"
	// se crea el cliente que hara la peticion (se usa un puntero)
	client := &http.Client{}
	// se crea la request
	r, _ := http.NewRequest(http.MethodGet, urlStr, nil) // URL-encoded payload\
	// se asigna el header
	r.Header.Add("Content-Type", "application/json")

	// se ejecuta el request
	resp, err := client.Do(r)
	// handle error, la peticion fallo
	if err != nil {
		fmt.Println("PETICION FALLIDA")
		fmt.Println(err)
		return responseArray, err
	}

	// despues del flujo cierra la respuesta
	defer resp.Body.Close()
	// leo todo el body
	body, _ := ioutil.ReadAll(resp.Body)
	// creo una variable de tipo AppleResponse
	var tvmazeResponse []TvmazeResponse
	// formateo el body de json a go y se asigna a appleResponse
	json.Unmarshal(body, &tvmazeResponse)

	// recorro los resultados
	for _, iterator := range tvmazeResponse {
		Show := iterator.Show
		responseArray = append(responseArray, StandardResponse{Category: "tv-show", Name: Show.Name, Author: "Not available", PreviewURL: Show.URL, Origin: "tvmaze"})
	}

	// responseArrayJSON, _ := json.MarshalIndent(responseArray, "", "  ")
	// fmt.Println(string(responseArrayJSON))
	return responseArray, nil
}
