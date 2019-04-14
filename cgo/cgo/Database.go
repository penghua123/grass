package cgo

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

//数据库的配置
const (
	username   = "usgmtr"
	password   = ""
	ip         = "localhost"
	port       = 5432
	dbName     = "cgo"
	driverName = "postgres"
)

//DB数据库连接池
var DB *sql.DB

func InitDB() {
	//构建连接："用户名:密码@tcp(IP:端口)/数据库?charset=uft8"
	//注意：要想解析time.Time类型，必须要设置parseTime=True
	path := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", ip, port, username, password, dbName)
	//打开数据库，前者是驱动名，所以要导入:_"github.com/lib/pq"
	DB, err := sql.Open(driverName, path)
	if err != nil {
		log.Panic(err)
	}
	/*stmt, err := DB.Prepare("select count(*) from \"Vm\"")
	if err != nil {
		log.Panic(err)
	}
	res, err := stmt.Exec()
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(res)*/
	//设置数据库最大连接数
	DB.SetConnMaxLifetime(100)
	//设置数据库最大闲置连接数
	DB.SetMaxIdleConns(10)
	//验证连接
	if err := DB.Ping(); err != nil {
		log.Panic(err)
	}
	log.Println("database connect success")
}

func CreateTable() {
	userTable := "CREATE TABLE IF NOT EXISTS \"user\"(" +
		"id serial," +
		"username VARCHAR(20) NOT NULL," +
		"password VARCHAR(40) NOT NULL," +
		"\"create_time\" timestamp without time zone," +
		"PRIMARY KEY ( id )" +
		");"

	feedbackTable := "CREATE TABLE IF NOT EXISTS feedback(" +
		"id serial," +
		"\"user_id\" INT UNSIGNED NOT NULL," +
		"title VARCHAR(50) NOT NULL," +
		"content VARCHAR(200) NOT NULL," +
		"\"create_time\" timestamp without time zone," +
		"PRIMARY KEY ( id )" +
		");"

	pictureTable := "CREATE TABLE IF NOT EXISTS picture(" +
		"id serial," +
		"\"feedback_id\" INT UNSIGNED NOT NULL," +
		"address VARCHAR(200) NOT NULL," +
		"\"create_time\" timestamp without time zone," +
		"PRIMARY KEY ( id )" +
		");"
	fmt.Println(userTable)
	err := execDb(userTable)
	if err != nil {
		log.Panic(err)
	}
	err = execDb(feedbackTable)
	if err != nil {
		log.Panic(err)
	}
	err = execDb(pictureTable)
	if err != nil {
		log.Panic(err)
	}
}

func execDb(sStmt string) error {
	stmtPre, err := DB.Prepare(sStmt)
	if err != nil {
		return err
	}
	_, err = stmtPre.Exec()
	if err != nil {
		return err
	}
	return nil
}
