package models

// demo
type Device struct {
	Id          int    `json:"id" gorm:"comment:id"`
	DeviceName  string `json:"device_name" gorm:"comment:设备编号"`
	FirmwareVer int    `json:"firmware_ver" gorm:"comment:固件版本"`
	HardwareVer int    `json:"hardware_ver" gorm:"comment:硬件版本"`
	DeviceGroup int    `json:"device_group" gorm:"comment:分组id"`
	LastTime    int    `json:"last_time" gorm:"comment:更新时间"`
	Note        string `json:"note" gorm:"comment:备注"`
	Status      int8   `json:"status" gorm:"comment:升级状态：1可升级"`
}
