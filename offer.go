package affise

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

// Offer ...
type Offer struct {
	ID         int       `json:"id"`
	OfferID    string    `json:"offer_id"`
	Advertiser string    `json:"advertiser"`
	Title      string    `json:"title"`
	URL        string    `json:"url"`
	PreviewURL string    `json:"preview_url"`
	Sources    []string  `json:"sources"`
	Logo       string    `json:"logo"`
	Status     string    `json:"status"`
	Payments   []Payment `json:"payments"`
	Landings   []Landing `json:"landings"`
}

// Payment ...
type Payment struct {
	Countries []string `json:"countries"`
	Goal      string   `json:"goal"`
	Revenue   float64  `json:"total"`
	Payout    float64  `json:"revenue"`
	Currency  string   `json:"currency"`
	Type      string   `json:"type"`
}

// Landing ...
type Landing struct {
	ID         uint64 `json:"id"`
	Title      string `json:"title"`
	URL        string `json:"url"`
	PreviewURL string `json:"url_preview"`
	Type       string `json:"type"`
}

// OfferGetResponse ...
type OfferGetResponse struct {
	Status int    `json:"status"`
	Offer  Offer  `json:"offer"`
	Error  string `json:"error"`
}

// ErrOfferGet ...
var ErrOfferGet = errors.New("OfferGet Error")

// OfferGet ...
func (a *API) OfferGet(ID int) (*Offer, error) {
	url := fmt.Sprintf("http://api.%s.affise.com/3.0/offer/%d", a.NetworkID, ID)

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, errors.Wrap(err, "OfferGet")
	}

	req.Header.Set("API-Key", a.Token)

	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "OfferGet")
	}
	defer resp.Body.Close()

	var r OfferGetResponse
	err = json.NewDecoder(resp.Body).Decode(&r)
	if err != nil {
		return nil, errors.Wrap(err, "OfferGet")
	}

	if r.Status != 1 {
		return nil, errors.Wrap(ErrOfferGet, r.Error)
	}

	return &r.Offer, nil
}

// OfferEditResponse ...
type OfferEditResponse struct {
	Status int    `json:"status"`
	Error  string `json:"error"`
}

// ErrOfferEdit ...
var ErrOfferEdit = errors.New("Offer Edit Error")

// OfferEdit ...
func (a *API) OfferEdit(ID int, data *url.Values) error {
	url := fmt.Sprintf("http://api.%s.affise.com/3.0/admin/offer/%d", a.NetworkID, ID)

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPost, url, strings.NewReader(data.Encode()))
	if err != nil {
		return errors.Wrap(err, "OfferEdit")
	}

	req.Header.Set("API-Key", a.Token)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	resp, err := client.Do(req)
	if err != nil {
		return errors.Wrap(err, "OfferEdit")
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return errors.Wrap(err, "OfferEdit")
	}
	var r OfferEditResponse
	err = json.Unmarshal(b, &r)
	if err != nil {
		return errors.Wrap(err, "OfferEdit")
	}

	if r.Status != 1 {
		return errors.Wrap(ErrOfferEdit, r.Error)
	}

	return nil
}
