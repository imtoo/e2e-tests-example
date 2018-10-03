package addarticle

import (
	"net/http"
	"net/url"
	"strings"
	"testing"
	"time"

	"github.com/gorilla/mux"
	"github.com/imtoo/e2e-tests-example/config"
	"github.com/imtoo/e2e-tests-example/models"
	"github.com/imtoo/e2e-tests-example/testhelpers"
	"github.com/stretchr/testify/assert"
)

func init() {
	testhelpers.RegisterTxDB("txdb")
}

func TestHandler(t *testing.T) {
	assert := assert.New(t)
	db, _ := testhelpers.PrepareTestDB("txdb")
	defer testhelpers.CleanTestDB(db)

	store := &models.StoreType{DB: db}
	handlerType := StoreType{Store: store}

	router := mux.NewRouter()
	data := url.Values{}
	data.Add("channel_name", "channel")
	data.Add("surname", "bar")
	data.Add("user_name", "user")
	data.Add("text", "someurl.com")
	req, _ := http.NewRequest("POST", config.RouteArticleAdd, strings.NewReader(data.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	router.HandleFunc(config.RouteArticleAdd, handlerType.Handler)

	rr := testhelpers.ExecuteRequest(req, router)
	// StatusCreated because Slack doesn't display different status codes
	assert.Equal(http.StatusCreated, rr.Code, "Status code should be StatusCreated")
	assert.Equal(
		"Ooops. You have to provide url in correct format (e.g. https://medium.com/article).",
		rr.Body.String(),
		"Error message should start with Ooops",
	)

	article := &models.Article{
		URL: "https://someurl.com",
	}
	db.Create(&article)
	data.Set("text", "https://someurl.com")
	req, _ = http.NewRequest("POST", config.RouteArticleAdd, strings.NewReader(data.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	rr = testhelpers.ExecuteRequest(req, router)
	assert.Equal(http.StatusCreated, rr.Code, "Status code should be StatusCreated")
	assert.Contains(
		rr.Body.String(),
		"added same article within last week",
		"Error message should contain that the article was already added",
	)

	data.Set("text", "https://someUniqueUrl.com")
	req, _ = http.NewRequest("POST", config.RouteArticleAdd, strings.NewReader(data.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	rr = testhelpers.ExecuteRequest(req, router)
	assert.Equal(http.StatusCreated, rr.Code, "Status code should be StatusCreated")
	assert.Equal(
		"Article has been successfully added!",
		rr.Body.String(),
		"Text messsage should be article successfully added!",
	)
	createdArticle, err := handlerType.Store.ArticleFindByURL("https://someUniqueUrl.com")
	assert.Nil(err, "error should be nil")
	assert.Equal("https://someUniqueUrl.com", createdArticle.URL, "a new article should be created")
}

func TestStartsBeforeIntervalEnd(t *testing.T) {
	assert := assert.New(t)
	now := time.Date(2017, 11, 21, 11, 47, 0, 0, time.UTC)

	startTrue := time.Date(2017, 11, 18, 11, 47, 0, 0, time.UTC)
	resultTrue := startsBeforeIntervalEnd(startTrue, now)
	assert.Equal(true, resultTrue, "resultTrue should be true")

	startFalse := time.Date(2017, 11, 12, 11, 47, 0, 0, time.UTC)
	resultFalse := startsBeforeIntervalEnd(startFalse, now)
	assert.Equal(false, resultFalse, "resultFalse should be false")
}
