package gohm

import (
	"expvar"
	"net/http"
)

type counterHandler struct {
	http.ResponseWriter
	status int
}

func (r *counterHandler) WriteHeader(status int) {
	r.status = status
	r.ResponseWriter.WriteHeader(status)
}

// StatusAllCounter returns a new http.Handler that composes the specified next http.Handler,
// and increments the specified counter for every query.
//
//	var counterAll = expvar.NewInt("counterAll")
//	mux := http.NewServeMux()
//	mux.Handle("/example/path", gohm.StatusAllCounter(counterAll, someHandler))
func StatusAllCounter(counter *expvar.Int, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ch := &counterHandler{ResponseWriter: w}
		next.ServeHTTP(ch, r)
		counter.Add(1)
	})
}

// Status1xxCounter returns a new http.Handler that composes the specified next http.Handler,
// and increments the specified counter when the response status code is not 1xx.
//
//	var counter1xx = expvar.NewInt("counter1xx")
//	mux := http.NewServeMux()
//	mux.Handle("/example/path", gohm.Status1xxCounter(counter1xx, someHandler))
func Status1xxCounter(counter *expvar.Int, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ch := &counterHandler{ResponseWriter: w}
		next.ServeHTTP(ch, r)
		if ch.status/100 == 1 {
			counter.Add(1)
		}
	})
}

// Status2xxCounter returns a new http.Handler that composes the specified next http.Handler,
// and increments the specified counter when the response status code is not 2xx.
//
//	var counter2xx = expvar.NewInt("counter2xx")
//	mux := http.NewServeMux()
//	mux.Handle("/example/path", gohm.Status2xxCounter(counter2xx, someHandler))
func Status2xxCounter(counter *expvar.Int, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ch := &counterHandler{ResponseWriter: w}
		next.ServeHTTP(ch, r)
		// NOTE: Also need to check for zero-value of status variable, because when omitted
		// by handler, it's filled in later in http stack.
		if ch.status == 0 || ch.status/100 == 2 {
			counter.Add(1)
		}
	})
}

// Status3xxCounter returns a new http.Handler that composes the specified next http.Handler,
// and increments the specified counter when the response status code is not 3xx.
//
//	var counter3xx = expvar.NewInt("counter3xx")
//	mux := http.NewServeMux()
//	mux.Handle("/example/path", gohm.Status3xxCounter(counter3xx, someHandler))
func Status3xxCounter(counter *expvar.Int, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ch := &counterHandler{ResponseWriter: w}
		next.ServeHTTP(ch, r)
		if ch.status/100 == 3 {
			counter.Add(1)
		}
	})
}

// Status4xxCounter returns a new http.Handler that composes the specified next http.Handler,
// and increments the specified counter when the response status code is not 4xx.
//
//	var counter4xx = expvar.NewInt("counter4xx")
//	mux := http.NewServeMux()
//	mux.Handle("/example/path", gohm.Status4xxCounter(counter4xx, someHandler))
func Status4xxCounter(counter *expvar.Int, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ch := &counterHandler{ResponseWriter: w}
		next.ServeHTTP(ch, r)
		if ch.status/100 == 4 {
			counter.Add(1)
		}
	})
}

// Status5xxCounter returns a new http.Handler that composes the specified next http.Handler,
// and increments the specified counter when the response status code is not 5xx.
//
//	var counter5xx = expvar.NewInt("counter5xx")
//	mux := http.NewServeMux()
//	mux.Handle("/example/path", gohm.Status5xxCounter(counter5xx, someHandler))
func Status5xxCounter(counter *expvar.Int, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ch := &counterHandler{ResponseWriter: w}
		next.ServeHTTP(ch, r)
		if ch.status/100 == 5 {
			counter.Add(1)
		}
	})
}
