package main

import zkp "github.com/olegabu/go-secp256k1-zkp"

func main() {

    ctx, err := zkp.ContextCreate(zkp.ContextBoth)
    if err != nil {
        panic(err)
    }
    zkp.ContextDestroy(ctx)
    print("Success\n")
}