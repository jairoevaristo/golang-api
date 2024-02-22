package entities

import (
	"time"

	"github.com/jairoevaristo/cmd/pkg"
)

type Link struct {
	Id          string    `json:"id"`
	OriginalUrl string    `json:"original_url"`
	ShortUrl    string    `json:"short_url"`
	Hash        string    `json:"hash"`
	CreatedAt   time.Time `json:"created_at"`
}

func NewLink(url string) (*Link, error) {
	hash := pkg.NewHash(5)

	uuid, err := pkg.NewUUID()
	if err != nil {
		return nil, err
	}

	link := &Link{
		Id:          uuid,
		OriginalUrl: url,
		ShortUrl:    "http://localhost:3333/" + hash,
		Hash:        hash,
		CreatedAt:   time.Now(),
	}

	return link, nil
}
