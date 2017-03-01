package gohm_test

import (
	"expvar"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/karrick/gohm"
)

func TestStatus1xxCounterHandlerGood(t *testing.T) {
	req := httptest.NewRequest("GET", "/some/url", nil)

	counter := expvar.NewInt("counter-status-continue")
	response := "gimme more!!!"
	status := http.StatusContinue

	handler := gohm.Status1xxCounterHandler(counter, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(status)
		w.Write([]byte(response))
	}))

	rr := httptest.NewRecorder()

	handler.ServeHTTP(rr, req)

	if actual, expected := counter.Value(), int64(0); actual != expected {
		t.Fatalf("Actual: %#v; Expected: %#v", actual, expected)
	}

	if actual, expected := rr.Body.String(), response; actual != expected {
		t.Errorf("Actual: %#v; Expected: %#v", actual, expected)
	}

	if actual, expected := rr.Code, status; actual != expected {
		t.Errorf("Actual: %#v; Expected: %#v", actual, expected)
	}
}

func TestStatus1xxCounterHandlerBad(t *testing.T) {
	req := httptest.NewRequest("GET", "/some/url", nil)

	counter := expvar.NewInt("counter-status-created")
	response := "record created"
	status := http.StatusCreated

	handler := gohm.Status1xxCounterHandler(counter, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(status)
		w.Write([]byte(response))
	}))

	rr := httptest.NewRecorder()

	handler.ServeHTTP(rr, req)

	if actual, expected := counter.Value(), int64(1); actual != expected {
		t.Fatalf("Actual: %#v; Expected: %#v", actual, expected)
	}

	if actual, expected := rr.Body.String(), response; actual != expected {
		t.Errorf("Actual: %#v; Expected: %#v", actual, expected)
	}

	if actual, expected := rr.Code, status; actual != expected {
		t.Errorf("Actual: %#v; Expected: %#v", actual, expected)
	}
}

func TestStatus2xxCounterHandlerWriteHeaderElided(t *testing.T) {
	req := httptest.NewRequest("GET", "/some/url", nil)

	counter := expvar.NewInt("counter-status-elided")
	response := "elided"
	status := http.StatusOK

	handler := gohm.Status2xxCounterHandler(counter, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(response))
	}))

	rr := httptest.NewRecorder()

	handler.ServeHTTP(rr, req)

	if actual, expected := counter.Value(), int64(0); actual != expected {
		t.Fatalf("Actual: %#v; Expected: %#v", actual, expected)
	}

	if actual, expected := rr.Body.String(), response; actual != expected {
		t.Errorf("Actual: %#v; Expected: %#v", actual, expected)
	}

	if actual, expected := rr.Code, status; actual != expected {
		t.Errorf("Actual: %#v; Expected: %#v", actual, expected)
	}
}

func TestStatus2xxCounterHandlerGood(t *testing.T) {
	req := httptest.NewRequest("GET", "/some/url", nil)

	counter := expvar.NewInt("counter-status-accepted")
	response := "gimme more!!!"
	status := http.StatusAccepted

	handler := gohm.Status2xxCounterHandler(counter, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(status)
		w.Write([]byte(response))
	}))

	rr := httptest.NewRecorder()

	handler.ServeHTTP(rr, req)

	if actual, expected := counter.Value(), int64(0); actual != expected {
		t.Fatalf("Actual: %#v; Expected: %#v", actual, expected)
	}

	if actual, expected := rr.Body.String(), response; actual != expected {
		t.Errorf("Actual: %#v; Expected: %#v", actual, expected)
	}

	if actual, expected := rr.Code, status; actual != expected {
		t.Errorf("Actual: %#v; Expected: %#v", actual, expected)
	}
}

func TestStatus2xxCounterHandlerBad(t *testing.T) {
	req := httptest.NewRequest("GET", "/some/url", nil)

	counter := expvar.NewInt("counter-status-processing")
	response := "record created"
	status := http.StatusProcessing

	handler := gohm.Status2xxCounterHandler(counter, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(status)
		w.Write([]byte(response))
	}))

	rr := httptest.NewRecorder()

	handler.ServeHTTP(rr, req)

	if actual, expected := counter.Value(), int64(1); actual != expected {
		t.Fatalf("Actual: %#v; Expected: %#v", actual, expected)
	}

	if actual, expected := rr.Body.String(), response; actual != expected {
		t.Errorf("Actual: %#v; Expected: %#v", actual, expected)
	}

	if actual, expected := rr.Code, status; actual != expected {
		t.Errorf("Actual: %#v; Expected: %#v", actual, expected)
	}
}

func TestStatus3xxCounterHandlerGood(t *testing.T) {
	req := httptest.NewRequest("GET", "/some/url", nil)

	counter := expvar.NewInt("counter-status-found")
	response := "I found it"
	status := http.StatusFound

	handler := gohm.Status3xxCounterHandler(counter, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(status)
		w.Write([]byte(response))
	}))

	rr := httptest.NewRecorder()

	handler.ServeHTTP(rr, req)

	if actual, expected := counter.Value(), int64(0); actual != expected {
		t.Fatalf("Actual: %#v; Expected: %#v", actual, expected)
	}

	if actual, expected := rr.Body.String(), response; actual != expected {
		t.Errorf("Actual: %#v; Expected: %#v", actual, expected)
	}

	if actual, expected := rr.Code, status; actual != expected {
		t.Errorf("Actual: %#v; Expected: %#v", actual, expected)
	}
}

func TestStatus3xxCounterHandlerBad(t *testing.T) {
	req := httptest.NewRequest("GET", "/some/url", nil)

	counter := expvar.NewInt("counter-status-already-reported")
	response := "record created"
	status := http.StatusAlreadyReported

	handler := gohm.Status3xxCounterHandler(counter, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(status)
		w.Write([]byte(response))
	}))

	rr := httptest.NewRecorder()

	handler.ServeHTTP(rr, req)

	if actual, expected := counter.Value(), int64(1); actual != expected {
		t.Fatalf("Actual: %#v; Expected: %#v", actual, expected)
	}

	if actual, expected := rr.Body.String(), response; actual != expected {
		t.Errorf("Actual: %#v; Expected: %#v", actual, expected)
	}

	if actual, expected := rr.Code, status; actual != expected {
		t.Errorf("Actual: %#v; Expected: %#v", actual, expected)
	}
}

func TestStatus4xxCounterHandlerGood(t *testing.T) {
	req := httptest.NewRequest("GET", "/some/url", nil)

	counter := expvar.NewInt("counter-status-bad-request")
	response := "gimme more!!!"
	status := http.StatusBadRequest

	handler := gohm.Status4xxCounterHandler(counter, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(status)
		w.Write([]byte(response))
	}))

	rr := httptest.NewRecorder()

	handler.ServeHTTP(rr, req)

	if actual, expected := counter.Value(), int64(0); actual != expected {
		t.Fatalf("Actual: %#v; Expected: %#v", actual, expected)
	}

	if actual, expected := rr.Body.String(), response; actual != expected {
		t.Errorf("Actual: %#v; Expected: %#v", actual, expected)
	}

	if actual, expected := rr.Code, status; actual != expected {
		t.Errorf("Actual: %#v; Expected: %#v", actual, expected)
	}
}

func TestStatus4xxCounterHandlerBad(t *testing.T) {
	req := httptest.NewRequest("GET", "/some/url", nil)

	counter := expvar.NewInt("counter-status-see-other")
	response := "record created"
	status := http.StatusSeeOther

	handler := gohm.Status4xxCounterHandler(counter, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(status)
		w.Write([]byte(response))
	}))

	rr := httptest.NewRecorder()

	handler.ServeHTTP(rr, req)

	if actual, expected := counter.Value(), int64(1); actual != expected {
		t.Fatalf("Actual: %#v; Expected: %#v", actual, expected)
	}

	if actual, expected := rr.Body.String(), response; actual != expected {
		t.Errorf("Actual: %#v; Expected: %#v", actual, expected)
	}

	if actual, expected := rr.Code, status; actual != expected {
		t.Errorf("Actual: %#v; Expected: %#v", actual, expected)
	}
}

func TestStatus5xxCounterHandlerGood(t *testing.T) {
	req := httptest.NewRequest("GET", "/some/url", nil)

	counter := expvar.NewInt("counter-status-internal-server-error")
	response := "gimme more!!!"
	status := http.StatusInternalServerError

	handler := gohm.Status5xxCounterHandler(counter, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(status)
		w.Write([]byte(response))
	}))

	rr := httptest.NewRecorder()

	handler.ServeHTTP(rr, req)

	if actual, expected := counter.Value(), int64(0); actual != expected {
		t.Fatalf("Actual: %#v; Expected: %#v", actual, expected)
	}

	if actual, expected := rr.Body.String(), response; actual != expected {
		t.Errorf("Actual: %#v; Expected: %#v", actual, expected)
	}

	if actual, expected := rr.Code, status; actual != expected {
		t.Errorf("Actual: %#v; Expected: %#v", actual, expected)
	}
}

func TestStatus5xxCounterHandlerBad(t *testing.T) {
	req := httptest.NewRequest("GET", "/some/url", nil)

	counter := expvar.NewInt("counter-status-unauthorized")
	response := "record created"
	status := http.StatusUnauthorized

	handler := gohm.Status5xxCounterHandler(counter, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(status)
		w.Write([]byte(response))
	}))

	rr := httptest.NewRecorder()

	handler.ServeHTTP(rr, req)

	if actual, expected := counter.Value(), int64(1); actual != expected {
		t.Fatalf("Actual: %#v; Expected: %#v", actual, expected)
	}

	if actual, expected := rr.Body.String(), response; actual != expected {
		t.Errorf("Actual: %#v; Expected: %#v", actual, expected)
	}

	if actual, expected := rr.Code, status; actual != expected {
		t.Errorf("Actual: %#v; Expected: %#v", actual, expected)
	}
}
