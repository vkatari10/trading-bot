package risk

// Declares structs to represent a users own trading account

import (
	"sync"
)

// Account represents a users Alpaca account
type Account struct {
	Cash float64 `json:"cash"`
	Positions []Position
	mu sync.Mutex // may be unused 
} // Account

// Position represents the position of a single asset
type Position struct {
	Price float64 `json:"current_price"`
	Quantity int `json:"qty"`
	TotalCost float64  `json:"cost_basis"`
	AvgCost float64 `json:"entry_price"`
	MarketValue float64 `json:"market_value"`
	TotalPL float64 `json:"unrealized_pl"`
} // Position


func (acct *Account) addPosition(ticker string) {
	acct.mu.Lock()
	defer acct.mu.Lock() 	
} // addPosition


/* Sample Response
ALL VALUES ARE STRINGS (use import strconv)
[
  {
    "asset_id": "b0b6dd9d-8b9b-48a9-ba46-b9d54906e415",
    "symbol": "AAPL",
    "exchange": "NASDAQ",
    "asset_class": "us_equity",
    "asset_marginable": true,
    "qty": "-1",
    "avg_entry_price": "197.91",
    "side": "short",
    "market_value": "-207.6399",
    "cost_basis": "-197.91",
    "unrealized_pl": "-9.7299",
    "unrealized_plpc": "-0.0491632560254661",
    "unrealized_intraday_pl": "2.3701",
    "unrealized_intraday_plpc": "0.0112856530641398",
    "current_price": "207.6399",
    "lastday_price": "210.01",
    "change_today": "-0.0112856530641398",
    "qty_available": "-1"
  }
]
*/

/*

Get this array into Go objects
1. create parser method
2. create new method based on parser
3. accept input from eventloop
4. develop runtime style integration 
5. 



*/
