package service

import (
	"context"
	"gocv.io/x/gocv"
	v1 "htp-platform/api/capture/service/v1"
)

func (s *CaptureService) mat2bytes(mat *gocv.Mat) ([]byte, error) {
	buffer, err := gocv.IMEncode(".jpg", *mat)
	if err != nil {
		return nil, err
	}
	return buffer.GetBytes(), nil
}

func (s *CaptureService) ReadOne(ctx context.Context, in *v1.ReadOneRequest) (*v1.ImageReply, error) {
	capture, err := s.uu.ReadOne(ctx, int(in.GetId()))
	if err != nil {
		return nil, v1.ErrorReadDeviceError("read device %d failed, err: %v", in.GetId(), err)
	}

	bytes, err := s.mat2bytes(capture.Mat)
	if err != nil {
		return nil, v1.ErrorConvertError("convert mat to bytes failed, err: %v", err)
	}

	return &v1.ImageReply{Image: &v1.Image{
		Data: bytes,
	}}, nil
}

func (s *CaptureService) ReadAll(ctx context.Context, _ *v1.ReadAllRequest) (*v1.ImagesReply, error) {
	captures, err := s.uu.ReadAll(ctx)
	if err != nil {
		return nil, v1.ErrorReadDeviceError("read all devices failed, err: %v", err)
	}

	var images []*v1.Image
	for _, capture := range captures {
		bytes, err := s.mat2bytes(capture.Mat)
		if err != nil {
			return nil, v1.ErrorConvertError("convert mat to bytes failed, err: %v", err)
		}
		images = append(images, &v1.Image{Data: bytes})
	}

	return &v1.ImagesReply{
		Images: images,
	}, nil
}

func (s *CaptureService) ReadOneWithBinary(ctx context.Context, in *v1.ReadOneWithBinaryRequest) (*v1.ImageReply, error) {
	capture, err := s.uu.ReadOne(ctx, int(in.GetId()))
	if err != nil {
		return nil, v1.ErrorReadDeviceError("read device %d failed, err: %v", in.GetId(), err)
	}

	binary, err := s.uu.Binary(ctx, capture)
	if err != nil {
		return nil, v1.ErrorBinaryError("binary mat failed, err: %v", err)
	}

	bytes, err := s.mat2bytes(binary.Mat)
	if err != nil {
		return nil, v1.ErrorConvertError("convert binary mat to bytes failed, err: %v", err)
	}

	return &v1.ImageReply{Image: &v1.Image{
		Data: bytes,
	}}, nil
}

func (s *CaptureService) ReadAllWithBinary(ctx context.Context, _ *v1.ReadAllWithBinaryRequest) (*v1.ImagesReply, error) {
	captures, err := s.uu.ReadAll(ctx)
	if err != nil {
		return nil, v1.ErrorReadDeviceError("read all devices failed, err: %v", err)
	}

	var images []*v1.Image
	for _, capture := range captures {
		binary, err := s.uu.Binary(ctx, capture)
		if err != nil {
			return nil, v1.ErrorBinaryError("binary mat failed, err: %v", err)
		}

		bytes, err := s.mat2bytes(binary.Mat)
		if err != nil {
			return nil, v1.ErrorConvertError("convert binary mat to bytes failed, err: %v", err)
		}
		images = append(images, &v1.Image{Data: bytes})
	}

	return &v1.ImagesReply{
		Images: images,
	}, nil
}

func (s *CaptureService) ReadOneWithBinaryAndCalArea(ctx context.Context, in *v1.ReadOneWithBinaryAndCalAreaRequest) (*v1.ImageWithAreaReply, error) {
	capture, err := s.uu.ReadOne(ctx, int(in.GetId()))
	if err != nil {
		return nil, v1.ErrorReadDeviceError("read device %d failed, err: %v", in.GetId(), err)
	}

	binary, err := s.uu.Binary(ctx, capture)
	if err != nil {
		return nil, v1.ErrorBinaryError("binary mat failed, err: %v", err)
	}

	pixels := s.uu.Pixels(s.uu.BaweraOpen(*binary.Mat, 180))

	area := s.uu.Area(float64(pixels), 50)

	bytes, err := s.mat2bytes(binary.Mat)
	if err != nil {
		return nil, v1.ErrorConvertError("convert binary mat to bytes failed, err: %v", err)
	}

	return &v1.ImageWithAreaReply{
		Image: &v1.Image{
			Data: bytes,
		},
		Pixels: int64(pixels),
		Area:   area,
	}, nil
}

func (s *CaptureService) ReadAllWithBinaryAndCalArea(ctx context.Context, _ *v1.ReadAllWithBinaryAndCalAreaRequest) (*v1.ImagesWithAreaReply, error) {
	captures, err := s.uu.ReadAll(ctx)
	if err != nil {
		return nil, v1.ErrorReadDeviceError("read all devices failed, err: %v", err)
	}

	var data []*v1.ImagesWithAreaReply_Data
	for _, capture := range captures {
		binary, err := s.uu.Binary(ctx, capture)
		if err != nil {
			return nil, v1.ErrorBinaryError("binary mat failed, err: %v", err)
		}

		pixels := s.uu.Pixels(s.uu.BaweraOpen(*binary.Mat, 180))

		area := s.uu.Area(float64(pixels), 50)

		bytes, err := s.mat2bytes(binary.Mat)
		if err != nil {
			return nil, v1.ErrorConvertError("convert binary mat to bytes failed, err: %v", err)
		}

		data = append(data, &v1.ImagesWithAreaReply_Data{
			Image:  &v1.Image{Data: bytes},
			Pixels: int64(pixels),
			Area:   area,
		})
	}

	return &v1.ImagesWithAreaReply{
		Data: data,
	}, nil
}
