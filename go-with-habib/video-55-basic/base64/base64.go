package main

import (
	
	"encoding/base64"
	"fmt"
	
)

func main() {

	a := "a";
	b := "b";
	c := "c";
	

	byteArr := []byte(a)  // this is the byte type . without array it's not possiable. its uint8 , mane negative hobe na . 0 to 8 bit number supported. 8 bit = 2^8-1=255. mane 0 theke 255 porjonto number hote parbe

	byteArr2 :=[]byte{a[0], b[0], c[0]}

	fmt.Println(byteArr);
	fmt.Println(byteArr2);

	enc :=base64.URLEncoding;
	enc =enc.WithPadding(base64.NoPadding);
    b64 := enc.EncodeToString(byteArr); // return type base64

	fmt.Println(b64);

	//base64 to byteArray Or DecodeString

    b64TobyteArr, err :=enc.DecodeString(b64);

	if (err != nil) {
		fmt.Println("Error decoding base64:", err)
		return;
	}

	fmt.Println(b64TobyteArr)




}