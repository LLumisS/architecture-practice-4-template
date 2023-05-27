package main

import (
	"testing"

	. "gopkg.in/check.v1"
)

func Test(t *testing.T) { TestingT(t) }

type TestSuite struct{}

var _ = Suite(&TestSuite{})

func (s *TestSuite) TestBalancer(c *C) {
	// TODO: Реалізуйте юніт-тест для балансувальникка.
	address1 := getIndex("127.0.0.1:8080")
	address2 := getIndex("192.168.0.0:80")
	address3 := getIndex("26.143.218.9:80")

	c.Assert(address1, Equals, 2)
	c.Assert(address2, Equals, 0)
	c.Assert(address3, Equals, 1)
}
