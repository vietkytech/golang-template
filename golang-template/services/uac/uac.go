package uac

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/vietkytech/golang-template/golang-template/config"
	"git.chotot.org/go-common/kit/logger"
)

var log = logger.GetLogger("sendy-http-client")

type UserAccountResponse struct {
	AccountID             int           `json:"account_id"`
	AccountOid            string        `json:"account_oid"`
	Address               string        `json:"address"`
	BPhoneVerified        bool          `json:"b_phone_verified"`
	CreateTime            int           `json:"create_time"`
	CreatedAt             time.Time     `json:"created_at"`
	Email                 string        `json:"email"`
	EmailVerified         string        `json:"email_verified"`
	FacebookToken         string        `json:"facebook_token"`
	Favorites             []interface{} `json:"favorites"`
	FullName              string        `json:"full_name"`
	ID                    int           `json:"id"`
	IsActive              bool          `json:"is_active"`
	IsPayooLinked         bool          `json:"is_payoo_linked"`
	LongTermFacebookToken string        `json:"long_term_facebook_token"`
	OldPhone              string        `json:"old_phone"`
	Password              string        `json:"password"`
	PayooPhone            string        `json:"payoo_phone"`
	Phone                 string        `json:"phone"`
	PhoneVerified         string        `json:"phone_verified"`
	StartTime             int           `json:"start_time"`
	UpdateTime            int           `json:"update_time"`
	UpdatedAt             time.Time     `json:"updated_at"`
}

type IUacClient interface {
	Get(accountID int) (*UserAccountResponse, error)
}

type UacClient struct {
	Config *config.UacConfig
}

func NewUacClient(cfg *config.UacConfig) IUacClient {
	return &UacClient{
		Config: cfg,
	}
}

func (svc *UacClient) Get(accountID int) (*UserAccountResponse, error) {
	apiUrl := fmt.Sprintf("%s/%d", svc.Config.GetProfileApiURL, accountID)

	log.Debugf("request url %+v", apiUrl)
	req, _ := http.NewRequest("GET", apiUrl, nil)

	req.Header.Add("content-type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	if res.StatusCode > 299 {
		return nil, fmt.Errorf("Request error: %v", res.Status)
	}
	var result UserAccountResponse
	err = json.NewDecoder(res.Body).Decode(&result)

	log.Debugf("Response status %+v Error %+v", res.Status, err)

	return &result, err
}
