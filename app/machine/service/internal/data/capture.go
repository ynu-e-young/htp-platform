package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	captureV1 "htp-platform/api/capture/service/v1"
	machineV1 "htp-platform/api/machine/service/v1"
	"htp-platform/app/machine/service/internal/biz"
	"htp-platform/app/machine/service/internal/data/ent"
	"htp-platform/app/machine/service/internal/data/ent/capturelog"
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

func (r *captureRepo) FindLogsByMachineId(ctx context.Context, machineId int64) ([]*biz.CaptureLog, error) {
	targets, err := r.data.db.CaptureLog.
		Query().Where(capturelog.MachineIDEQ(machineId)).
		All(ctx)

	if err != nil && ent.IsNotFound(err) {
		return nil, machineV1.ErrorNotFoundError("find machineId: %d not found, err: %v", machineId, err)
	}
	if err != nil {
		r.log.Errorf("unknown err: %v", err)
		return nil, machineV1.ErrorUnknownError("unknown err: %v", err)
	}

	var logs []*biz.CaptureLog
	for _, target := range targets {
		logs = append(logs, &biz.CaptureLog{
			Id:        target.ID,
			MachineId: target.MachineID,
			Pixels:    target.Pixels,
			Area:      target.Area,
			ImageName: target.ImageName,
			OssUrl:    target.OssURL,
		})
	}

	return logs, nil
}

func (r *captureRepo) CreateLog(ctx context.Context, captureLog *biz.CaptureLog) (*biz.CaptureLog, error) {
	po, err := r.data.db.CaptureLog.
		Create().
		SetMachineID(captureLog.MachineId).
		SetPixels(captureLog.Pixels).
		SetArea(captureLog.Area).
		SetImageName(captureLog.ImageName).
		SetOssURL(captureLog.OssUrl).
		Save(ctx)
	if err != nil {
		r.log.Errorf("unknown err: %v", err)
		return nil, machineV1.ErrorUnknownError("unknown err: %v", err)
	}

	return &biz.CaptureLog{
		Id:        po.ID,
		MachineId: po.MachineID,
		Pixels:    po.Pixels,
		Area:      po.Area,
		ImageName: po.ImageName,
		OssUrl:    po.OssURL,
	}, nil
}

func (r *captureRepo) GetLog(ctx context.Context, id int64) (*biz.CaptureLog, error) {
	po, err := r.data.db.CaptureLog.Get(ctx, id)
	if err != nil && ent.IsNotFound(err) {
		return nil, machineV1.ErrorNotFoundError("find log id: %s not found, err: %v", id, err)
	}
	if err != nil {
		r.log.Errorf("unknown err: %v", err)
		return nil, machineV1.ErrorUnknownError("unknown err: %v", err)
	}

	return &biz.CaptureLog{
		Id:        po.ID,
		MachineId: po.MachineID,
		Pixels:    po.Pixels,
		Area:      po.Area,
		ImageName: po.ImageName,
		OssUrl:    po.OssURL,
	}, nil
}
