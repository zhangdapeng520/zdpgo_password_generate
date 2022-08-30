package main

import (
	"fmt"
	"github.com/zhangdapeng520/zdpgo_password_generate"
)

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(zdpgo_password_generate.DefaultGenerator.GenerateByLength(uint((i + 1) * 10)))
		fmt.Println(zdpgo_password_generate.DefaultGenerator.GenerateByWeak())
		fmt.Println(zdpgo_password_generate.DefaultGenerator.GenerateByOK())
		fmt.Println(zdpgo_password_generate.DefaultGenerator.GenerateByStrong())
		fmt.Println(zdpgo_password_generate.DefaultGenerator.GenerateByVeryStrong())
		fmt.Println("================")
	}
}
