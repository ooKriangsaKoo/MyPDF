package rest

import (
	"context"
	"net/http"

	"github.com/bxcodec/go-clean-arch/domain"
	"github.com/labstack/echo/v4"
	validator "gopkg.in/go-playground/validator.v9"
)

type IlovepdfService interface {
	Process(ctx context.Context, ilovepdf *domain.IlovepdfRequest) error
}

type ILovePdfHandler struct {
	Service IlovepdfService
}

func NewILovePdfHandlerHandler(e *echo.Echo, svc IlovepdfService) {
	handler := &ILovePdfHandler{
		Service: svc,
	}
	e.POST("/v1/process", handler.Process)
}

// @Summary Process pdf
// @Description Process pdf type [split,compress,editpdf]
// @Tags process
// @Accept json
// @Produce json
// @param Process body domain.IlovepdfRequest true "Process model"
// @Success 200 {object} domain.IlovepdfReponse
// @Router /process [post]
func (a *ILovePdfHandler) Process(c echo.Context) error {
	var ilovepdf domain.IlovepdfRequest
	err := c.Bind(&ilovepdf)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	var ok bool
	if ok, err = isRequestValid2(&ilovepdf); !ok {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	ctx := c.Request().Context()
	err = a.Service.Process(ctx, &ilovepdf)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"download_filename": "output.zip",
		"filesize":          0,
		"output_filesize":   0,
		"output_filenumber": 2,
		"output_extensions": "[\"pdf\"]",
		"timer":             "0.028",
		"status":            "TaskSuccess",
	})
}

func isRequestValid2(m *domain.IlovepdfRequest) (bool, error) {
	validate := validator.New()
	err := validate.Struct(m)
	if err != nil {
		return false, err
	}
	return true, nil
}
