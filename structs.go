package testpage

import (
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	//"math"
)

var engine *xorm.Engine

func init() {
	var err error
	engine, err = xorm.NewEngine("mysql", "root:111111@/luowen?charset=utf8")
	if err != nil {
		panic(err)
	}
}

func GetSession() (session *xorm.Session) {
	return engine.NewSession()
}

func Paginate(pageNo, pageSize int32, session *xorm.Session, entity *interface{}) (pageResult *PageResult) {
	pageResult = &PageResult{Pagination: &PageInfo{}, Data: &Result{}}
	var err error
	entitySingle := User{}
	pageResult.Pagination.Total, err = session.Count(&entitySingle)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(pageResult.Pagination.Total / int64(pageSize))
	pageResult.Pagination.TotalPage = 111
	entityList := make([]User, 0)
	session.Find(&entityList)
	pageResult.Data = entityList
	return pageResult
}

func Success(data interface{}) (b []byte) {
	resp := Response{0, "success", nil}
	resp.Data = data
	b, err := json.Marshal(resp)
	if err != nil {
		eb, _ := Error(err, data)
		return eb
	}
	return b
}

func Error(err error, data interface{}) (b []byte, err1 error) {
	var code int32 = -1
	/*
		if err.Code != nil {
			code = err.Code
		}
	*/
	resp := Response{code, err.Error(), data}
	return json.Marshal(resp)
}

type HttpErr struct {
	code int32
	err  error
}

// 对应用户表机构
type User struct {
	Id     uint32 `xorm:"id" json:"id"`
	Name   string `xorm:"name" json:"name"`
	Gender int16  `xorm:"gender" json "gender"`
	Mobile string `xorm:"mobile" json:"mobile"`
}

// 分页信息
type PageInfo struct {
	Total     int64 `json:"total"`
	Page      int64 `json:"page"`
	TotalPage int64 `json:"totalPage"`
}

// 不分也信息
type Result struct {
	Data interface{} `json:"data"`
}

// 分页结果信息
type PageResult struct {
	Pagination *PageInfo   `json:"pagination"`
	Data       interface{} `json:"data"`
}

// 返回结果信息
type Response struct {
	Code int32       `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}
