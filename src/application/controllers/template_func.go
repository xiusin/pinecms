package controllers

import (
	"fmt"
	"reflect"

	"github.com/CloudyKit/jet"
)

//{dede:arclist typeid='26' row='15' titlelen='50' orderby='pubdate'}
func ArcList()  {

}


func Channel()  {

}

//{dede:pagelist listsize='1' listitem='index,end,pre,next,pageno'/}
func PageList()  {

}

// todo 列表页面用与接收参数，应该是控制器内暴露的方法
// {dede:list pagesize ='5'}
// {dede:list perpage='20'}
func List()  {

}

// {dede:tag row='60' sort='new'}
func Tag()  {

}

//  {dede:type typeid='26'}
func Type()  {

}

func Prenext()  {

}

func Global()  {

}

func ChannelArtList()  {

}


func Flink(args jet.Arguments) reflect.Value  {
	//orm := pine.Make("*xorm.Engine").(*xorm.Engine)
	//models.NewLinkModel(orm).GetList()

	fmt.Println("调用了flink", args.NumOfArguments())
	return reflect.Value{}
}
