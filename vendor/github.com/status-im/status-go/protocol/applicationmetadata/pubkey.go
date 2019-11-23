package applicationmetadata

import (
	"crypto/ecdsa"
	"github.com/status-im/status-go/eth-node/crypto"
)

func (m *Message) RecoverKey() (*ecdsa.PublicKey, error) {
	if m.Signature == nil {
		return nil, nil
	}

	recoveredKey, err := crypto.SigToPub(
		crypto.Keccak256(m.Payload),
		m.Signature,
	)
	if err != nil {
		return nil, err
	}

	return recoveredKey, nil
}
