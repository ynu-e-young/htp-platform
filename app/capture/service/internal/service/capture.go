package service

import (
	"context"

	captureV1 "github.com/ynu-e-young/apis-go/htpp/capture/service/v1"
	"gocv.io/x/gocv"
)

func (s *CaptureService) mat2bytes(mat *gocv.Mat) ([]byte, error) {
	buffer, err := gocv.IMEncode(".jpg", *mat)
	if err != nil {
		return nil, err
	}
	return buffer.GetBytes(), nil
}

func (s *CaptureService) ReadOne(ctx context.Context, in *captureV1.ReadOneRequest) (*captureV1.ImageReply, error) {
	capture, err := s.uu.ReadOne(ctx, int(in.GetId()))
	if err != nil {
		return nil, captureV1.ErrorReadDeviceError("read device %d failed, err: %v", in.GetId(), err)
	}

	bytes, err := s.mat2bytes(capture.Mat)
	if err != nil {
		return nil, captureV1.ErrorConvertError("convert mat to bytes failed, err: %v", err)
	}

	return &captureV1.ImageReply{Image: &captureV1.Image{
		Data: bytes,
	}}, nil
}

func (s *CaptureService) ReadAll(ctx context.Context, _ *captureV1.ReadAllRequest) (*captureV1.ImagesReply, error) {
	captures, err := s.uu.ReadAll(ctx)
	if err != nil {
		return nil, captureV1.ErrorReadDeviceError("read all devices failed, err: %v", err)
	}

	var images []*captureV1.Image
	for _, capture := range captures {
		bytes, err := s.mat2bytes(capture.Mat)
		if err != nil {
			return nil, captureV1.ErrorConvertError("convert mat to bytes failed, err: %v", err)
		}
		images = append(images, &captureV1.Image{Data: bytes})
	}

	return &captureV1.ImagesReply{
		Images: images,
	}, nil
}

func (s *CaptureService) ReadOneWithBinary(ctx context.Context, in *captureV1.ReadOneWithBinaryRequest) (*captureV1.ImageReply, error) {
	capture, err := s.uu.ReadOne(ctx, int(in.GetId()))
	if err != nil {
		return nil, captureV1.ErrorReadDeviceError("read device %d failed, err: %v", in.GetId(), err)
	}

	binary, err := s.uu.Binary(ctx, capture)
	if err != nil {
		return nil, captureV1.ErrorBinaryError("binary mat failed, err: %v", err)
	}

	bytes, err := s.mat2bytes(binary.Mat)
	if err != nil {
		return nil, captureV1.ErrorConvertError("convert binary mat to bytes failed, err: %v", err)
	}

	return &captureV1.ImageReply{Image: &captureV1.Image{
		Data: bytes,
	}}, nil
}

func (s *CaptureService) ReadOneWithBinaryAndSrc(ctx context.Context, in *captureV1.ReadOneWithBinaryAndSrcRequest) (*captureV1.ImageWithSrcReply, error) {
	capture, err := s.uu.ReadOne(ctx, int(in.GetId()))
	if err != nil {
		return nil, captureV1.ErrorReadDeviceError("read device %d failed, err: %v", in.GetId(), err)
	}

	binary, err := s.uu.Binary(ctx, capture)
	if err != nil {
		return nil, captureV1.ErrorBinaryError("binary mat failed, err: %v", err)
	}

	sBytes, err := s.mat2bytes(capture.Mat)
	if err != nil {
		return nil, captureV1.ErrorConvertError("convert capture mat to bBytes failed, err: %v", err)
	}

	bBytes, err := s.mat2bytes(binary.Mat)
	if err != nil {
		return nil, captureV1.ErrorConvertError("convert binary mat to bBytes failed, err: %v", err)
	}

	return &captureV1.ImageWithSrcReply{
		ImageSrc: &captureV1.ImageSrc{
			Proc: bBytes,
			Src:  sBytes,
		},
	}, nil
}

func (s *CaptureService) ReadAllWithBinary(ctx context.Context, _ *captureV1.ReadAllWithBinaryRequest) (*captureV1.ImagesReply, error) {
	captures, err := s.uu.ReadAll(ctx)
	if err != nil {
		return nil, captureV1.ErrorReadDeviceError("read all devices failed, err: %v", err)
	}

	var images []*captureV1.Image
	for _, capture := range captures {
		binary, err := s.uu.Binary(ctx, capture)
		if err != nil {
			return nil, captureV1.ErrorBinaryError("binary mat failed, err: %v", err)
		}

		bytes, err := s.mat2bytes(binary.Mat)
		if err != nil {
			return nil, captureV1.ErrorConvertError("convert binary mat to bytes failed, err: %v", err)
		}
		images = append(images, &captureV1.Image{Data: bytes})
	}

	return &captureV1.ImagesReply{
		Images: images,
	}, nil
}

func (s *CaptureService) ReadAllWithBinaryAndSrc(ctx context.Context, _ *captureV1.ReadAllWithBinaryAndSrcRequest) (*captureV1.ImagesWithSrcReply, error) {
	captures, err := s.uu.ReadAll(ctx)
	if err != nil {
		return nil, captureV1.ErrorReadDeviceError("read all devices failed, err: %v", err)
	}

	var images []*captureV1.ImageSrc
	for _, capture := range captures {
		binary, err := s.uu.Binary(ctx, capture)
		if err != nil {
			return nil, captureV1.ErrorBinaryError("binary mat failed, err: %v", err)
		}

		sBytes, err := s.mat2bytes(capture.Mat)
		if err != nil {
			return nil, captureV1.ErrorConvertError("convert capture mat to bytes failed, err: %v", err)
		}

		bBytes, err := s.mat2bytes(binary.Mat)
		if err != nil {
			return nil, captureV1.ErrorConvertError("convert binary mat to bytes failed, err: %v", err)
		}

		images = append(images, &captureV1.ImageSrc{
			Proc: bBytes,
			Src:  sBytes,
		})
	}

	return &captureV1.ImagesWithSrcReply{
		ImageSrc: images,
	}, nil
}

func (s *CaptureService) ReadOneWithBinaryAndCalArea(ctx context.Context, in *captureV1.ReadOneWithBinaryAndCalAreaRequest) (*captureV1.ImageWithAreaReply, error) {
	capture, err := s.uu.ReadOne(ctx, int(in.GetId()))
	if err != nil {
		return nil, captureV1.ErrorReadDeviceError("read device %d failed, err: %v", in.GetId(), err)
	}

	binary, err := s.uu.Binary(ctx, capture)
	if err != nil {
		return nil, captureV1.ErrorBinaryError("binary mat failed, err: %v", err)
	}

	pixels := s.uu.Pixels(s.uu.BaweraOpen(*binary.Mat, 180))

	area := s.uu.Area(float64(pixels), 50)

	bytes, err := s.mat2bytes(binary.Mat)
	if err != nil {
		return nil, captureV1.ErrorConvertError("convert binary mat to bytes failed, err: %v", err)
	}

	return &captureV1.ImageWithAreaReply{
		Image: &captureV1.Image{
			Data: bytes,
		},
		Pixels: int64(pixels),
		Area:   area,
	}, nil
}

func (s *CaptureService) ReadOneWithBinaryAndCalAreaAndSrc(ctx context.Context, in *captureV1.ReadOneWithBinaryAndCalAreaAndSrcRequest) (*captureV1.ImageWithAreaAndSrcReply, error) {
	capture, err := s.uu.ReadOne(ctx, int(in.GetId()))
	if err != nil {
		return nil, captureV1.ErrorReadDeviceError("read device %d failed, err: %v", in.GetId(), err)
	}

	binary, err := s.uu.Binary(ctx, capture)
	if err != nil {
		return nil, captureV1.ErrorBinaryError("binary mat failed, err: %v", err)
	}

	pixels := s.uu.Pixels(s.uu.BaweraOpen(*binary.Mat, 180))

	area := s.uu.Area(float64(pixels), 50)

	sBytes, err := s.mat2bytes(capture.Mat)
	if err != nil {
		return nil, captureV1.ErrorConvertError("convert capture mat to bytes failed, err: %v", err)
	}

	bBytes, err := s.mat2bytes(binary.Mat)
	if err != nil {
		return nil, captureV1.ErrorConvertError("convert binary mat to bytes failed, err: %v", err)
	}

	return &captureV1.ImageWithAreaAndSrcReply{
		ImageSrc: &captureV1.ImageSrc{
			Proc: bBytes,
			Src:  sBytes,
		},
		Pixels: int64(pixels),
		Area:   area,
	}, nil
}

func (s *CaptureService) ReadAllWithBinaryAndCalArea(ctx context.Context, _ *captureV1.ReadAllWithBinaryAndCalAreaRequest) (*captureV1.ImagesWithAreaReply, error) {
	captures, err := s.uu.ReadAll(ctx)
	if err != nil {
		return nil, captureV1.ErrorReadDeviceError("read all devices failed, err: %v", err)
	}

	var data []*captureV1.ImagesWithAreaReply_Data
	for _, capture := range captures {
		binary, err := s.uu.Binary(ctx, capture)
		if err != nil {
			return nil, captureV1.ErrorBinaryError("binary mat failed, err: %v", err)
		}

		pixels := s.uu.Pixels(s.uu.BaweraOpen(*binary.Mat, 180))

		area := s.uu.Area(float64(pixels), 50)

		bytes, err := s.mat2bytes(binary.Mat)
		if err != nil {
			return nil, captureV1.ErrorConvertError("convert binary mat to bytes failed, err: %v", err)
		}

		data = append(data, &captureV1.ImagesWithAreaReply_Data{
			Image:  &captureV1.Image{Data: bytes},
			Pixels: int64(pixels),
			Area:   area,
		})
	}

	return &captureV1.ImagesWithAreaReply{
		Data: data,
	}, nil
}

func (s *CaptureService) ReadAllWithBinaryAndCalAreaAndSrc(ctx context.Context, _ *captureV1.ReadAllWithBinaryAndCalAreaAndSrcRequest) (*captureV1.ImagesWithAreaAndSrcReply, error) {
	captures, err := s.uu.ReadAll(ctx)
	if err != nil {
		return nil, captureV1.ErrorReadDeviceError("read all devices failed, err: %v", err)
	}

	var data []*captureV1.ImagesWithAreaAndSrcReply_Data
	for _, capture := range captures {
		binary, err := s.uu.Binary(ctx, capture)
		if err != nil {
			return nil, captureV1.ErrorBinaryError("binary mat failed, err: %v", err)
		}

		pixels := s.uu.Pixels(s.uu.BaweraOpen(*binary.Mat, 180))

		area := s.uu.Area(float64(pixels), 50)

		sBytes, err := s.mat2bytes(capture.Mat)
		if err != nil {
			return nil, captureV1.ErrorConvertError("convert capture mat to bytes failed, err: %v", err)
		}

		bBytes, err := s.mat2bytes(binary.Mat)
		if err != nil {
			return nil, captureV1.ErrorConvertError("convert binary mat to bytes failed, err: %v", err)
		}

		data = append(data, &captureV1.ImagesWithAreaAndSrcReply_Data{
			ImageSrc: &captureV1.ImageSrc{
				Proc: bBytes,
				Src:  sBytes,
			},
			Pixels: int64(pixels),
			Area:   area,
		})
	}

	return &captureV1.ImagesWithAreaAndSrcReply{Data: data}, nil
}
