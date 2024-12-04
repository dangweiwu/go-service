package memodel

import "strconv"

// redis login id
func GetAdminRedisLoginId(id int) string {
	return "lgn:" + strconv.Itoa(id)
}
