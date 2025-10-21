# Democratic Counter 

A distributed system demonstrating leader election using the Bully Algorithm.

## How It Works

### The Bully Algorithm
- Each node has a unique ID (1, 2, 3)
- When election starts, nodes "bully" others with lower IDs
- Highest ID node becomes the leader
- Leader manages the shared counter

### Nodes
- **Node 1**: http://localhost:8001
- **Node 2**: http://localhost:8002  
- **Node 3**: http://localhost:8003

## API Endpoints
- `GET /ping` - Health check
- `GET /status` - Node status (ID, leader, state)
- `POST /election` - Trigger election
- `POST /victory` - Announce new leader
## Run with Docker
Pull the image
```bash
docker pull omardawoud/democratic-quorum:latest
```
Run the container
```bash
docker run -p 8001:8001 omardawoud/democratic-quorum
```
### Docker Hub Link
[https://hub.docker.com/r/omardawoud/democratic-quorum](https://hub.docker.com/r/omardawoud/democratic-quorum)