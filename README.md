# dolarpy-go

Go Package to easily consume the [Dolarpy API](https://dolar.melizeche.com).

<https://github.com/melizeche/dolarPy> - by [melizeche](https://github.com/melizeche/)

## Install

```go
go get github.com/bitebait/dolarpy-go
```

## Usage Example

```go
package main

import (
    "encoding/json"
    "fmt"

    "github.com/bitebait/dolarpy-go"
)

func main() {
    allData, err := dolarpy.All()
    if err != nil {
        fmt.Printf("Failed to get all providers data: %v\n", err)
        return
    }
    fmt.Println(allData)

    providers, err := dolarpy.Providers()
    if err != nil {
        fmt.Printf("Failed to get providers names: %v\n", err)
        return
    }
    fmt.Println(providers)

    reference, err := dolarpy.Reference()
    if err != nil {
        fmt.Printf("Failed to get reference value: %v\n", err)
        return
    }
    fmt.Println(reference)

    purchase, err := dolarpy.Purchase("cambioschaco")
    if err != nil {
        fmt.Printf("Failed to get purchase value from 'cambioschaco': %v\n", err)
        return
    }
    fmt.Println(purchase)

    sale, err := dolarpy.Sale("cambioschaco")
    if err != nil {
        fmt.Printf("Failed to get sale value from 'cambioschaco': %v\n", err)
        return
    }
    fmt.Println(sale)

    data, err := dolarpy.All()
    if err != nil {
        fmt.Printf("Failed to get all providers data: %v\n", err)
        return
    }

    j, _ := json.Marshal(data)
    fmt.Println(string(j))

    for provider, values := range data {
        purchase := values["compra"]
        sale := values["venta"]
        fmt.Printf("Provider: %s - [Purchase: %.2f Sale: %.2f]\n", provider, purchase, sale)
    }
}

```

## Contributing

1. Fork it!
2. Create your feature branch: `git checkout -b my-new-feature`
3. Commit your changes: `git commit -am 'Add some feature'`
4. Push to the branch: `git push origin my-new-feature`
5. Submit a pull request :D

## License

This project is licensed under the terms of the MIT License - see the [LICENSE](LICENSE) file for details
