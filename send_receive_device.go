package MockDevices

import (
	"context"
	"time"

	"github.com/NaKa2355/pirem/pkg/driver_module/v1"
)

type SendReceiveDevice struct {
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

func NewSendReceiveDevice(config *Config) SendReceiveDevice {
	return SendReceiveDevice{
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
}

func (m *SendReceiveDevice) SendIR(ctx context.Context, irdata *driver_module.IRData) error {
	time.Sleep(m.SendTime)
	return convertError(m.SendErrorCode, m.SendErrorMessage)
}

func (m *SendReceiveDevice) ReceiveIR(ctx context.Context) (*driver_module.IRData, error) {
	irdata := &driver_module.IRData{
		CarrierFreqKiloHz: m.ReceivingIrData.CarrierFreqKiloHz,
		PluseNanoSec:      m.ReceivingIrData.PluseNanoSec,
	}
	time.Sleep(m.ReceiveTime)
	return irdata, convertError(m.ReceiveErrorCode, m.ReceiveErrorMessage)
}

func (m *SendReceiveDevice) GetInfo(ctx context.Context) (*driver_module.DeviceInfo, error) {
	return &driver_module.DeviceInfo{
		DriverVersion:   m.DriverVersion,
		FirmwareVersion: m.FirmwareVersion,
	}, nil
}

func (m *SendReceiveDevice) Drop() error {
	return nil
}
