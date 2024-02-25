package endpoints

import (
	"net/http"

	"github.com/charmingruby/clize/internal/domain/application"
	"github.com/gin-gonic/gin"
)

func NewFetchApplicationsHandler(svc *application.ApplicationService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		apps, err := svc.FetchApplication()
		if err != nil {
			res := WrapResponse[error](
				&err,
				http.StatusBadRequest,
				err.Error(),
			)

			ctx.JSON(http.StatusBadRequest, res)
			return
		}

		res := WrapResponse[[]*application.Application](
			&apps,
			http.StatusOK,
			NewFetchedResponse("application", len(apps)),
		)

		ctx.JSON(http.StatusOK, res)
	}
}
