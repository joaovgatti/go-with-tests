package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

type StubPlayerStore struct {
	scores map[string]int
	winCalls []string
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	score := s.scores[name]
	return score
}

func (s *StubPlayerStore) RecordWin(name string){
	s.winCalls = append(s.winCalls, name)
}

func TestGETPlayers(t *testing.T){

	store := StubPlayerStore{
		map[string]int{
			"Pele":90,
			"Romario":60,
		},nil,
	}

	server := &PlayerServer{&store}

	t.Run("return Pele's score", func (t *testing.T){
		request := newGetScoreRequest("Pele")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertResponseBody(t, response.Body.String(), "90")
		assertStatus(t, response.Code, http.StatusOK)
	})

	t.Run("returns Romario score", func (t *testing.T){
		request := newGetScoreRequest("Romario")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertResponseBody(t, response.Body.String(), "60")
		assertStatus(t, response.Code, http.StatusOK)

	})

	t.Run("returns 404 on missing players", func (t *testing.T){
		request := newGetScoreRequest("Bebeto")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		got := response.Code
		want := http.StatusNotFound

		if got != want {
			t.Errorf("got status %d wand %d", got, want)
		}

		assertStatus(t, response.Code, http.StatusNotFound)
	})
}

func TestStoreWins(t *testing.T){
	store := StubPlayerStore{
		map[string]int{},nil,
	}
	server := &PlayerServer{&store}

	t.Run("it returns accepted on POST", func (t *testing.T){
		request, _ := http.NewRequest(http.MethodPost, "/players/Pele/", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusAccepted)
	})

	t.Run("it records wins when POST", func (t *testing.T){

		player := "Pele"

		request := newPostWinRequest("Pele")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusAccepted)

		if len(store.winCalls) != 1 {
			t.Errorf("got %d calls to RecordWin, want %d", len(store.winCalls),1)
		}

		if store.winCalls[0] != player {
			t.Errorf("did not get the correct winner, got %q want %q", store.winCalls[0], player)
		}
	})
}


func newPostWinRequest(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/players/%s", name), nil)
		return req
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

func assertStatus(t testing.TB, got, want int ){
	t.Helper()
	if got != want {
		t.Errorf("did not get the correct status, got %d wand %d", got , want)
	}
}