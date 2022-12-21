package service

import (
	"goblog/config"
	"goblog/dao"
	"goblog/models"
	"html/template"
	"log"
)

func UpdatePost(post *models.Post) {
	dao.UpdatePost(post)
}
func SavePost(post *models.Post) {
	dao.SavePost(post)

}

func GetPostDetail(pid int) (*models.PostRes, error) {
	post, err := dao.GetPostByPid(pid)

	if err != nil {
		return nil, err
	}
	categoryName := dao.GetCategoryNameById(post.CategoryId)
	userName := dao.GetUserNameById(post.UserId)

	PostMore := models.PostMore{
		post.Pid,
		post.Title,
		post.Slug,
		template.HTML(post.Content),
		post.CategoryId,
		categoryName,
		post.UserId,
		userName,
		post.ViewCount,
		post.Type,
		models.DateDay(post.CreateAt),
		models.DateDay(post.UpdateAt),
	}

	PostMore.Pid = pid
	var postRes = &models.PostRes{
		config.Cfg.Viewer,
		config.Cfg.System,
		PostMore,
	}
	return postRes, nil
}

func Writing() (wr models.WritingRes) {
	wr.Title = config.Cfg.Viewer.Title
	wr.CdnURL = config.Cfg.System.CdnURL
	category, err := dao.GetAllCategory()
	if err != nil {
		log.Println(err)
		return wr
	}
	wr.Categorys = category
	return wr
}
