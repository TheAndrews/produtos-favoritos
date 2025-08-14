package controllers

import (
	"net/http"
	"reflect"

	//"github.com/dgrijalva/jwt-go"
	"produtos-favoritos/src/api/forms"
	"produtos-favoritos/src/internals/exceptions"

	"github.com/gin-gonic/gin"
)

type BaseController struct {
}

// func (b *BaseController) getCurrentUser(ctx *gin.Context) (string, error) {
// 	c, err := ctx.Request.Cookie("token")
// 	if err != nil {
// 		if err == http.ErrNoCookie {
// 			ctx.JSON(http.StatusUnauthorized, nil)
// 			return "", err
// 		}
// 		ctx.JSON(http.StatusBadRequest, nil)
// 		return "", err
// 	}

// 	tknStr := c.Value
// 	claims := &Claims{}
// 	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
// 		return jwtKey, nil
// 	})
// 	if !tkn.Valid {
// 		return "", err
// 	}
// 	if err != nil {
// 		return "", err
// 	}
// 	return claims.Username, nil
// }

func (b *BaseController) respond(ctx *gin.Context, result interface{}) {
	value := reflect.ValueOf(result)

	if value.Kind() == reflect.Struct {
		field := value.FieldByName("BaseForm")

		if field.IsValid() {
			value := field.Interface()
			form, ok := value.(forms.BaseForm)

			if ok && !form.IsValid() {
				ctx.JSON(http.StatusUnprocessableEntity, form.GetErrors())
				return
			}
		}
	}

	ctx.JSON(http.StatusOK, result)
}

func (b *BaseController) respondSuccessNoContent(ctx *gin.Context) {
	ctx.Status(http.StatusNoContent)
}

func (b *BaseController) respondError(ctx *gin.Context, err error) {
	switch err.(type) {
	case *exceptions.BadRequestError:
		ctx.JSON(http.StatusBadRequest, err.Error())
	case *exceptions.EmailAlreadyRegisteredErr:
		ctx.JSON(http.StatusUnprocessableEntity, err.Error())
	case *exceptions.AlreadyWishlistedErr:
		ctx.JSON(http.StatusBadRequest, err.Error())
	case *exceptions.InvalidEntityError:
		ctx.JSON(http.StatusUnprocessableEntity, err.Error())
	case *exceptions.InvalidCredentialsError:
		ctx.JSON(http.StatusUnauthorized, err.Error())
	case *exceptions.NotFoundEntityError:
		ctx.JSON(http.StatusNotFound, err.Error())
	default:
		ctx.JSON(http.StatusInternalServerError, err.Error())
	}
}
