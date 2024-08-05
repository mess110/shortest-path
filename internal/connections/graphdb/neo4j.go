package graphdb

import (
	"context"
	"fmt"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

type GraphDB struct {
	Ctx    context.Context
	Driver neo4j.DriverWithContext
}

func NewConnection(ctx context.Context) (*GraphDB, error) {
	dbUri := "neo4j://0.0.0.0"
	dbUser := "neo4j"
	dbPassword := "neo4jneo4j"
	driver, err := neo4j.NewDriverWithContext(
		dbUri,
		neo4j.BasicAuth(dbUser, dbPassword, ""))
	if err != nil {
		return nil, err
	}

	err = driver.VerifyConnectivity(ctx)
	if err != nil {
		driver.Close(ctx)
		return nil, err
	}
	fmt.Println("Connection established.")

	graphDB := &GraphDB{
		Driver: driver,
		Ctx:    ctx,
	}

	return graphDB, nil
}

func (g *GraphDB) ExecuteQuery(query string, params map[string]any) (*neo4j.EagerResult, error) {
	result, err := neo4j.ExecuteQuery(g.Ctx, g.Driver,
		query,
		params,
		neo4j.EagerResultTransformer,
		neo4j.ExecuteQueryWithDatabase("neo4j"))
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (g *GraphDB) Close() {
	g.Driver.Close(g.Ctx)
	fmt.Println("Connection closed.")
}
