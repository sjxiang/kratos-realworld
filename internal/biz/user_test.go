package biz

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
)



func TestHashPassword(t  *testing.T) {
	s := hashPassword("123")

	spew.Dump(s)
	panic(1)
} 


func TestVerifyPassword(t  *testing.T) {
	a := assert.New(t)
	a.True(verifyPassword( "$2a$10$dfPui4.TG0DoOXSbHXJmsOTtzpMiN.7afIhqeMWYRIHlz8HMfyHO2", "123"))
}