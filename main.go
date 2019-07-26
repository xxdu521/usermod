package main

import (
	"flag"
	"fmt"
	"strconv"

	"github.com/xxdu521/usermod/users"
)

func main() {

	//认证
	auth := flag.Bool("N", false, "no auth")
	flag.Parse()
	fmt.Println(*auth)
	if !users.Auth(*auth) {
		return
	}

	//功能测试
	metu := `
1. 新建用户
2. 修改用户
3. 删除用户
4. 查询用户
5. 退出`

	callbacks := map[int]func(){
		1: users.Add,
		2: users.Query,
		3: users.Del,
		4: users.Update,
	}

	for {
		fmt.Println(metu)

		if callback, ok := callbacks[strconv.Atoi(Inputstring("请选择功能项: "))]; ok {
			//if callback, ok := callbacks[strconv.Atoi(users.Inputstring("请选择功能项: "))]; ok {
			callback()
		} else {
			print("输入错误，请重新输入!!!")
		}
	}
	users.Test()
}
