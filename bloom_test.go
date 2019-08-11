package bloom

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"testing"
)

func TestNewBloom(t *testing.T) {
	b := NewBloom(nil)
	b.Add([]byte("pct"))
	if b.Test([]byte("pct")) != true {
		t.Fatal("error test pct")
	}
	if b.Test([]byte("pct2")) == true {
		t.Fatal("error test pct2")
	}

	data := b.Bytes()
	b2 := NewBloom(data)
	if b2.Test([]byte("pct")) != true {
		t.Fatal("error test pct")
	}
	if b2.Test([]byte("pct2")) == true {
		t.Fatal("error test pct2")
	}
}

func TestBloomCycle(t *testing.T) {
	b := NewBloom(nil)
	for i := 0; i < 100000; i++ {
		key := new(big.Int).SetInt64(int64(i))
		b.Add(key.Bytes())
	}
	for i := 0; i < 100000; i++ {
		key := new(big.Int).SetInt64(int64(i))
		if !b.Test(key.Bytes()) {
			t.Fatal("test error", i)
		}
	}
}

func TestEthAddress(t *testing.T) {
	b := NewBloom(nil)
	for i := int64(0); i < 100; i++ {
		key := common.HexToAddress(fmt.Sprintf("%x", new(big.Int).SetInt64(i)))
		b.Add(key.Bytes())
	}
	for i := int64(0); i < 100; i++ {
		key := common.HexToAddress(fmt.Sprintf("%x", new(big.Int).SetInt64(i)))
		if !b.Test(key.Bytes()) {
			t.Fatal("test error", i)
		}
	}
	for i := int64(101); i < 200; i++ {
		key := common.HexToAddress(fmt.Sprintf("%x", new(big.Int).SetInt64(i)))
		if b.Test(key.Bytes()) {
			t.Fatal("test error", i)
		}
	}
}
