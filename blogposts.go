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
	ExtID                int64       `json:"id" db:"ext_id"`
	Title                string      `json:"title"`
	URL                  string      `json:"url" db:"url"`
	PreviewURL           string      `json:"preview_url"`
	Body                 string      `json:"body"`
	Tags                 []string    `json:"tags"`
	Summary              string      `json:"summary"`
	IsPublished          bool        `json:"is_published" db:"published"`
	PublishedDate        interface{} `json:"published_date" db:"published_date"`
	PublishedDateISO8601 string      `json:"published_date_iso8601"`
	MetaDescription      string      `json:"meta_description"`
	MetaKeywords         string      `json:"meta_keywords"`
	Author               string      `json:"author" db:"author"`
	ThumbnailPath        string      `json:"thumbnail_path" db:"thumbnail_path"`
	Featured             bool        `json:"-" db:"featured"`
	Assigned             []Product   `json:"assigned"`
	AssignedCats         []Category  `json:"categories"`
	AssignedBrands       []Brand     `json:"brands"`
}

func (bc *BigCommerce) GetAllPosts(context, client, token string, cid int64) ([]Post, error) {
	cs := []Post{}
	var csp []Post
	page := 1
	more := true
	var err error
	retries := 0
	for more {
		csp, more, err = bc.GetPosts(context, client, token, cid, page)
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

func (bc *BigCommerce) GetPosts(context, client, token string, cid int64, page int) ([]Post, bool, error) {
	url := context + "/v2/blog/posts?limit=250&page=" + strconv.Itoa(page)

	req := bc.getAPIRequest(http.MethodGet, url, client, token)
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
