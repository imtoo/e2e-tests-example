package models

import (
	"errors"
)

// ArticleCreate - creates an article
func (store *StoreType) ArticleCreate(article *Article) error {
	err := store.DB.Create(article).Error

	return err
}

// ArticleFindByURL - finds an article by URL
func (store *StoreType) ArticleFindByURL(URL string) (*Article, error) {
	var article Article

	notFound := store.DB.Where(Article{URL: URL}).First(&article).RecordNotFound()

	if notFound {
		return &article, errors.New("Record not found")
	}

	return &article, store.DB.Error
}
