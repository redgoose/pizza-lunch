# Pizza lunch 🍕

## Overview

This tool takes an exported Excel document from [TDSB's SchoolCash Online](https://tdsb.schoolcashonline.com/) and generates a user friendly report for ordering and distribution of pizza to hungry students on pizza lunch day. It is used by school council members.

![Sample PDF](docs/resources/pizza_lunch.png)

[View a sample PDF](https://github.com/redgoose/pizza-lunch/raw/main/docs/resources/pizza_lunch.pdf)

If school council had our way we would integrate directly with SchoolCash Online, but we lack access since we are not school staff. Instead, an Excel document is shared with us in this [format](https://github.com/redgoose/pizza-lunch/raw/main/docs/resources/pizza_lunch.xlsx).

## Quick start

1. Install [Go](https://golang.org/doc/install)
2. Install `pizza-lunch`:

	```
	go install github.com/redgoose/pizza-lunch@latest
	```
3. Rename `pizza-lunch.sample.yml` to `pizza-lunch.yml` and modify the configuration to meet your needs. Place this file and the exported Excel document in the same directory as the `pizza-lunch` executable.
4. Execute `pizza-lunch` which will then generate a PDF report.

## Testing

Run all tests from the root folder by running:

```
go test -v ./...
```

## License

MIT © redgoose, see [LICENSE](https://github.com/redgoose/pizza-lunch/blob/main/LICENSE) for details.