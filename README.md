# shortest-path

Do not use this code for anything, it is hacked together for a quick experiment

Start API
```
go run cmd/server/api.go
```

Flush db
```
curl -X POST -d "{\"query\": \"MATCH (n) OPTIONAL MATCH (n)-[r]-() DELETE n,r\"}" -H "Content-Type: application/json" http://localhost:3000/query
```

Create node with generated internal uuids
```
curl -X POST -d "{\"query\": \"CREATE (a:Node {value: '\$uuid'}) CREATE (b:Node {value: '\$uuid'}) CREATE (a)-[:RELATED_TO]->(b);\"}" -H "Content-Type: application/json" http://localhost:3000/query
```

Get shortest path
```
curl http://localhost:3000/shortest_path/:start/:end
```

Get parent
```
curl -X POST -d "{\"query\": \"MATCH (child:Node {value: '6a5753f1-c7d3-4763-8c73-22f542b86c74'})<-[:RELATED_TO*]-(parent:Node) RETURN parent\"}" -H "Content-Type: application/json" http://localhost:3000/query
```