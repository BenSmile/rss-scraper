package models

import (
	"time"

	"github.com/bensmile/rssaggregator/internal/database"
	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	ApiKey    string    `json:"api_key"`
}

type Feed struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	Url       string    `json:"url"`
	UserID    uuid.UUID `json:"user_id"`
}

type FeedFollow struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	UserID    uuid.UUID `json:"user_id"`
	FeedID    uuid.UUID `json:"feed_id"`
}

type Post struct {
	ID          uuid.UUID `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Publishedat time.Time `json:"published_at"`
	Url         string    `json:"url"`
	FeedID      uuid.UUID `json:"feed_id"`
}

func DbUser2User(dbUser database.User) User {
	return User{
		ID:        dbUser.ID,
		Name:      dbUser.Name,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
		ApiKey:    dbUser.ApiKey,
	}
}

func DbPost2Post(dbPost database.Post) Post {
	return Post{
		ID:          dbPost.ID,
		Title:       dbPost.Title,
		CreatedAt:   dbPost.CreatedAt,
		UpdatedAt:   dbPost.UpdatedAt,
		Description: dbPost.Description.String,
		Publishedat: dbPost.Publishedat,
		FeedID:      dbPost.FeedID,
		Url:         dbPost.Url,
	}
}

func DbFeed2Feed(dbFeed database.Feed) Feed {
	return Feed{
		ID:        dbFeed.ID,
		Name:      dbFeed.Name,
		CreatedAt: dbFeed.CreatedAt,
		UpdatedAt: dbFeed.UpdatedAt,
		Url:       dbFeed.Url,
		UserID:    dbFeed.UserID,
	}
}

func DbFeedFollow2FeedFollow(dbFeedFollow database.FeedFollow) FeedFollow {
	return FeedFollow{
		ID:        dbFeedFollow.ID,
		CreatedAt: dbFeedFollow.CreatedAt,
		UpdatedAt: dbFeedFollow.UpdatedAt,
		UserID:    dbFeedFollow.UserID,
		FeedID:    dbFeedFollow.FeedID,
	}
}

func DbFeeds2Feeds(feedsDB []database.Feed) []Feed {

	feeds := []Feed{}

	for _, f := range feedsDB {
		feeds = append(feeds, DbFeed2Feed(f))
	}
	return feeds
}

func DbFeedFollows2FeedFollows(feedFollowsDB []database.FeedFollow) []FeedFollow {

	feeds := []FeedFollow{}

	for _, f := range feedFollowsDB {
		feeds = append(feeds, DbFeedFollow2FeedFollow(f))
	}
	return feeds
}

func DbPosts2Posts(postsDB []database.Post) []Post {

	posts := []Post{}

	for _, f := range postsDB {
		posts = append(posts, DbPost2Post(f))
	}
	return posts
}
