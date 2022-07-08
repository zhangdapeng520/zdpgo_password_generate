package main

import (
	"fmt"
	"github.com/zhangdapeng520/zdpgo_password_generate"
	"os"
)

/*
@Time : 2022/6/21 17:16
@Author : 张大鹏
@File : main.go
@Software: Goland2021.3.1
@Description: 测试生成密码
*/

func main() {
	config := zdpgo_password_generate.Config{
		Length:                   128,
		IncludeSymbols:           true,
		IncludeNumbers:           true,
		IncludeLowercaseLetters:  true,
		IncludeUppercaseLetters:  true,
		ExcludeSimilarCharacters: true,
	}
	g := zdpgo_password_generate.NewWithConfig(&config)

	// 生成1个密码
	data, err := g.Generate()
	if err != nil {
		panic(err)
	}
	fmt.Println("生成1个密码：", *data)

	// 生成10个密码
	fmt.Println("生成10个密码")
	pwds, err := g.GenerateMany(10)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	for _, pwd := range pwds {
		fmt.Println(pwd)
	}
}
