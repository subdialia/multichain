package digibyte_test

import (
	"github.com/subdialia/multichain/api/address"
	"github.com/subdialia/multichain/chain/bitcoin"
	"github.com/subdialia/multichain/chain/digibyte"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("DigiByte", func() {
	Context("when decoding an address", func() {
		It("should work without errors", func() {
			_, err := bitcoin.NewAddressDecoder(&digibyte.MainNetParams).DecodeAddress(address.Address("DBLsEv4FdFPGrMWzcagDQvoKgUL2CikhMf"))
			Expect(err).NotTo(HaveOccurred())
		})
	})
})
