package dao

import (
	"goblog/models"
	"log"
)

func UpdatePost(post *models.Post) {
	_, err := DB.Exec("update blog_post set title=?,content=?,markdown=?,category_id=?,type=?,slug=?,update_at=? where user_id=?",
		post.Title, post.Content, post.Markdown, post.CategoryId, post.Type, post.Slug, post.UpdateAt, post.UserId)
	if err != nil {
		log.Println(err)
	}

}

func SavePost(post *models.Post) {
	ret, err := DB.Exec("insert into blog_post (title, content ,markdown,category_id,user_id,view_count,type,slug,create_at,update_at)"+
		"values(?,?,?,?,?,?,?,?,?,?)", post.Title, post.Content, post.Markdown, post.CategoryId, post.UserId, post.ViewCount, post.Type, post.Slug, post.CreateAt, post.UpdateAt)
	if err != nil {
		log.Println(err)
	}
	pid, _ := ret.LastInsertId()
	post.Pid = int(pid)
}

func CountGetPostsByCategoryId(cid int) (count int) {
	rows := DB.QueryRow("select count(1) from blog_post where category_id = ?", cid)
	_ = rows.Scan(&count)
	return

}

func CountGetAllPost() (count int) {
	rows := DB.QueryRow("select count(1) from blog_post")
	_ = rows.Scan(&count)
	return

}
func GetPosts(page, pageSize int) ([]models.Post, error) {
	page = (page - 1) * pageSize

	rows, err := DB.Query("select * from blog_post limit?,?", page, pageSize)
	if err != nil {
		return nil, err
	}
	var posts []models.Post
	for rows.Next() {
		var post models.Post
		err := rows.Scan(
			&post.Pid,
			&post.Title,
			&post.Content,
			&post.Markdown,
			&post.CategoryId,
			&post.UserId,
			&post.ViewCount,
			&post.Type,
			&post.Slug,
			&post.CreateAt,
			&post.UpdateAt,
		)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	return posts, err
}

func GetPostAll() ([]models.Post, error) {

	rows, err := DB.Query("select * from blog_post ")
	if err != nil {
		return nil, err
	}
	var posts []models.Post
	for rows.Next() {
		var post models.Post
		err := rows.Scan(
			&post.Pid,
			&post.Title,
			&post.Content,
			&post.Markdown,
			&post.CategoryId,
			&post.UserId,
			&post.ViewCount,
			&post.Type,
			&post.Slug,
			&post.CreateAt,
			&post.UpdateAt,
		)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, err
}

func GetPostsByCategolyId(cid, page, pageSize int) ([]models.Post, error) {
	page = (page - 1) * pageSize

	rows, err := DB.Query("select * from blog_post where category_id =? limit?,?", cid, page, pageSize)
	if err != nil {
		return nil, err
	}
	var posts []models.Post
	for rows.Next() {
		var post models.Post
		err := rows.Scan(
			&post.Pid,
			&post.Title,
			&post.Content,
			&post.Markdown,
			&post.CategoryId,
			&post.UserId,
			&post.ViewCount,
			&post.Type,
			&post.Slug,
			&post.CreateAt,
			&post.UpdateAt,
		)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, err
}

func GetPostByPid(pid int) (models.Post, error) {
	row := DB.QueryRow("select * from blog_post where pid=?", pid)
	var post models.Post
	if row.Err() != nil {
		return post, row.Err()
	}
	err := row.Scan(
		&post.Pid,
		&post.Title,
		&post.Content,
		&post.Markdown,
		&post.CategoryId,
		&post.UserId,
		&post.ViewCount,
		&post.Type,
		&post.Slug,
		&post.CreateAt,
		&post.UpdateAt,
	)
	if err != nil {
		return post, err
	}

	return post, err
}
