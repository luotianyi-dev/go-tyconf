# TyConf
TyConf is Tianyi Network's custom configuration library for Go.

Featured in parsing configuration from **environment variables** and **command-line arguments**, TyConf is designed to be simple and easy to use.

Please note that only following types are supported: `string`, `bool`, `int`, `int64`, `uint`, `uint64`, `float64`.

See [go.dev](https://pkg.go.dev/github.com/luotianyi-dev/go-tyconf) for detailed documentation.

## Installation
Run the following command to install the package:
```bash
go get github.com/luotianyi-dev/go-tyconf
```

## Usage
Here is an example of how to use TyConf:

```go
import (
	"fmt"

	"github.com/luotianyi-dev/go-tyconf"
)

type Config struct {
	EnableTCP         bool   `env:"ENABLE_TCP"           cli:"tcp"       description:"Enable TCP server"`
	ServerName        string `env:"SERVER_NAME"          cli:"sname"     description:"Name of the server"`
	UpdateIntervalSec int    `env:"UPDATE_INTERVAL_SEC"  cli:"interval"  description:"Update interval in seconds"`
}

func main() {
	defaultConfig := Config{
		EnableTCP:         false,
		ServerName:        "default",
		UpdateIntervalSec: 30,
	}
	config := tyconf.Parse(defaultConfig).(Config)
	fmt.Printf("Config: %+v\n", config)
}
```


### Command Line Helps
Run `go run main.go` to test the code. TyConf use `flag` package as underlying implementation, so you can pass `-h` or `--help` to see the help message:
```bash
go run main.go --help
```

### Passing Configuration Values
You can set environment variables or pass command-line arguments to override the default values:
```bash
ENABLE_TCP=true SERVER_NAME=example go run main.go --interval 10
```

### Configuration Priority
**CLI is always prioritized over environment variables.** For example, if you run the following command:
```bash
ENABLE_TCP=false SERVER_NAME=example go run main.go --tcp
```
The `EnableTCP` field will be set to `true` because the `--tcp` flag is passed.

## License
This project is licensed under the [MIT License](LICENSE).
