package memberlog

import (
	"fmt"
	"server/model/dao"
	"server/model/dto"
	"strconv"
)

func QueryLog(account string) []string {
	results, err := dao.MysqlConn.Query("SELECT login_time FROM logintime WHERE account=?", account)
	if err != nil {
		panic(err)

	}

	var querydb dto.Memberlog
	log := []string{}
	for results.Next() {
		err = results.Scan(&querydb.Logintime)
		if err != nil {
			fmt.Println(err.Error())
		}
		Logintime := strconv.Itoa(querydb.Logintime)
		log = append(log, Logintime)
	}
	defer results.Close()
	return log
}
