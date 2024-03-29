package cryptor

import (
	"testing"

	"lib/utils/internal"
)

//Hot0zQ23v/8IY2JRo8hcXe5rLQspnD1ZKLStwPAdfWZDH55oYFr9zeEKDrxlGVNOzR5CoSCKa4tAxhDAYHaEQvOvHpEcrY50gZXMByHPCXmX9R8eCZ+Us/Er8TKm2GKC
//uDhqSNy4Ghah7WmTJ8/CuZV1W7vRiraYWdjEBgNZOG9mA/EWD2sqKhpirI6sFs86UH7nmSqbJwHVG8hQ+8QCQk6Rbc0vf8KN1m3VKL+N3WrmoS6xjuKAH5OUWOkZfnZU
//PRqbV91Wze4pTcKYjvIpSiy7n++D2+Ug3oMoCqBjKk02kHbsxQydOOG0D+NL52qOyRyI7I3qgWc3kJRxnGR8mzyzZOFG49mha91QcXv37sP0+ib3J51LdexHwGRaSrbW
func TestAesEcbEncrypt(t *testing.T) {
	data := "hello world"
	key := "abcdefghijklmnop"

	aesEcbEncrypt := AesEcbEncrypt([]byte(data), []byte(key))
	aesEcbDecrypt := AesEcbDecrypt(aesEcbEncrypt, []byte(key))

	assert := internal.NewAssert(t, "TestAesEcbEncrypt")
	assert.Equal(data, string(aesEcbDecrypt))
}

func TestAesCbcEncrypt(t *testing.T) {
	data := "hello world"
	key := "abcdefghijklmnop"

	aesCbcEncrypt := AesCbcEncrypt([]byte(data), []byte(key))
	aesCbcDecrypt := AesCbcDecrypt(aesCbcEncrypt, []byte(key))

	assert := internal.NewAssert(t, "TestAesCbcEncrypt")
	assert.Equal(data, string(aesCbcDecrypt))
}

func TestAesCtrCrypt(t *testing.T) {
	data := "hello world"
	key := "abcdefghijklmnop"

	aesCtrCrypt := AesCtrCrypt([]byte(data), []byte(key))
	aesCtrDeCrypt := AesCtrCrypt(aesCtrCrypt, []byte(key))

	assert := internal.NewAssert(t, "TestAesCtrCrypt")
	assert.Equal(data, string(aesCtrDeCrypt))
}

func TestAesCfbEncrypt(t *testing.T) {
	data := "hello world"
	key := "abcdefghijklmnop"

	aesCfbEncrypt := AesCfbEncrypt([]byte(data), []byte(key))
	aesCfbDecrypt := AesCfbDecrypt(aesCfbEncrypt, []byte(key))

	assert := internal.NewAssert(t, "TestAesCfbEncrypt")
	assert.Equal(data, string(aesCfbDecrypt))
}

func TestAesOfbEncrypt(t *testing.T) {
	data := "hello world"
	key := "abcdefghijklmnop"

	aesOfbEncrypt := AesOfbEncrypt([]byte(data), []byte(key))
	aesOfbDecrypt := AesOfbDecrypt(aesOfbEncrypt, []byte(key))

	assert := internal.NewAssert(t, "TestAesOfbEncrypt")
	assert.Equal(data, string(aesOfbDecrypt))
}
