/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package bbs

import (
	"crypto/rand"
	"errors"
	"fmt"
	"hash"
	"io"

	ml "github.com/IBM/mathlib"
	"golang.org/x/crypto/hkdf"
)

var (
	// nolint:gochecknoglobals
	seedSize = frCompressedSize

	// nolint:gochecknoglobals
	generateKeySalt = "BBS-SIG-KEYGEN-SALT-"
)

// PublicKey defines BLS Public Key.
type PublicKey struct {
	PointG2 *ml.G2
	curve   *ml.Curve
}

// PrivateKey defines BLS Public Key.
type PrivateKey struct {
	FR    *ml.Zr
	curve *ml.Curve
}

// PublicKeyWithGenerators extends PublicKey with a blinding generator h0, a commitment to the secret key w,
// and a generator for each message h.
type PublicKeyWithGenerators struct {
	H0 *ml.G1
	H  []*ml.G1

	w *ml.G2

	MessagesCount int
}

// ToPublicKeyWithGenerators creates PublicKeyWithGenerators from the PublicKey.
func (pk *PublicKey) ToPublicKeyWithGenerators(messagesCount int) (*PublicKeyWithGenerators, error) {
	offset := pk.curve.G2ByteSize + 1

	data := calcData(pk, messagesCount)

	h0 := hashToG1(data, pk.curve)

	h := make([]*ml.G1, messagesCount)

	for i := 1; i <= messagesCount; i++ {
		dataCopy := make([]byte, len(data))
		copy(dataCopy, data)

		iBytes := uint32ToBytes(uint32(i))

		for j := 0; j < len(iBytes); j++ {
			dataCopy[j+offset] = iBytes[j]
		}

		h[i-1] = hashToG1(dataCopy, pk.curve)
	}

	return &PublicKeyWithGenerators{
		H0:            h0,
		H:             h,
		w:             pk.PointG2,
		MessagesCount: messagesCount,
	}, nil
}

func calcData(key *PublicKey, messagesCount int) []byte {
	data := key.PointG2.Bytes()

	data = append(data, 0, 0, 0, 0, 0, 0)

	mcBytes := uint32ToBytes(uint32(messagesCount))

	data = append(data, mcBytes...)

	return data
}

func hashToG1(data []byte, curve *ml.Curve) *ml.G1 {
	var dstG1 = []byte("BLS12381G1_XMD:BLAKE2B_SSWU_RO_BBS+_SIGNATURES:1_0_0")

	return curve.HashToG1WithDomain(data, dstG1)
}

// UnmarshalPrivateKey unmarshals PrivateKey.
func (b *BBSLib) UnmarshalPrivateKey(privKeyBytes []byte) (*PrivateKey, error) {
	if len(privKeyBytes) != frCompressedSize {
		return nil, errors.New("invalid size of private key")
	}

	fr := b.parseFr(privKeyBytes)

	return &PrivateKey{
		FR:    fr,
		curve: b.curve,
	}, nil
}

// Marshal marshals PrivateKey.
func (k *PrivateKey) Marshal() ([]byte, error) {
	bytes := k.FR.Bytes()
	return bytes, nil
}

// PublicKey returns a Public Key as G2 point generated from the Private Key.
func (k *PrivateKey) PublicKey() *PublicKey {
	pointG2 := k.curve.GenG2.Mul(FrToRepr(k.FR))

	return &PublicKey{
		curve:   k.curve,
		PointG2: pointG2,
	}
}

// UnmarshalPublicKey parses a PublicKey from bytes.
func (b *BBSLib) UnmarshalPublicKey(pubKeyBytes []byte) (*PublicKey, error) {
	if len(pubKeyBytes) != b.bls12381G2PublicKeyLen {
		return nil, errors.New("invalid size of public key")
	}

	pointG2, err := b.curve.NewG2FromCompressed(pubKeyBytes)
	if err != nil {
		return nil, fmt.Errorf("deserialize public key: %w", err)
	}

	return &PublicKey{
		PointG2: pointG2,
		curve:   b.curve,
	}, nil
}

// Marshal marshals PublicKey.
func (pk *PublicKey) Marshal() ([]byte, error) {
	pkBytes := pk.PointG2.Compressed()

	return pkBytes, nil
}

// GenerateKeyPair generates BBS+ PublicKey and PrivateKey pair.
func (b *BBSLib) GenerateKeyPair(h func() hash.Hash, seed []byte) (*PublicKey, *PrivateKey, error) {
	if len(seed) != 0 && len(seed) != seedSize {
		return nil, nil, errors.New("invalid size of seed")
	}

	okm, err := generateOKM(seed, h, b.curve)
	if err != nil {
		return nil, nil, err
	}

	privKeyFr := FrFromOKM(okm, b.curve)

	privKey := &PrivateKey{
		FR:    privKeyFr,
		curve: b.curve,
	}
	pubKey := privKey.PublicKey()

	return pubKey, privKey, nil
}

func generateOKM(ikm []byte, h func() hash.Hash, curve *ml.Curve) ([]byte, error) {
	salt := []byte(generateKeySalt)
	info := make([]byte, 2)

	if ikm != nil {
		ikm = append(ikm, 0)
	} else {
		ikm = make([]byte, seedSize+1)

		_, err := rand.Read(ikm)
		if err != nil {
			return nil, err
		}

		ikm[seedSize] = 0
	}

	return newHKDF(h, ikm, salt, info, curve.ScalarByteSize)
}

func newHKDF(h func() hash.Hash, ikm, salt, info []byte, length int) ([]byte, error) {
	reader := hkdf.New(h, ikm, salt, info)
	result := make([]byte, length)

	_, err := io.ReadFull(reader, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
