package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
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

func (r *captureRepo) ReadOneWithBinaryAndSrc(ctx context.Context, device int64) (*biz.CaptureSrc, error) {
	reply, err := r.data.cc.ReadOneWithBinaryAndSrc(ctx, &captureV1.ReadOneWithBinaryAndSrcRequest{Id: device})
	if err != nil {
		return nil, captureV1.ErrorReadDeviceError("read device %d failed, err: %v", device, err)
	}

	return &biz.CaptureSrc{
		Proc: reply.GetImageSrc().GetProc(),
		Src:  reply.GetImageSrc().GetSrc(),
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

func (r *captureRepo) ReadAllWithBinaryAndSrc(ctx context.Context) ([]*biz.CaptureSrc, error) {
	reply, err := r.data.cc.ReadAllWithBinaryAndSrc(ctx, &captureV1.ReadAllWithBinaryAndSrcRequest{})
	if err != nil {
		return nil, captureV1.ErrorReadDeviceError("read all devices failed, err: %v", err)
	}

	var rets []*biz.CaptureSrc
	for _, imageSrc := range reply.GetImageSrc() {
		rets = append(rets, &biz.CaptureSrc{
			Proc: imageSrc.GetProc(),
			Src:  imageSrc.GetSrc(),
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

func (r *captureRepo) ReadOneWithBinaryAndCalAreaAndSrc(ctx context.Context, device int64) (*biz.CaptureSrc, error) {
	reply, err := r.data.cc.ReadOneWithBinaryAndCalAreaAndSrc(ctx, &captureV1.ReadOneWithBinaryAndCalAreaAndSrcRequest{Id: device})
	if err != nil {
		return nil, captureV1.ErrorReadDeviceError("read device %d failed, err: %v", device, err)
	}

	return &biz.CaptureSrc{
		Proc:   reply.GetImageSrc().GetProc(),
		Src:    reply.GetImageSrc().GetSrc(),
		Pixels: reply.GetPixels(),
		Area:   reply.GetArea(),
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

func (r *captureRepo) ReadAllWithBinaryAndCalAreaAndSrc(ctx context.Context) ([]*biz.CaptureSrc, error) {
	reply, err := r.data.cc.ReadAllWithBinaryAndCalAreaAndSrc(ctx, &captureV1.ReadAllWithBinaryAndCalAreaAndSrcRequest{})
	if err != nil {
		return nil, captureV1.ErrorReadDeviceError("read all devices failed, err: %v", err)
	}

	var rets []*biz.CaptureSrc
	for _, data := range reply.GetData() {
		rets = append(rets, &biz.CaptureSrc{
			Proc:   data.GetImageSrc().GetProc(),
			Src:    data.GetImageSrc().GetSrc(),
			Pixels: data.GetPixels(),
			Area:   data.GetArea(),
		})
	}

	return rets, nil
}

func (r *captureRepo) FindLogsByMachineId(ctx context.Context, machineId string) ([]*biz.CaptureLog, error) {
	u, err := uuid.Parse(machineId)
	if err != nil {
		return nil, machineV1.ErrorUuidParseFailed("update machine conflict, err: %v", err)
	}

	targets, err := r.data.db.CaptureLog.
		Query().Where(capturelog.MachineIDEQ(u)).
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
			Id:         target.ID,
			MachineId:  target.MachineID.String(),
			Pixels:     target.Pixels,
			Area:       target.Area,
			SrcName:    target.SrcName,
			ProcName:   target.ProcName,
			SrcOssUrl:  target.SrcOssURL,
			ProcOssUrl: target.ProcOssURL,
		})
	}

	return logs, nil
}

func (r *captureRepo) CreateLog(ctx context.Context, captureLog *biz.CaptureLog) (*biz.CaptureLog, error) {
	u, err := uuid.Parse(captureLog.MachineId)
	if err != nil {
		return nil, machineV1.ErrorUuidParseFailed("update machine conflict, err: %v", err)
	}

	po, err := r.data.db.CaptureLog.
		Create().
		SetMachineID(u).
		SetPixels(captureLog.Pixels).
		SetArea(captureLog.Area).
		SetSrcName(captureLog.SrcName).
		SetProcName(captureLog.ProcName).
		SetSrcOssURL(captureLog.SrcOssUrl).
		SetProcOssURL(captureLog.ProcOssUrl).
		Save(ctx)
	if err != nil {
		r.log.Errorf("unknown err: %v", err)
		return nil, machineV1.ErrorUnknownError("unknown err: %v", err)
	}

	return &biz.CaptureLog{
		Id:         po.ID,
		MachineId:  po.MachineID.String(),
		Pixels:     po.Pixels,
		Area:       po.Area,
		SrcName:    po.SrcName,
		ProcName:   po.ProcName,
		SrcOssUrl:  po.SrcOssURL,
		ProcOssUrl: po.ProcOssURL,
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
		Id:         po.ID,
		MachineId:  po.MachineID.String(),
		Pixels:     po.Pixels,
		Area:       po.Area,
		SrcName:    po.SrcName,
		ProcName:   po.ProcName,
		SrcOssUrl:  po.SrcOssURL,
		ProcOssUrl: po.ProcOssURL,
	}, nil
}
