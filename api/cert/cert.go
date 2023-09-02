package cert

import (
	"io/ioutil"
	"os"
)

var (
	Appid string
	// 应用私钥
	AppPrivateKey string
	// 支付宝公钥证书
	AlipayPublicContentRSA2 []byte
	// 支付宝根证书
	AlipayRootContent []byte
	// 应用公钥证书
	AppPublicContent []byte
)

func init() {
	RSA2, _ := os.Open("./alipayCertPublicKey_RSA2.crt")
	rootKey, _ := os.Open("./alipayRootCert.crt")
	appPublicKey, _ := os.Open("./appCertPublicKey_2021004108683556.crt")

	AlipayPublicContentRSA2, _ = ioutil.ReadAll(RSA2)
	AlipayRootContent, _ = ioutil.ReadAll(rootKey)
	AppPublicContent, _ = ioutil.ReadAll(appPublicKey)

	AppPrivateKey = "MIIEpAIBAAKCAQEAoX6Z==................."
	Appid = "2021000116674577"
}
