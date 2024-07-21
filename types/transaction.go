package types

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Transaction struct {
	senderBlockChainAddress    string
	recipientBlockChainAddress string
	value                      float32
}

func NewTransaction(sender, recipient string, value float32) *Transaction {
	return &Transaction{sender, recipient, value}
}

func (t *Transaction) Print() {
	fmt.Printf("%s [Transaction] %s\n",strings.Repeat("-", 32), strings.Repeat("-", 32))
	fmt.Printf(" sender_blockchain_address       %s\n", t.senderBlockChainAddress)
	fmt.Printf(" recipient_blockchain_address    %s\n", t.recipientBlockChainAddress)
	fmt.Printf(" value                           %s.1f\n", t.recipientBlockChainAddress)
}

func (t *Transaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Sender    string  `json:"sender_blockchain_address"`
		Recipient string  `json:"recipient_blockchain_address"`
		Value     float32 `json:"value"`
	}{
		Sender:    t.senderBlockChainAddress,
		Recipient: t.recipientBlockChainAddress,
		Value:     t.value,
	})
}
