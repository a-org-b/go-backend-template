package publish

import (
	"net/http"

	"github.com/TheLazarusNetwork/go-helpers/httpo"
	"github.com/gin-gonic/gin"
)

// ApplyRoutes applies router to gin Router
func ApplyRoutes(r *gin.RouterGroup) {
	g := r.Group("/publish")
	{
		g.POST("", publish)
	}
}

func publish(c *gin.Context) {
	var body PublishRequest
	err := c.BindJSON(&body)
	if err != nil {
		httpo.NewErrorResponse(http.StatusBadRequest, "failed to validate body").
			Send(c, http.StatusBadRequest)
		return
	}

	httpo.NewSuccessResponse(http.StatusOK, "site deployed successfully").
		SendD(c)

}
