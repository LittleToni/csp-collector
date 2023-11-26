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

This project is structured according to the [Standard Go Project Layout](https://github.com/golang-standards/project-layout/tree/master). For more details on the project structure and its components, you can refer to the [Standard Go Project Layout documentation](https://github.com/golang-standards/project-layout/tree/master).

### Test

Test CSP Collector by running:

```sh
make test
```

### Example

Post a sample csp violation report:

```sh
curl  -X POST \
  'http://localhost:8080/report' \
  --header 'Content-Type: application/csp-report' \
  --data-raw '{
    "csp-report": {
      "blocked-uri": "http://example.com/css/style.css",
      "disposition": "report",
      "document-uri": "http://example.com/signup.html",
      "effective-directive": "style-src-elem",
      "original-policy": "default-src 'none'; style-src cdn.example.com; report-to /_/csp-reports",
      "referrer": "",
      "status-code": 200,
      "violated-directive": "style-src-elem"
    }
  }'
```

## License

This project is licensed under the MIT License.
