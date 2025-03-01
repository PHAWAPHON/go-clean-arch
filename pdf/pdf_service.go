package pdf

import (
	"context"

	pdfRepo "github.com/PHAWAPHON/go-clean-arch/internal/repository/pdf_repo"
	"github.com/sirupsen/logrus"
)

type Service struct {
	repo pdfRepo.PDFRepository
}

func NewService(repo pdfRepo.PDFRepository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) Merge(ctx context.Context, files []string, output string) error {
	logrus.Infof("Merging files: %v into output: %s", files, output)
	if err := s.repo.MergePDF(files, output); err != nil {
		logrus.Errorf("Merge failed: %v", err)
		return err
	}
	return nil
}

func (s *Service) Split(ctx context.Context, file string, outputDir string) error {
	logrus.Infof("Splitting file: %s into directory: %s", file, outputDir)
	if _, err := s.repo.SplitPDF(file, outputDir); err != nil {
		logrus.Errorf("Split failed: %v", err)
		return err
	}
	return nil
}

func (s *Service) Compress(ctx context.Context, file string, output string) error {
	logrus.Infof("Compressing file: %s into output: %s", file, output)
	if err := s.repo.CompressPDF(file, output); err != nil {
		logrus.Errorf("Compress failed: %v", err)
		return err
	}
	return nil
}
