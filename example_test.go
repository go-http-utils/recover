package recover_test

import (
	"net/http"

	"github.com/go-http-utils/recover"
)

func Example() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		res.Write([]byte("Hello World"))
	})

	http.ListenAndServe(":8080", recover.Handler(mux, recover.DefaultRecoverHandler))
}
