package aion

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/trustwallet/blockatlas/models"
	"net/http"
	"net/url"
	"strconv"
)

type Client struct {
	HTTPClient *http.Client
	BaseURL    string
}

func (c *Client) GetTxsOfAddress(address string) (*TxPage, error) {
	uri := fmt.Sprintf("%s/getTransactionsByAddress?%s",
		c.BaseURL,
		url.Values{
			"accountAddress": {address},
			"size":           {strconv.FormatInt(models.TxPerPage, 10)},
		}.Encode())

	res, err := c.HTTPClient.Get(uri)
	if err != nil {
		logrus.WithError(err).Errorf("Aion: Failed to get transactions for address %s", address)
	}
	defer res.Body.Close()

	txPage := new(TxPage)
	err = json.NewDecoder(res.Body).Decode(txPage)
	return txPage, err
}

