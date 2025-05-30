package api

import (
	"io"
	"net/http"
    "fmt"
)

// TODO implement

/*

This file should talk to the broker api to get account information

*/


// Returns aspects of the Broe
func Acct(data string) any {
	url := "https://paper-api.alpaca.markets/v2/account"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("accept", "application/json")
	req.Header.Add("APCA-API-KEY-ID", alpacaApi)
	req.Header.Add("APCA-API-SECRET-KEY", alpacaSec)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

    return body

} // Acct

/*

Account -> return entire account JSON

Ordering
1. place market order
2. place limit order
3. cancel all orders
4. cancel order

Positions
ssss1. get all positions
2. close all positions
3. get position
4. close position
4. close position by shares
*/
