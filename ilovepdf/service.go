package ilovepdf

import (
	"context"
	"strings"

	"github.com/bxcodec/go-clean-arch/domain"
	"gitlab.com/go-emat/pdfcpu-mattex/pkg/api"
	"gitlab.com/go-emat/pdfcpu-mattex/pkg/pdfcpu"
)

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (a *Service) Process(ctx context.Context, request *domain.IlovepdfRequest) error {
	switch request.Tool {
	case "split":
		SplitProcess(request)
	case "compress":
		CompressProcess(request)
	case "editpdf":
		EditProcess(request)
	}

	return nil
}

func EditProcess(request *domain.IlovepdfRequest) {
	panic("unimplemented")
}

func CompressProcess(request *domain.IlovepdfRequest) {
	panic("unimplemented")
}

func SplitProcess(request *domain.IlovepdfRequest) error {
	conf := pdfcpu.NewDefaultConfiguration()
	spans := strings.Split(request.Ranges, ",")
    err := api.ExtractPagesFile("simple.pdf", "output_dir", spans, conf)
    if err != nil {
        return err
    }

    return nil
}
