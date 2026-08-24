package components

import (
	"context"
)

type DetectorService struct {
	ctx context.Context
}

func NewDetectorService() *DetectorService {
	return &DetectorService{}
}

func (this *DetectorService) Startup(ctx context.Context) error {
	this.ctx = ctx
	return nil
}
