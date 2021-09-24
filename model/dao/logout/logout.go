package logout

import (
	"fmt"
	"server/model/dao"
	"strconv"
)

//判斷是否存在JWT，如果存在就刪除
func JwtIsExistAndDeleteIfExist(jwt string) bool {
	results := dao.MysqlConn.QueryRow("SELECT COUNT(jwt) FROM member Where jwt=?", jwt)
	var count int
	err := results.Scan(&count)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("DBcount :" + strconv.Itoa(count))
	if count != 0 {
		results, err := dao.MysqlConn.Query("UPDATE member SET jwt='NULL' WHERE jwt=?", jwt)
		if err != nil {
			panic(err)

		}
		defer results.Close()
		return true
	}

	return false
}

func LogoutJwt(k string) {
	results, err := dao.MysqlConn.Query("UPDATE member SET jwt='NULL' WHERE account=?", k)
	if err != nil {
		panic(err)

	}
	defer results.Close()

}
