package MockDevices

import (
	"context"
	"time"

	"github.com/NaKa2355/pirem/pkg/driver_module/v1"
)

type SendOnlyDevice struct {
	SendTime         time.Duration
	FirmwareVersion  string
	DriverVersion    string
	SendErrorCode    string
	SendErrorMessage string
}

func NewSendOnlyDevice(config *Config) SendOnlyDevice {
	return SendOnlyDevice{
		FirmwareVersion:  config.FirmwareVersion,
		DriverVersion:    config.DriverVersion,
		SendTime:         time.Duration(time.Millisecond * time.Duration(config.SendTimeMs)),
		SendErrorCode:    config.SendErrorCode,
		SendErrorMessage: config.SendErrorMessage,
	}
}

func (m *SendOnlyDevice) SendIR(ctx context.Context, irdata *driver_module.IRData) error {
	time.Sleep(m.SendTime)
	return convertError(m.SendErrorCode, m.SendErrorMessage)
}

func (m *SendOnlyDevice) GetInfo(ctx context.Context) (*driver_module.DeviceInfo, error) {
	return &driver_module.DeviceInfo{
		DriverVersion:   m.DriverVersion,
		FirmwareVersion: m.FirmwareVersion,
	}, nil
}

func (m *SendOnlyDevice) Drop() error {
	return nil
}
