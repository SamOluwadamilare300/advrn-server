package utils

import (
	"advrn-server/models/models"
	"advrn-server/models/storage"
	"strconv"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/jwt"
)

func UserIDMiddleware(ctx iris.Context) {
	params := ctx.Params()
	id := params.Get("id")

	claims := jwt.Get(ctx).(*AccessToken)

	userID := strconv.FormatUint(uint64(claims.ID), 10)

	if userID != id {
		ctx.StatusCode(iris.StatusForbidden)
		return
	}
	ctx.Next()
}

func RoleMiddleware(allowedRoles ...string) iris.Handler {
	return func(ctx iris.Context) {
		claims := jwt.Get(ctx).(*AccessToken)
		
		var user models.User
		if err := storage.DB.First(&user, claims.ID).Error; err != nil {
			CreateError(iris.StatusUnauthorized, "Authentication Error", "User not found", ctx)
			return
		}

		// Check if user's role is in allowed roles
		roleAllowed := false
		for _, role := range allowedRoles {
			if user.Role == role {
				roleAllowed = true
				break
			}
		}

		if !roleAllowed {
			CreateError(iris.StatusForbidden, "Authorization Error", "You don't have permission to access this resource", ctx)
			return
		}

		// Store user in context for later use
		ctx.Values().Set("user", user)
		ctx.Next()
	}
}