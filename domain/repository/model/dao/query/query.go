package query

import (
	"fmt"
	"server/domain/repository/model/dao"
	"server/domain/repository/model/dto"
	"strconv"
)

func QueryInfoIsExist(account string) bool {
	results := dao.MysqlConn.QueryRow("SELECT COUNT(account) FROM member WHERE account=?", account)
	var count int
	err := results.Scan(&count)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("DBcount :" + strconv.Itoa(count))
	return count != 0 //true=有人，false=沒人
}

func QueryInfo(account string) (querydb dto.Response) {
	results, err := dao.MysqlConn.Query("SELECT password,jwt FROM member WHERE account=?", account)
	if err != nil {
		panic(err)

	}
	// var querydb Response
	for results.Next() {
		err = results.Scan(&querydb.Password, &querydb.JWT)
		if err != nil {
			fmt.Println(err.Error())
			return dto.Response{}
		}
	}
	defer results.Close()
	return querydb
}
