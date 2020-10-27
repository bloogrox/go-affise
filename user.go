package affise

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

// User ...
type User struct {
	ID        string   `json:"id"`
	FirstName string   `json:"first_name"`
	LastName  string   `json:"last_name"`
	Email     string   `json:"email"`
	Skype     string   `json:"skype"`
	Roles     []string `json:"roles"`
	APIKey    string   `json:"api_key"`
	UpdatedAt string   `json:"updated_at"`
	CreatedAt string   `json:"created_at"`
}

// Pagination ...
type Pagination struct {
	PerPage    int `json:"per_page"`
	TotalCount int `json:"total_count"`
	Page       int `json:"page"`
}

// UsersListResponse ...
type UsersListResponse struct {
	Status     int        `json:"status"`
	Users      []User     `json:"users"`
	Pagination Pagination `json:"pagination"`
	Error      string     `json:"error"`
}

// UsersListParams ...
type UsersListParams struct {
	Page      int
	Limit     int
	UpdatedAt string
	Q         string
}

// ErrUsersList ...
var ErrUsersList = errors.New("Get UsersList Error")

// UsersList ...
func (a *API) UsersList(params UsersListParams) (*UsersListResponse, error) {
	url := fmt.Sprintf("http://api.%s.affise.com/3.0/admin/users", a.NetworkID)

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, errors.Wrap(err, "UsersList")
	}

	req.Header.Set("API-Key", a.Token)

	q := req.URL.Query()
	if params.Page != 0 {
		q.Add("page", fmt.Sprintf("%d", params.Page))
	}
	if params.Limit != 0 {
		q.Add("limit", fmt.Sprintf("%d", params.Limit))
	}
	if params.UpdatedAt != "" {
		q.Add("updated_at", params.UpdatedAt)
	}
	if params.Q != "" {
		q.Add("q", params.Q)
	}
	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "UsersList")
	}
	defer resp.Body.Close()

	var r UsersListResponse
	err = json.NewDecoder(resp.Body).Decode(&r)
	if err != nil {
		return nil, errors.Wrap(err, "UsersList")
	}

	if r.Status != 1 {
		return nil, errors.Wrap(ErrUsersList, r.Error)
	}

	return &r, nil
}
