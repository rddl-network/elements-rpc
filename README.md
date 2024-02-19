# elements-rpc

## Example usage

### Reissue Asset

```
package main

import (
        "fmt"
        "log"

        elements "github.com/rddl-network/elements-rpc"
)

func main() {
        url := "http://user:password@127.0.0.1:18891/wallet/foowallet"
        assetID := "7add40beb27df701e02ee85089c5bc0021bc813823fedb5f1dcb5debda7f3da9"

        result, err := elements.ReissueAsset(url, []string{assetID, "100.00000000"})
        if err != nil {
                log.Fatal(err)
        }
        fmt.Printf("%+v\n", result)
}
```
Output:
```
{TxID:f1a0465275947a84ffca16bb51cb6efa94616d25d507bb93bd18d106acfa8365 Vin:0}
```

### Send To Address


```
package main

import (
        "fmt"
        "log"

        elements "github.com/rddl-network/elements-rpc"
)

func main() {
        url := "http://user:password@127.0.0.1:18891/wallet/foowallet"
        assetID := "7add40beb27df701e02ee85089c5bc0021bc813823fedb5f1dcb5debda7f3da9"

        hex, err := elements.SendToAddress(url, []string{
                "tlq1qqg4eudh5kes0pw903stdzvzdwrgrtg6nvqynvz9p9pler7xeugaau9s5h0plnqwchl5cgj2el03y5e0m65usqyyawjg9dar6t",
                `"100"`,
                `""`,
                `""`,
                "false",
                "true",
                "null",
                `"unset"`,
                "false",
                `"` + assetID + `"`,
        })
        if err != nil {
                log.Fatal(err)
        }
        fmt.Printf("%+v\n", hex)
}
```
Output:
```
62eefd1ab286bd4c344f8806fc507d5effd26176e91add5afb7f4f6531cf543e
```

### Wallet management
```
package main

import (
        "fmt"
        "log"

        elements "github.com/rddl-network/elements-rpc"
)

func main() {
        url := "http://user:password@127.0.0.1:18891/wallet/foowallet"
        loadWalletResult, err := elements.LoadWallet(url, []string{"foowallet", "true"})
        if err != nil {
                log.Fatal(err)
        }
        fmt.Printf("%+v\n", loadWalletResult)

        // wallet was created with passphrase "foobar"
        err = elements.Walletpassphrase(url, []string{"foobar", "60"})
        if err != nil {
                log.Fatal(err)
        }

        unloadWalletResult, err := elements.UnloadWallet(url, []string{"foowallet", "true"})
        if err != nil {
                log.Fatal(err)
        }
        fmt.Printf("%+v\n", unloadWalletResult)
}
```
Output:
```
{Name:foowallet Warning:}
{Warning:}
```

### Complex Example

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

        address, err := elements.GetNewAddress(url, []string{
                ``,
        })
        if err != nil {
                log.Fatal(err)
        }

        addressInfo, err := elements.GetAddressInfo(url, []string{
                address,
        })
        if err != nil {
                log.Fatal(err)
        }

        hex, err := elements.CreateRawTransaction(url, []string{
                `[]`,
                `[{"data":"00"}]`,
        })
        if err != nil {
                log.Fatal(err)
        }

        fundRawTransactionResult, err := elements.FundRawTransaction(url, []string{
                hex,
                `{"feeRate":0.00001000}`,
        })
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

        rawIssueAssetResults, err := elements.RawIssueAsset(url, []string{
                fundRawTransactionResult.Hex,
                `[{"asset_amount":0.00000001, "asset_address":"` + address + `", "blind":false, "contract_hash":"` + fmt.Sprintf("%+x", hash) + `"}]`,
        })
        if err != nil {
                log.Fatal(err)
        }

        rawIssueAssetResult := rawIssueAssetResults[len(rawIssueAssetResults)-1]
        hex, err = elements.BlindRawTransaction(url, []string{
                rawIssueAssetResult.Hex,
                `true, [], false`,
        })
        if err != nil {
                log.Fatal(err)
        }

        signRawTransactionWithWalletResult, err := elements.SignRawTransactionWithWallet(url, []string{
                hex,
        })
        if err != nil {
                log.Fatal(err)
        }

        testMempoolAcceptResults, err := elements.TestMempoolAccept(url, []string{
                `["` + signRawTransactionWithWalletResult.Hex + `"]`,
        })
        if err != nil {
                log.Fatal(err)
        }

        testMempoolAcceptResult := testMempoolAcceptResults[len(testMempoolAcceptResults)-1]
        if !testMempoolAcceptResult.Allowed {
                log.Fatalln("not accepted by mempool")
        }

        hex, err = elements.SendRawTransaction(url, []string{
                signRawTransactionWithWalletResult.Hex,
        })
        if err != nil {
                log.Fatal(err)
        }
        fmt.Println(hex)
}
```
