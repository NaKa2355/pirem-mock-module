package MockDevices

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/NaKa2355/pirem/pkg/driver_module/v1"
)

type MockIRData struct {
	CarrierFreqKiloHz uint32   `json:"carrier_freq_kilo_hz"`
	PluseNanoSec      []uint32 `json:"pluse_nano_sec"`
}

type Config struct {
	CanSend             bool       `json:"can_send"`
	CanReceive          bool       `json:"can_receive"`
	ReceivingIrData     MockIRData `json:"receiving_ir_data"`
	ReceiveTimeMs       uint32     `json:"receive_time_ms"`
	SendTimeMs          uint32     `json:"send_time_ms"`
	FirmwareVersion     string     `json:"firmware_version"`
	DriverVersion       string     `json:"driver_version"`
	SendErrorCode       string     `json:"send_error_code"`
	SendErrorMessage    string     `json:"send_error_message"`
	ReceiveErrorCode    string     `json:"receive_error_code"`
	ReceiveErrorMessage string     `json:"receive_error_message"`
}

type Module struct{}

type MockDriver struct {
	CanSend             bool
	CanReceive          bool
	ReceivingIrData     MockIRData
	ReceiveTime         time.Duration
	SendTime            time.Duration
	FirmwareVersion     string
	DriverVersion       string
	SendErrorCode       string
	SendErrorMessage    string
	ReceiveErrorCode    string
	ReceiveErrorMessage string
}

func (p *Module) LoadDevice(conf json.RawMessage) (driver_module.Device, error) {
	config := Config{}
	err := json.Unmarshal(conf, &config)
	d := &MockDriver{
		CanSend:    config.CanSend,
		CanReceive: config.CanReceive,
		ReceivingIrData: MockIRData{
			CarrierFreqKiloHz: config.ReceivingIrData.CarrierFreqKiloHz,
			PluseNanoSec:      config.ReceivingIrData.PluseNanoSec,
		},
		FirmwareVersion:     config.FirmwareVersion,
		DriverVersion:       config.DriverVersion,
		ReceiveTime:         time.Duration(time.Millisecond * time.Duration(config.ReceiveTimeMs)),
		SendTime:            time.Duration(time.Millisecond * time.Duration(config.SendTimeMs)),
		ReceiveErrorCode:    config.ReceiveErrorCode,
		ReceiveErrorMessage: config.ReceiveErrorMessage,
		SendErrorCode:       config.SendErrorCode,
		SendErrorMessage:    config.SendErrorMessage,
	}
	return d, err
}

func convertError(errCode string, errMessage string) error {
	switch errCode {
	case "invaild_input":
		return driver_module.WrapErr(driver_module.CodeInvaildInput, fmt.Errorf(errMessage))
	case "timeout":
		return driver_module.WrapErr(driver_module.CodeTimeout, fmt.Errorf(errMessage))
	case "busy":
		return driver_module.WrapErr(driver_module.CodeBusy, fmt.Errorf(errMessage))
	case "unknown":
		return driver_module.WrapErr(driver_module.CodeUnknown, fmt.Errorf(errMessage))
	default:
		return nil
	}
}

func (m *MockDriver) SendIR(ctx context.Context, irdata *driver_module.IRData) error {
	time.Sleep(m.SendTime)
	return convertError(m.SendErrorCode, m.SendErrorMessage)
}

func (m *MockDriver) ReceiveIR(ctx context.Context) (*driver_module.IRData, error) {
	irdata := &driver_module.IRData{
		CarrierFreqKiloHz: m.ReceivingIrData.CarrierFreqKiloHz,
		PluseNanoSec:      m.ReceivingIrData.PluseNanoSec,
	}
	time.Sleep(m.ReceiveTime)
	return irdata, convertError(m.ReceiveErrorCode, m.ReceiveErrorMessage)
}

func (m *MockDriver) GetInfo(ctx context.Context) (*driver_module.DeviceInfo, error) {
	return &driver_module.DeviceInfo{
		DriverVersion:   m.DriverVersion,
		FirmwareVersion: m.FirmwareVersion,
	}, nil
}

func (m *MockDriver) Drop() error {
	return nil
}
