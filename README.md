# Go net/http package

- Form go version 1.22
- YouTube [video](https://www.youtube.com/watch?v=H7tbjKFSg58&t=405s)

## Path parameters

```go
package main

import (
	"log"
	"net/http"
)

func main() {
	router := http.NewServeMux()
	router.HandleFunc("/item/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		w.Write([]byte("Item ID: " + id))
	})

	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	log.Println("Starting server on :8080")
	server.ListenAndServe()
}
```

## Method based router

