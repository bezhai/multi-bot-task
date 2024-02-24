// Code generated by thriftgo (0.3.6). DO NOT EDIT.

package data_trans

import (
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
)

type UploadTosFileToLarkRequest struct {
	PixivAddr string `thrift:"pixiv_addr,1,required" form:"pixiv_addr,required" json:"pixiv_addr,required" query:"pixiv_addr,required"`
}

func NewUploadTosFileToLarkRequest() *UploadTosFileToLarkRequest {
	return &UploadTosFileToLarkRequest{}
}

func (p *UploadTosFileToLarkRequest) GetPixivAddr() (v string) {
	return p.PixivAddr
}

var fieldIDToName_UploadTosFileToLarkRequest = map[int16]string{
	1: "pixiv_addr",
}

func (p *UploadTosFileToLarkRequest) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16
	var issetPixivAddr bool = false

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}

		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRING {
				if err = p.ReadField1(iprot); err != nil {
					goto ReadFieldError
				}
				issetPixivAddr = true
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		default:
			if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		}
		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	if !issetPixivAddr {
		fieldId = 1
		goto RequiredFieldNotSetError
	}
	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_UploadTosFileToLarkRequest[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
RequiredFieldNotSetError:
	return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("required field %s is not set", fieldIDToName_UploadTosFileToLarkRequest[fieldId]))
}

func (p *UploadTosFileToLarkRequest) ReadField1(iprot thrift.TProtocol) error {

	if v, err := iprot.ReadString(); err != nil {
		return err
	} else {
		p.PixivAddr = v
	}
	return nil
}

func (p *UploadTosFileToLarkRequest) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("UploadTosFileToLarkRequest"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField1(oprot); err != nil {
			fieldId = 1
			goto WriteFieldError
		}
	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldError:
	return thrift.PrependError(fmt.Sprintf("%T write field %d error: ", p, fieldId), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *UploadTosFileToLarkRequest) writeField1(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("pixiv_addr", thrift.STRING, 1); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteString(p.PixivAddr); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 end error: ", p), err)
}

func (p *UploadTosFileToLarkRequest) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("UploadTosFileToLarkRequest(%+v)", *p)

}

type UploadTosFileToLarkResponseData struct {
	ImageKey string `thrift:"image_key,1,required" form:"image_key,required" json:"image_key,required" query:"image_key,required"`
	Width    int32  `thrift:"width,2,required" form:"width,required" json:"width,required" query:"width,required"`
	Height   int32  `thrift:"height,3,required" form:"height,required" json:"height,required" query:"height,required"`
}

func NewUploadTosFileToLarkResponseData() *UploadTosFileToLarkResponseData {
	return &UploadTosFileToLarkResponseData{}
}

func (p *UploadTosFileToLarkResponseData) GetImageKey() (v string) {
	return p.ImageKey
}

func (p *UploadTosFileToLarkResponseData) GetWidth() (v int32) {
	return p.Width
}

func (p *UploadTosFileToLarkResponseData) GetHeight() (v int32) {
	return p.Height
}

var fieldIDToName_UploadTosFileToLarkResponseData = map[int16]string{
	1: "image_key",
	2: "width",
	3: "height",
}

func (p *UploadTosFileToLarkResponseData) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16
	var issetImageKey bool = false
	var issetWidth bool = false
	var issetHeight bool = false

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}

		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRING {
				if err = p.ReadField1(iprot); err != nil {
					goto ReadFieldError
				}
				issetImageKey = true
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		case 2:
			if fieldTypeId == thrift.I32 {
				if err = p.ReadField2(iprot); err != nil {
					goto ReadFieldError
				}
				issetWidth = true
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		case 3:
			if fieldTypeId == thrift.I32 {
				if err = p.ReadField3(iprot); err != nil {
					goto ReadFieldError
				}
				issetHeight = true
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		default:
			if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		}
		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	if !issetImageKey {
		fieldId = 1
		goto RequiredFieldNotSetError
	}

	if !issetWidth {
		fieldId = 2
		goto RequiredFieldNotSetError
	}

	if !issetHeight {
		fieldId = 3
		goto RequiredFieldNotSetError
	}
	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_UploadTosFileToLarkResponseData[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
RequiredFieldNotSetError:
	return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("required field %s is not set", fieldIDToName_UploadTosFileToLarkResponseData[fieldId]))
}

func (p *UploadTosFileToLarkResponseData) ReadField1(iprot thrift.TProtocol) error {

	if v, err := iprot.ReadString(); err != nil {
		return err
	} else {
		p.ImageKey = v
	}
	return nil
}
func (p *UploadTosFileToLarkResponseData) ReadField2(iprot thrift.TProtocol) error {

	if v, err := iprot.ReadI32(); err != nil {
		return err
	} else {
		p.Width = v
	}
	return nil
}
func (p *UploadTosFileToLarkResponseData) ReadField3(iprot thrift.TProtocol) error {

	if v, err := iprot.ReadI32(); err != nil {
		return err
	} else {
		p.Height = v
	}
	return nil
}

func (p *UploadTosFileToLarkResponseData) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("UploadTosFileToLarkResponseData"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField1(oprot); err != nil {
			fieldId = 1
			goto WriteFieldError
		}
		if err = p.writeField2(oprot); err != nil {
			fieldId = 2
			goto WriteFieldError
		}
		if err = p.writeField3(oprot); err != nil {
			fieldId = 3
			goto WriteFieldError
		}
	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldError:
	return thrift.PrependError(fmt.Sprintf("%T write field %d error: ", p, fieldId), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *UploadTosFileToLarkResponseData) writeField1(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("image_key", thrift.STRING, 1); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteString(p.ImageKey); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 end error: ", p), err)
}

func (p *UploadTosFileToLarkResponseData) writeField2(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("width", thrift.I32, 2); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteI32(p.Width); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 2 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 2 end error: ", p), err)
}

func (p *UploadTosFileToLarkResponseData) writeField3(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("height", thrift.I32, 3); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteI32(p.Height); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 3 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 3 end error: ", p), err)
}

func (p *UploadTosFileToLarkResponseData) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("UploadTosFileToLarkResponseData(%+v)", *p)

}

type UploadTosFileToLarkResponse struct {
	Code int32                            `thrift:"code,1,required" form:"code,required" json:"code,required" query:"code,required"`
	Msg  string                           `thrift:"msg,2,required" form:"msg,required" json:"msg,required" query:"msg,required"`
	Data *UploadTosFileToLarkResponseData `thrift:"data,3,optional" form:"data" json:"data,omitempty" query:"data"`
}

func NewUploadTosFileToLarkResponse() *UploadTosFileToLarkResponse {
	return &UploadTosFileToLarkResponse{}
}

func (p *UploadTosFileToLarkResponse) GetCode() (v int32) {
	return p.Code
}

func (p *UploadTosFileToLarkResponse) GetMsg() (v string) {
	return p.Msg
}

var UploadTosFileToLarkResponse_Data_DEFAULT *UploadTosFileToLarkResponseData

func (p *UploadTosFileToLarkResponse) GetData() (v *UploadTosFileToLarkResponseData) {
	if !p.IsSetData() {
		return UploadTosFileToLarkResponse_Data_DEFAULT
	}
	return p.Data
}

var fieldIDToName_UploadTosFileToLarkResponse = map[int16]string{
	1: "code",
	2: "msg",
	3: "data",
}

func (p *UploadTosFileToLarkResponse) IsSetData() bool {
	return p.Data != nil
}

func (p *UploadTosFileToLarkResponse) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16
	var issetCode bool = false
	var issetMsg bool = false

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}

		switch fieldId {
		case 1:
			if fieldTypeId == thrift.I32 {
				if err = p.ReadField1(iprot); err != nil {
					goto ReadFieldError
				}
				issetCode = true
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		case 2:
			if fieldTypeId == thrift.STRING {
				if err = p.ReadField2(iprot); err != nil {
					goto ReadFieldError
				}
				issetMsg = true
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		case 3:
			if fieldTypeId == thrift.STRUCT {
				if err = p.ReadField3(iprot); err != nil {
					goto ReadFieldError
				}
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		default:
			if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		}
		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	if !issetCode {
		fieldId = 1
		goto RequiredFieldNotSetError
	}

	if !issetMsg {
		fieldId = 2
		goto RequiredFieldNotSetError
	}
	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_UploadTosFileToLarkResponse[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
RequiredFieldNotSetError:
	return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("required field %s is not set", fieldIDToName_UploadTosFileToLarkResponse[fieldId]))
}

func (p *UploadTosFileToLarkResponse) ReadField1(iprot thrift.TProtocol) error {

	if v, err := iprot.ReadI32(); err != nil {
		return err
	} else {
		p.Code = v
	}
	return nil
}
func (p *UploadTosFileToLarkResponse) ReadField2(iprot thrift.TProtocol) error {

	if v, err := iprot.ReadString(); err != nil {
		return err
	} else {
		p.Msg = v
	}
	return nil
}
func (p *UploadTosFileToLarkResponse) ReadField3(iprot thrift.TProtocol) error {
	p.Data = NewUploadTosFileToLarkResponseData()
	if err := p.Data.Read(iprot); err != nil {
		return err
	}
	return nil
}

func (p *UploadTosFileToLarkResponse) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("UploadTosFileToLarkResponse"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField1(oprot); err != nil {
			fieldId = 1
			goto WriteFieldError
		}
		if err = p.writeField2(oprot); err != nil {
			fieldId = 2
			goto WriteFieldError
		}
		if err = p.writeField3(oprot); err != nil {
			fieldId = 3
			goto WriteFieldError
		}
	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldError:
	return thrift.PrependError(fmt.Sprintf("%T write field %d error: ", p, fieldId), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *UploadTosFileToLarkResponse) writeField1(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("code", thrift.I32, 1); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteI32(p.Code); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 end error: ", p), err)
}

func (p *UploadTosFileToLarkResponse) writeField2(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("msg", thrift.STRING, 2); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteString(p.Msg); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 2 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 2 end error: ", p), err)
}

func (p *UploadTosFileToLarkResponse) writeField3(oprot thrift.TProtocol) (err error) {
	if p.IsSetData() {
		if err = oprot.WriteFieldBegin("data", thrift.STRUCT, 3); err != nil {
			goto WriteFieldBeginError
		}
		if err := p.Data.Write(oprot); err != nil {
			return err
		}
		if err = oprot.WriteFieldEnd(); err != nil {
			goto WriteFieldEndError
		}
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 3 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 3 end error: ", p), err)
}

func (p *UploadTosFileToLarkResponse) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("UploadTosFileToLarkResponse(%+v)", *p)

}

type ProxyRequest struct {
	URL     string `thrift:"url,1,required" form:"url,required" json:"url,required" query:"url,required"`
	Referer string `thrift:"referer,2,required" form:"referer,required" json:"referer,required" query:"referer,required"`
}

func NewProxyRequest() *ProxyRequest {
	return &ProxyRequest{}
}

func (p *ProxyRequest) GetURL() (v string) {
	return p.URL
}

func (p *ProxyRequest) GetReferer() (v string) {
	return p.Referer
}

var fieldIDToName_ProxyRequest = map[int16]string{
	1: "url",
	2: "referer",
}

func (p *ProxyRequest) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16
	var issetURL bool = false
	var issetReferer bool = false

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}

		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRING {
				if err = p.ReadField1(iprot); err != nil {
					goto ReadFieldError
				}
				issetURL = true
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		case 2:
			if fieldTypeId == thrift.STRING {
				if err = p.ReadField2(iprot); err != nil {
					goto ReadFieldError
				}
				issetReferer = true
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		default:
			if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		}
		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	if !issetURL {
		fieldId = 1
		goto RequiredFieldNotSetError
	}

	if !issetReferer {
		fieldId = 2
		goto RequiredFieldNotSetError
	}
	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_ProxyRequest[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
RequiredFieldNotSetError:
	return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("required field %s is not set", fieldIDToName_ProxyRequest[fieldId]))
}

func (p *ProxyRequest) ReadField1(iprot thrift.TProtocol) error {

	if v, err := iprot.ReadString(); err != nil {
		return err
	} else {
		p.URL = v
	}
	return nil
}
func (p *ProxyRequest) ReadField2(iprot thrift.TProtocol) error {

	if v, err := iprot.ReadString(); err != nil {
		return err
	} else {
		p.Referer = v
	}
	return nil
}

func (p *ProxyRequest) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("ProxyRequest"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField1(oprot); err != nil {
			fieldId = 1
			goto WriteFieldError
		}
		if err = p.writeField2(oprot); err != nil {
			fieldId = 2
			goto WriteFieldError
		}
	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldError:
	return thrift.PrependError(fmt.Sprintf("%T write field %d error: ", p, fieldId), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *ProxyRequest) writeField1(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("url", thrift.STRING, 1); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteString(p.URL); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 end error: ", p), err)
}

func (p *ProxyRequest) writeField2(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("referer", thrift.STRING, 2); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteString(p.Referer); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 2 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 2 end error: ", p), err)
}

func (p *ProxyRequest) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("ProxyRequest(%+v)", *p)

}

type ProxyResponse struct {
}

func NewProxyResponse() *ProxyResponse {
	return &ProxyResponse{}
}

var fieldIDToName_ProxyResponse = map[int16]string{}

func (p *ProxyResponse) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		if err = iprot.Skip(fieldTypeId); err != nil {
			goto SkipFieldTypeError
		}
		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
SkipFieldTypeError:
	return thrift.PrependError(fmt.Sprintf("%T skip field type %d error", p, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *ProxyResponse) Write(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteStructBegin("ProxyResponse"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *ProxyResponse) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("ProxyResponse(%+v)", *p)

}