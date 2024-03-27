package middleware

import (
	"PopcornMovie/internal/utils"
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

var (
	authorizationHeaderKey  = "authorization"
	authorizationTypeBearer = "bearer"
	authorizationPayloadKey = "authorization_payload"
)

// AuthMiddleware is a middleware verify the accessToken
func AuthMiddleware(jwtSecret string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorizationHeader := ctx.GetHeader(authorizationHeaderKey)
		if len(authorizationHeader) == 0 {
			ctx.Next()
			return
		}

		fields := strings.Fields(authorizationHeader)
		if len(fields) < 2 {
			err := errors.New("invalid authorization header format")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, err)
			return
		}

		authorizationType := fields[0]
		if strings.ToLower(authorizationType) != authorizationTypeBearer {
			err := fmt.Sprintf("unsupported authorization type %s", authorizationTypeBearer)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errors.New(err))
			return
		}

		accessToken := fields[1]
		payload, err := utils.VerifyToken(accessToken, jwtSecret)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errors.New(err.Error()))
			return
		}

		c := context.WithValue(ctx.Request.Context(), authorizationPayloadKey, payload)
		ctx.Request = ctx.Request.WithContext(c)
		ctx.Next()
	}
}

func GetPayload(ctx context.Context) *utils.Payload {
	raw, _ := ctx.Value(authorizationPayloadKey).(*utils.Payload)
	return raw
}
