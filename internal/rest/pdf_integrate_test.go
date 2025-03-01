package rest_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	pdfRepo "github.com/PHAWAPHON/go-clean-arch/internal/repository/pdf_repo"
	"github.com/PHAWAPHON/go-clean-arch/internal/rest"
	"github.com/PHAWAPHON/go-clean-arch/pdf"
	"github.com/PHAWAPHON/go-clean-arch/pdf/domain"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func setupIntegrationServer() *echo.Echo {
	e := echo.New()
	repo := pdfRepo.NewPDFRepository()
	svc := pdf.NewService(repo)
	rest.NewPDFHandler(e, svc)
	return e
}

func TestIntegration_PDFMerge(t *testing.T) {
	mergeReq := domain.MergeRequest{
		Files:  []string{"file1.pdf", "file2.pdf"},
		Output: "merged.pdf",
	}
	body, err := json.Marshal(mergeReq)
	assert.NoError(t, err)
	e := setupIntegrationServer()
	req := httptest.NewRequest(http.MethodPost, "/pdf/merge", bytes.NewReader(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusOK, rec.Code)
	_, err = os.Stat(mergeReq.Output)
	assert.NoError(t, err)
	os.Remove(mergeReq.Output)
}

func TestIntegration_PDFSplit(t *testing.T) {
	splitReq := domain.SplitRequest{
		File:      "file1.pdf",
		OutputDir: "split_output",
	}
	body, err := json.Marshal(splitReq)
	assert.NoError(t, err)
	e := setupIntegrationServer()
	req := httptest.NewRequest(http.MethodPost, "/pdf/split", bytes.NewReader(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusOK, rec.Code)
	files, err := os.ReadDir(splitReq.OutputDir)
	assert.NoError(t, err)
	var pdfCount int
	for _, f := range files {
		if filepath.Ext(f.Name()) == ".pdf" {
			pdfCount++
		}
	}
	assert.Greater(t, pdfCount, 0)
	os.RemoveAll(splitReq.OutputDir)
}

func TestIntegration_PDFCompress(t *testing.T) {
	compressReq := domain.CompressRequest{
		File:   "file1.pdf",
		Output: "compressed.pdf",
	}
	body, err := json.Marshal(compressReq)
	assert.NoError(t, err)
	e := setupIntegrationServer()
	req := httptest.NewRequest(http.MethodPost, "/pdf/compress", bytes.NewReader(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusOK, rec.Code)
	info, err := os.Stat(compressReq.Output)
	assert.NoError(t, err)
	assert.Greater(t, info.Size(), int64(0))
	os.Remove(compressReq.Output)
}
