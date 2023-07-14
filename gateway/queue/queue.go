package queue

import "time"

type Partner struct{}

type PaymentInfo struct {
	MegaCC struct {
		ReferenceID string
		CCNumber    string
		RawResponse string
	}
}

type Request struct {
	ID          string
	Source      string
	Amount      string
	Partner     Partner
	PaymentInfo PaymentInfo
}

type Response struct {
	ID          string
	Source      string
	Amount      string
	Partner     Partner
	PaymentInfo PaymentInfo
}

type QueuePayload struct {
	ID          string    `json:"id"`
	CreatedTime time.Time `json:"createdTime"`
	// UpdatedTime        time.Time              `json:"updatedTime"`
	Currency           string                 `json:"currency"`
	Amount             int                    `json:"amount"`
	InquiryID          string                 `json:"inquiryId"`
	MerchantID         string                 `json:"merchantId"`
	Type               string                 `json:"type"`
	PaymentSource      string                 `json:"paymentSource"`
	PaymentSourceData  map[string]interface{} `json:"paymentSourceData"`
	Status             string                 `json:"status"`
	VoidStatus         string                 `json:"voidStatus,omitempty"`
	StatusCode         string                 `json:"statusCode"`
	StatusData         map[string]interface{} `json:"statusData"`
	NetworkReferenceID interface{}            `json:"networkReferenceId"`
	NetworkData        map[string]interface{} `json:"networkData"`
	Inquiry            struct {
		ID string `json:"id"`
		// CreatedTime  time.Time `json:"createdTime"`
		// UpdatedTime  time.Time `json:"updatedTime"`
		Currency     string `json:"currency"`
		Amount       int    `json:"amount"`
		LockedAmount int    `json:"lockedAmount"`
		State        string `json:"state"`
		Status       string `json:"status"`
		// MerchantID   string    `json:"merchantId"`
		ReferenceID  string `json:"referenceId"`
		ReferenceURL string `json:"referenceUrl"`
		Order        struct {
			ID    string `json:"id"`
			Items []struct {
				Name     string `json:"name"`
				Quantity int    `json:"quantity"`
				Amount   int    `json:"amount"`
				ID       string `json:"id,omitempty"`
				Url      string `json:"url,omitempty"`
				Type     string `json:"type,omitempty"`
			} `json:"items"`
		} `json:"order"`
		Customer struct {
			Name        string `json:"name"`
			Email       string `json:"email"`
			PhoneNumber string `json:"phoneNumber"`
			Country     string `json:"country"`
			PostalCode  string `json:"postalCode"`
			Address     string `json:"address,omitempty"`
			Locality    string `json:"locality,omitempty"`
		} `json:"customer"`
	} `json:"inquiry"`
	Merchant struct {
		ID   string `json:"id"`
		Name string `json:"name"`
		// Status                  string `json:"status"`
		PartnerID string `json:"partnerId"`
		// WebhookURL              string `json:"webhookUrl"`
		// WebhookUseSecurity      bool   `json:"webhookUseSecurity"`
		// WebhookExtendedSecurity bool   `json:"webhookExtendedSecurity"`
		// WebhookAutoRedirect     bool   `json:"webhookAutoRedirect"`
		// LogoURL                 string `json:"logoUrl"`
	} `json:"merchant"`
	Token struct {
		Data string `json:"data"`
	} `json:"token,omitempty"`
}
