# elements-rpc

## Example usage

```
package main

import (
        "crypto/sha256"
        "encoding/json"
        "fmt"
        "log"

        elements "github.com/rddl-network/elements-rpc"
)

type Entity struct {
        Domain string `json:"domain"`
}

type Contract struct {
        Entity       Entity `json:"entity"`
        IssuerPubkey string `json:"issuer_pubkey"`
        MachineAddr  string `json:"machine_addr"`
        Name         string `json:"name"`
        Precision    uint64 `json:"precision"`
        Version      uint64 `json:"version"`
}

func main() {
        url := "http://user:password@127.0.0.1:18891/wallet/foowallet"

        address, err := elements.GetNewAddress(url, ``)
        if err != nil {
                log.Fatal(err)
        }

        addressInfo, err := elements.GetAddressInfo(url, `"`+address+`"`)
        if err != nil {
                log.Fatal(err)
        }

        hex, err := elements.CreateRawTransaction(url, `[], [{"data":"00"}]`)
        if err != nil {
                log.Fatal(err)
        }

        fundRawTransactionResult, err := elements.FundRawTransaction(url, `"`+hex+`", {"feeRate":0.00001000}`)
        if err != nil {
                log.Fatal(err)
        }

        contract := Contract{
                Entity: Entity{
                        Domain: "testnet-assets.rddl.io",
                },
                IssuerPubkey: addressInfo.Pubkey,
                MachineAddr:  "plmnt18sgaelxxxc2crmkxx9vnj4g7ht3g2cwgy6nx6d",
                Name:         "plmnt18sgaelxxxc2crmkxx9vnj4g7ht3g2cwgy6nx6d",
                Precision:    0,
                Version:      0,
        }
        b, err := json.Marshal(contract)
        if err != nil {
                log.Fatal(err)
        }

        h := sha256.New()
        _, err = h.Write(b)
        if err != nil {
                log.Fatal(err)
        }
        hash := h.Sum(nil)

        // Reverse hash
        for i, j := 0, len(hash)-1; i < j; i, j = i+1, j-1 {
                hash[i], hash[j] = hash[j], hash[i]
        }

        rawIssueAssetResults, err := elements.RawIssueAsset(url, `"`+fundRawTransactionResult.Hex+`", [{"asset_amount":0.00000001, "asset_address":"`+address+`", "blind":false, "contract_hash":"`+fmt.Sprintf("%+x", hash)+`"}]`)
        if err != nil {
                log.Fatal(err)
        }

        rawIssueAssetResult := rawIssueAssetResults[len(rawIssueAssetResults)-1]
        hex, err = elements.BlindRawTransaction(url, `"`+rawIssueAssetResult.Hex+`", true, [], false`)
        if err != nil {
                log.Fatal(err)
        }

        signRawTransactionWithWalletResult, err := elements.SignRawTransactionWithWallet(url, `"`+hex+`"`)
        if err != nil {
                log.Fatal(err)
        }

        testMempoolAcceptResults, err := elements.TestMempoolAccept(url, `["`+signRawTransactionWithWalletResult.Hex+`"]`)
        if err != nil {
                log.Fatal(err)
        }

        testMempoolAcceptResult := testMempoolAcceptResults[len(testMempoolAcceptResults)-1]
        if !testMempoolAcceptResult.Allowed {
                log.Fatalln("not accepted by mempool")
        }

        hex, err = elements.SendRawTransaction(url, `"`+signRawTransactionWithWalletResult.Hex+`"`)
        if err != nil {
                log.Fatal(err)
        }
        fmt.Println(hex)
}
```
