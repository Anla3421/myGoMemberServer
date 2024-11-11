package dao

import (
	"fmt"
	"server/domain/repository/model/dto"
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
