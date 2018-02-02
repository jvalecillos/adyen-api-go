package adyen

/*********
* Amount *
*********/

// Amount value/currency representation
type Amount struct {
	Value    float32 `json:"value"`
	Currency string  `json:"currency"`
}

/*******
* Card *
*******/

// Card structure representation
type Card struct {
	Number      string `json:"number"`
	ExpireMonth string `json:"expiryMonth"`
	ExpireYear  string `json:"expiryYear"`
	Cvc         string `json:"cvc"`
	HolderName  string `json:"holderName"`
}
