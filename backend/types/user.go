package types

type UserStore interface {
	RegisterOrUpdateUser(wcaUser WCAUser) error
	GetWCAUser(accessToken string) (WCAUser, error)
}

type WCAAvater struct {
	url string
}

type WCAUser struct {
	Name        string    `json:"name"`
	WcaId       string    `json:"wca_id"`
	Gender      string    `json:"gender"`
	CountryIso2 string    `json:"country_iso2"`
	Avatar      WCAAvater `json:"avatar"`
	Email       string    `json:"email"`
}

type WCAMe struct {
	Me WCAUser `json:"me"`
}
