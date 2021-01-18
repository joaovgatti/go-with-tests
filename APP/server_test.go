package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

type StubPlayerStore struct {
	scores map[string]int
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	score := s.scores[name]
	return score
}

func TestGETPlayers(t *testing.T){

	store := StubPlayerStore{
		map[string]int{
			"Pele":90,
			"Romario":60,
		},
	}

	server := &PlayerServer{&store}

	t.Run("return Pele's score", func (t *testing.T){
		request := newGetScoreRequest("Pele")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertResponseBody(t, response.Body.String(), "90")
	})

	t.Run("returns Romario score", func (t *testing.T){
		request := newGetScoreRequest("Romario")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertResponseBody(t, response.Body.String(), "60")

	})
}

func newGetScoreRequest(name string) *http.Request{
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s",name), nil)
	return req
}

func assertResponseBody(t testing.TB, got, want string){
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}