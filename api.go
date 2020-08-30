package affise

// API ...
type API struct {
	NetworkID string
	Token     string
}

// New ...
func New(NetworkID string, Token string) *API {
	return &API{
		NetworkID: NetworkID,
		Token:     Token,
	}
}
