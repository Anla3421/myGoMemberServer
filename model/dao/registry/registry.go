package registry

import (
	"fmt"
	"server/model/dao"
	"strconv"
)

func IsExist(account string) bool {
	results := dao.MysqlConn.QueryRow("SELECT COUNT(account) FROM member WHERE account=?", account)
	var count int
	err := results.Scan(&count)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("DBcount :" + strconv.Itoa(count))
	return count != 0

}

func RegMember(account string, password string) {
	results, err := dao.MysqlConn.Query("INSERT INTO member (account,password) VALUES (?,?)", account, password)
	if err != nil {
		panic(err)
	}
	defer results.Close()

}
