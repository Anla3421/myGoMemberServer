package logdel

import (
	"fmt"
	"server/domain/repository/model/dao"

	"strconv"
)

func DeleteLog(account string) {
	results, err := dao.MysqlConn.Query("DELETE FROM logintime WHERE account=?", account)
	if err != nil {
		panic(err)

	}
	defer results.Close()
}

func IFDeleteLog(account string) bool {
	results := dao.MysqlConn.QueryRow("SELECT COUNT(account) FROM member Where account=?", account)
	var count int
	err := results.Scan(&count)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("DBcount :" + strconv.Itoa(count))
	if count != 0 {

		results, err := dao.MysqlConn.Query("DELETE FROM logintime WHERE account=?", account)
		if err != nil {
			panic(err)

		}
		defer results.Close()
		return true
	}
	return false
}
