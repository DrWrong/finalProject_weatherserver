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

// a weather server class
// it contains two attributes
//  + ExeDir the execute directory of the server
//  + server the thrift server
type WeatherServer struct {
	ExeDir string
	server *thrift.TSimpleServer
}

//  create a new weather server and have it initialized
func NewWeatherServer(exeDir string) (weatherserver *WeatherServer) {
	weatherserver = &WeatherServer{ExeDir: exeDir}
	weatherserver.init()
	return weatherserver
}

// initialize the weather server and register a signal handler to handle terminate signal
func (self *WeatherServer) init() {
	self.registerSignalHandler()
	log.Info("weather server has been inited")
}

//  when received the terminated signal stop the server
func (s *WeatherServer) registerSignalHandler() {
	go func() {
		for {
			c := make(chan os.Signal)
			signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
			sig := <-c
			log.Infof("Signal %d received", sig)
			s.stop()
			time.Sleep(time.Second)
			os.Exit(0)
		}
	}()
}

// run the server
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

// stop the server
func (s *WeatherServer) stop() {
	log.Info("Weather server is about to stop...")
	s.server.Stop()
	log.Info("Weather server has gone away")
}
