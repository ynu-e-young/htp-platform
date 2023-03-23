package data

import (
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"gocv.io/x/gocv"

	"github.com/ynu-e-young/htp-platform/app/capture/service/internal/conf"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewVideoCaptures, NewCaptureRepo)

// Data .
type Data struct {
	cameras  []*gocv.VideoCapture
	captures []*gocv.Mat
	stop     <-chan struct{}
}

func NewData(cameras []*gocv.VideoCapture, logger log.Logger) (*Data, func(), error) {
	helper := log.NewHelper(log.With(logger, "module", "cv-service/data"))

	stop := make(chan struct{})
	var captures []*gocv.Mat
	for _, device := range cameras {
		mat := gocv.NewMat()
		captures = append(captures, &mat)
		go func(device *gocv.VideoCapture) {
			for {
				device.Read(&mat)
				time.Sleep(17 * time.Millisecond)
			}
		}(device)
	}

	d := &Data{
		cameras:  cameras,
		captures: captures,
		stop:     stop,
	}

	return d, func() {
		for i, camera := range cameras {
			if err := camera.Close(); err != nil {
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
