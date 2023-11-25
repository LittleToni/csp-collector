# CSP Collector

CSP Collector is a Go-based tool designed to collect and process Content Security Policy (CSP) violation reports. It implements CSP's report-to functionality, enabling web developers to receive detailed reports on CSP violations occurring on their websites.

## Installation

To install CSP Collector, follow these steps:

```sh
git clone https://github.com/LittleToni/csp-collector
cd csp-collector
```

## Configuration

Run the following command to initialize CSP Collector:

```sh
make init
```

Edit the `.env` file to customize CSP Collector.

## Usage

Build CSP Collector by running:

```sh
make build
```

After build, you can start CSP Collector by running:

```sh
make start
```

To use CSP Collector in your web application, add the appropriate CSP report-to to your CSP header.

## Development

### Test

Test CSP Collector by running:

```sh
make test
```

## License

This project is licensed under the MIT License.
