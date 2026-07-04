package util;

import (
	"encoding/base64"
	"encoding/json"
	"crypto/hmac"
	"crypto/sha256"

)

type Header struct{
	ALG string `json:"alg"`;
	TYP string `json:"typ"`;
}


type Payload struct{
	Sub string `json:"sub"`;  //user id
	Name string `json:"name"`;
    Email string `json:"email"`;
	
}

func convertBase64(data []byte) string {
	enc :=base64.URLEncoding;
	enc =enc.WithPadding(base64.NoPadding);
    b64 := enc.EncodeToString(data);
	return b64;
}

func convertHMAC_SHA256(data []byte,secret []byte)[]byte{
	
h := hmac.New(sha256.New, secret)
h.Write(data);
return h.Sum(nil)

}

func CreateJwtToken(secret string,payload Payload) string {
	header := Header{
		ALG: "HS256",
		TYP: "JWT",
	}

	headerBytes, err:= json.Marshal(header)
	if err != nil {
		panic(err)
	}
	payloadBytes, err2:= json.Marshal(payload)
	if err2 != nil {
		panic(err2)
	}	
	headerEncoded:= convertBase64(headerBytes)
	payloadEncoded:= convertBase64(payloadBytes)
    signature:= convertHMAC_SHA256([]byte(headerEncoded + "." + payloadEncoded), []byte(secret));

	return headerEncoded + "." + payloadEncoded + "." + convertBase64(signature);

}