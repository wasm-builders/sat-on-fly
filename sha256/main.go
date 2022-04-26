package main

import (
    "crypto/sha256"
    "encoding/base64"

    "github.com/suborbital/reactr/api/tinygo/runnable"
)

type Sha256 struct{}

func (h Sha256) Run(input []byte) ([]byte, error) {
    hasher := sha256.New()
    hasher.Write(input)

    hashBytes := hasher.Sum(nil)

    hashString := base64.URLEncoding.EncodeToString(hashBytes)

    return []byte(hashString), nil
}

// initialize runnable, do not edit //
func main() {
    runnable.Use(Sha256{})
}