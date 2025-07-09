package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

type PaystackService struct {
	SecretKey string
	BaseURL   string
}

type PaystackInitializeRequest struct {
	Email     string `json:"email"`
	Amount    int    `json:"amount"` // Amount in kobo
	Reference string `json:"reference"`
	Currency  string `json:"currency"`
	Metadata  map[string]interface{} `json:"metadata,omitempty"`
}

type PaystackInitializeResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    struct {
		AuthorizationURL string `json:"authorization_url"`
		AccessCode       string `json:"access_code"`
		Reference        string `json:"reference"`
	} `json:"data"`
}

type PaystackVerifyResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    struct {
		ID              int       `json:"id"`
		Domain          string    `json:"domain"`
		Status          string    `json:"status"`
		Reference       string    `json:"reference"`
		Amount          int       `json:"amount"`
		Message         string    `json:"message"`
		GatewayResponse string    `json:"gateway_response"`
		PaidAt          time.Time `json:"paid_at"`
		CreatedAt       time.Time `json:"created_at"`
		Channel         string    `json:"channel"`
		Currency        string    `json:"currency"`
		IPAddress       string    `json:"ip_address"`
		Metadata        map[string]interface{} `json:"metadata"`
		Customer        struct {
			ID           int    `json:"id"`
			FirstName    string `json:"first_name"`
			LastName     string `json:"last_name"`
			Email        string `json:"email"`
			CustomerCode string `json:"customer_code"`
			Phone        string `json:"phone"`
		} `json:"customer"`
		Authorization struct {
			AuthorizationCode string `json:"authorization_code"`
			Bin               string `json:"bin"`
			Last4             string `json:"last4"`
			ExpMonth          string `json:"exp_month"`
			ExpYear           string `json:"exp_year"`
			Channel           string `json:"channel"`
			CardType          string `json:"card_type"`
			Bank              string `json:"bank"`
			CountryCode       string `json:"country_code"`
			Brand             string `json:"brand"`
			Reusable          bool   `json:"reusable"`
			Signature         string `json:"signature"`
		} `json:"authorization"`
	} `json:"data"`
}

type PaystackRefundRequest struct {
	Transaction string `json:"transaction"`
	Amount      int    `json:"amount,omitempty"` // Optional: partial refund amount in kobo
}

type PaystackRefundResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    struct {
		Transaction struct {
			ID        int    `json:"id"`
			Reference string `json:"reference"`
			Amount    int    `json:"amount"`
			Status    string `json:"status"`
		} `json:"transaction"`
		Integration int       `json:"integration"`
		Domain      string    `json:"domain"`
		Amount      int       `json:"amount"`
		Currency    string    `json:"currency"`
		Status      string    `json:"status"`
		RefundedBy  string    `json:"refunded_by"`
		RefundedAt  time.Time `json:"refunded_at"`
		CreatedAt   time.Time `json:"created_at"`
		UpdatedAt   time.Time `json:"updated_at"`
	} `json:"data"`
}

func NewPaystackService() *PaystackService {
	return &PaystackService{
		SecretKey: os.Getenv("PAYSTACK_SECRET_KEY"),
		BaseURL:   "https://api.paystack.co",
	}
}

func (p *PaystackService) InitializeTransaction(req PaystackInitializeRequest) (*PaystackInitializeResponse, error) {
	url := fmt.Sprintf("%s/transaction/initialize", p.BaseURL)
	
	jsonData, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	httpReq, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}

	httpReq.Header.Set("Authorization", "Bearer "+p.SecretKey)
	httpReq.Header.Set("Content-Type", "application/json")

	client := &http.Client{Timeout: 30 * time.Second}
	resp, err := client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var paystackResp PaystackInitializeResponse
	if err := json.Unmarshal(body, &paystackResp); err != nil {
		return nil, err
	}

	if !paystackResp.Status {
		return nil, fmt.Errorf("paystack error: %s", paystackResp.Message)
	}

	return &paystackResp, nil
}

func (p *PaystackService) VerifyTransaction(reference string) (*PaystackVerifyResponse, error) {
	url := fmt.Sprintf("%s/transaction/verify/%s", p.BaseURL, reference)

	httpReq, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	httpReq.Header.Set("Authorization", "Bearer "+p.SecretKey)

	client := &http.Client{Timeout: 30 * time.Second}
	resp, err := client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var paystackResp PaystackVerifyResponse
	if err := json.Unmarshal(body, &paystackResp); err != nil {
		return nil, err
	}

	if !paystackResp.Status {
		return nil, fmt.Errorf("paystack error: %s", paystackResp.Message)
	}

	return &paystackResp, nil
}

func (p *PaystackService) RefundTransaction(req PaystackRefundRequest) (*PaystackRefundResponse, error) {
	url := fmt.Sprintf("%s/refund", p.BaseURL)
	
	jsonData, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	httpReq, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}

	httpReq.Header.Set("Authorization", "Bearer "+p.SecretKey)
	httpReq.Header.Set("Content-Type", "application/json")

	client := &http.Client{Timeout: 30 * time.Second}
	resp, err := client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var paystackResp PaystackRefundResponse
	if err := json.Unmarshal(body, &paystackResp); err != nil {
		return nil, err
	}

	if !paystackResp.Status {
		return nil, fmt.Errorf("paystack error: %s", paystackResp.Message)
	}

	return &paystackResp, nil
}
