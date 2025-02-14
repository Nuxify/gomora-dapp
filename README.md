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

> You can easily listen to the example Greeter contract (https://sepolia.etherscan.io/address/0x0e10e90f67C67c2cB9DD5071674FDCfb7853a6F5#code) deployed in Ethereum Testnet (Sepolia)

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
NAME=<create_users_schema> make migrate-schema
```

To migrate up, run:

```bash
STEPS=<remove STEPS to apply all or specify step number> make migrate-up
```

To migrate down, run:

```bash
STEPS=<remove STEPS to apply all or specify step number> make migrate-down
```

To check migrate version, run:

```bash
make migrate-version
````

To force migrate, run:
```bash
STEPS=<specify step number> make migrate-force
```

## License

[MIT](https://choosealicense.com/licenses/mit/)

Made with ❤️ at [Nuxify](https://nuxify.tech)
