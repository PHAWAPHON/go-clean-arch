package domain

type PDFRequest struct {
	InputFiles []string `json:"input_files,omitempty"`
	InputFile  string   `json:"input_file,omitempty"`
	OutputFile string   `json:"output_file,omitempty"`
	OutputDir  string   `json:"output_dir,omitempty"`
}

type PDFUseCase interface {
	MergePDF(req PDFRequest) (string, error)
	SplitPDF(req PDFRequest) ([]string, error)
	CompressPDF(req PDFRequest) (string, error)
}

type MergeRequest struct {
	Files  []string `json:"files" validate:"required,min=1"`
	Output string   `json:"output" validate:"required"`
}

type SplitRequest struct {
	File      string `json:"file" validate:"required"`
	OutputDir string `json:"outputDir" validate:"required"`
}
type CompressRequest struct {
	File   string `json:"file" validate:"required"`
	Output string `json:"output" validate:"required"`
}
