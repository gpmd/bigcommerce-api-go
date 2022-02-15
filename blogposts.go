package bigcommerce

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

// Post is a BC blog post
type Post struct {
	ID                   int64       `json:"id"`
	Title                string      `json:"title"`
	URL                  string      `json:"url"`
	PreviewURL           string      `json:"preview_url"`
	Body                 string      `json:"body"`
	Tags                 []string    `json:"tags"`
	Summary              string      `json:"summary"`
	IsPublished          bool        `json:"is_published"`
	PublishedDate        interface{} `json:"published_date"`
	PublishedDateISO8601 string      `json:"published_date_iso8601"`
	MetaDescription      string      `json:"meta_description"`
	MetaKeywords         string      `json:"meta_keywords"`
	Author               string      `json:"author"`
	ThumbnailPath        string      `json:"thumbnail_path"`
}

// GetAllPosts downloads all posts from BigCommerce, handling pagination
// context: the BigCommerce context (e.g. stores/23412341234) where 23412341234 is the store hash
// xAuthClient: the BigCommerce Store's X-Auth-Client coming from store credentials (see AuthContext)
// xAuthToken: the BigCommerce Store's X-Auth-Token coming from store credentials (see AuthContext)
func (bc *BigCommerce) GetAllPosts(context, xAuthClient, xAuthToken string) ([]Post, error) {
	cs := []Post{}
	var csp []Post
	page := 1
	more := true
	var err error
	retries := 0
	for more {
		csp, more, err = bc.GetPosts(context, xAuthClient, xAuthToken, page)
		if err != nil {
			log.Println(err)
			retries++
			if retries > bc.MaxRetries {
				log.Println("Max retries reached")
				return cs, fmt.Errorf("max retries reached")
			}
			break
		}
		log.Println("More posts:", more, " count:", len(csp))
		cs = append(cs, csp...)
		page++
	}
	return cs, nil
}

// GetPosts downloads all posts from BigCommerce, handling pagination
// context: the BigCommerce context (e.g. stores/23412341234) where 23412341234 is the store hash
// xAuthClient: the BigCommerce Store's X-Auth-Client coming from store credentials (see AuthContext)
// xAuthToken: the BigCommerce Store's X-Auth-Token coming from store credentials (see AuthContext)
// page: the page number to download
func (bc *BigCommerce) GetPosts(context, xAuthClient, xAuthToken string, page int) ([]Post, bool, error) {
	url := context + "/v2/blog/posts?limit=250&page=" + strconv.Itoa(page)

	req := bc.getAPIRequest(http.MethodGet, url, xAuthClient, xAuthToken)
	var c = &http.Client{
		Timeout: time.Second * 10,
	}
	res, err := c.Do(req)
	if err != nil {
		return nil, false, err
	}

	defer res.Body.Close()

	if res.StatusCode == http.StatusNoContent {
		return []Post{}, false, fmt.Errorf("no content from BigCommerce (status: %d)", res.StatusCode)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, false, err
	}

	var pp []Post
	err = json.Unmarshal(body, &pp)
	if err != nil {
		return nil, false, err
	}
	return pp, len(pp) > 0, nil
}
