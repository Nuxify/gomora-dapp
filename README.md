# Gomora dApp
A Gomora template for building EVM-compatible API and smart contract indexers and listeners

## Local Development

Setup the .env file first

```bash
cp .env.example .env
```

To bootstrap everything, run:
```bash
make
```
The command above will install, build, and run the binary

For manual install:
```bash
make install
```

For lint:
```bash
make install
```

Just ensure you installed golangci-lint.

To test:
```bash
make test
```

For manual build:
```bash
make build

# The output for this is in bin/
```

For contract build:

```bash
make contract-build

# the output for this is in infrastructures/smartcontracts
```

> You can easily listen to the example Greeter contract (https://mumbai.polygonscan.com/address/0x927ec7f1f1CA6b09d0c448868aAB2C56d465a6e8#code) deployed in Polygon Testnet (Mumbai)

## Docker Build

To build, run:
```bash
make run
```

To run the container:
```bash
make up
```

## Database Migration

Gomora uses go-migrate (https://github.com/golang-migrate/migrate) to handle migration. Download and change your migrate database command accordingly.

To create a schema, run:

```bash
make schema NAME=<init_schema>
```

To migrate up, run:

```bash
make migrate-up STEPS=<remove STEPS to apply all or specify step number>
```

To migrate down, run:

```bash
make migrate-down STEPS=<remove STEPS to apply all or specify step number>
```

## License

[MIT](https://choosealicense.com/licenses/mit/)

Made with ❤️ at [Nuxify](https://nuxify.tech)
