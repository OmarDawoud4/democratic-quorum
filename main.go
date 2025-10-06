package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

// Single Peer
type Node struct {
	id       int
	port     int
	peers    []string
	state    int
	isLeader bool
	leaderID int
	mu       sync.Mutex
}

func newNode(id int, port int, peers []string) *Node {
	return &Node{
		id:       id,
		port:     port,
		peers:    peers,
		state:    0,
		isLeader: false,
		leaderID: -1,
	}
}

// HTTP Stuff
// trigger server => handlers => Registration
func (n *Node) handlePing(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func (n *Node) handleStatus(w http.ResponseWriter, r *http.Request) {
	n.mu.Lock()
	defer n.mu.Unlock()
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(map[string]interface{}{
		"id":       n.id,
		"isLeader": n.isLeader,
		"state":    n.state,
		"leaderID": n.leaderID,
	})
}
func (n *Node) triggerServer() {

	fmt.Printf("Node %d starting on port %d\n", n.id, n.port)
	mux := http.NewServeMux()
	mux.HandleFunc("/ping", n.handlePing)
	mux.HandleFunc("/status", n.handleStatus)

	addr := fmt.Sprintf(":%d", n.port)

	if err := http.ListenAndServe(addr, mux); err != nil {
		fmt.Printf("Node %d failed to start: %v\n", n.id, err)
	}
}

func main() {
	node1 := newNode(1, 8001, []string{"http://localhost:8002", "http://localhost:8003"})
	node2 := newNode(2, 8002, []string{"http://localhost:8001", "http://localhost:8003"})
	node3 := newNode(3, 8003, []string{"http://localhost:8001", "http://localhost:8002"})

	go node1.triggerServer()
	go node2.triggerServer()
	go node3.triggerServer()

	select {}
}
