package service

import (
	"fmt"
	"goblog/config"
	"goblog/dao"
	"goblog/models"
)

func FindPostPigeonhole() models.Pigeonhole {
	// 查询所有的文章 进行月份的整理
	fmt.Println("已经调用了这个功能")
	panic("已经调用了这个功能")

	posts, _ := dao.GetPostAll()
	pigeonholeMap := make(map[string][]models.Post)
	for _, post := range posts {
		at := post.CreateAt
		month := at.Format("2006-01")
		pigeonholeMap[month] = append(pigeonholeMap[month], post)
	}

	categorys, _ := dao.GetAllCategory()
	fmt.Println("categorys ==", categorys)
	return models.Pigeonhole{
		config.Cfg.Viewer,
		config.Cfg.System,
		categorys,
		pigeonholeMap,
	}
}
