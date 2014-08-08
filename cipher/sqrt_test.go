package cipher_test

import (
	. "github.com/bbengfort/avaux/cipher"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Sqrt", func() {

	var (
		result float64
		err    error
	)

	Context("Positive Numbers", func() {

		BeforeEach(func() {
			result, err = Sqrt(4)
		})

		It("should be able to find the root of 4", func() {
			Expect(result).To(Equal(2.0))
		})

		It("should not error", func() {
			Expect(err).NotTo(HaveOccurred())
		})
	})

	Context("Negative Numbers", func() {

		BeforeEach(func() {
			result, err = Sqrt(-4)
		})

		It("should return the zero-value for result", func() {
			Expect(result).To(BeZero())
		})

		It("should error", func() {
			Expect(err).To(HaveOccurred())
		})
	})

})
