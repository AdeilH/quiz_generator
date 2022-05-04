package main

import (
	"net/http"
)

func main() {
	handle_flags()
	initialize_vars()
	register_handles()

	http.ListenAndServe(":8000", nil)
}
