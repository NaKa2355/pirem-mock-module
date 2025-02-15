package MockDevices

import (
	"encoding/json"
	"fmt"

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

var _ driver_module.DriverModule = &Module{}

func (p *Module) LoadDevice(conf json.RawMessage) (driver_module.Device, error) {
	config := Config{}
	err := json.Unmarshal(conf, &config)
	if err != nil {
		return nil, err
	}

	if config.CanSend && config.CanReceive {
		d := NewSendReceiveDevice(&config)
		return &d, nil
	}
	if config.CanSend {
		d := NewSendOnlyDevice(&config)
		return &d, nil
	}
	if config.CanReceive {
		d := NewReceiveOnlyDevice(&config)
		return &d, nil
	}
	return nil, fmt.Errorf("failed to create device")
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
