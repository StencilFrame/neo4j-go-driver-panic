package main

import (
	"context"
	"crypto/rand"
	"errors"
	"flag"
	"fmt"
	"math/big"

	"github.com/google/uuid"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

const (
	uri      = "bolt://localhost:7687"
	username = "neo4j"
	password = "Pa$$word"
)

func getDB(ctx context.Context) (neo4j.DriverWithContext, error) {
	db, err := neo4j.NewDriverWithContext(uri, neo4j.BasicAuth(username, password, ""))
	if err != nil {
		return nil, err
	}
	err = db.VerifyConnectivity(ctx)
	if err != nil {
		fmt.Println("Failed to verify connectivity", err, uri)
		return nil, err
	}
	return db, nil
}

func populateDB(ctx context.Context, uuid string, amount int) error {
	fmt.Println("Populating DB", uuid)
	db, err := getDB(ctx)
	if err != nil {
		fmt.Println("Failed to verify connectivity", err, uri)
		return err
	}

	session := db.NewSession(ctx, neo4j.SessionConfig{
		AccessMode:   neo4j.AccessModeWrite,
		DatabaseName: "neo4j",
	})
	defer session.Close(ctx)

	// Prepare data
	records := []map[string]interface{}{}
	for i := 1; i < amount; i++ {
		records = append(records, map[string]interface{}{
			"id":  i,
			"typ": getRandomType(),
		})
	}

	// Populate the DB
	_, err = session.ExecuteWrite(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		result, err := tx.Run(ctx, `
			WITH $records AS records
			UNWIND records AS record
			MERGE (p:Record {uuid: $uuid, eid: record.id - 1})
			MERGE (n:Record {uuid: $uuid, eid: record.id})
			MERGE (p)-[:NEXT]->(n)
			WITH n, record
			CALL apoc.create.addLabels(n, [record.typ]) YIELD node
			RETURN node
			`,
			map[string]interface{}{
				"uuid":    uuid,
				"records": records,
			},
		)
		if err != nil {
			return nil, err
		}

		records, err := result.Collect(ctx)
		if err != nil {
			return nil, err
		}
		return records, nil
	})
	if err != nil {
		fmt.Println("Failed to execute write", err)
		return err
	}

	return nil
}

func getRandomType() string {
	types := []string{"Item", "NotItem"}
	n, _ := rand.Int(rand.Reader, big.NewInt(2))
	return types[n.Int64()]
}

func queryDB(ctx context.Context, uuid string, id1, id2 int) (any, error) {
	db, err := getDB(ctx)
	if err != nil {
		fmt.Println("Failed to verify connectivity", err, uri)
		return nil, err
	}

	session := db.NewSession(ctx, neo4j.SessionConfig{
		AccessMode:   neo4j.AccessModeRead,
		DatabaseName: "neo4j",
	})
	defer session.Close(ctx)

	node1, err := getNextItem(ctx, uuid, id1)
	if err != nil {
		fmt.Println("Failed to get next Item node", err)
		return nil, err
	}
	n1 := node1.([]*neo4j.Record)
	eid1 := n1[0].Values[0].(neo4j.Node).Props["eid"]
	fmt.Println("Node1", eid1)

	node2, err := getNextItem(ctx, uuid, id2)
	if err != nil {
		fmt.Println("Failed to get next Item node", err)
		return nil, err
	}
	n2 := node2.([]*neo4j.Record)
	eid2 := n2[0].Values[0].(neo4j.Node).Props["eid"]
	fmt.Println("Node2", eid2)

	// Query the DB
	result, err := session.ExecuteRead(ctx, func(tx neo4j.ManagedTransaction) (interface{}, error) {
		result, err := tx.Run(ctx, `
			MATCH (s1:Item {uuid: $uuid, eid: $eid1})
			WITH s1
			MATCH (s2:Item {uuid: $uuid, eid: $eid2})
			WITH s1, s2
			MATCH p = shortestPath((s1)-[r:NEXT*0..500]->(s2)),
				p1 = (before:Item)-[:NEXT*0..500]->(s1),
				p2 = (s2)-[:NEXT*0..500]->(after:Item),
				path = shortestPath((before)-[:NEXT*0..500]->(after))
			WHERE
				NONE(n IN nodes(p1)[1..-1] WHERE n:Item)
				AND NONE(n IN nodes(p2)[1..-1] WHERE n:Item)
			RETURN path
			`,
			map[string]interface{}{
				"eid1": eid1,
				"eid2": eid2,
				"uuid": uuid,
			})
		if err != nil {
			return nil, err
		}
		if result == nil {
			return nil, errors.New("no path found between the Item nodes")
		}
		records, err := result.Collect(ctx)
		if err != nil {
			return nil, err
		}
		return records, nil
	})
	if err != nil {
		fmt.Println("Failed to execute read", err)
		return nil, err
	}

	return result, nil
}

// Find the next Item node after the given eid
func getNextItem(ctx context.Context, uuid string, eid int) (any, error) {
	db, err := getDB(ctx)
	if err != nil {
		fmt.Println("Failed to verify connectivity", err, uri)
		return nil, err
	}

	session := db.NewSession(ctx, neo4j.SessionConfig{
		AccessMode:   neo4j.AccessModeRead,
		DatabaseName: "neo4j",
	})
	defer session.Close(ctx)

	// Query the DB
	result, err := session.ExecuteRead(ctx, func(tx neo4j.ManagedTransaction) (interface{}, error) {
		result, err := tx.Run(ctx, `
			MATCH (n:Record {uuid: $uuid, eid: $eid})
			WITH n
			MATCH p = shortestPath((n)-[r:NEXT*0..500]->(i:Item))
			WHERE NONE(x IN nodes(p)[0..-1] WHERE x:Item)
			RETURN i
			`,
			map[string]interface{}{
				"eid":  eid,
				"uuid": uuid,
			})
		if err != nil {
			return nil, err
		}
		if result == nil {
			return nil, errors.New("no path found between the Item nodes")
		}
		records, err := result.Collect(ctx)
		if err != nil {
			return nil, err
		}
		return records, nil
	})
	if err != nil {
		fmt.Println("Failed to execute read", err)
		return nil, err
	}

	return result, nil
}

func main() {
	ctx := context.Background()

	flag.Usage = func() {
		fmt.Println("Usage: main [populate|query]")
		flag.PrintDefaults()
	}
	uuidFlag := flag.String("uuid", "", "UUID of the data")
	flag.Parse()
	command := flag.Arg(0)

	if command == "" || (command != "populate" && command != "query") {
		flag.Usage()
		return
	}

	if command == "populate" {
		uuid := uuid.New().String()

		err := populateDB(ctx, uuid, 2000)
		if err != nil {
			fmt.Println("Failed to populate DB", err)
			return
		}
	}

	if command == "query" {
		// result, err := queryDB(ctx, *uuidFlag, 122, 600) <- works
		result, err := queryDB(ctx, *uuidFlag, 916, 1015) // <- panic
		if err != nil {
			fmt.Println("Failed to query DB", err)
			return
		}

		records := result.([]*neo4j.Record)
		if len(records) == 0 {
			fmt.Println("No path found between the Item nodes")
			return
		}
		for _, record := range records {
			for _, value := range record.Values {
				node := value.(neo4j.Path).Nodes
				for _, n := range node {
					fmt.Print(n.Props["eid"], " ")
				}
			}
			fmt.Println()
		}
	}
}
