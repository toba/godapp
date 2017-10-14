# Overview
This is an adaptation of Lior Schefer's sample [tvdapp](https://github.com/liors/tvdapp) to leverage
potential performance and security advantages in native [Go contract bindings](https://github.com/ethereum/go-ethereum/wiki/Native-DApps:-Go-bindings-to-Ethereum-contracts).

# Setup

## Install solc
Needed only if generating Go directly from `sol` rather than from `ABI`.

#### OSX
```
brew tap ethereum/ethereum
brew install solidity
brew linkapps solidity
```

#### Windows
[Download the latest release](https://github.com/ethereum/solidity/releases) zip and copy `solc.exe` to the path (such as `$GOPATH/bin`). The other files in the archive aren't needed
for Go bindings.

## [Install Geth](https://ethereum.github.io/go-ethereum/downloads/)
The Windows installer puts it in the `path`. For other operating systems, move the binary to an existing path location such as `$GOPATH/bin`.

Alternatively, build from the source downloaded below.

## Install Solidity language support
In VSCode, the [JuanBlanco.solidity](https://marketplace.visualstudio.com/items?itemName=JuanBlanco.solidity) plugin provides syntax support and compilation.

## Install Go bindings generator

```
go get -u github.com/ethereum/go-ethereum
cd $GOPATH/src/github.com/ethereum/go-ethereum
go install ./cmd/abigen
```

# Deployment
Uses binary Heroku buildpack
https://github.com/ph3nx/heroku-binary-buildpack
https://github.com/ph3nx/heroku-binary-buildpack.git

add `PATH` env
```
/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin:/app/bin
```