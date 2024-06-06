package main

import (
	// "bytes"
	// "fmt"
	"net/http"
)

// secureHeaders → servemux → application handler
func secureHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-XSS-Protection", "1; mode=block")
		w.Header().Set("X-Frame-Options", "deny")
		next.ServeHTTP(w, r)
	})
}

// logRequest ↔ secureHeaders ↔ servemux ↔ application handler
func (app *application) logRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		app.infoLog.Printf("%s - %s %s %s", r.RemoteAddr, r.Proto, r.Method, r.URL.String())
		next.ServeHTTP(w, r)
	})
}

// func (app *application) logResponse(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		//write the code to return the response
// 		rw := &responseWriter{
// 			ResponseWriter: w,
// 			body: new(bytes.Buffer),
// 		}

// 		// Call the next handler in the chain
// 		next.ServeHTTP(rw, r)

// 		// Log the response body
// 		fmt.Println(rw.body.String())
// 	})
// }

// type responseWriter struct {
// 	http.ResponseWriter  //inheritance occurs here
// 	body *bytes.Buffer
// }

// func (rw *responseWriter) Write(b []byte) (int, error) {
// 	rw.body.Write(b)
// 	return rw.ResponseWriter.Write(b)
// }
