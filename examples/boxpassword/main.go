package main

import (
	"fmt"
	"github.com/zhangdapeng520/zdpgo_password_generate"
)

// 这里是管理密码，应该从配置文件或者环境变量获取
var mastPw0 = "masterpassword"
var mastPw1 = "masterpassword1"

func main() {
	// 用户密码
	userPw := "userpassword"

	// 对用户密码进行加密
	pwHash, err := zdpgo_password_generate.Hash(userPw, mastPw0, 0, zdpgo_password_generate.ScryptParams{N: 32768, R: 16, P: 1}, zdpgo_password_generate.DefaultParams)
	if err != nil {
		fmt.Println("Hash fail. ", err)
	}

	// 将用户密码存储到数据库

	// -------- Verify -------------
	// 从数据库获取用户密码
	// 从环境变量获取管理密码
	// 校验用户密码是否正确
	err = zdpgo_password_generate.Verify(userPw, mastPw0, pwHash)
	if err != nil {
		fmt.Println("Verify fail. ", err)
	}
	fmt.Println("Success")

	// --------- Update ------------
	// 从数据库获取密码，用新的管理密码加密
	updated, err := zdpgo_password_generate.UpdateMaster(mastPw1, mastPw0, 1, pwHash, zdpgo_password_generate.DefaultParams)
	if err != nil {
		fmt.Println("Update fail. ", err)
	}

	// 使用新的管理密码校验
	err = zdpgo_password_generate.Verify(userPw, mastPw1, updated)
	if err != nil {
		fmt.Println("Verify fail. ", err)
	}
	fmt.Println("Success verifying updated hash")
}
