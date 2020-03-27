package client

import (
	"context"
	log "github.com/sirupsen/logrus"
	realtypb "gitlab.bj.sensetime.com/SenseRealEstate/Schema/realty.git"
	"time"
)

func Recognize(c *gin.Context) {
	var resp *realtypb.RecognizeResponse
	var err error
	defer commonDefer(c, resp, err)

	logBuf := c.MustGet("logEntry").(*log.Entry)
	logBuf.Info("CommonProcess begin.")

	req := &realtypb.RecognizeRequest{}
	err = c.MustBindWith(req, binding.FormMultipart)
	if err != nil {
		logBuf.Errorf("MustBindWith err:%v", err)
		return
	}
	logBuf.WithField("request", req).Debug("Recognize")

	client, err := util.RealtyClient()  //返回链接grpc服务器句柄
	if err != nil {
		logBuf.Errorf("get realtyclient err:%v", err)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	resp, err = client.Recognize(ctx, req)
	if err != nil {
		logBuf.Errorf("recognize: %v", err)
		return
	}
}