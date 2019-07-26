package main

import (
	"fmt"

	"github.com/howeyc/gopass"

	//. "github.com/JevonWei/usermanager/auth"
	. "github.com/xxdu521/usermod/users"
)

func main() {
	users := make(map[string]map[string]string)
	id := 0
	fmt.Println("马哥用户管理系统")

	if !Auth() {
		fmt.Println(".............密码错误")
		return
	}

	fmt.Print("请输入JevonWei用户系统密码: ")
	bytes, _ := gopass.GetPasswd()
	fmt.Println(bytes)

	for {
		var op string
		fmt.Print(`
----------------
1.add
2.update
3.del
4.query
5.exit

请输入指令：`)

		fmt.Scan(&op)
		if op == "1" {
			//fmt.Println("add")
			id++
			Add(id, users)
		} else if op == "2" {
			//fmt.Println("update")
			Update(users)
		} else if op == "3" {
			//fmt.Println("del")
			Del(users)
		} else if op == "4" {
			//fmt.Println("query")
			Query(users)
		} else if op == "5" {
			fmt.Println("\n\n退出管理系统，下次再见！")
			break
		} else {
			fmt.Printf("\n\n输入错误，请重新输入")
		}
	}
}
