package middleware

import (
	"food-delivery-service/common"
	"food-delivery-service/component"
	"github.com/gin-gonic/gin"
)

func RequiredRoles(appCtx component.AppContext, roles ...string) func(*gin.Context) {
	return func(c *gin.Context) {
		u := c.MustGet(common.CurrentUser).(common.Requester)

		for i := range roles {
			if u.GetRole() == roles[i] {
				c.Next()
				return
			}
		}

		panic(common.ErrNoPermission(nil))
	}
}
