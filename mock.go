package mockdevice

import (
	"context"
	"encoding/json"

	"github.com/NaKa2355/pirem/pkg/module/v1"
)

type MockIRData struct {
	CarrierFreqKiloHz uint32   `json:"carrier_freq_kilo_hz"`
	PluseNanoSec      []uint32 `json:"pluse_nano_sec"`
}

type Config struct {
	CanSend         bool       `json:"can_send"`
	CanReceive      bool       `json:"can_receive"`
	ReceivingIrData MockIRData `json:"receiving_ir_data"`
	FirmwareVersion string     `json:"firmware_version"`
	DriverVersion   string     `json:"driver_version"`
}

type Module struct{}

type MockDev struct {
	CanSend         bool
	CanReceive      bool
	ReceivingIrData MockIRData
	FirmwareVersion string
	DriverVersion   string
}

func (p *Module) NewDriver(conf json.RawMessage) (module.Driver, error) {
	config := Config{}
	err := json.Unmarshal(conf, &config)
	dev := &MockDev{
		CanSend:    config.CanSend,
		CanReceive: config.CanReceive,
		ReceivingIrData: MockIRData{
			CarrierFreqKiloHz: config.ReceivingIrData.CarrierFreqKiloHz,
			PluseNanoSec:      config.ReceivingIrData.PluseNanoSec,
		},
		FirmwareVersion: config.FirmwareVersion,
		DriverVersion:   config.DriverVersion,
	}
	return dev, err
}

func (m *MockDev) SendIR(ctx context.Context, irdata *module.IRData) error {
	return nil
}

func (m *MockDev) ReceiveIR(ctx context.Context) (*module.IRData, error) {
	irdata := &module.IRData{
		CarrierFreqKiloHz: m.ReceivingIrData.CarrierFreqKiloHz,
		PluseNanoSec:      m.ReceivingIrData.PluseNanoSec,
	}
	return irdata, nil
}

func (m *MockDev) GetInfo(ctx context.Context) (*module.DeviceInfo, error) {
	return &module.DeviceInfo{
		CanSend:         m.CanSend,
		CanReceive:      m.CanReceive,
		DriverVersion:   m.DriverVersion,
		FirmwareVersion: m.FirmwareVersion,
	}, nil
}

func (m *MockDev) Drop() error {
	return nil
}
