package testdata_test

import (
	"bytes"
	"fmt"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/mokiat/hacktherank/internal/testdata"
)

var _ = Describe("Parser", func() {
	var validData *bytes.Buffer

	BeforeEach(func() {
		validData = &bytes.Buffer{}
		fmt.Fprintln(validData, "input:")
		fmt.Fprintln(validData, "first-sample-first-input")
		fmt.Fprintln(validData, "first-sample-second-input")
		fmt.Fprintln(validData, "output:")
		fmt.Fprintln(validData, "first-sample-first-output")
		fmt.Fprintln(validData, "first-sample-second-output")
		fmt.Fprintln(validData, "input:")
		fmt.Fprintln(validData, "second-sample-first-input")
		fmt.Fprintln(validData, "output:")
		fmt.Fprintln(validData, "second-sample-first-output")
		fmt.Fprintln(validData, "second-sample-second-output")
		fmt.Fprintln(validData, "second-sample-third-output")
	})

	It("parses all samples", func() {
		samples, err := testdata.Parse(validData)
		Expect(err).ToNot(HaveOccurred())
		Expect(samples).To(HaveLen(2))
		Expect(samples[0]).To(Equal(testdata.Sample{
			InputLines: []string{
				"first-sample-first-input",
				"first-sample-second-input",
			},
			OutputLines: []string{
				"first-sample-first-output",
				"first-sample-second-output",
			},
		}))
		Expect(samples[1]).To(Equal(testdata.Sample{
			InputLines: []string{
				"second-sample-first-input",
			},
			OutputLines: []string{
				"second-sample-first-output",
				"second-sample-second-output",
				"second-sample-third-output",
			},
		}))
	})
})
