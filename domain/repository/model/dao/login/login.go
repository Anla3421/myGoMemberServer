package login

import (
	"crypto/md5"
	"fmt"
	"server/domain/repository/model/dao"
)

// 判斷密碼是否正確
func IsRight(account string, password string) bool {
	results, err := dao.MysqlConn.Query("SELECT password FROM member WHERE account=?", account)
	//sql injection test
	// results, err := dao.MysqlConn.Query("SELECT password FROM member WHERE account=" + account)
	if err != nil {
		panic(err)

	}
	var pwd string
	for results.Next() {
		err = results.Scan(&pwd)
		if err != nil {
			fmt.Println(err.Error())
		}

	}
	defer results.Close()
	fmt.Println("DBpwd :" + pwd)
	fmt.Printf("%x\n", md5.Sum([]byte(password)))
	return pwd == fmt.Sprintf("%x", md5.Sum([]byte(password)))
}

func LoginJwt(ss string, account string) {
	results, err := dao.MysqlConn.Query("UPDATE member SET jwt=? WHERE account=?", ss, account)
	if err != nil {
		panic(err)

	}
	defer results.Close()

}

func Loginlog(account string, loginlog int64) {
	results, err := dao.MysqlConn.Query("INSERT INTO logintime (account,login_time) VALUES (?,?)", account, loginlog)
	if err != nil {
		panic(err)

	}
	defer results.Close()

}
