package user

import (
	"database/sql"
	"fmt"
	_"github.com/go-sql-driver/mysql" //mysql的go驱动包
	"go/src/server/common"
)

type UserDao struct {

}

func (UserDao) GetUserInfoById(useId string) (id int, name string, sex int, age int) {
	//函数用来返回准备要执行的sql操作，然后返回准备完毕的执行状态。
	sql := "select id,name,sex,age from user where id = ?"
	stmt, err := common.DB.Prepare(sql)
	delError(err)
	//关闭查询连接
	defer stmt.Close()
	//查询的结果将保存在name中
	//执行之前预执行的命令
	err = stmt.QueryRow(useId).Scan(&id, &name, &sex, &age)
	delError(err)

	return id, name, sex, age
}

func (UserDao) CreateUser(user User) bool {
	sql := "insert into user(name,sex,age) values(?, ?, ? )"

	stmt, err := common.DB.Prepare(sql)
	delError(err)
	defer stmt.Close()

	rs, err := stmt.Exec(user.name, user.sex, user.age)
	delError(err)

	id, err := rs.LastInsertId()
	delError(err)
	if id > 0 {
		fmt.Printf("本次执行后，生效的数据表id是%d", id)
		return true
	}
	return false
}

func (UserDao) connectDB() *sql.DB {
	fmt.Println("database connecting...")
	//连接数据库
	db, err := sql.Open("mysql", "web:123456@tcp(192.168.1.204:3306)/test")
	if err != nil {
		fmt.Println("database connect error", err)
	}

	db.SetMaxIdleConns(2)
	db.SetMaxOpenConns(100)

	//如果没有访问并不会真正的建立链接 可以用ping方法来验证连接
	err = db.Ping()
	if err != nil {
		fmt.Println("database connect error", err)
	} else {
		fmt.Println("database connect success")
	}

	return db
}

func delError(err error) {
	if (err != nil) {
		fmt.Println("error: ", err)
	}
}