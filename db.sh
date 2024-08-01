#!/usr/bin/env bash

# Create a database
# -v $PWD/data:/data -v $PWD/plugins:/plugins -v $PWD/logs:/logs
docker run -d --rm --name neo4j \
    -p 7474:7474 -p 7687:7687 \
    -e NEO4J_AUTH='neo4j/Pa$$word' \
    -e NEO4J_apoc_export_file_enabled=true \
    -e NEO4J_apoc_import_file_enabled=true \
    -e NEO4J_apoc_import_file_use__neo4j__config=true \
    -e NEO4J_PLUGINS=\[\"apoc\"\] \
    neo4j:5.21

# Wait for the database to start
sleep 10

# Create constraints
docker exec -it neo4j cypher-shell -u neo4j -p 'Pa$$word' -d neo4j \
    "CREATE CONSTRAINT FOR (n:Record) REQUIRE (n.eid, n.uuid) IS UNIQUE"
