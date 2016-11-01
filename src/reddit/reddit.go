package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"golang.org/x/net/context"
	"golang.org/x/oauth2/clientcredentials"
)

var clientID = "-7lHkwD_f6RFbA"
var clientSecret = "-Xb6gT2mP3l5u-yu4i5hoCHVq-8"
var redirectURL = "http://localhost:3000/redirect"

//authURL, tokenURL are used by oauth2
var authURL = "https://www.reddit.com/api/v1/authorize"
var tokenURL = "https://www.reddit.com/api/v1/access_token"

var baseURL = "http://reddit.com"
var userAgent = "golang-lucy v.02 for machine learning and subreddit modelling (by anmousyony)"
var scopes = []string{""}

func oAuth2(clientID, clientSecret, tokenURL string, scopes []string) *http.Client {
	var NoContext = context.TODO()
	config := &clientcredentials.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		TokenURL:     tokenURL,
		Scopes:       scopes,
	}
	client := config.Client(NoContext)
	return client
}

type Post struct {
	ApprovedBy          string        `json:"approved_by"`
	Archived            bool          `json:"archived"`
	Author              string        `json:"author"`
	AuthorFlairCSSClass string        `json:"author_flair_css_class"`
	AuthorFlairText     string        `json:"author_flair_text"`
	BannedBy            string        `json:"banned_by"`
	Clicked             bool          `json:"clicked"`
	ContestMode         bool          `json:"contest_mode"`
	Created             int           `json:"created"`
	CreatedUtc          int           `json:"created_utc"`
	Distinguished       string        `json:"distinguished"`
	Domain              string        `json:"domain"`
	Downs               int           `json:"downs"`
	Edited              bool          `json:"edited"`
	Gilded              int           `json:"gilded"`
	Hidden              bool          `json:"hidden"`
	HideScore           bool          `json:"hide_score"`
	ID                  string        `json:"id"`
	IsSelf              bool          `json:"is_self"`
	Likes               bool          `json:"likes"`
	LinkFlairCSSClass   string        `json:"link_flair_css_class"`
	LinkFlairText       string        `json:"link_flair_text"`
	Locked              bool          `json:"locked"`
	Media               Media         `json:"media"`
	MediaEmbed          interface{}   `json:"media_embed"`
	ModReports          []interface{} `json:"mod_reports"`
	Name                string        `json:"name"`
	NumComments         int           `json:"num_comments"`
	NumReports          int           `json:"num_reports"`
	Over18              bool          `json:"over_18"`
	Permalink           string        `json:"permalink"`
	Quarantine          bool          `json:"quarantine"`
	RemovalReason       interface{}   `json:"removal_reason"`
	ReportReasons       []interface{} `json:"report_reasons"`
	Saved               bool          `json:"saved"`
	Score               int           `json:"score"`
	SecureMedia         interface{}   `json:"secure_media"`
	SecureMediaEmbed    interface{}   `json:"secure_media_embed"`
	SelftextHTML        string        `json:"selftext_html"`
	Selftext            string        `json:"selftext"`
	Stickied            bool          `json:"stickied"`
	Subreddit           string        `json:"subreddit"`
	SubredditID         string        `json:"subreddit_id"`
	SuggestedSort       string        `json:"suggested_sort"`
	Thumbnail           string        `json:"thumbnail"`
	Title               string        `json:"title"`
	URL                 string        `json:"url"`
	Ups                 int           `json:"ups"`
	UserReports         []interface{} `json:"user_reports"`
	Visited             bool          `json:"visited"`
}

type postListing struct {
	Kind string `json:"kind"`
	Data struct {
		Modhash  string `json:"modhash"`
		Children []struct {
			Kind string `json:"kind"`
			Data Link   `json:"data"`
		} `json:"children"`
		After  string      `json:"after"`
		Before interface{} `json:"before"`
	} `json:"data"`
}

func getPosts() {
	client := oAuth2(clientID, clientSecret, tokenURL, scopes)
	var sort = "hot"
	var subreddit = "news"
	url := fmt.Sprintf("%s/r/%s/%s.json", baseURL, subreddit, sort)
	resp, err := client.Get(URL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result postListing
	err = json.NewDecoder(resp.body).Decode(&result)
	if err != nil {
		return nil, err
	}
	var posts *[]Posts
	for _, postt := range result.Data.Children {
		posts = append(posts, &post.Data)
	}
	return posts, nil
}

func main() {
	posts := getPosts()
	for _, post := range posts {
		fmt.Printf(post.Title, "\n")
	}
}
