package usergin

import (
	"food-delivery-service/common"
	"food-delivery-service/component"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Profile(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		u := c.MustGet(common.CurrentUser)

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(u))
	}
}
