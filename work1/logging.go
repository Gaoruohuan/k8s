package main

import (
	"log"
	"net/http"
)

type NewResponseWriter struct {
	http.ResponseWriter
	status int
	size   int
}

func (rw *NewResponseWriter) WriteHeader(s int) {
	rw.ResponseWriter.WriteHeader(s)
	rw.status = s
}

func (rw *NewResponseWriter) Write(b []byte) (int, error) {
	if !rw.Written() {
		rw.WriteHeader(http.StatusOK)
	}
	size, err := rw.ResponseWriter.Write(b)
	rw.size += size
	return size, err
}

func (rw *NewResponseWriter) Status() int {
	return rw.status
}

func (rw *NewResponseWriter) Size() int {
	return rw.size
}

func (rw *NewResponseWriter) Written() bool {
	return rw.status != 0
}

func Logging(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		nrw := &NewResponseWriter{
			ResponseWriter: w,
		}
		f(nrw, r)
		log.Printf("host: %s path: %s statusCode: %d size: %d\n", r.RemoteAddr, r.RequestURI, nrw.Status(), nrw.Size())
	}
}
