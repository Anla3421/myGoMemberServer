package query

import (
	"fmt"
	"server/domain/repository/model/dao"
	"server/domain/repository/model/dto"

	"strconv"
)

func GetInfoByJWT(JWT string) (res dto.GetInfoByJWT, err error) {
	results := dao.MysqlConn.QueryRow("SELECT ID,account FROM member Where jwt=?", JWT)
	err = results.Scan(&res.ID, &res.Account)
	if err != nil {
		return res, err
	}
	fmt.Println("Account ID :" + strconv.FormatInt(int64(res.ID), 10))
	return res, nil
}
