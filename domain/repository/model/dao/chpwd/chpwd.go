package chpwd

import (
	"fmt"
	"server/domain/repository/model/dao"
)

func GetAccount(jwt string) string {
	results, err := dao.MysqlConn.Query("SELECT account FROM member WHERE jwt=?", jwt)
	if err != nil {
		panic(err)
	}
	var account string
	for results.Next() {
		err = results.Scan(&account)
		if err != nil {
			fmt.Println(err.Error())
			return account
		}
	}
	defer results.Close()
	return account
}

func Chpwd(password string, account string) {
	results, err := dao.MysqlConn.Query("UPDATE member SET password=? WHERE account=?", password, account)
	if err != nil {
		panic(err)

	}
	defer results.Close()

}
