package auth

import (
	"technical-test-backend/helper"

	"technical-test-backend/exception"

	auth "technical-test-backend/verification_auth"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

const (
	RoleAdministrator = "Administrator"
	RoleClient        = "Client"
)

type AccessDetails struct {
	ID   uint
	Role string
	Name string
	Nip  string
}

func Auth(next func(c *gin.Context, auth *AccessDetails), roles []string) gin.HandlerFunc {
	return func(c *gin.Context) {

		// Extracting the token metadata from the request.
		tokenAuth, err := auth.ExtractTokenMetadata(c.Request, nil, func(claims jwt.MapClaims) interface{} {
			return AccessDetails{
				ID:   uint(claims["id"].(float64)),
				Role: claims["role"].(string),
				Name: claims["name"].(string),
				Nip:  claims["nip"].(string),
			}
		})
		if err != nil {
			helper.PanicIfError(exception.ErrUnauthorized)
		}

		// Checking if the type of tokenAuth is AccessDetails. If it is not, it will panic.
		accessDetails, ok := tokenAuth.(AccessDetails)
		if !ok {
			helper.PanicIfError(exception.ErrInternalServer)
		}

		// Checking if the role of the user is in the list of roles that are allowed to access the route.
		if !helper.Contains(roles, accessDetails.Role) {
			helper.PanicIfError(exception.ErrPermissionDenied)
		}

		// Calling the next function in the chain.
		next(c, &accessDetails)
	}
}
