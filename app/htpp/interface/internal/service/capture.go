package service

import (
	"context"
	"os"
	"time"

	interfaceV1 "github.com/ynu-e-young/apis-go/htpp/htpp/interface/v1"
)

func (s *InterfaceService) saveImage(data []byte) error {
	if err := os.WriteFile(s.dcf.Images.Dir+"/"+time.Now().String()+".jpg", data, 0644); err != nil {
		return err
	}
	return nil
}

func (s *InterfaceService) ReadOne(ctx context.Context, in *interfaceV1.ReadOneRequest) (*interfaceV1.ImageReply, error) {
	capture, err := s.cu.ReadOne(ctx, in.GetId())
	if err != nil {
		return nil, err
	}

	if err = s.saveImage(capture.Data); err != nil {
		return nil, err
	}

	return &interfaceV1.ImageReply{
		Image: &interfaceV1.Image{
			Data: capture.Data,
		}}, nil
}

func (s *InterfaceService) ReadAll(ctx context.Context, _ *interfaceV1.ReadAllRequest) (*interfaceV1.ImagesReply, error) {
	captures, err := s.cu.ReadAll(ctx)
	if err != nil {
		return nil, err
	}

	var images []*interfaceV1.Image
	for _, capture := range captures {
		if err = s.saveImage(capture.Data); err != nil {
			return nil, err
		}

		images = append(images, &interfaceV1.Image{
			Data: capture.Data,
		})
	}

	return &interfaceV1.ImagesReply{
		Images: images,
	}, nil
}

func (s *InterfaceService) ReadOneWithBinary(ctx context.Context, in *interfaceV1.ReadOneWithBinaryRequest) (*interfaceV1.ImageReply, error) {
	capture, err := s.cu.ReadOneWithBinary(ctx, in.GetId())
	if err != nil {
		return nil, err
	}

	if err = s.saveImage(capture.Data); err != nil {
		return nil, err
	}

	return &interfaceV1.ImageReply{
		Image: &interfaceV1.Image{
			Data: capture.Data,
		}}, nil
}

func (s *InterfaceService) ReadAllWithBinary(ctx context.Context, _ *interfaceV1.ReadAllWithBinaryRequest) (*interfaceV1.ImagesReply, error) {
	captures, err := s.cu.ReadAllWithBinary(ctx)
	if err != nil {
		return nil, err
	}

	var images []*interfaceV1.Image
	for _, capture := range captures {
		if err = s.saveImage(capture.Data); err != nil {
			return nil, err
		}

		images = append(images, &interfaceV1.Image{
			Data: capture.Data,
		})
	}

	return &interfaceV1.ImagesReply{
		Images: images,
	}, nil
}

func (s *InterfaceService) ReadOneWithBinaryAndCalArea(ctx context.Context, in *interfaceV1.ReadOneWithBinaryAndCalAreaRequest) (*interfaceV1.ImageWithAreaReply, error) {
	capture, err := s.cu.ReadOneWithBinaryAndCalArea(ctx, in.GetId())
	if err != nil {
		return nil, err
	}

	if err = s.saveImage(capture.Data); err != nil {
		return nil, err
	}

	return &interfaceV1.ImageWithAreaReply{
		Image: &interfaceV1.Image{
			Data: capture.Data,
		},
		Pixels: capture.Pixels,
		Area:   capture.Area,
	}, nil
}

func (s *InterfaceService) ReadAllWithBinaryAndCalArea(ctx context.Context, _ *interfaceV1.ReadAllWithBinaryAndCalAreaRequest) (*interfaceV1.ImagesWithAreaReply, error) {
	captures, err := s.cu.ReadAllWithBinaryAndCalArea(ctx)
	if err != nil {
		return nil, err
	}

	var data []*interfaceV1.ImagesWithAreaReply_Data
	for _, capture := range captures {
		if err = s.saveImage(capture.Data); err != nil {
			return nil, err
		}

		data = append(data, &interfaceV1.ImagesWithAreaReply_Data{
			Image:  &interfaceV1.Image{Data: capture.Data},
			Pixels: capture.Pixels,
			Area:   capture.Area,
		})
	}

	return &interfaceV1.ImagesWithAreaReply{
		Data: data,
	}, nil
}
