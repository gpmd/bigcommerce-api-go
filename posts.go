package bigcommerce

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
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
	PublishedDate        interface{} `json:"publisheddate"`
	PublishedDateISO8601 string      `json:"publisheddate_iso8601"`
	MetaDescription      string      `json:"meta_description"`
	MetaKeywords         string      `json:"meta_keywords"`
	Author               string      `json:"author"`
	ThumbnailPath        string      `json:"thumbnailpath"`
}

// GetAllPosts downloads all posts from BigCommerce, handling pagination
// context: the BigCommerce context (e.g. stores/23412341234) where 23412341234 is the store hash
// xAuthToken: the BigCommerce Store's X-Auth-Token coming from store credentials (see AuthContext)
func (bc *BigCommerce) GetAllPosts(context, xAuthToken string) ([]Post, error) {
	cs := []Post{}
	var csp []Post
	page := 1
	more := true
	var err error
	retries := 0
	for more {
		csp, more, err = bc.GetPosts(page)
		if err != nil {
			retries++
			if retries > bc.MaxRetries {
				log.Println("Max retries reached")
				return cs, fmt.Errorf("max retries reached")
			}
			break
		}
		cs = append(cs, csp...)
		page++
	}
	return cs, err
}

// GetPosts downloads all posts from BigCommerce, handling pagination
// context: the BigCommerce context (e.g. stores/23412341234) where 23412341234 is the store hash
// xAuthToken: the BigCommerce Store's X-Auth-Token coming from store credentials (see AuthContext)
// page: the page number to download
func (bc *BigCommerce) GetPosts(page int) ([]Post, bool, error) {
	url := "/v2/blog/posts?limit=250&page=" + strconv.Itoa(page)

	req := bc.getAPIRequest(http.MethodGet, url, nil)
	res, err := bc.HTTPClient.Do(req)
	if err != nil {
		return nil, false, err
	}

	defer res.Body.Close()
	body, err := processBody(res)
	if err != nil {
		return nil, false, err
	}

	var pp []Post
	err = json.Unmarshal(body, &pp)
	if err != nil {
		log.Printf("Error unmarshalling posts: %s %s", err, string(body))
		return nil, false, err
	}
	return pp, len(pp) == 250, nil
}