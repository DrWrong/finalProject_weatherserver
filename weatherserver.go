package main

import (
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"github.com/DrWrong/finalProject_weatherserver/handler"
	"github.com/DrWrong/finalProject_weatherserver/thrift_interface"
	"github.com/DrWrong/finalProject_weatherserver/utils"
	log "github.com/Sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type WeatherServer struct {
	ExeDir string
	server *thrift.TSimpleServer
}

func NewWeatherServer(exeDir string) (weatherserver *WeatherServer) {
	weatherserver = &WeatherServer{ExeDir: exeDir}
	weatherserver.init()
	return weatherserver
}

func (self *WeatherServer) init() {
	self.registerSignalHandler()
	log.Info("weather server has been inited")
}

func (s *WeatherServer) registerSignalHandler() {
	go func() {
		for {
			c := make(chan os.Signal)
			signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
			//sig is blocked as c is 没缓冲
			sig := <-c
			log.Infof("Signal %d received", sig)
			s.stop()
			time.Sleep(time.Second)
			os.Exit(0)
		}
	}()
}
func (s *WeatherServer) run() {
	transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	networkAddr := fmt.Sprintf(":%s", utils.IniConf.String("serverport"))
	serverTransport, err := thrift.NewTServerSocket(networkAddr)
	if err != nil {
		log.Error("Error! %s", err)
		os.Exit(1)
	}
	handler := handler.NewWeatherServerImpl()
	processor := thrift_interface.NewCityWeatherInfoServiceProcessor(handler)
	s.server = thrift.NewTSimpleServer4(processor, serverTransport, transportFactory, protocolFactory)
	log.Infof("thrift server in %s", networkAddr)
	err = s.server.Serve()
	if err != nil {
		log.Errorf("Server start error: %s", err)
	}
}

func (s *WeatherServer) stop() {
	log.Info("Weather server is about to stop...")
	s.server.Stop()
	log.Info("Weather server has gone away")
}
