# Go Memory Monitoring Example

## Overview

This Go program demonstrates a simple memory monitoring tool using the `github.com/shirou/gopsutil/mem` package. It continuously monitors the available system memory and allocates a 128 MiB array when the available memory exceeds a certain threshold.

## Prerequisites

Before running this program, ensure that you have Go installed on your system.

```bash
# Install the required package
go get -u github.com/shirou/gopsutil/mem
```

## Usage

1. Clone the repository:

```bash
git clone https://github.com/john-h3/mem-test-cmd.git
cd mem-test-cmd
```

2. Build the executable:

```bash
go build -o memory-monitor ./main
```

3. Run the executable:

```bash
./memory-monitor
```

The program will continuously monitor the available memory and allocate a 128 MiB array when the available memory is above a certain threshold.

## Configuration

You can customize the program by adjusting the following parameters in the `main` function:

- `if virtualMemory.Available>>20 < 256`: The threshold for available memory (in MiB) below which memory allocation is skipped.
- `time.Sleep(time.Second)`: The duration of the sleep between memory checks and allocations.

## License

This code is provided under the [AGPL-3.0](https://github.com/john-h3/mem-test-cmd/blob/main/LICENSE).

## Acknowledgments

- [gopsutil](https://github.com/shirou/gopsutil) - A Go cross-platform library for retrieving system information.

## Author

[john-h3]