package service

import (
	"context"
	interfaceV1 "htp-platform/api/htpp/interface/v1"
	"os"
	"time"
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
