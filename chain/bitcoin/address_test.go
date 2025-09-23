package bitcoin_test

import (
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/subdialia/multichain/api/address"
	"github.com/subdialia/multichain/chain/bitcoin"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Bitcoin", func() {
	Context("when decoding address with invalid characters", func() {
		It("should not panic", func() {
			decoder := bitcoin.NewAddressDecoder(&chaincfg.MainNetParams)
			Expect(func() {
				_, err := decoder.DecodeAddress(address.Address(rune(256)))
				Expect(err).To(HaveOccurred())
			}).ToNot(Panic())
		})
	})
})
