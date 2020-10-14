package models

import (
	"errors"
	"github.com/jinzhu/gorm"
	"time"
)

type Article struct {
	ID      uint64 `gorm:"primary_key"`
	Title   string `gorm:"size:255;unique;not null"`
	Content string `gorm:"size:2000"`
	//Published bool    `gorm:"default:false"`
	Comments  []Comment `gorm:"foreignkey:ArticleId"`
	Tags      []*Tag `gorm:"many2many:tags_articles;"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

// FindArticleByID allows you to find a specific article using its id
func FindArticleByID(uid uint64) (Article, error) {
	var err error
	var article Article
	err = db.Debug().Preload("Comments").First(&article, uid).Error
	if err != nil {
		return Article{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return Article{}, errors.New("Article Not Found")
	}
	return article, nil
}

// FindArticles returns you a list of articles
func FindArticles() ([]Article, error) {
	var err error
	var articles []Article
	err = db.Debug().Find(&articles).Error
	if err != nil {
		return nil, err
	}
	return articles, nil
}

// DeleteArticleByID allows you to remove an article from the db using its id
func DeleteArticleByID(id uint64) error {
	var err error
	var article Article

	err = db.Debug().First(&article, id).Error
	if gorm.IsRecordNotFoundError(err) {
		return errors.New("Article Not Found")
	}
	err = db.Debug().Delete(&article, id).Error
	if err != nil {
		return err
	}

	return nil
}

// EditArticleByID allow you to modify an article using its id
func EditArticleByID(article *Article, id uint64) error {
	var err error
	var old Article
	err = db.Debug().Where("id = ?", id).First(&old).Error
	if gorm.IsRecordNotFoundError(err) {
		return errors.New("Article Not Found")
	}
	article.ID = id
	article.UpdatedAt = time.Now()

	err = db.Debug().Save(&article).Error
	if err != nil {
		return errors.New("Could'nt update article")
	}
	return nil
}

// CreateArticle creates an article row in database
func CreateArticle(article *Article) error {
	var err error
	article.CreatedAt = time.Now()
	err = db.Debug().Create(article).Error
	if err != nil {
		return err
	}
	return nil
}

func AssociateTagsToArticle(article *Article) error {
	err := db.Debug().Association("Tag").Append(&article.Tags).Error
	if err != nil {
		return err
	}
	return nil
}
// FindArticlesByName returns you an article from its name
func FindArticleByName(name string) (Article, error) {
	var err error
	var article Article
	err = db.Debug().Where("name = ?", name).First(&article).Error
	if err != nil {
		return Article{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return Article{}, errors.New("Article Not Found")
	}
	return article, nil
}

// FindArticleByOrderedDate returns you a list of articles sorted by descendant date
func FindArticleByOrderedDate() ([]Article, error) {
	var err error
	var articles []Article
	err = db.Debug().Order("created_at desc").Limit(5).Find(&articles).Error
	if err != nil {
		return nil, err
	}
	return articles, nil
}
