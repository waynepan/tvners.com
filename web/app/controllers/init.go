package controllers

import (
    "github.com/revel/revel"
    "github.com/coopernurse/gorp"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    "fmt"
    "strings"
    "tvners.com/web/app/models"
)

func init(){
    revel.OnAppStart(InitDb)
    revel.InterceptMethod((*GorpController).Begin, revel.BEFORE)
    revel.InterceptMethod((*GorpController).Commit, revel.AFTER)
    revel.InterceptMethod((*GorpController).Rollback, revel.FINALLY)
}

var InitDb func() = func(){
    connectionString := getConnectionString()
//    revel.ERROR.Printf("%s", connectionString)
    if db, err := sql.Open("mysql", connectionString); err != nil {
        revel.ERROR.Fatal(err)
    } else {
        Dbm = &gorp.DbMap{
            Db: db, 
            Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}
    }
    // Defines the table for use by GORP
    // This is a function we will create soon.
    defineCarsTable(Dbm)
    if err := Dbm.CreateTablesIfNotExists(); err != nil {
        revel.ERROR.Fatal(err)
    }
}

func defineCarsTable(dbm *gorp.DbMap){
    // set "id" as primary key and autoincrement
    t := dbm.AddTable(models.Car{}).SetKeys(true, "id") 
    t.ColMap("name").SetMaxSize(500)
}

func getParamString(param string, defaultValue string) string {
    p, found := revel.Config.String(param)
    if !found {
        if defaultValue == "" {
            revel.ERROR.Fatal("Cound not find parameter: " + param)
        } else {
            return defaultValue
        }
    }
    return p
}

func getConnectionString() string {
    host := getParamString("db.host", "")
    port := getParamString("db.port", "3306")
    user := getParamString("db.user", "")
    pass := getParamString("db.password", "")
    dbname := getParamString("db.name", "tvners")
    protocol := getParamString("db.protocol", "tcp")
    dbargs := getParamString("dbargs", " ")

    if strings.Trim(dbargs, " ") != "" {
        dbargs = "?" + dbargs
    } else {
        dbargs = ""
    }

    return fmt.Sprintf("%s:%s@%s([%s]:%s)/%s%s", user, pass, protocol, host, port, dbname, dbargs)
}
