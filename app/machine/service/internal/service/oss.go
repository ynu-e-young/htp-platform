package service

import (
	"bytes"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

func (s *MachineService) OssUpload(fileName string, data []byte) {
	client, err := oss.New(s.dcf.Oss.GetEndpoint(), s.dcf.Oss.GetAccessKeyID(), s.dcf.Oss.GetAccessKeySecret())
	if err != nil {
		s.log.Error(err)
		return
	}

	bucket, err := client.Bucket(s.dcf.Oss.GetBucket())
	if err != nil {
		s.log.Error(err)
		return
	}

	err = bucket.PutObject(fileName, bytes.NewReader(data))
	if err != nil {
		s.log.Error(err)
		return
	}
}
