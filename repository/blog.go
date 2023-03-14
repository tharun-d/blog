package repository

import (
	"github.com/tharun-d/blog/models"
)

const (
	saveBlog = `
	INSERT INTO blog.blog (id, title, content, author)
	VALUES ($1, $2, $3, $4)
	`

	getBlogByID = `
	SELECT id, title, content, author FROM blog.blog
	WHERE id = $1
	`

	getAllBlogDetails = `
	SELECT id, title, content, author FROM blog.blog
	`
)

func (r *Repository) SaveBlog(blogData models.Blog) error {

	_, err := r.conn.Exec(saveBlog, blogData.Id,
		blogData.Title,
		blogData.Content,
		blogData.Author,
	)

	if err != nil {
		return err
	}

	return err

}

func (r *Repository) GetBlogByID(id string) (models.Blog, error) {
	var blogData models.Blog

	row := r.conn.QueryRow(getBlogByID, id)
	err := row.Scan(
		&blogData.Id,
		&blogData.Title,
		&blogData.Content,
		&blogData.Author,
	)
	if err != nil {
		return blogData, err
	}

	return blogData, nil

}

func (r *Repository) GetAllBlogDetails() ([]models.Blog, error) {
	results := make([]models.Blog, 0)

	row, err := r.conn.Query(getAllBlogDetails)
	if err != nil {
		return results, err
	}

	for row.Next() {
		var blogData models.Blog
		err := row.Scan(
			&blogData.Id,
			&blogData.Title,
			&blogData.Content,
			&blogData.Author,
		)

		if err != nil {
			return results, err
		}

		results = append(results, blogData)
	}
	return results, nil
}
