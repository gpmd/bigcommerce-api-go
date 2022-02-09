package bigcommerce

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
)

// GetClientRequest returns a ClientRequest object from the BigCommerce API
// Call it with r.URL.Query() - will return BigCommerce Client Request or error
func (bc *BigCommerce) GetClientReqest(requestURLQuery url.Values) (*ClientRequest, error) {
	s := requestURLQuery.Get("signed_payload")
	decoded, err := bc.CheckSignature(s)
	if err != nil {
		return nil, err
	}
	var clrq ClientRequest
	err = json.Unmarshal(decoded, &clrq)
	if err != nil {
		return nil, err
	}
	return &clrq, nil
}

// CheckSignature checks the signature of the request whith SHA256 HMAC
func (bc *BigCommerce) CheckSignature(signed_payload string) ([]byte, error) {
	ss := strings.Split(signed_payload, ".")
	if signed_payload == "" {
		return nil, fmt.Errorf("no signed payload")
	}
	decoded, err := base64.StdEncoding.DecodeString(ss[0])
	if err != nil {
		return nil, fmt.Errorf("can't decode signed payload %v", err)
	}
	decodedSig, err := base64.StdEncoding.DecodeString(ss[1])
	if err != nil {
		return nil, fmt.Errorf("can't decode signature %v", err)
	}
	hms := hmac.New(sha256.New, []byte(bc.AppClientSecret))
	hms.Write(decoded)
	if !hmac.Equal([]byte(hex.EncodeToString(hms.Sum(nil))), decodedSig) {
		return nil, fmt.Errorf("signature mismatch")
	}
	return decoded, nil
}
