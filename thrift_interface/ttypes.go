// Autogenerated by Thrift Compiler (0.9.1)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package thrift_interface

import (
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"math"
)

// (needed to ensure safety because of naive import list construction.)
var _ = math.MinInt32
var _ = thrift.ZERO
var _ = fmt.Printf

var GoUnusedProtection__ int

type CommonRequest struct {
	Requester string `thrift:"requester,1"`
	RequestId int32  `thrift:"requestId,2"`
}

func NewCommonRequest() *CommonRequest {
	return &CommonRequest{}
}

func (p *CommonRequest) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return fmt.Errorf("%T read error", p)
	}
	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return fmt.Errorf("%T field %d read error: %s", p, fieldId, err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.readField1(iprot); err != nil {
				return err
			}
		case 2:
			if err := p.readField2(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return fmt.Errorf("%T read struct end error: %s", p, err)
	}
	return nil
}

func (p *CommonRequest) readField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return fmt.Errorf("error reading field 1: %s")
	} else {
		p.Requester = v
	}
	return nil
}

func (p *CommonRequest) readField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return fmt.Errorf("error reading field 2: %s")
	} else {
		p.RequestId = v
	}
	return nil
}

func (p *CommonRequest) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("CommonRequest"); err != nil {
		return fmt.Errorf("%T write struct begin error: %s", p, err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := p.writeField2(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return fmt.Errorf("%T write field stop error: %s", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return fmt.Errorf("%T write struct stop error: %s", err)
	}
	return nil
}

func (p *CommonRequest) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("requester", thrift.STRING, 1); err != nil {
		return fmt.Errorf("%T write field begin error 1:requester: %s", p, err)
	}
	if err := oprot.WriteString(string(p.Requester)); err != nil {
		return fmt.Errorf("%T.requester (1) field write error: %s", p)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 1:requester: %s", p, err)
	}
	return err
}

func (p *CommonRequest) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("requestId", thrift.I32, 2); err != nil {
		return fmt.Errorf("%T write field begin error 2:requestId: %s", p, err)
	}
	if err := oprot.WriteI32(int32(p.RequestId)); err != nil {
		return fmt.Errorf("%T.requestId (2) field write error: %s", p)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 2:requestId: %s", p, err)
	}
	return err
}

func (p *CommonRequest) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("CommonRequest(%+v)", *p)
}

type CityWeatherInfoRequest struct {
	CityName string `thrift:"cityName,1"`
}

func NewCityWeatherInfoRequest() *CityWeatherInfoRequest {
	return &CityWeatherInfoRequest{}
}

func (p *CityWeatherInfoRequest) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return fmt.Errorf("%T read error", p)
	}
	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return fmt.Errorf("%T field %d read error: %s", p, fieldId, err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.readField1(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return fmt.Errorf("%T read struct end error: %s", p, err)
	}
	return nil
}

func (p *CityWeatherInfoRequest) readField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return fmt.Errorf("error reading field 1: %s")
	} else {
		p.CityName = v
	}
	return nil
}

func (p *CityWeatherInfoRequest) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("CityWeatherInfoRequest"); err != nil {
		return fmt.Errorf("%T write struct begin error: %s", p, err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return fmt.Errorf("%T write field stop error: %s", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return fmt.Errorf("%T write struct stop error: %s", err)
	}
	return nil
}

func (p *CityWeatherInfoRequest) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("cityName", thrift.STRING, 1); err != nil {
		return fmt.Errorf("%T write field begin error 1:cityName: %s", p, err)
	}
	if err := oprot.WriteString(string(p.CityName)); err != nil {
		return fmt.Errorf("%T.cityName (1) field write error: %s", p)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 1:cityName: %s", p, err)
	}
	return err
}

func (p *CityWeatherInfoRequest) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("CityWeatherInfoRequest(%+v)", *p)
}
