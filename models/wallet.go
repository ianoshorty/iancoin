package models

type Keypair struct {
	Index     int
	SecretKey string
	PublicKey string
}

type Wallet struct {
	Id       string
	Password string
	Secret   string
	KeyPairs []Keypair
}
