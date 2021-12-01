# dolarpy-go

Go Package to easily consume the [Dolarpy API](https://dolar.melizeche.com).

<https://github.com/melizeche/dolarPy> - by [melizeche](https://github.com/melizeche/)

## Install

```go
import "github.com/bitebait/dolarpy-go"
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
	dolarpy.All()                    // returns all providers data
	dolarpy.Providers()              // returns providers names
	dolarpy.Reference()              // returns reference value from 'bcp'
	dolarpy.Purchase("cambioschaco") // returns purchase value from 'cambioschaco' (default: 'bcp)
	dolarpy.Sale("cambioschaco")     // returns sale value from 'cambioschaco' (default: 'bcp)

	// JSON Example
	j, _ := json.Marshal(dolarpy.All())
	fmt.Println(string(j))

	// For Loop Example
	data := dolarpy.All()
	for key, value := range data {
		fmt.Printf("Provider: %s - [Purchase: %.2f Sale: %.2f]\n", key, value["compra"], value["venta"])
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
