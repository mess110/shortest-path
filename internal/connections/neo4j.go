package connections

import (
	"context"
	"fmt"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func NewNeo4jConnection(ctx context.Context) (neo4j.DriverWithContext, error) {
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

	return driver, nil
}

func Neo4jExecuteQuery(ctx context.Context, driver neo4j.DriverWithContext, query string, params map[string]any) (*neo4j.EagerResult, error) {
	result, err := neo4j.ExecuteQuery(ctx, driver,
		query,
		params,
		neo4j.EagerResultTransformer,
		neo4j.ExecuteQueryWithDatabase("neo4j"))
	if err != nil {
		return nil, err
	}

	return result, nil
}
