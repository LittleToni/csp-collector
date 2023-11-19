# CSP Collector

CSP Collector is a Go-based tool designed to collect and process Content Security Policy (CSP) violation reports. It implements CSP's report-to functionality, enabling web developers to receive detailed reports on CSP violations occurring on their websites.

## Installation

To install CSP Collector, ensure you have Go installed on your system and follow these steps:

```sh
git clone https://github.com/LittleToni/csp-collector
cd csp-collector
go build
```

## Usage

After installation, you can start CSP Collector by running:

```sh
./csp-collector
```

To use CSP Collector in your web application, add the appropriate CSP report-to to your CSP header.

## License

This project is licensed under the MIT License.
