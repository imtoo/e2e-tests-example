package database

import (
	"database/sql"

	"github.com/imtoo/e2e-tests-example/models"
	"github.com/jinzhu/gorm"
	"syreclabs.com/go/faker"
)

func runSeeds(db *gorm.DB) {
	// Articles
	seedsArticlesAllDefault(db)
	seedsArticlesApprovedForTwitter(db)
	seedsArticlesSentToTwitter(db)
	seedsArticlesSentToSlack(db)
}

func seedsArticlesAllDefault(db *gorm.DB) {
	sum := 0
	for sum < 100 {
		db.Create(&models.Article{
			URL:      faker.Internet().Url(),
			Channel:  faker.Internet().DomainWord(),
			Username: faker.Internet().UserName(),
		})
		sum++
	}
}

func seedsArticlesApprovedForTwitter(db *gorm.DB) {
	approvedForTwitter := true
	sum := 0
	for sum < 100 {
		db.Create(&models.Article{
			URL:                faker.Internet().Url(),
			Channel:            faker.Internet().DomainWord(),
			Username:           faker.Internet().UserName(),
			ApprovedForTwitter: &approvedForTwitter,
			TwitterMessage:     sql.NullString{Valid: true, String: faker.Lorem().Characters(117)},
		})
		sum++
	}
}

func seedsArticlesSentToTwitter(db *gorm.DB) {
	approvedForTwitter := true
	sentToTwitter := true
	sum := 0
	for sum < 100 {
		db.Create(&models.Article{
			URL:                faker.Internet().Url(),
			Channel:            faker.Internet().DomainWord(),
			Username:           faker.Internet().UserName(),
			ApprovedForTwitter: &approvedForTwitter,
			SentToTwitter:      &sentToTwitter,
			TwitterMessage:     sql.NullString{Valid: true, String: faker.Lorem().Characters(117)},
		})
		sum++
	}
}

func seedsArticlesSentToSlack(db *gorm.DB) {
	trueValue := true
	sum := 0
	for sum < 100 {
		db.Create(&models.Article{
			URL:         faker.Internet().Url(),
			Channel:     faker.Internet().DomainWord(),
			Username:    faker.Internet().UserName(),
			SentToSlack: &trueValue,
		})
		sum++
	}
}
