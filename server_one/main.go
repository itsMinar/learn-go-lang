package main

import "my_server/cmd"

func DemoEnc() {
	// var s string
	// s = "hello world"

	// byteArr := []byte(s)
	// fmt.Println("byteArr", byteArr)

	// enc := base64.URLEncoding
	// enc = enc.WithPadding(base64.NoPadding)
	// b64Str := enc.EncodeToString(byteArr)
	// fmt.Println("b64Str", b64Str)

	// decoded, err := enc.DecodeString(b64Str)
	// if err != nil {
	// 	fmt.Println("err", err)
	// 	return
	// }
	// fmt.Println("decoded", string(decoded))

	// data := []byte("Hello")
	// hash := sha256.Sum256(data)
	// fmt.Println("hash", hash)

	// secret := []byte("my-secret")
	// message := []byte("Hello")

	// h := hmac.New(sha256.New, secret)
	// h.Write(message)
	// text := h.Sum(nil)
	// fmt.Println("sum", text)

}

func main() {
	cmd.Serve()
}
