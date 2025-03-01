package pdf_repo

import (
	"os"
	"path/filepath"

	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu"
)

type PDFRepository interface {
	MergePDF(inputFiles []string, outputFile string) error
	SplitPDF(inputFile, outputDir string) ([]string, error)
	CompressPDF(inputFile, outputFile string) error
}

type pdfRepository struct{}

func NewPDFRepository() PDFRepository {
	return &pdfRepository{}
}

func (r *pdfRepository) MergePDF(inputFiles []string, outputFile string) error {
	if err := api.MergeCreateFile(inputFiles, outputFile, nil); err != nil {
		return err
	}
	return nil
}
func (r *pdfRepository) SplitPDF(inputFile, outputDir string) ([]string, error) {

	if err := os.MkdirAll(outputDir, os.ModePerm); err != nil {
		return nil, err
	}

	cfg := pdfcpu.NewDefaultConfiguration()

	if err := api.SplitFile(inputFile, outputDir, 1, cfg); err != nil {
		return nil, err
	}

	splittedFiles := []string{}
	files, err := os.ReadDir(outputDir)
	if err != nil {
		return nil, err
	}

	for _, f := range files {
		if filepath.Ext(f.Name()) == ".pdf" {
			splittedFiles = append(splittedFiles, filepath.Join(outputDir, f.Name()))
		}
	}
	return splittedFiles, nil
}

func (r *pdfRepository) CompressPDF(inputFile, outputFile string) error {
	cfg := pdfcpu.NewDefaultConfiguration()
	if err := api.OptimizeFile(inputFile, outputFile, cfg); err != nil {
		return err
	}
	return nil
}
