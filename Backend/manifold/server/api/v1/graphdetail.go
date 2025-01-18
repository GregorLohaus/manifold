package v1

import (
	"encoding/json"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/surrealdb/surrealdb.go"
	"gitlab.com/manifold555112/manifold/lib/graph/types"
	"gitlab.com/manifold555112/manifold/server/api/middleware"
)

type Unmarshalable struct {
	Nodes []*types.BaseNode `json:"nodes,omitempty"`
	Edges []*types.BaseEdge `json:"edges,omitempty"`
}

func GraphDetail(c *fiber.Ctx) error {
	if db, ok := c.Locals(middleware.USER_DB_KEY).(*surrealdb.DB); ok {
		id := c.Params("id")
		res, err := db.Query("select * from $id", map[string]interface{}{"id": "graphs:" + id})
		if err != nil {
			fmt.Println(err.Error())
			c.SendStatus(fiber.StatusNoContent)
			return nil
		}
		graphs := []Unmarshalable{}
		_, err = surrealdb.UnmarshalRaw(res, &graphs)
		if err != nil {
			fmt.Println(err.Error())
			c.SendStatus(fiber.StatusInternalServerError)
			return nil
		}
		if len(graphs) < 1 {
			c.SendStatus(fiber.StatusNoContent)
			return nil
		}
		graphjson, err := json.Marshal(graphs[0])
		if err != nil {
			fmt.Println(err.Error())
			c.SendStatus(fiber.StatusInternalServerError)
			return nil
		}
		fmt.Println("Sending json")
		fmt.Println(string(graphjson))
		c.SendString(string(graphjson))
		c.SendStatus(fiber.StatusAccepted)
		return nil

	}
	c.SendStatus(fiber.StatusUnauthorized)
	return nil
}
