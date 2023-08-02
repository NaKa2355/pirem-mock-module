package MockDriverice

import (
	"context"
	"encoding/json"
	"time"

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
	ReceiveTimeMs   uint32     `json:"receive_time_ms"`
	SendTimeMs      uint32     `json:"send_time_ms"`
	FirmwareVersion string     `json:"firmware_version"`
	DriverVersion   string     `json:"driver_version"`
}

type Module struct{}

type MockDriver struct {
	CanSend         bool
	CanReceive      bool
	ReceivingIrData MockIRData
	ReceiveTime     time.Duration
	SendTime        time.Duration
	FirmwareVersion string
	DriverVersion   string
}

func (p *Module) NewDriver(conf json.RawMessage) (module.Driver, error) {
	config := Config{}
	err := json.Unmarshal(conf, &config)
	d := &MockDriver{
		CanSend:    config.CanSend,
		CanReceive: config.CanReceive,
		ReceivingIrData: MockIRData{
			CarrierFreqKiloHz: config.ReceivingIrData.CarrierFreqKiloHz,
			PluseNanoSec:      config.ReceivingIrData.PluseNanoSec,
		},
		FirmwareVersion: config.FirmwareVersion,
		DriverVersion:   config.DriverVersion,
		ReceiveTime:     time.Duration(time.Millisecond * time.Duration(config.ReceiveTimeMs)),
		SendTime:        time.Duration(time.Millisecond + time.Duration(config.SendTimeMs)),
	}
	return d, err
}

func (m *MockDriver) SendIR(ctx context.Context, irdata *module.IRData) error {
	time.Sleep(m.SendTime)
	return nil
}

func (m *MockDriver) ReceiveIR(ctx context.Context) (*module.IRData, error) {
	irdata := &module.IRData{
		CarrierFreqKiloHz: m.ReceivingIrData.CarrierFreqKiloHz,
		PluseNanoSec:      m.ReceivingIrData.PluseNanoSec,
	}
	time.Sleep(m.ReceiveTime)
	return irdata, nil
}

func (m *MockDriver) GetInfo(ctx context.Context) (*module.DeviceInfo, error) {
	return &module.DeviceInfo{
		CanSend:         m.CanSend,
		CanReceive:      m.CanReceive,
		DriverVersion:   m.DriverVersion,
		FirmwareVersion: m.FirmwareVersion,
	}, nil
}

func (m *MockDriver) Drop() error {
	return nil
}
