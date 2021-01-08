# tvmazeServiceGo
service to consulting tvmaze API

# Install
```
go get github.com/serter95/tvmazeServiceGo
```
# Use

```go
package main

import "github.com/serter95/tvmazeServiceGo"

func main() {
    sliceWithData, err := tvmazeServiceGo.FindResults("your criteria")
    // do wath you want
}

// the slice data will content this struct:

type StandardResponse struct {
	Category   string `json:"category"`
	Name       string `json:"name"`
	Author     string `json:"author"`
	PreviewURL string `json:"previewUrl"`
	Origin     string `json:"origin"`
}
```
