package service

import (
	"goblog/config"
	"goblog/dao"
	"goblog/models"
	"html/template"
)

func GetPostsByCategolyId(cId, page, pageSize int) (*models.CategoryResponse, error) {
	categorys, err := dao.GetAllCategory()
	if err != nil {
		return nil, err
	}

	posts, err := dao.GetPostsByCategolyId(cId, page, pageSize)
	var PostMores []models.PostMore
	for _, post := range posts {
		categoryName := dao.GetCategoryNameById(post.CategoryId)
		userName := dao.GetUserNameById(post.UserId)

		content := []rune(post.Content)
		if len(content) > 100 {
			content = content[0:100]
		}

		PostMore := models.PostMore{
			post.Pid,
			post.Title,
			post.Slug,
			template.HTML(content),
			post.CategoryId,
			categoryName,
			post.UserId,
			userName,
			post.ViewCount,
			post.Type,
			models.DateDay(post.CreateAt),
			models.DateDay(post.UpdateAt),
		}
		PostMores = append(PostMores, PostMore)

	}
	total := dao.CountGetPostsByCategoryId(cId)
	pagesCount := (total-1)/10 + 1
	var pages []int
	for i := 0; i < pagesCount; i++ {
		pages = append(pages, i+1)
	}
	var hr = &models.HomeResponse{
		config.Cfg.Viewer,
		categorys,
		PostMores,
		total,
		page,
		pages,
		page != pagesCount,
	}

	categporyName := dao.GetCategoryNameById(cId)
	var categoryResponse = &models.CategoryResponse{
		hr,
		categporyName,
	}
	return categoryResponse, nil
}
