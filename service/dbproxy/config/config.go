package config

import "fmt"

var (
	MySQLSource = "root:13628@tcp(127.0.0.1:3306)/fileserver?charset=utf8"
)

func UpdateDBHost(host string) {
	MySQLSource = fmt.Sprintf("root:13628@tcp(%s)/fileserver?charset=utf8", host)
}

