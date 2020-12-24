package data

import (
	"encoding/json"

	"todo/core/utils"
	gql "todo/data/graphql"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"
	"github.com/graphql-go/graphql"
)

type reqBody struct {
	Query    string `json:"query"`
	Mutation string `json:"mutation"`
}

// Instance type struct - serves to maintain graphql
// server instance.
type Instance struct {
	GqlSchema graphql.Schema
}

// GraphQL function returns GraphQL handler.
func (s *Instance) GraphQL(c *gin.Context) {
	if c.Request.Body == nil {
		errorStats := utils.GetRequestErrorStats("no-query", "")

		c.JSON(errorStats.Code, map[string]string{
			"message": errorStats.Message,
		})

		return
	}

	var body reqBody

	err := json.NewDecoder(c.Request.Body).Decode(&body)

	if err != nil {
		errorStats := utils.GetRequestErrorStats(err.Error(), "")

		c.JSON(errorStats.Code, map[string]string{
			"message": errorStats.Message,
		})

		return
	}

	result, err := gql.ExecuteQuery(body.Query, *&s.GqlSchema)

	if err != nil {
		errorStats := utils.GetRequestErrorStats(err.Error(), err.Error())

		c.JSON(errorStats.Code, map[string]string{
			"message": errorStats.Message,
		})

		return
	}

	render.WriteJSON(c.Writer, result)
}
