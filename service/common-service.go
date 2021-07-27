package service

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type (
	CommonService interface {
		GetUserId(context *gin.Context) string
	}

	commonService struct {
		jwtService JWTService
	}
)

func NewCommonService(jwtService JWTService) CommonService {
	return &commonService{
		jwtService: jwtService,
	}
}

func (service *commonService) GetUserId(context *gin.Context) string {
	authHeader := context.GetHeader("Authorization")
	aToken, err := service.jwtService.ValidateToken(authHeader)
	if err != nil {
		panic(err.Error())
	}
	claims := aToken.Claims.(jwt.MapClaims)
	id := fmt.Sprintf("%v", claims["user_id"])
	return id

}
