package v1

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/surrealdb/surrealdb.go"
	"gitlab.com/manifold555112/manifold/lib/graph"
	"gitlab.com/manifold555112/manifold/lib/graph/types"
	"gitlab.com/manifold555112/manifold/server/api/middleware"
)

type schema struct {
	Nodes []*types.BaseNode
	Edges []*types.BaseEdge
}

func PostGraph(c *fiber.Ctx) error {
	fmt.Println(string(c.Body()))
	schema := schema{}
	err := c.BodyParser(&schema)
	if err != nil {
		fmt.Println(err.Error())
		c.SendStatus(fiber.StatusInternalServerError)
		return nil
	}
	graph := graph.Graph{
		Nodes:    []types.Node{},
		Edges:    []types.Edge{},
		Channels: []chan types.Message{},
	}
	for _, bn := range schema.Nodes {
		graph.PushNode(bn)
	}
	for _, e := range schema.Edges {
		graph.PushEdge(e)
	}
	fmt.Printf("%v", graph)
	if buildError := graph.Build(); buildError != nil {
		fmt.Println(buildError.Error)
	}
	if validateError := graph.Validate(); validateError != nil {
		fmt.Println(validateError.Error)
	}
	if db, ok := c.Locals(middleware.USER_DB_KEY).(*surrealdb.DB); ok {
		graphId := c.Params("id")
		if graphId == "new" {
			newId := make([]byte, 16)
			rand.Read(newId)
			newIdString := hex.EncodeToString(newId)
			res, err := db.Create("graphs:"+newIdString, graph)
			if err != nil {
				c.SendStatus(fiber.StatusInternalServerError)
				return nil
			}
			fmt.Printf("%v", res)
			apiResponse := make(map[string]interface{}, 0)
			apiResponse["id"] = newIdString
			apiResponse["graph"] = graph
			responseString, err := json.Marshal(apiResponse)
			if err != nil {
				c.SendStatus(fiber.StatusInternalServerError)
				return nil
			}
			c.SendString(string(responseString))
			c.SendStatus(fiber.StatusAccepted)
			return nil
		}
		_, err := db.Update("graphs:"+graphId, graph)
		if err != nil {
			fmt.Println(err.Error())
			c.SendStatus(fiber.StatusInternalServerError)
		}
		c.SendStatus(fiber.StatusAccepted)
		return nil

	}
	c.SendStatus(fiber.StatusUnauthorized)
	return nil
}
