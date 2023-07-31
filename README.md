# pirem-mock-module
pirem用にモックデバイスを追加するためのモジュール

設定ファイルによって、モックデバイスのファームウェアバージョン、ドライバーバージョン、受信するデータ、受信機能、送信機能の有無を設定できます。

設定例
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
        "receiving_ir_data": {
          "varrier_freq_kilo_hz": 40,
          "pluse_nano_sec": [10,20,30,40,50]
        }
      }
    }
  ]
}
```
