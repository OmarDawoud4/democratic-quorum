package main

import (
	"fmt"
	"net/http"
	"sync"
)

type Peer struct {
	id  int
	url string
}

type Node struct {
	id       int
	port     int
	peers    []Peer // updated for election comparison
	state    int
	isLeader bool
	leaderID int
	mu       sync.Mutex
}

func newNode(id int, port int, peers []Peer) *Node {
	return &Node{
		id:       id,
		port:     port,
		peers:    peers,
		state:    0,
		isLeader: false,
		leaderID: -1,
	}
}

// register all endpoints
func (n *Node) triggerServer() {
	fmt.Printf("Node %d starting on port %d\n", n.id, n.port)
	mux := http.NewServeMux()
	mux.HandleFunc("/ping", n.handlePing)
	mux.HandleFunc("/status", n.handleStatus)
	mux.HandleFunc("/election", n.handleElection)
	mux.HandleFunc("/victory", n.handleVictory)

	addr := fmt.Sprintf(":%d", n.port)

	if err := http.ListenAndServe(addr, mux); err != nil {
		fmt.Printf("Node %d failed to start: %v\n", n.id, err)
	}
}
