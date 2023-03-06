package dao

import (
	"iot_backend/model"
	"iot_backend/util"
)

// QuerySensorByDeviceTag
// @Description: 以设备tag查询该设备所拥有的所有传感器信息
// @param data 需要绑定的数据
// @param tag 传入设备tag
// @return int64 查询到的数据条数
func QuerySensorByDeviceTag(data *[]model.Sensor, tag string) int64 {
	db, _ := util.GetOrm()
	result := db.Where("device_tag = ?", tag).Find(&data)
	counts := result.RowsAffected

	return counts
}

// UpdateSensor
// @Description: 更新传感器数据
// @param deviceTag 该传感器属于的设备Tag
// @param tag 传感器tag
// @param value 需要更新的值
// @return error
// @return int64 找到的传感器条数
func UpdateSensor(deviceTag string, tag string, value string) (error, int64) {
	db, _ := util.GetOrm()
	var count int64
	result := db.Model(&model.Sensor{}).Where("device_tag = ? AND tag = ?", deviceTag, tag).Count(&count).Update("value", value)
	err := result.Error

	return err, count
}