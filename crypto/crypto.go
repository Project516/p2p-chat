package crypto

import (
	"crypto/rand"
	"fmt"
	"io"

	"golang.org/x/crypto/nacl/box"
)

// inital key exchange

func ExchangeKeys(r io.Reader, w io.Writer) (*[32]byte, error) {
	pub, priv, err := box.GenerateKey(rand.Reader)
	if err != nil {
		return nil, err
	}
	_, err = w.Write(pub[:])
	if err != nil {
		return nil, err
	}
	var peerPub [32]byte
	_, err = io.ReadFull(r, peerPub[:])
	if err != nil {
		return nil, err
	}
	var sharedKey [32]byte
	box.Precompute(&sharedKey, &peerPub, priv)
	fmt.Println("Encryption established.")
	return &sharedKey, nil
}

// encrypt message function

func Encrypt(plaintext []byte, sharedKey *[32]byte) ([]byte, error) {
	var nonce [24]byte
	_, err := io.ReadFull(rand.Reader, nonce[:])
	if err != nil {
		return nil, err
	}
	ciphertext := box.SealAfterPrecomputation(nonce[:], plaintext, &nonce, sharedKey)
	return ciphertext, nil
}

// decrypt message function

func Decrypt(encrypted []byte, sharedKey *[32]byte) ([]byte, error) {
	if len(encrypted) < 24 {
		return nil, fmt.Errorf("encrypted message to short")
	}
	var nonce [24]byte
	copy(nonce[:], encrypted[:24])
	ciphertext := encrypted[24:]
	plaintext, ok := box.OpenAfterPrecomputation(nil, ciphertext, &nonce, sharedKey)
	if !ok {
		return nil, fmt.Errorf("decryption failed")
	}
	return plaintext, nil
}
