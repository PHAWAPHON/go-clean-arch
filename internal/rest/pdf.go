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

// NewPDFHandler registers the PDF endpoints
func NewPDFHandler(e *echo.Echo, svc PDFService) {
	handler := &PDFHandler{
		Service: svc,
	}
	e.POST("/pdf/merge", handler.Merge)
	e.POST("/pdf/split", handler.Split)
	e.POST("/pdf/compress", handler.Compress)
}

// Merge godoc
// @Summary Merge PDF files
// @Description Merges multiple PDF files into one PDF file.
// @Tags PDF
// @Accept json
// @Produce json
// @Param mergeRequest body domain.MergeRequest true "Merge Request Body"
// @Success 200 {object} map[string]string "Merge successful"
// @Failure 400 {object} ResponseErrorPDF "Invalid request payload"
// @Failure 500 {object} ResponseErrorPDF "Internal server error"
// @Router /pdf/merge [post]
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

// Split godoc
// @Summary Split PDF file
// @Description Splits a single PDF file into multiple pages stored in a specified directory.
// @Tags PDF
// @Accept json
// @Produce json
// @Param splitRequest body domain.SplitRequest true "Split Request Body"
// @Success 200 {object} map[string]string "Split successful"
// @Failure 400 {object} ResponseErrorPDF "Invalid request payload"
// @Failure 404 {object} ResponseErrorPDF "Not Found"
// @Failure 409 {object} ResponseErrorPDF "Conflict error"
// @Failure 500 {object} ResponseErrorPDF "Internal server error"
// @Router /pdf/split [post]
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

// Compress godoc
// @Summary Compress PDF file
// @Description Compresses a PDF file to reduce its size.
// @Tags PDF
// @Accept json
// @Produce json
// @Param compressRequest body domain.CompressRequest true "Compress Request Body"
// @Success 200 {object} map[string]string "Compress successful"
// @Failure 400 {object} ResponseErrorPDF "Invalid request payload"
// @Failure 404 {object} ResponseErrorPDF "Not Found"
// @Failure 409 {object} ResponseErrorPDF "Conflict error"
// @Failure 500 {object} ResponseErrorPDF "Internal server error"
// @Router /pdf/compress [post]
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
