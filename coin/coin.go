package coin

//go:generate rm -f slip44.go
//go:generate go run gen.go

// Coin is the native currency of a blockchain
type Coin struct {
	// SLIP-44 index
	Index    uint   `json:"index"`
	// Symbol of native currency
	Symbol   string `json:"symbol"`
	// Full name of native currency
	Title    string `json:"name"`
	// Project website
	Website  string `json:"link"`
	// Number of decimals
	Decimals uint   `json:"decimals"`
}
