package resolvers

import (
	"reflect"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = Describe("Resolver", func() {
	DescribeTable("Invalid function",
		func(r interface{}, message string) {
			err := validators.run(reflect.TypeOf(r))

			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal(message))
		},

		Entry("Not a function, but boolean", true, "resolver is not a function, got bool"),
		Entry("Not a function, but integer", 1234, "resolver is not a function, got int"),
		Entry("Not a function, but string", "123", "resolver is not a function, got string"),

		Entry("Parameter is string", func(args string) (interface{}, error) { return nil, nil }, "resolver argument must be struct"),
		Entry("Too many parameters", func(args struct{}, foo struct{}, foo2 struct{}) (interface{}, error) { return nil, nil }, "resolver must not have more than two arguments, got 3"),

		Entry("No return value", func() {}, "resolver must have at least one return value"),
		Entry("Non-error return value", func(args struct{}) interface{} { return nil }, "last return value must be an error"),
		Entry("Multiple non-error return values", func(args struct{}) (interface{}, interface{}) { return nil, nil }, "last return value must be an error"),
		Entry("Too many return values", func(args struct{}) (interface{}, error, error) { return nil, nil, nil }, "resolver must not have more than two return values, got 3"),
	)

	DescribeTable("Valid function",
		func(r interface{}) {
			Expect(validators.run(reflect.TypeOf(r))).NotTo(HaveOccurred())
		},

		Entry("With parameter and multiple return values", func(args struct{}) (interface{}, error) { return nil, nil }),
		Entry("With parameter and single return value", func(args struct{}) error { return nil }),
		Entry("Without parameter, but multiple return values", func() (interface{}, error) { return nil, nil }),
		Entry("Without parameter, but single return value", func() error { return nil }),
	)
})
