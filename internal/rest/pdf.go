package rest

import (
	"context"
	"net/http"

	"github.com/PHAWAPHON/go-clean-arch/pdf/domain"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type ResponseErrorPDF struct {
	Message string `json:"message"`
}

type PDFService interface {
	Merge(ctx context.Context, files []string, output string) error
	Split(ctx context.Context, file string, outputDir string) error
	Compress(ctx context.Context, file string, output string) error
}

type PDFHandler struct {
	Service PDFService
}

func NewPDFHandler(e *echo.Echo, svc PDFService) {
	handler := &PDFHandler{
		Service: svc,
	}
	e.POST("/pdf/merge", handler.Merge)
	e.POST("/pdf/split", handler.Split)
	e.POST("/pdf/compress", handler.Compress)
}

func (h *PDFHandler) Merge(c echo.Context) error {
	var req domain.MergeRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, ResponseErrorPDF{Message: "Invalid request payload"})
	}

	ctx := c.Request().Context()
	if err := h.Service.Merge(ctx, req.Files, req.Output); err != nil {
		logrus.Errorf("Merge failed: %v", err)
		return c.JSON(getStatusCode(err), ResponseErrorPDF{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Merge successful"})
}

func (h *PDFHandler) Split(c echo.Context) error {
	var req domain.SplitRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, ResponseErrorPDF{Message: "Invalid request payload"})
	}

	ctx := c.Request().Context()
	if err := h.Service.Split(ctx, req.File, req.OutputDir); err != nil {
		logrus.Errorf("Split failed: %v", err)
		return c.JSON(getStatusCodePDF(err), ResponseErrorPDF{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Split successful"})
}

func (h *PDFHandler) Compress(c echo.Context) error {
	var req domain.CompressRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, ResponseErrorPDF{Message: "Invalid request payload"})
	}

	ctx := c.Request().Context()
	if err := h.Service.Compress(ctx, req.File, req.Output); err != nil {
		logrus.Errorf("Compress failed: %v", err)
		return c.JSON(getStatusCodePDF(err), ResponseErrorPDF{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Compress successful"})
}

func getStatusCodePDF(err error) int {
	if err == nil {
		return http.StatusOK
	}

	logrus.Error(err)
	switch err {
	case domain.ErrInternalServerError:
		return http.StatusInternalServerError
	case domain.ErrNotFound:
		return http.StatusNotFound
	case domain.ErrConflict:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}
