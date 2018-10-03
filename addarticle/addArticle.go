package addarticle

import (
	"fmt"
	"html"
	"net/http"
	"regexp"
	"time"

	"github.com/imtoo/e2e-tests-example/models"
)

// StoreType is store type of the package
type StoreType struct {
	Store *models.StoreType
}

// Handler adds a new article into database.
// This is connected to Slack so we can't use 4xx or 5xx status codes.
// Also we cannot use JSON response.
func (local StoreType) Handler(w http.ResponseWriter, r *http.Request) {
	articleURL := html.EscapeString(r.FormValue("text"))

	channel := html.EscapeString(r.FormValue("channel_name"))
	username := html.EscapeString(r.FormValue("user_name"))

	// Check if url begins with http or https
	regex, errURL := regexp.Compile("^https?://(.*)")
	match := regex.MatchString(articleURL)

	if match == false || errURL != nil {
		message := "Ooops. You have to provide url in correct format (e.g. https://medium.com/article)."
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(message))
		return
	}

	// Check if article with same url was already added within one week
	existingArticle, _ := local.Store.ArticleFindByURL(articleURL)

	if startsBeforeIntervalEnd(existingArticle.CreatedAt, time.Now()) {
		formatted := existingArticle.CreatedAt.Format("02/01/2006 15:04")
		message := fmt.Sprintf("Ooops. <@%s> added same article within last week. Added on %s.", existingArticle.Username, formatted)
		// Because Slack doesn't display 4xx 5xx messages
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(message))
		return
	}

	err := local.Store.ArticleCreate(&models.Article{
		URL:      articleURL,
		Channel:  channel,
		Username: username,
	})

	if err != nil {
		message := "Ooops. Could not save the article."
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(message))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Article has been successfully added!"))
}

// Check if there is an article with same url added within last 7 days
func startsBeforeIntervalEnd(start time.Time, now time.Time) bool {
	end := start.AddDate(0, 0, 7)

	return now.Before(end)
}
