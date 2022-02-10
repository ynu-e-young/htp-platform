package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"gocv.io/x/gocv"
	"htp-platform/app/capture/service/internal/conf"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewVideoCaptures, NewCaptureRepo)

// Data .
type Data struct {
	captures []*gocv.VideoCapture
}

func NewData(captures []*gocv.VideoCapture, logger log.Logger) (*Data, func(), error) {
	helper := log.NewHelper(log.With(logger, "module", "cv-service/data"))

	for i, capture := range captures {
		if !capture.IsOpened() {
			helper.Fatalf("recheck device: %d is not open", i)
		}
	}

	d := &Data{
		captures: captures,
	}
	return d, func() {
		for i, capture := range captures {
			if err := capture.Close(); err != nil {
				helper.Errorf("capture %d close failed with error: %v", i, err)
			}
		}
	}, nil
}

func NewVideoCaptures(conf *conf.Data, logger log.Logger) []*gocv.VideoCapture {
	helper := log.NewHelper(log.With(logger, "module", "cv-service/data/captures"))

	var captures []*gocv.VideoCapture
	for _, device := range conf.Capture.Devices {
		capture, err := gocv.VideoCaptureDevice(int(device))
		if err != nil {
			helper.Fatalf("error opening device: %d", device)
		}
		captures = append(captures, capture)
	}

	return captures
}
