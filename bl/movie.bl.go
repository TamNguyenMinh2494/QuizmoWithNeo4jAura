package bl

import (
	"net/http"
	"quizmo/core"

	"github.com/labstack/echo/v4"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

type MovieBusinessLogic struct {
	Server *core.Server
}

func NewMovieBusinessLogic(server *core.Server) *MovieBusinessLogic {
	return &MovieBusinessLogic{
		Server: server,
	}
}

func (bl *MovieBusinessLogic) RelationshipBetweenPeople() (interface{}, error) {

	data := make(map[string]interface{})

	records, err := bl.Server.Session.WriteTransaction(
		func(tx neo4j.Transaction) (interface{}, error) {
			createRelationshipBetweenPeopleQuery := `
				MERGE (p1:Person { name: $person1_name })
				MERGE (p2:Person { name: $person2_name })
				MERGE (p1)-[:KNOWS]->(p2)
				RETURN p1, p2`
			result, err := tx.Run(createRelationshipBetweenPeopleQuery, map[string]interface{}{
				"person1_name": "Alice",
				"person2_name": "David",
			})
			if err != nil {
				// Return the error received from driver here to indicate rollback,
				// the error is analyzed by the driver to determine if it should try again.
				return nil, err
			}

			return result.Collect()
		})
	if err != nil {
		panic(err)
	}
	for _, record := range records.([]*neo4j.Record) {
		firstPerson := record.Values[0].(neo4j.Node)
		data["firstPerson"] = firstPerson.Props["name"]
		secondPerson := record.Values[1].(neo4j.Node)
		data["secondPerson"] = secondPerson.Props["name"]
	}

	_, _ = bl.Server.Session.ReadTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		// Code within this function might be invoked more than once in case of
		// transient errors.
		readPersonByName := `
				MATCH (p:Person)
				WHERE p.name = $person_name
				RETURN p.name AS name`
		result, err := tx.Run(readPersonByName, map[string]interface{}{
			"person_name": "Alice",
		})
		if err != nil {
			return nil, err
		}

		for result.Next() {
			record := result.Record().Values[0].(string)
			data["person"] = record
		}

		return nil, echo.NewHTTPError(http.StatusInternalServerError, result.Err())
	})
	// fmt.Printf("data: %v", data)
	// if err != nil {
	// 	return nil, echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	// }
	// fmt.Printf("data: %v", data)
	return data, nil
}
