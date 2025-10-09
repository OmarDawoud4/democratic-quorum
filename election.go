package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func (n *Node) startElection() {
	n.mu.Lock()
	// Cooldown: Don't start election if one happened in last 2 seconds
	if time.Since(n.lastElection) < 2*time.Second {
		n.mu.Unlock()
		return
	}
	n.lastElection = time.Now()
	n.mu.Unlock()

	fmt.Printf("Node %d: Election triggered \n", n.id)

	bullies := []Peer{}
	for _, peer := range n.peers {
		if peer.id > n.id {
			bullies = append(bullies, peer)
		}
	}

	if len(bullies) == 0 {
		n.declareVictory()
		return
	}

	// contact bullies
	responseCount := 0
	for _, peer := range bullies {
		resp, err := http.Get(peer.url + "/election")
		if err == nil && resp.StatusCode == http.StatusOK {
			responseCount++
			resp.Body.Close()
		}
	}

	if responseCount == 0 {
		// bullies are unresponsive
		n.declareVictory()
	}
}

func (n *Node) declareVictory() {
	n.mu.Lock()
	n.isLeader = true
	n.leaderID = n.id
	n.mu.Unlock()

	fmt.Printf("Node %d: I AM THE LEADER!\n", n.id)

	// announce to all peers
	for _, peer := range n.peers {
		go func(p Peer) {
			body, _ := json.Marshal(map[string]int{"leaderId": n.id})
			http.Post(p.url+"/victory", "application/json", bytes.NewBuffer(body))
		}(peer)
	}
}
