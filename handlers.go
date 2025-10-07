package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// health check
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

// bully behaviour
func (n *Node) handleElection(w http.ResponseWriter, r *http.Request) {
	n.mu.Lock()
	defer n.mu.Unlock()

	go n.startElection()
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]int{"id": n.id})
}

func (n *Node) handleVictory(w http.ResponseWriter, r *http.Request) {
	var msg struct {
		LeaderID int `json:"leaderId"`
	}

	json.NewDecoder(r.Body).Decode(&msg)

	n.mu.Lock()
	n.leaderID = msg.LeaderID
	n.isLeader = false
	n.mu.Unlock()

	fmt.Printf("Node %d: Acknowledged leader %d\n", n.id, msg.LeaderID)
	w.WriteHeader(http.StatusOK)
}
