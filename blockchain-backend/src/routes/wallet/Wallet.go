package wallet

import (
	"fmt"
	"net/http"
)

// GetHandler handles the get method on the app.go
func GetHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	fmt.Fprintf(w, "hello !!!!")
}
