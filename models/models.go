package models

import (
	_ "beego-admin/routers"
	"fmt"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gogo/protobuf/test/data"
	"github.com/syndtr/goleveldb/leveldb/table"
)

type User struct {
    Id int
    Name string
    Profile *Profile `orm:"rel(one)"`
    Post []*Post `orm:"reverse(many)"` // 设置一对多反向关系
}

type Profile struct {
    Id int
    Age int16
    User *User `orm:"reverse(one)"`//设置一对一反向关系(可选)
}

type Post struct {
    Id int
    Title string
    User *User `orm:"rel(fk)"`//设置一对多关系
    Tags []*Tag `orm:"rel(m2m)"`
}

type Tag struct {
    Id int
    Name string
    Posts []*Post `orm:"reverse(many)"`//设置多对多反向关系
}

type Users struct {
	Id int 
	Name string `orm:"size(100)"`
	Number int64 `orm:"size(100)"`
	Authority string `orm:"size(100)"`
	Description string `orm:"size(100)"`
	Token string `orm:"size(100)"`
}

func init() {
    // orm.RegisterDriver("mysql",orm.DRMySQL)
	orm.RegisterDataBase("default","mysql","root:root@/beego-admin?charset=utf8", 30)
	orm.RegisterModel(new(Users))
	orm.RunSyncdb("default", false, true)
}

func Insert(table,data) bool{
    o := orm.NewOrm()
	id, err := o.Insert(&table)
	if err == nil {
		return true
	} 
    return false
}

func Update() {

}

func Delete() {

}

func Select() {

}

func Error() {

}
