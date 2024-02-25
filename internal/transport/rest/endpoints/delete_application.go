package endpoints

import (
	"net/http"

	"github.com/charmingruby/clize/internal/domain/application"
	"github.com/charmingruby/clize/internal/validation"
	"github.com/gin-gonic/gin"
)

func NewDeleteApplicationHandler(svc *application.ApplicationService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		name := ctx.Param("name")

		err := svc.DeleteApplication(name)
		if err != nil {
			rnf, ok := err.(*validation.ResourceNotFoundError)

			if !ok {
				res := WrapResponse[validation.ResourceNotFoundError](
					rnf,
					http.StatusNotFound,
					err.Error(),
				)
				ctx.JSON(http.StatusNotFound, res)
				return
			}

		}

		res := WrapResponse[string](
			nil,
			http.StatusOK,
			NewDeletedResponse(name),
		)

		ctx.JSON(http.StatusOK, res)
	}
}
