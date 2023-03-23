package data

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"

	"github.com/ynu-e-young/htp-platform/app/capture/service/internal/biz"
)

var _ biz.CaptureRepo = (*captureRepo)(nil)

type captureRepo struct {
	data *Data
	log  *log.Helper
}

// NewCaptureRepo .
func NewCaptureRepo(data *Data, logger log.Logger) biz.CaptureRepo {
	return &captureRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "data/cv")),
	}
}

func (r *captureRepo) ReadOne(ctx context.Context, device int) (*biz.Capture, error) {
	mat := r.data.captures[device].Clone()

	return &biz.Capture{
		Mat: &mat,
	}, nil
}

func (r *captureRepo) ReadAll(ctx context.Context) ([]*biz.Capture, error) {
	var rets []*biz.Capture

	for _, capture := range r.data.captures {
		mat := capture.Clone()
		rets = append(rets, &biz.Capture{
			Mat: &mat,
		})
	}

	return rets, nil
}
