package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	// t.Run("with fast servers", func(t *testing.T) {
	// 	slowServer := makeDelayServer(20 * time.Millisecond)
	// 	fastServer := makeDelayServer(0 * time.Millisecond)

	// 	slowUrl := slowServer.URL
	// 	fastUrl := fastServer.URL

	// 	defer slowServer.Close()
	// 	defer fastServer.Close()

	// 	want := fastUrl
	// 	got, _ := Racer(slowUrl, fastUrl)

	// 	if got != want {
	// 		t.Errorf("got %q, want %q", got, want)
	// 	}
	// })

	t.Run("with slow servers", func(t *testing.T) {
		server := makeDelayServer(9 * time.Second)

		defer server.Close()

		_, err := ConfigurableRacer(10*time.Millisecond, server.URL, server.URL)

		// if err != fmt.Errorf("yo") {
		// 	t.Errorf("not yo obviously, got %d", err)
		// }

		if err == nil {
			t.Error("expected a timeout error, but didn't get one")
		}
	})

}

func makeDelayServer(delay time.Duration) *httptest.Server {
	slowServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
	return slowServer
}
