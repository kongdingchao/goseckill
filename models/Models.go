package models

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"net/url"
	"os"
)

func init() {
	//如果存在app.conf文件，则表示程序已安装，执行数据库初始化
	if _, err := os.Stat("conf/app.conf"); err == nil {
		Init()
	}
}

//初始化数据库注册
func Init() {
	//初始化数据库
	RegisterDB()

	runmode := beego.AppConfig.String("runmode")
	if runmode == "prod" {
		orm.Debug = false
		orm.RunSyncdb("default", false, false)
	} else {
		orm.Debug = true
		orm.RunSyncdb("default", false, true)
	}

	//安装初始数据
	install()
}

//注册数据库
func RegisterDB() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	models := []interface{}{
		NewSeckill(),
	}
	orm.RegisterModelWithPrefix(beego.AppConfig.DefaultString("db::prefix", "kdc_"), models...)
	dbUser := beego.AppConfig.String("db::user")
	dbPassword := beego.AppConfig.String("db::password")
	if envpass := os.Getenv("MYSQL_PASSWORD"); envpass != "" {
		dbPassword = envpass
	}
	dbDatabase := beego.AppConfig.String("db::database")
	if envdatabase := os.Getenv("MYSQL_DATABASE"); envdatabase != "" {
		dbDatabase = envdatabase
	}
	dbCharset := beego.AppConfig.String("db::charset")
	dbHost := beego.AppConfig.String("db::host")
	if envhost := os.Getenv("MYSQL_HOST"); envhost != "" {
		dbHost = envhost
	}
	dbPort := beego.AppConfig.String("db::port")
	if envport := os.Getenv("MYSQL_PORT"); envport != "" {
		dbPort = envport
	}
	loc := "Local"
	if timezone := beego.AppConfig.String("db::timezone"); timezone != "" {
		loc = url.QueryEscape(timezone)
	}
	dbLink := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&loc=%v", dbUser, dbPassword, dbHost, dbPort, dbDatabase, dbCharset, loc)
	maxIdle := beego.AppConfig.DefaultInt("db::maxIdle", 50)
	maxConn := beego.AppConfig.DefaultInt("db::maxConn", 300)
	if err := orm.RegisterDataBase("default", "mysql", dbLink, maxIdle, maxConn); err != nil {
		panic(err)
	}
}
