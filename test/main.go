package main

import (
	"context"
	"fmt"

	"github.com/17535250630/solana-go"
	"github.com/17535250630/solana-go/rpc"
)

func main() {
	v := uint64(0)
	r := rpc.New("https://necessary-spring-energy.solana-mainnet.quiknode.pro/27e21f80b9860dea852873d0e34f893eed291036/")
	tx, e := r.GetConfirmedTransactionWithOpts(context.Background(), solana.MustSignatureFromBase58("2NtY4utK3PwthKmQacYEqb32x2v81sfqU7nQGFb4ahCjo625Zn9PQJbufu4GUiDunPRGbb1u55AxJzMQg7wFXLqp"), &rpc.GetTransactionOpts{
		MaxSupportedTransactionVersion: &v,
	})
	if e != nil {
		panic(e)
	}
	fmt.Println(tx)
}
