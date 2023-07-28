package mockdevice

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/NaKa2355/pirem/pkg/module/v1"
)

type MockIRData struct {
	CarrierFreqKiloHz uint32   `json:"carrier_freq_kilo_hz"`
	PluseNanoSec      []uint32 `json:"pluse_nano_sec"`
}

type Config struct {
	CanSend         bool       `json:"can_send"`
	CanReceive      bool       `json:"can_receive"`
	SendIrData      MockIRData `json:"ir_data"`
	FirmwareVersion string     `json:"firmware_version"`
	DriverVersion   string     `json:"driver_version"`
}

type Module struct{}

type MockDev struct {
	CanSend         bool
	CanReceive      bool
	SendIrData      MockIRData
	FirmwareVersion string
	DriverVersion   string
}

func (p *Module) NewDriver(conf json.RawMessage) (module.Driver, error) {
	config := Config{}
	err := json.Unmarshal(conf, &config)
	dev := &MockDev{
		CanSend:    config.CanSend,
		CanReceive: config.CanReceive,
		SendIrData: MockIRData{
			CarrierFreqKiloHz: config.SendIrData.CarrierFreqKiloHz,
			PluseNanoSec:      config.SendIrData.PluseNanoSec,
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
	fmt.Println("receive ir")
	return &module.IRData{CarrierFreqKiloHz: m.SendIrData.CarrierFreqKiloHz, PluseNanoSec: m.SendIrData.PluseNanoSec}, nil
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
