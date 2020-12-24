package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"
	"github.com/graphql-go/graphql"

	"products/core/utils"
	gql "products/data/graphql"
)

type reqBody struct {
	Query string `json:"query"`
}

// Instance Graphql Schema type.
type Instance struct {
	GqlSchema graphql.Schema
}

// GraphQL function initiates graphql operations with the provided query.
func (s *Instance) GraphQL(context *gin.Context) {
	if context.Request.Body == nil {
		context.JSON(http.StatusBadRequest, map[string]string{
			"message": "the request need a body",
		})

		return
	}

	var body reqBody

	err := json.NewDecoder(context.Request.Body).Decode(&body)

	if err != nil {
		context.JSON(http.StatusBadRequest, map[string]string{
			"message": "mal formed body",
		})

		return
	}

	result, err := gql.ExecuteQuery(body.Query, *&s.GqlSchema)

	if err != nil {
		stats := utils.GetRequestErrorStats(err.Error(), "")

		context.JSON(stats.Code, map[string]string{
			"message": stats.Message,
		})

		return
	}

	render.WriteJSON(context.Writer, result)
}
