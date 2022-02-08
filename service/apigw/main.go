package main

import (
	"LeiliNetdisk/service/apigw/route"
)

func main() {
	r := route.Router()
	r.Run(":8086")
}
