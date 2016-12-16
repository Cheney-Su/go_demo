package user

import (
	"github.com/kataras/iris"
	"go/src/server/entity"
	"fmt"
	"github.com/bitly/go-simplejson"
)

type UserController struct {
}

var (
	myDao = &UserDao{}
)

func (UserController) GetUserInfoById(ctx *iris.Context) {
	userId := ctx.URLParam("userId")
	id, name, sex, age := myDao.GetUserInfoById(userId)

	m := make(map[string]interface{})
	m["id"] = id
	m["name"] = name
	m["sex"] = sex
	m["age"] = age

	ctx.JSON(iris.StatusOK, entity.Result{0, m, "success"})
	return
}

func (UserController) CreateUser(ctx *iris.Context) {
	//将获取的body转换为json
	json, err := simplejson.NewJson(ctx.PostBody())
	if err != nil {
		fmt.Println("error :", err)
	}
	fmt.Println("jsonBody :", json)
	name := json.Get("name").MustString()
	sex := json.Get("sex").MustInt()
	age := json.Get("age").MustInt()
	user := User{0, name, sex, age}

	result := myDao.CreateUser(user)
	if result {
		ctx.JSON(iris.StatusOK, entity.Result{0, "", "success"})
	} else {
		ctx.JSON(iris.StatusOK, entity.Result{-1, "", "error"})
	}

	return
}