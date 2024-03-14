package main

import (
	"disbursement-service/models"
	_ "disbursement-service/routers"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/config"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
	"log"
	"os/user"
	"path/filepath"
)

const (
	keyDbUser     = "DBUsername"
	keyDbPassword = "DBPassword"
	keyDbHost     = "DBHost"
	keyDbPort     = "DBPort"
	keyDbName     = "DBName"
)

func init() {
	cfgLog := `{"filename":"logs/app.log","level":7,"daily":true,"maxdays":365,"rotate":true}`
	beego.SetLogger("file", cfgLog)

	orm.Debug = beego.AppConfig.DefaultBool("orm.query.debug", false)

	dbUser, dbPassword, dbHost, dbPort, dbName := getDatabaseConfig()
	orm.RegisterDriver("postgres", orm.DRPostgres)

	orm.RegisterDataBase("default", "postgres",
		fmt.Sprintf("user=%s password=%s host=%s port=%s  dbname=%s sslmode=%s search_path=%s ",
			dbUser, dbPassword, dbHost, dbPort, dbName, "disable", "disbursement"))
	orm.SetMaxOpenConns("default", beego.AppConfig.DefaultInt("DBMaxOpenCon", 20))
	orm.SetMaxIdleConns("default", beego.AppConfig.DefaultInt("DBMaxIdleCon", 10))

	orm.RegisterModel(new(models.Bank))
	orm.RegisterModel(new(models.Disbursement))
}

func getDatabaseConfig() (string, string, string, string, string) {
	dbUser := beego.AppConfig.String(keyDbUser)
	dbPassword := beego.AppConfig.String(keyDbPassword)
	dbHost := beego.AppConfig.String(keyDbHost)
	dbPort := beego.AppConfig.String(keyDbPort)
	dbName := beego.AppConfig.String(keyDbName)
	usr, err := user.Current()
	if err == nil {
		configPath := filepath.Join(usr.HomeDir, "conf", "royalty-service.conf")
		conf, err := config.NewConfig("ini", configPath)
		if err == nil {
			log.Println("External configuration found in ", configPath)
			dbUser = conf.DefaultString(keyDbUser, beego.AppConfig.String(keyDbUser))
			dbPassword = conf.DefaultString(keyDbPassword, beego.AppConfig.String(keyDbPassword))
			dbHost = conf.DefaultString(keyDbHost, beego.AppConfig.String(keyDbHost))
			dbPort = conf.DefaultString(keyDbPort, beego.AppConfig.String(keyDbPort))
			dbName = conf.DefaultString(keyDbName, beego.AppConfig.String(keyDbName))
		}
	}
	return dbUser, dbPassword, dbHost, dbPort, dbName
}

func main() {
	var ok = true
	if ok = models.RegisterErrorCode(); ok {
		beego.Run()
	} else {
		logs.Emergency("Failed to startup ")
	}
}
