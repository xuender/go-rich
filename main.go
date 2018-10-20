package main

import "./cmd"

//go:generate mkdir keys -p
//go:generate go run $GOROOT/src/crypto/tls/generate_cert.go --host=localhost
//go:generate mv cert.pem key.pem keys
//go:generate go-bindata -nomemcopy -pkg keys -o ./keys/bindata.go keys/cert.pem keys/key.pem
//go:generate go-bindata -nomemcopy -pkg rich -o ./rich/bindata.go www/...

func main() {
	cmd.Execute()
}
