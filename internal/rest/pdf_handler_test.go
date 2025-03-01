package rest_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/PHAWAPHON/go-clean-arch/internal/rest"
	"github.com/PHAWAPHON/go-clean-arch/internal/rest/mocks"
	"github.com/PHAWAPHON/go-clean-arch/pdf/domain"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestPDFMerge_Success(t *testing.T) {

	mergeReq := domain.MergeRequest{
		Files:  []string{"file1.pdf", "file2.pdf"},
		Output: "merged.pdf",
	}

	mockPDFService := new(mocks.PDFService)
	mockPDFService.
		On("Merge", mock.Anything, mergeReq.Files, mergeReq.Output).
		Return(nil)

	e := echo.New()
	reqBody, err := json.Marshal(mergeReq)
	assert.NoError(t, err)
	req := httptest.NewRequest(http.MethodPost, "/pdf/merge", strings.NewReader(string(reqBody)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	handler := rest.PDFHandler{Service: mockPDFService}
	err = handler.Merge(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	mockPDFService.AssertExpectations(t)
}

func TestPDFMerge_Error(t *testing.T) {
	mergeReq := domain.MergeRequest{
		Files:  []string{"file1.pdf", "file2.pdf"},
		Output: "merged.pdf",
	}

	mockPDFService := new(mocks.PDFService)
	mockPDFService.
		On("Merge", mock.Anything, mergeReq.Files, mergeReq.Output).
		Return(assert.AnError)

	e := echo.New()
	reqBody, err := json.Marshal(mergeReq)
	assert.NoError(t, err)
	req := httptest.NewRequest(http.MethodPost, "/pdf/merge", strings.NewReader(string(reqBody)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	handler := rest.PDFHandler{Service: mockPDFService}
	err = handler.Merge(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusInternalServerError, rec.Code)
	mockPDFService.AssertExpectations(t)
}

func TestPDFSplit_Success(t *testing.T) {
	splitReq := domain.SplitRequest{
		File:      "input.pdf",
		OutputDir: "output",
	}

	mockPDFService := new(mocks.PDFService)
	mockPDFService.
		On("Split", mock.Anything, splitReq.File, splitReq.OutputDir).
		Return(nil)

	e := echo.New()
	reqBody, err := json.Marshal(splitReq)
	assert.NoError(t, err)
	req := httptest.NewRequest(http.MethodPost, "/pdf/split", strings.NewReader(string(reqBody)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	handler := rest.PDFHandler{Service: mockPDFService}
	err = handler.Split(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	mockPDFService.AssertExpectations(t)
}

func TestPDFCompress_Success(t *testing.T) {
	compressReq := domain.CompressRequest{
		File:   "input.pdf",
		Output: "compressed.pdf",
	}

	mockPDFService := new(mocks.PDFService)
	mockPDFService.
		On("Compress", mock.Anything, compressReq.File, compressReq.Output).
		Return(nil)

	e := echo.New()
	reqBody, err := json.Marshal(compressReq)
	assert.NoError(t, err)
	req := httptest.NewRequest(http.MethodPost, "/pdf/compress", strings.NewReader(string(reqBody)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	handler := rest.PDFHandler{Service: mockPDFService}
	err = handler.Compress(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	mockPDFService.AssertExpectations(t)
}
