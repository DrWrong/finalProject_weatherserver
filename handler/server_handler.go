package handler

import (
	"fmt"
	"github.com/DrWrong/finalProject_weatherserver/thrift_interface"
	log "github.com/Sirupsen/logrus"
	"github.com/parnurzeal/gorequest"
)

type WeatherServerImpl struct {
	request *gorequest.SuperAgent
}

func NewWeatherServerImpl() *WeatherServerImpl {
	request := gorequest.New()
	return &WeatherServerImpl{request}
}

func (h *WeatherServerImpl) log(commonRequest *thrift_interface.CommonRequest) {
	log.WithFields(log.Fields{
		"requester": commonRequest.Requester,
		"requestId": commonRequest.RequestId,
	}).Info("processing request")
}

func (h *WeatherServerImpl) Ping(
	commonRequest *thrift_interface.CommonRequest) (
	bool, error) {
	h.log(commonRequest)
	return true, nil
}

func (h *WeatherServerImpl) GetCityWeatherInfo(
	commonRequest *thrift_interface.CommonRequest,
	cityInfo *thrift_interface.CityWeatherInfoRequest) (string, error) {
	h.log(commonRequest)
	_, body, _ := h.request.Get(fmt.Sprintf(
		"http://api.openweathermap.org/data/2.5/weather?q=%s", cityInfo.CityName)).End()
	return body, nil
}
