package models

// JSON Data shape
// { // Block
//     "index": 0,
//     "previousHash": "0",
//		 "hash": "1b3035ffe0a93491afd4fbb9a84efd8fd5c0970ea2ae1d0c16415d41d7d02245"
//     "timestamp": 1634845388,
//     "nonce": 0,
//     "transactions": [
//         {
//             "id": "1d0c16415...",
//             "hash": "abcdef123...",
//             "type": "regular",
//             "data": {
//                 "inputs": [],
//                 "outputs": []
//             }
//         }
//     ],
// }

// Block is a single entry in the blockchain
type Block struct {
	Index        int
	PreviousHash string
	Hash         string
	Timestamp    string
	Nonce        int
	Transactions
}

// Blocks are a collection of entries in the blockchain, or the whole chain
type Blocks []Block
