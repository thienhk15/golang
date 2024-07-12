package middleware

import (
	"main/component/api"
	"main/component/responses"
	"main/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RouteConfig struct {
	RequestUrl string
	Method     string
}

// PermissionMiddleWare
// Check access permission by route
func (m *ApiMiddleware) PermissionMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		oauth, exists := c.Get(utils.ContextUserAuthorized)
		if !exists {
			api.Custom(c, http.StatusForbidden, "Access Denied", nil)
			c.Abort()
			return
		}

		response := oauth.(responses.UserLoginResponse)
		curUser := response.User

		mapRoles := m.configRoutePermission()
		currentRole := ""
		for _, role := range curUser.Role {
			if len(mapRoles[utils.Permission(role)]) == 0 {
				c.Next()
				return
			}

			currentRole = string(role)
		}

		url := c.FullPath()
		method := c.Request.Method

		for _, val := range mapRoles[utils.Permission(currentRole)] {
			if val.Method == method && val.RequestUrl == url {
				c.Next()
				return
			}
		}

		api.Custom(c, http.StatusForbidden, "Access Denied", nil)
		c.Abort()
	}
}

func (m *ApiMiddleware) configRoutePermission() map[utils.Permission][]RouteConfig {
	return map[utils.Permission][]RouteConfig{
		utils.AdminPermission: {
			{RequestUrl: "/api/v1/users", Method: http.MethodGet},
		},
	}
}

func Scopes(permissions ...utils.Permission) gin.HandlerFunc {
	return (func(c *gin.Context) {
		curUser := utils.GetCurrentUser(c).User

		set := make(map[utils.Permission]bool)
		for _, p := range permissions {
			set[p] = true
		}

		for _, role := range curUser.Role {
			if string(role) != "" && set[utils.Permission(role)] {
				c.Next()
				return
			}
		}

		api.Custom(c, http.StatusForbidden, "Access Denied", nil)
		c.Abort()
	})
}
