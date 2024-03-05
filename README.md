# Finazon Go gRPC Client

This is the official Go library for Finazon, offering access to:
- Lists of datasets, publishers, markets, and tickers.
- Market data: ticker snapshots, time series, trades, and technical indicators.
- Data from specific datasets such as Benzinga, Binance, Crypto, Forex, SEC, and SIP.
- Full compatibility with both JavaScript and TypeScript.

🔑 **API key** is essential. If you haven't got one yet, [sign up here](https://finazon.io/).

## Requirements

Ensure you have Go 1.18 or later.

## Installation

```bash
go get github.com/mivxxx/finazon-grpc-go-client
```

## Updating to last version

```bash
go get -u github.com/mivxxx/finazon-grpc-go-client
```

## Quick start

### 1. Set up a new project
```bash
mkdir hello-finazon && cd hello-finazon
go mod init example/hello
go get github.com/mivxxx/finazon-grpc-go-client
go mod tidy
go mod vendor
```

### 2. Create `hello-world.go` script
```go
package main

import (
	"fmt"
	finazon_grpc_go_client "github.com/mivxxx/finazon-grpc-go-client/pb"
)

const API_KEY = "your_api_key"

func main() {
	con, err := finazon_grpc_go_client.GetConnection(API_KEY)
	if err != nil {
		fmt.Println("%s", err)
		return
	}

	timeSeriesClient := con.GetTimeSeriesClient()
	timeSeriesRequest := finazon_grpc_go_client.GetTimeSeriesRequest{
		Dataset: "sip_non_pro",
		Ticker:  "AAPL",
	}
	data, err := timeSeriesClient.GetTimeSeries(&timeSeriesRequest)
	if err != nil {
		fmt.Println("%s", err)
		return
	}
	fmt.Println("%s", data.String())
}
```

### 3. Input your API key
Replace `'your_api_key'` with your actual key.

### 4. Run the example
```bash
go run hello-world.go
```

📝 Expected output:
```
result:{timestamp:1709326680  open:179.655  close:179.695  high:179.73  low:179.64  volume:563082}
result:{timestamp:1709326620  open:179.67  close:179.66  high:179.672  low:179.6  volume:242826}
...
result:{timestamp:1709324940  open:179.71  close:179.68  high:179.715  low:179.64  volume:107197}
```

## RPC support

The following table outlines the supported rpc calls:
<!--rpc_table_boundary-->
<!--rpc_table_boundary-->
Here's how you can import `client` and `request` objects:
```go
import finazon_grpc_go_client "github.com/mivxxx/finazon-grpc-go-client/pb"

con, _ := finazon_grpc_go_client.GetConnection(API_KEY)

client := con.GetServiceNameClient()

request := finazon_grpc_go_client.GetServiceNameRequest{}
data, err := client.RpcName(&request)
```

## Documentation
Delve deeper with our [official documentation](https://finazon.io/docs).

## Support
- 🌐 Visit our [contact page](https://finazon.io/contact-sales).
- 🛠 Or our [support center](https://support.finazon.io/en/).

## Stay updated
- Follow us on [𝕏](https://twitter.com/finazon_io).
- Join our community on [Discord](https://discord.gg/D5u4ZpB7w7).
- Follow us on [LinkedIn](https://www.linkedin.com/company/finazon).

## Contributing
1. Fork and clone: `$ git clone https://github.com/mivxxx/finazon-grpc-go-client.git`.
2. Branch out: `$ cd finazon-grpc-node && git checkout -b my-feature`.
3. Commit changes and test.
4. Push your branch and open a pull request with a comprehensive description.

For more guidance on contributing, see the [GitHub Docs](https://docs.github.com/en/get-started/quickstart/contributing-to-projects) on GitHub.

## License

This project is licensed under the MIT License. See the `LICENSE.txt` file in this repository for more details.
