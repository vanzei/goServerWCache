// Create a Server cache implementation
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"sync"
)

type User struct {
	ID       int
	Username string
}

type Server struct {
	db    map[int]*User
	dbhit int
	cache map[int]*User
	mu    sync.Mutex
}

func NewServer() *Server {
	db := make(map[int]*User)

	for i := 0; i < 100; i++ {

		db[i+1] = &User{
			ID:       i + 1,
			Username: fmt.Sprintf("user_%d", i+1),
		}
	}
	return &Server{
		db:    db,
		cache: make(map[int]*User),
	}
}

func (s *Server) tryCache(id int) (*User, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	user, ok := s.cache[id]
	return user, ok
}

func (s *Server) handleGetUser(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, _ := strconv.Atoi(idStr)

	user, ok := s.tryCache(id)
	if ok {
		json.NewEncoder(w).Encode(user)
		return
	}
	s.mu.Lock()
	user, ok = s.db[id]
	if !ok {
		s.mu.Unlock()
		panic("User not found with ID")
	}
	s.dbhit++

	// Cache user that was taken from db

	s.cache[id] = user
	s.mu.Unlock()
	json.NewEncoder(w).Encode(user)
}

func main() {

}
