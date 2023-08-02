# pirem-mock-module
PiRem用にモックデバイスを追加するためのモジュール

設定ファイル(/etc/piremd.json)によって、モックデバイスのファームウェアバージョン、ドライバーバージョン、受信するデータ、受信機能、送信機能の有無を設定できます。

## 設定例
```json
{
  "enable_reflection": true,
  "devices":[
    {
      "name": "mock_device",
      "id": "1",
      "module_name": "mock",
      "config": {
        "can_send": true,
        "can_receive": true,
        "firmware_version": "0.1.0",
        "driver_version": "0.1.1",
        "receive_time_ms": 3000,
        "send_time_ms": 500,
        "receiving_ir_data": {
          "carrier_freq_kilo_hz": 40,
          "pluse_nano_sec": [10,20,30,40,50]
        }
      }
    }
  ]
}
```

| キー | 意味 |
| ---- | ---- |
| `"can_send"` | 送信機能の有効化 |
| `"can_receive"` | 受信機能の有効化 |
| `"firmware_version"` | ファームウェアバージョン |
| `"driver_version"` | ドライバーバージョン(Moduleのバージョン) |
| `"receiving_ir_data"` | 赤外線を受信時のモックデータ |
| `"varrier_freq_kilo_hz"` | 赤外線のキャリア周波数 |
| `"pluse_nano_sec"` | 赤外線のデータ |
| `"receive_time_ms"` | 赤外線受信にかかる時間(ms) |
| `"send_time_ms"` | 赤外線送信にかかる時間(ms) |


