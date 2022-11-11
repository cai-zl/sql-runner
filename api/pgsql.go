package api

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type PgsqlInfo struct {
	SqlInfo
}

func (receiver *PgsqlInfo) Run() {
	url := receiver.GetDriver()
	db, err := sql.Open("postgres", url)
	if err != nil {
		panic(err)
	}
	receiver.DoRun(db)
}

func (receiver *PgsqlInfo) GetDriver() string {
	url := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", receiver.Host, receiver.Port, receiver.User, receiver.Password, receiver.DbName)
	return url
}

func init() {
	SqlInfoCache[PgsqlDb] = func(info *SqlInfo) Runner {
		return &PgsqlInfo{
			*info,
		}
	}
}
