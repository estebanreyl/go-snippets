package main

import (
	"encoding"
	"fmt"

	_ "crypto/sha256"

	"github.com/opencontainers/go-digest"
)

func main() {
	digester := digest.Canonical.Digester()
	h, ok := digester.Hash().(encoding.BinaryMarshaler)
	if !ok {
		panic("Hash is not BinaryMarshaler")
	}

	data, err := h.MarshalBinary()
	if err != nil {
		panic(err)
	}

	// Typically there would be a write here but its not necessary for this example e.g. n, err := bw.digester.Hash().Write(p)

	// This is a dummy example to show how to use the BinaryUnmarshaler interface
	// is expected to work. This will panic in the fips image
	h2, ok := digest.Canonical.Digester().Hash().(encoding.BinaryUnmarshaler)
	if !ok {
		panic("Hash is not BinaryUnmarshaler")
	}
	err = h2.UnmarshalBinary(data)
	if err != nil {
		panic(err)
	}

	fmt.Print("Successfully marshaled and unmarshaled with sha256")
}
