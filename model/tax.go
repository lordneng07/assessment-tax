package model

type Tax struct {
	Tax      float64    `json:"tax"`
	Refund   float64    `json:"totalRefund,omitempty"`
	TaxLevel []TaxLevel `json:"taxLevel"`
}
