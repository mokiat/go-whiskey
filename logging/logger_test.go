package logging_test

import (
	"log"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/mokiat/go-whiskey/logging"
	"github.com/mokiat/go-whiskey/logging/logging_stubs"
)

var _ = Describe("Logger", func() {
	It("root Printf delegates to DefaultLogger", func() {
		originalLogger := logging.DefaultLogger
		defer func() { logging.DefaultLogger = originalLogger }()
		fakeLogger := new(logging_stubs.LoggerStub)
		logging.DefaultLogger = fakeLogger

		logging.Printf("hello %s", "world")

		Ω(fakeLogger.PrintfCallCount()).Should(Equal(1))
		argMsg, argArgs := fakeLogger.PrintfArgsForCall(0)
		Ω(argMsg).Should(Equal("hello %s"))
		Ω(argArgs).Should(Equal([]interface{}{"world"}))
	})

	It("default implementation is NopLogger", func() {
		_, ok := logging.DefaultLogger.(logging.NopLogger)
		Ω(ok).Should(BeTrue())
	})

	It("native go Logger implements Logger", func() {
		var logger interface{} = log.New(nil, "", 0)
		_, ok := logger.(logging.Logger)
		Ω(ok).Should(BeTrue())
	})
})
