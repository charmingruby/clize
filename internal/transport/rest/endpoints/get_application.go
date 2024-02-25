package endpoints

import (
	"net/http"

	"github.com/charmingruby/clize/internal/domain/application"
	"github.com/charmingruby/clize/internal/validation"
	"github.com/gin-gonic/gin"
)

func NewGetApplicationHandler(svc *application.ApplicationService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		name := ctx.Param("name")

		app, err := svc.GetApplication(name)
		if err != nil {
			rnf, ok := err.(*validation.ResourceNotFoundError)
			if ok {
				res := WrapResponse[validation.ResourceNotFoundError](
					rnf,
					http.StatusNotFound,
					rnf.Error(),
				)

				ctx.JSON(http.StatusNotFound, res)
				return
			}

			res := WrapResponse[string](
				nil,
				http.StatusBadRequest,
				rnf.Error(),
			)

			ctx.JSON(http.StatusBadRequest, res)
			return
		}

		res := WrapResponse[application.Application](
			app,
			http.StatusOK,
			NewFetchedResponse("application", 1),
		)

		ctx.JSON(http.StatusOK, res)
	}
}
