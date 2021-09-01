package main

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type SpyStore struct {
	response string
	t        *testing.T
}

func (s *SpyStore) assertWasCancelled() {
	// s.t.Helper()
	// if !s.cancelled {
	// 	s.t.Error("Store was not told to cancel")
	// }
}

func (s *SpyStore) assertWasNotCancelled() {
	// s.t.Helper()
	// if s.cancelled {
	// 	s.t.Error("store was told to cancel")
	// }
}

func (s *SpyStore) Fetch(ctx context.Context) (string, error) {
	// Make 1 string channel
	data := make(chan string, 1)

	// Set up a goroutine
	go func() {
		// Set up result string
		var result string

		// Loop thru all the letters in the response (see default case)
		// We could just do the select without the loop, but we'd miss out on the slow string
		for _, c := range s.response {
			// Create a select to race the two conditions
			select {
			// if ctx.Done gets called in this channel, it will stop work in it
			case <-ctx.Done():
				s.t.Log("spy store got cancelled")
				return
				// write the response one (c char) at a time, with 10ms delay
			default:
				time.Sleep(10 * time.Millisecond)
				result += string(c)
			}
		}
		// Put the result into the data chanel
		data <- result
	}()

	// Create another race
	select {
	// If it gets cancelled first return err
	case <-ctx.Done():
		return "", ctx.Err()
		// Otherwise set the data from the channel to var res and return it
	case res := <-data:
		return res, nil
	}
}

func (s *SpyStore) Cancel() {
	// s.cancelled = true
}

type SpyResposeWriter struct {
	written bool
}

func (s *SpyResposeWriter) Header() http.Header {
	s.written = true
	return nil
}
func (s *SpyResposeWriter) Write([]byte) (int, error) {
	s.written = true
	return 0, errors.New("not implemented")
}
func (s *SpyResposeWriter) WriteHeader(statusCode int) {
	s.written = true
}
func TestServer(t *testing.T) {
	t.Run("returns data from store", func(t *testing.T) {
		data := "hello, world"
		store := &SpyStore{response: data, t: t}
		svr := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		svr.ServeHTTP(response, request)

		if response.Body.String() != data {
			t.Errorf(`got "%s" want "%s"`, response.Body.String(), data)
		}
	})

	t.Run("Tells store to cancel work if request is cancelled", func(t *testing.T) {
		data := "hello, world"
		store := &SpyStore{response: data, t: t}
		svr := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)

		cancellingCtx, cancel := context.WithCancel(request.Context())
		time.AfterFunc(5*time.Millisecond, cancel)
		request = request.WithContext(cancellingCtx)

		response := &SpyResposeWriter{}

		svr.ServeHTTP(response, request)

		if response.written {
			t.Error("a response should not have been written")
		}
	})
}
