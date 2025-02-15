package MockDevices

import (
	"context"
	"time"

	"github.com/NaKa2355/pirem/pkg/driver_module/v1"
)

type ReceiveOnlyDevice struct {
	ReceivingIrData     MockIRData
	ReceiveTime         time.Duration
	FirmwareVersion     string
	DriverVersion       string
	ReceiveErrorCode    string
	ReceiveErrorMessage string
}

func NewReceiveOnlyDevice(config *Config) ReceiveOnlyDevice {
	return ReceiveOnlyDevice{
		ReceivingIrData: MockIRData{
			CarrierFreqKiloHz: config.ReceivingIrData.CarrierFreqKiloHz,
			PluseNanoSec:      config.ReceivingIrData.PluseNanoSec,
		},
		FirmwareVersion:     config.FirmwareVersion,
		DriverVersion:       config.DriverVersion,
		ReceiveTime:         time.Duration(time.Millisecond * time.Duration(config.ReceiveTimeMs)),
		ReceiveErrorCode:    config.ReceiveErrorCode,
		ReceiveErrorMessage: config.ReceiveErrorMessage,
	}
}

func (m *ReceiveOnlyDevice) ReceiveIR(ctx context.Context) (*driver_module.IRData, error) {
	irdata := &driver_module.IRData{
		CarrierFreqKiloHz: m.ReceivingIrData.CarrierFreqKiloHz,
		PluseNanoSec:      m.ReceivingIrData.PluseNanoSec,
	}
	time.Sleep(m.ReceiveTime)
	return irdata, convertError(m.ReceiveErrorCode, m.ReceiveErrorMessage)
}
func (m *ReceiveOnlyDevice) GetInfo(ctx context.Context) (*driver_module.DeviceInfo, error) {
	return &driver_module.DeviceInfo{
		DriverVersion:   m.DriverVersion,
		FirmwareVersion: m.FirmwareVersion,
	}, nil
}

func (m *ReceiveOnlyDevice) Drop() error {
	return nil
}
