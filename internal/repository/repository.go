package repository

import (
"github.com/getclasslabs/go-tools/pkg/db"
"github.com/getclasslabs/{apiname}/internal/config"
_ "github.com/go-sql-driver/mysql"
)

func Start(){
	Db = &db.MySQL{}
	Db.Connect(config.Config.Mysql)
}

var Db db.Database