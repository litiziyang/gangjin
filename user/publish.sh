rover subgraph introspect \
  http://users:4001/query | \
  APOLLO_KEY=service:My-Graph-nojub:OTQ4lWKivvnbD0iypMdrzA \
  rover subgraph publish My-Graph-nojub@current \
  --name users --schema - \
  --routing-url http://users:4001/query