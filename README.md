# nife
simple multi purpose command line utilities
## Features

- **File Management**: Quickly copy, move, delete, and rename files and directories.
- **Text Processing**: Perform search, replace, and text manipulation tasks.
- **Network Utilities**: Check connectivity, perform DNS lookups, and more.
- **System Monitoring**: Monitor system resources like CPU, memory, and disk usage.
- **Automation**: Create and run scripts to automate repetitive tasks.

## Installation

To install `nife`, you can use the following command:

```sh
pip install nife
```

## Usage

Here are some examples of how to use `nife`:

### File Management

```sh
nife copy source.txt destination.txt
nife move oldname.txt newname.txt
```

### Text Processing

```sh
nife search "pattern" file.txt
nife replace "old" "new" file.txt
```

### Network Utilities

```sh
nife ping example.com
nife dnslookup example.com
```

### System Monitoring

```sh
nife sysinfo
nife diskusage
```

### Automation

```sh
nife runscript myscript.nife
```

For more detailed usage and options, refer to the [documentation](docs/USAGE.md).

## Contributing

Contributions are welcome! Please see the [CONTRIBUTING.md](CONTRIBUTING.md) for guidelines.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
## Building from Source

To build `nife` from source, ensure you have Go installed and set up on your machine. Then, run the following commands:

```sh
git clone https://github.com/yourusername/nife.git
cd nife
go build
```

This will generate an executable file named `nife` in the current directory.

## Installation using Go

You can also install `nife` directly using the `go` command:

```sh
go install github.com/yourusername/nife@latest
```

Make sure your `GOPATH` is set and included in your `PATH` to run `nife` from anywhere.

## Running Tests

To run the tests for `nife`, use the following command:

```sh
go test ./...
```

This will execute all the tests and provide a summary of the results.
## Installation

To install `nife`, you can use the following command:

```sh
go install github.com/yourusername/nife@latest
```

Make sure your `GOPATH` is set and included in your `PATH` to run `nife` from anywhere.