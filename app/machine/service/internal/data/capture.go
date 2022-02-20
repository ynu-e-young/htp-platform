package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	captureV1 "htp-platform/api/capture/service/v1"
	"htp-platform/app/machine/service/internal/biz"
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
		log:  log.NewHelper(log.With(logger, "module", "data/capture")),
	}
}

func (r *captureRepo) ReadOne(ctx context.Context, device int64) (*biz.Capture, error) {
	reply, err := r.data.cc.ReadOne(ctx, &captureV1.ReadOneRequest{Id: device})
	if err != nil {
		return nil, captureV1.ErrorReadDeviceError("read device %d failed, err: %v", device, err)
	}

	return &biz.Capture{
		Data: reply.Image.Data,
	}, nil
}

func (r *captureRepo) ReadAll(ctx context.Context) ([]*biz.Capture, error) {
	reply, err := r.data.cc.ReadAll(ctx, &captureV1.ReadAllRequest{})
	if err != nil {
		return nil, captureV1.ErrorReadDeviceError("read all devices failed, err: %v", err)
	}

	var rets []*biz.Capture
	for _, image := range reply.Images {
		rets = append(rets, &biz.Capture{
			Data: image.Data,
		})
	}
	return rets, nil
}

func (r *captureRepo) ReadOneWithBinary(ctx context.Context, device int64) (*biz.Capture, error) {
	reply, err := r.data.cc.ReadOneWithBinary(ctx, &captureV1.ReadOneWithBinaryRequest{Id: device})
	if err != nil {
		return nil, captureV1.ErrorReadDeviceError("read device %d failed, err: %v", device, err)
	}

	return &biz.Capture{
		Data: reply.Image.Data,
	}, nil
}

func (r *captureRepo) ReadAllWithBinary(ctx context.Context) ([]*biz.Capture, error) {
	reply, err := r.data.cc.ReadAllWithBinary(ctx, &captureV1.ReadAllWithBinaryRequest{})
	if err != nil {
		return nil, captureV1.ErrorReadDeviceError("read all devices failed, err: %v", err)
	}

	var rets []*biz.Capture
	for _, image := range reply.Images {
		rets = append(rets, &biz.Capture{
			Data: image.Data,
		})
	}
	return rets, nil
}

func (r *captureRepo) ReadOneWithBinaryAndCalArea(ctx context.Context, device int64) (*biz.Capture, error) {
	reply, err := r.data.cc.ReadOneWithBinaryAndCalArea(ctx, &captureV1.ReadOneWithBinaryAndCalAreaRequest{Id: device})
	if err != nil {
		return nil, captureV1.ErrorReadDeviceError("read device %d failed, err: %v", device, err)
	}

	return &biz.Capture{
		Data:   reply.Image.Data,
		Pixels: reply.Pixels,
		Area:   reply.Area,
	}, nil
}

func (r *captureRepo) ReadAllWithBinaryAndCalArea(ctx context.Context) ([]*biz.Capture, error) {
	reply, err := r.data.cc.ReadAllWithBinaryAndCalArea(ctx, &captureV1.ReadAllWithBinaryAndCalAreaRequest{})
	if err != nil {
		return nil, captureV1.ErrorReadDeviceError("read all devices failed, err: %v", err)
	}

	var rets []*biz.Capture
	for _, data := range reply.Data {
		rets = append(rets, &biz.Capture{
			Data:   data.Image.Data,
			Pixels: data.Pixels,
			Area:   data.Area,
		})
	}
	return rets, nil
}
