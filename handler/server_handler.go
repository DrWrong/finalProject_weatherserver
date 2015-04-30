package handler

import (
	"fmt"
	"github.com/DrWrong/finalProject_weatherserver/thrift_interface"
	log "github.com/Sirupsen/logrus"
	"github.com/parnurzeal/gorequest"
)

//  a weather server implementation to implement the pre-defined interface
//  it contains a attributes request, it is a go http client.
type WeatherServerImpl struct {
	request *gorequest.SuperAgent
}

//  create a new  weather server implementation instance
func NewWeatherServerImpl() *WeatherServerImpl {
	request := gorequest.New()
	return &WeatherServerImpl{request}
}

//  a commonly log
func (h *WeatherServerImpl) log(commonRequest *thrift_interface.CommonRequest) {
	log.WithFields(log.Fields{
		"requester": commonRequest.Requester,
		"requestId": commonRequest.RequestId,
	}).Info("processing request")
}

// when process ping request, it just return true, with error = nil
func (h *WeatherServerImpl) Ping(
	commonRequest *thrift_interface.CommonRequest) (
	bool, error) {
	h.log(commonRequest)
	return true, nil
}

// request the open weather api to get city weather info.
func (h *WeatherServerImpl) GetCityWeatherInfo(
	commonRequest *thrift_interface.CommonRequest,
	cityInfo *thrift_interface.CityWeatherInfoRequest) (string, error) {
	h.log(commonRequest)
	_, body, _ := h.request.Get(fmt.Sprintf(
		"http://api.openweathermap.org/data/2.5/weather?q=%s", cityInfo.CityName)).End()
	return body, nil
}
