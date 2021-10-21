package models

// JSON Data shape
// {
//     "id": "1234",
//     "hash": "1b3035ffe0a93491afd4fbb9a84efd8fd5c0970ea2ae1d0c16415d41d7d02245",
//     "type": "regular",
//     "data": {
//         "inputs": [
//             {
//                 "transaction": "1d0c16415...",
//                 "index": "0",
//                 "amount": 100,
//                 "address": "abcdef123...",
//                 "signature": "1234abcdef..."
//             }
//         ],
//         "outputs": [
//             {
//                 "amount": 100,
//                 "address": "09876abcd..."
//             }
//         ]
//     }
// }

type TransactionTransfer struct {
	Transaction string
	Index       int
	Amount      int
	Address     string
	Signature   string
}

type TransactionTarget struct {
	Amount  int
	Address string
}

type TransactionData struct {
	Inputs  []TransactionTransfer
	Outputs []TransactionTarget
}

type Transaction struct {
	Id   int
	Hash string
	Type string
	Data TransactionData
}

type Transactions []Transaction
