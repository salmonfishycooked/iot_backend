package tcp

import (
	"encoding/json"
	"errors"
	"iot_backend/dao"
	"iot_backend/model"
)

// authDevice
// @Description: 设备鉴权
// @param conn 已建立的连接
// @return bool 鉴权是否成功
func authDevice(conn *Connection) bool {
	err := checkTag(conn) // 等待客户端发送设备tag

	// 不发送tag或者这个设备没有在前端创建
	if err != nil {
		return false
	}
	return true
}

// checkTag
// @Description: 检查客户端是否发出tag
// @return error
func checkTag(conn *Connection) error {
	//conn.conn.SetReadDeadline(time.Now().Add(time.Duration(TIMEOUT) * time.Second)) // 对连接设置读数据的超时时间
	recStr, _ := readFromClient(conn.conn) // 从客户端读取DeviceTag
	data := DeviceTag{}
	if err := json.Unmarshal([]byte(recStr), &data); err == nil {
		counts := dao.QueryByTag(&model.Device{}, data.Tag)
		// 在数据库中找不到该设备
		if counts == 0 {
			return errors.New("error")
		}

		conn.DeviceTag = data.Tag // 鉴权成功，设置标识tag
	} else {
		return err
	}
	return nil
}