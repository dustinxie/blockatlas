package ontology

import (
	"bytes"
	"encoding/json"
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/models"
	"testing"
)

var ONTTransfer = `
     {
        "TxnType": 209,
        "ConfirmFlag": 1,
        "Fee": "0.010000000",
        "BlockIndex": 2,
        "TransferList": [
          {
            "FromAddress": "AUyL4TZ1zFEcSKDJrjFnD7vsq5iFZMZqT7",
            "Amount": "2.000000000",
            "ToAddress": "AQ9kzzHNLCcyrPwJuVMrSPgGzqmuQNVwMF",
            "AssetName": "ont"
          }
        ],
        "TxnTime": 1556952450,
        "TxnHash": "4804e1be63ebe1715d6b4a039cc9d84b86cde74c8a8c8411578e6dcadc1e5405",
        "Height": 3411115
      }
`

var expectedONTTransferTrx = models.Tx{
	ID:    "4804e1be63ebe1715d6b4a039cc9d84b86cde74c8a8c8411578e6dcadc1e5405",
	Coin:  coin.ONT,
	From:  "AUyL4TZ1zFEcSKDJrjFnD7vsq5iFZMZqT7",
	To:    "AQ9kzzHNLCcyrPwJuVMrSPgGzqmuQNVwMF",
	Fee:   "10000000",
	Date:  1556952450,
	Type:   "transfer",
	Status: "completed",
	Block: 3411115,
	Meta:  models.Transfer{
		Value: "2",
	},
}

var ONGTransferOfONTTransfer = `
      {
        "TxnType": 209,
        "ConfirmFlag": 1,
        "Fee": "0.010000000",
        "BlockIndex": 2,
        "TransferList": [
          {
            "FromAddress": "AUyL4TZ1zFEcSKDJrjFnD7vsq5iFZMZqT7",
            "Amount": "0.010000000",
            "ToAddress": "AFmseVrdL9f9oyCzZefL9tG6UbviEH9ugK",
            "AssetName": "ong"
          }
        ],
        "TxnTime": 1555341286,
        "TxnHash": "f494d7aeae2b88d313465d1f5f588b213795b38988a0bef182d9d3c2012f6e6e",
        "Height": 2863855
      }
`

var expectedONGTransferOfONTTransferTrx = models.Tx{
	ID:     "f494d7aeae2b88d313465d1f5f588b213795b38988a0bef182d9d3c2012f6e6e",
	Coin:   coin.ONT,
	From:   "AUyL4TZ1zFEcSKDJrjFnD7vsq5iFZMZqT7",
	To:     "AFmseVrdL9f9oyCzZefL9tG6UbviEH9ugK",
	Fee:    "10000000",
	Date:   1555341286,
	Type:   models.TxNativeTokenTransfer,
	Status: "completed",
	Block:  2863855,
	Meta:  models.NativeTokenTransfer{
		Name: "Ontology Gas",
		Symbol: "ONG",
		TokenID: "ong",
		Decimals: 9,
		Value: "0",
		From: "AUyL4TZ1zFEcSKDJrjFnD7vsq5iFZMZqT7",
		To: "AFmseVrdL9f9oyCzZefL9tG6UbviEH9ugK",
	},
}


var ONGTransfer = `
      {
        "TxnType": 209,
        "ConfirmFlag": 1,
        "Fee": "0.010000000",
        "BlockIndex": 1,
        "TransferList": [
          {
            "FromAddress": "AUyL4TZ1zFEcSKDJrjFnD7vsq5iFZMZqT7",
            "Amount": "10.000000000",
            "ToAddress": "AQ9kzzHNLCcyrPwJuVMrSPgGzqmuQNVwMF",
            "AssetName": "ong"
          },
          {
            "FromAddress": "AUyL4TZ1zFEcSKDJrjFnD7vsq5iFZMZqT7",
            "Amount": "0.010000000",
            "ToAddress": "AFmseVrdL9f9oyCzZefL9tG6UbviEH9ugK",
            "AssetName": "ong"
          }
        ],
        "TxnTime": 1556952520,
        "TxnHash": "eccbfd040925a22884d87e73f818f30ab42d06046460b86e9a042a1e9cba7561",
        "Height": 3411141
      }
`
var expectedONGTransferTrx = models.Tx{
	ID:     "eccbfd040925a22884d87e73f818f30ab42d06046460b86e9a042a1e9cba7561",
	Coin:   coin.ONT,
	From:   "AUyL4TZ1zFEcSKDJrjFnD7vsq5iFZMZqT7",
	To:     "AQ9kzzHNLCcyrPwJuVMrSPgGzqmuQNVwMF",
	Fee:    "10000000",
	Date:   1556952520,
	Type:   models.TxNativeTokenTransfer,
	Status: "completed",
	Block:  3411141,
	Meta:  models.NativeTokenTransfer{
		Name: "Ontology Gas",
		Symbol: "ONG",
		TokenID: "ong",
		Decimals: 9,
		Value: "10000000000",
		From: "AUyL4TZ1zFEcSKDJrjFnD7vsq5iFZMZqT7",
		To: "AQ9kzzHNLCcyrPwJuVMrSPgGzqmuQNVwMF",
	},
}

func TestNormalize(t *testing.T) {
	var tests = []struct {
		Transaction  string
		AssetName        string
		Expected 		 models.Tx
	}{
		{ONTTransfer, ONTAssetName, expectedONTTransferTrx},
		{ONGTransferOfONTTransfer, ONGAssetName, expectedONGTransferOfONTTransferTrx},
		{ONGTransfer, ONGAssetName, expectedONGTransferTrx},
	}
	

	for _, test := range tests {
		var trx Tx

		tErr := json.Unmarshal([]byte(test.Transaction), &trx)
		if tErr != nil {
			t.Fatal("Ontology: Can't unmarshall transaction", tErr)
		}

		var readyTx models.Tx
		normTx, ok := Normalize(&trx, test.AssetName)
		if !ok {
			t.Fatal("Ontology: Can't normalize transaction", readyTx)
		}
		readyTx = normTx

		actual, err := json.Marshal(&readyTx)
		if err != nil {
			t.Fatal(err)
		}

		expected, err := json.Marshal(&test.Expected)
		if err != nil {
			t.Fatal(err)
		}

		if !bytes.Equal(actual, expected) {
			println(string(actual))
			println(string(expected))
			t.Error("Transactions not equal")
		}
	}
}


