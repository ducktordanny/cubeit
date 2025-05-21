package types

import (
	"time"
)

type UserStore interface {
	RegisterOrUpdateUser(wcaUser WCAUser) (User, error)
	GetWCAUser(accessToken string) (WCAUser, error)
}

type User struct {
	Id         int64     `json:"id"`
	WcaId      *string   `json:"wcaId"`
	Name       string    `json:"name"`
	Email      string    `json:"email"`
	Gender     string    `json:"gender"`
	Bio        string    `json:"bio"`
	CountryISO string    `json:"countryISO"`
	AvatarURL  string    `json:"avatarURL"`
	Role       string    `json:"role"`
	CreatedAt  time.Time `json:"createdAt"`
}

type WCAAvatar struct {
	Url string
}

type WCAUser struct {
	Id          int64     `json:"id"`
	Name        string    `json:"name"`
	WcaId       *string   `json:"wca_id"`
	Gender      string    `json:"gender"`
	CountryIso2 string    `json:"country_iso2"`
	Avatar      WCAAvatar `json:"avatar"`
	Email       string    `json:"email"`
}

type WCAMe struct {
	Me WCAUser `json:"me"`
}
