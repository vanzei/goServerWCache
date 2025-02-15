package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandleGetUser(t *testing.T) {
	s := NewServer()
	ts := httptest.NewServer(http.HandlerFunc(s.handleGetUser))
	nreq := 1000

	for i := 0; i < nreq; i++ {
		id := i%100 + 1
		url := fmt.Sprintf("%s/?id=%d", ts.URL, id)
		resp, err := http.Get(url)
		if err != nil {
			t.Error(err)
		}
		user := &User{}
		if err := json.NewDecoder(resp.Body).Decode(user); err != nil {
			t.Error(err)
		}
		fmt.Printf("%+v\n", user)
	}

	fmt.Printf("%d DB Hits\n", s.dbhit)
}
