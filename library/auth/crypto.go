package auth

import (
	"encoding/base64"
	"encoding/json"

	"github.com/farmerx/gorsa"
)

var PubKey = `-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA9dLY5L0zL16T71oQbq/0
rjv2ahFeosH0ICDpE79lHWS8Ri8X+iom0klSuVTpJSskLn0/SQuI9HJmKiXunPLU
lO/CGp2doOQXtz09QJjPEP97tA6hhKDY9kZ2A0kiOj/UYeCzyynp87PZ1zMbPXCL
NRNNnWDvs07G45Dj088qV8Tch/DREXMGLbpHeAf737h2tSGHas3GZZ+ZjUOW5WIM
WfuI17j9J+KjLusK0sjUZ5xE+9MwvMt9tXx59O+ZgwtusbuD6Mvc8YmdLt5zBIhD
QRtUbCOUaVoWnMw2DLyuSEqRjxQaOH2cT+Kq7vMdiRhTeSxlnspX6o0GucJw16cC
ywIDAQAB
-----END PUBLIC KEY-----
`

var PrivateKey = `-----BEGIN PRIVATE KEY-----
MIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQD10tjkvTMvXpPv
WhBur/SuO/ZqEV6iwfQgIOkTv2UdZLxGLxf6KibSSVK5VOklKyQufT9JC4j0cmYq
Je6c8tSU78IanZ2g5Be3PT1AmM8Q/3u0DqGEoNj2RnYDSSI6P9Rh4LPLKenzs9nX
Mxs9cIs1E02dYO+zTsbjkOPTzypXxNyH8NERcwYtukd4B/vfuHa1IYdqzcZln5mN
Q5blYgxZ+4jXuP0n4qMu6wrSyNRnnET70zC8y321fHn075mDC26xu4Poy9zxiZ0u
3nMEiENBG1RsI5RpWhaczDYMvK5ISpGPFBo4fZxP4qru8x2JGFN5LGWeylfqjQa5
wnDXpwLLAgMBAAECggEBANdjgnWRrZqQxRBQmttRQxOzKGqxg3kadlv8Whzac694
h1r7gofIjTFt25TV6F6P2Wj+hwfqmnqNDiVaDSPpxQgmt614cxf5IyqD9vp1qeEb
6fjPZQh/ovrOrDKezuW47c9BAmg99mZ5LKNZfUGtLdw4aKY/vGcg3DWiy2zYNSKm
fHQ98nXS+Jcqpwye3Bmjd0qBFTGHzIJIPwT0Gbq6gEhHnxax60TCrP0OVLTVDXj4
O84Jf6sj2RPXfO0jVcCsSaSYL3KF8n8PirWkSQElDDLuGH2fYI7LIeWqQz0xyYXc
n6ywqEajHpH7Zr4n6FxnN7jkVEvuLlcsAkcu0Sjit0kCgYEA+yQr2aQIhvVwh0rx
mcrAdHUwH6netawJRVXl1H8YENaLXN016viDNdwTMKZzYvgwEKEl2DYK1UgzIlwm
KNZbwjmlpjg8UHe6Br33v0ZFY2aiOA77ruj2AxwipYduvbOn6ifjgDHiQ6FoRVtv
+P+bEWw10S6FnBju/lSVE9Z+zI8CgYEA+pRWzp5bB6EqMw6Li13Wzgpj3XInbfHC
zDBBRfOEe3DbaAYd06XLcKjUlwYde64R3+XeMeujWDsRQ6OucS2kbu3sgM0VCfIf
ZfwRDlpinESexFEaA1Wso2htW5QGbrGzqe8U3iGt6wmaSqRf4YZ7XyTc8voPBU9L
mGItK4nMvAUCgYBnxp2cFI01f2xda5mELGT0eoxFOUN/HLTEAueytEX0h0MVP4+N
Jd3KPsNrQgzCI75w0LY7rrExSuFeXGekTleiXYHWurwpoB+ts4gRcO8xUPLNaVuU
/kzSIikq71jSaM+FEbCPg4101tioeE0/vheMzoM6ihfJ1f/D9p2XkfXZfwKBgHW3
omBGvHUFHZIGzz+VwgfhkbDCpAtefCj4snFuSVrEVg60cOaxCLXQWq7oTImB6RvK
HWxOStp0RVQFXToGOy1x0J1huvSFLoL2u/yGMbU/92Y6w4G1ifjqYhWXoL339XNr
wd3o7I9yX22ZcwG779Fuu+3Z05ym99iKprXBctERAoGAcTwsb7M0U3YgXy+VU1kz
KlGWoDFdv2LPkKnSmlflK3TbfHe8dnoTPJY3LVcsBYCWC6wZgwEJOr0zZXVLXYVK
jOZXG1PWcRbrB7kjzvtvXlQwXVjBvYqi1YzQ1Rrgq9/x77JHJgke99eZc1XzpN7B
6HiAXC9lDsK4N+nAVvFH7wY=
-----END PRIVATE KEY-----
`

func init() {
	if err := gorsa.RSA.SetPublicKey(PubKey); err != nil {
		panic(err)
	}
	if err := gorsa.RSA.SetPrivateKey(PrivateKey); err != nil {
		panic(err)
	}
}

func EncodeUser(u User) (string, error) {
	data, err := json.Marshal(u)
	if err != nil {
		return "", err
	}
	pubKeyEncrypt, err := gorsa.RSA.PubKeyENCTYPT(data)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(pubKeyEncrypt), nil
}

func DecodeUser(s string) (User, error) {
	data, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return nil, err
	}
	priKeyDecrypt, err := gorsa.RSA.PriKeyDECRYPT(data)
	if err != nil {
		return nil, err
	}
	u := &user{}
	err = json.Unmarshal(priKeyDecrypt, u)
	return u, err
}
