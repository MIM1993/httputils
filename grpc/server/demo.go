package server

import (
	"context"
	"fmt"
	realtypb "gitlab.bj.sensetime.com/SenseGo/Schema/realty"
	"strings"
	"time"
)

//注册所用类
func NewGrpcService() *GrpcService {
	return &GrpcService{}
}

type GrpcService struct{}

//注册的函数,定义在pb文件中
func (gs *GrpcService) Trace(c context.Context, req *realtypb.TraceRequest) (*realtypb.TraceResponse, error) {
	return nil, nil
}
func (gs *GrpcService) Identify(context.Context, *realtypb.IdentifyRequest) (*realtypb.IdentifyResponse, error) {
	return nil, nil
}
func (gs *GrpcService) Recognize(context.Context, *realtypb.RecognizeRequest) (*realtypb.RecognizeResponse, error) {
	return nil, nil
}
func (gs *GrpcService) MultiGroupIdentify(context.Context, *realtypb.MultiGroupIdentifyRequest) (*realtypb.MultiGroupIdentifyResponse, error) {
	return nil, nil
}

func (gs *GrpcService) NVRTaskProxyGET(context.Context, *realtypb.NVRTaskProxyGETRequest) (*realtypb.NVRTaskProxyGETResponse, error) {
	return nil, nil
}
func (gs *GrpcService) NVRTaskProxyPOST(context.Context, *realtypb.NVRTaskProxyPOSTRequest) (*realtypb.NVRTaskProxyPOSTResponse, error) {
	return nil, nil
}
func (gs *GrpcService) MultiIdentify(context.Context, *realtypb.MultiIdentifyRequest) (*realtypb.MultiIdentifyResponse, error) {
	return nil, nil
}
func (gs *GrpcService) IdentifyTrace(context.Context, *realtypb.IdentifyTraceRequest) (*realtypb.IdentifyTraceResponse, error) {
	return nil, nil
}

func (gs *GrpcService) GetTraceListByFlock(context.Context, *realtypb.GetTraceListByFlockRequest) (*realtypb.GetTraceListByFlockResponse, error) {
	return nil, nil
}
func (gs *GrpcService) Summary(context.Context, *realtypb.SummaryRequest) (*realtypb.SummaryResponse, error) {
	return nil, nil
}
func (gs *GrpcService) Trend(context.Context, *realtypb.TrendRequest) (*realtypb.TrendResponse, error) {
	return nil, nil
}
func (gs *GrpcService) Total(context.Context, *realtypb.TotalRequest) (*realtypb.TotalResponse, error) {
	return nil, nil
}
func (gs *GrpcService) GetByHours(context.Context, *realtypb.GetByHoursRequest) (*realtypb.GetByHoursResponse, error) {
	return nil, nil
}
func (gs *GrpcService) GetByDays(context.Context, *realtypb.GetByDaysRequest) (*realtypb.GetByDaysResponse, error) {
	return nil, nil
}
func (gs *GrpcService) FrequencyStatistics(context.Context, *realtypb.FrequencyStatisticsRequest) (*realtypb.FrequencyStatisticsResponse, error) {
	return nil, nil
}
func (gs *GrpcService) GetCameraInfo(context.Context, *realtypb.GetCameraInfoRequest) (*realtypb.GetCameraInfoResponse, error) {
	return nil, nil
}
func (gs *GrpcService) GetStaffAttendance(context.Context, *realtypb.GetStaffAttendanceRequest) (*realtypb.GetStaffAttendanceResponse, error) {
	return nil, nil
}
func (gs *GrpcService) FaceCheckQuality(context.Context, *realtypb.FaceCheckQualityRequest) (*realtypb.FaceCheckQualityResponse, error) {
	return nil, nil
}

