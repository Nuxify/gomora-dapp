# Gomora dApp
A Gomora template for building dApps and web3-powered API and smart contract listeners

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

## Docker Build

To build, run:
- make build-docker

To run the container:
- make up

## License

[MIT](https://choosealicense.com/licenses/mit/)

Made with ❤️ at [Nuxify](https://nuxify.tech)