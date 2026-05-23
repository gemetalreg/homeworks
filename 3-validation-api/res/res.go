package res

import (
	"fmt"
	"net/http"
)

func TextAns(w http.ResponseWriter, data any) {
	w.Header().Set("Content-Type", "Content-Type: text/plain; charset=utf-8")
	fmt.Fprintf(w, "%s", fmt.Sprintf("%v", data))
}
