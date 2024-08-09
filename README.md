# goipinfo (Go IP info)

`goipinfo` is a command-line tool written in Go to fetch IP information from multiple providers and display it in a table format.

## Features

- Fetch your public IP information using different providers (ipinfo.io, ip-api.com, my-ip.io).
- Fetch information for a specific IP address using ip-api.com.
- Display fetched IP information in a nicely formatted table.

## Installation

### Prerequisites

- Go (version 1.22.2 or later)

### Install Dependencies

After head to the cloned repository issue the following command:

```sh
go mod tidy
```

### Build the Project

```sh
make build
```

**Note:** It will be built under the path `./build/bin/goipinfo`.

## Usage

### Get Your Public IP Information

```sh
./goipinfo myip --provider [ipinfo|ip-api|my-ip]
```

- The default provider is `ip-api`.
- Example:

  ```sh
  ./goipinfo myip --provider ipinfo
  ```

### Get Information for a Specific IP Address

```sh
./goipinfo info --ip 5.161.78.107
```

- Note: Only `ip-api` can be used as the provider for this command.
- Example:

  ```sh
  ./goipinfo info --ip 5.161.78.107
  ```

### Example Outputs

#### `myip` Command

```sh
./goipinfo myip --provider ipinfo
```

```sh
+----------+------------------------+
|  FIELD   |         VALUE          |
+----------+------------------------+
| IP       | 104.28.242.73          |
| City     | Düsseldorf             |
| Region   | North Rhine-Westphalia |
| Country  | Germany                |
| Location | 0.000000,0.000000      |
| Org      | Cloudflare WARP        |
| Timezone | Europe/Berlin          |
+----------+------------------------+
```

#### `info` Command

```sh
./goipinfo info --ip 104.28.242.73
```

```sh
+----------+------------------------+
|  FIELD   |         VALUE          |
+----------+------------------------+
| IP       | 104.28.242.73          |
| City     | Düsseldorf             |
| Region   | North Rhine-Westphalia |
| Country  | Germany                |
| Location | 0.000000,0.000000      |
| Org      | Cloudflare WARP        |
| Timezone | Europe/Berlin          |
+----------+------------------------+
```

## Development

### Running Tests

```sh
make test
```

### Cleaning the Build

```sh
make clean
```

### Formatting the Code

```sh
make fmt
```

### Linting the Code

```sh
make lint
```

## Contributing

Contributions are welcome! Please submit a pull request or open an issue to discuss any changes.

Get in touch with me @ <mohrezfadaei@gmail.com> 🙏

## License

This project is licensed under the MIT License. See the [LICENSE](./LICENSE) file for details.
