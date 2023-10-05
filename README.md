# Gomora dApp
A Gomora template for building EVM-compatible API and smart contract indexers and listeners

## Local Development

Setup the .env file first
- cp .env.example .env

To bootstrap everything, run:
- make

The command above will install, build, and run the binary

For manual install:
- make install

For lint:
- make lint

Just ensure you installed golangci-lint.

To test:
- make test

For manual build:
- make build
- NOTE: the output for this is in bin/

For contract build:
- make contract-build
- NOTE: the output for this is in infrastructures/smartcontracts

> You can easily listen to the example Greeter contract (https://mumbai.polygonscan.com/address/0x927ec7f1f1CA6b09d0c448868aAB2C56d465a6e8#code) deployed in Polygon Testnet (Mumbai)

## Docker Build

To build, run:
- make run

To run the container:
- make up

## License

[MIT](https://choosealicense.com/licenses/mit/)

Made with ❤️ at [Nuxify](https://nuxify.tech)
