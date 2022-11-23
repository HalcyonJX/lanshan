package main

import (
	"lanshan/work4/api"
	"lanshan/work4/dao"
)

func main() {
	dao.InitDB()
	api.InitRouter()
}
