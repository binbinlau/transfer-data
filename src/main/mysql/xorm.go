package mysql

import (
	"github.com/binsix/transfer-data/src/main/utils"
	"github.com/go-xorm/xorm"
)

func getXormSession(driverName string, dataSourceName string) (*xorm.Engine, error) {
	engine, err := xorm.NewEngine(driverName, dataSourceName)
	if err != nil {
		panic(err)
	}
	return engine, err
}

func getMysqlSession(user, passwd, db, charset string) (*xorm.Engine, error) {
	dataSourceName := user + ":" + passwd + "@/" + db + "?charset=" + charset
	return getXormSession("mysql", dataSourceName)
}

func GetMysqlSession() (*xorm.Engine, error) {
	return getMysqlSession(utils.Conf.Mysql.User, utils.Conf.Mysql.Password, utils.Conf.Mysql.Database, utils.Conf.Mysql.Charset)
}
