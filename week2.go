package main

import (
	"database/sql"
	"github.com/pkg/errors"
)

func main() {
	//情况一，记录根据逻辑可能存在也可能不存在，比如根据手机号查询用户，存在则返回用户信息
	//这种情况则应该屏蔽错误，直接返回nil

	//情况二，记录根据逻辑应该存在，比如用户id查询用户，正常情况应该返回用户信息
	//这种情况则应该返回错误抛个上层，代码如下
	uid := 1
	err = db.QueryRow(`SELECT * FROM users WHERE ID = ?`, uid).Scan()
	if err == sql.ErrNoRows {
		return errors.Wrapf(err, "get user info failed, uid: %d", uid)
	}
}
