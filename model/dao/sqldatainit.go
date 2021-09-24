package dao

import (
	"fmt"
	"server/model/dto"

	_ "github.com/go-sql-driver/mysql"
)

func QueryInfoAll() map[string]dto.Response {
	results, err := MysqlConn.Query("SELECT account,password,jwt FROM member")
	if err != nil {
		panic(err)

	}
	var querydb dto.Response
	member := make(map[string]dto.Response)
	for results.Next() {
		err = results.Scan(&querydb.Account, &querydb.Password, &querydb.JWT)
		if err != nil {
			fmt.Println(err.Error())
			// return Response
		}
		member[querydb.Account] = querydb
	}
	defer results.Close()
	return member
}
