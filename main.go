package main

func main() {

	// three-node cluster
	node1 := newNode(1, 8001, []Peer{{2, "http://localhost:8002"}, {3, "http://localhost:8003"}})
	node2 := newNode(2, 8002, []Peer{{1, "http://localhost:8001"}, {3, "http://localhost:8003"}})
	node3 := newNode(3, 8003, []Peer{{1, "http://localhost:8001"}, {2, "http://localhost:8002"}})

	go node1.triggerServer()
	go node2.triggerServer()
	go node3.triggerServer()

	select {}
}
