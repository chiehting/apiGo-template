package service

import "github.com/chiehting/go-template/pkg/snowflake"

func GetId() (id uint64, err error) {
	id, err = snowflake.NextID()
	return
}
