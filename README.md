# Implement a simple blockchain use dpos algorithm
 <a href="https://travis-ci.org/csunny/dpos"><img src="https://travis-ci.org/csunny/dpos.svg?branch=master" /></a>
[![](https://img.shields.io/github/issues/csunny/dpos)](https://github.com/csunny/dpos/issues)
[![](https://img.shields.io/github/stars/csunny/dpos)](https://github.com/csunny/dpos/star)
[![](https://img.shields.io/github/forks/csunny/dpos)](https://github.com/csunny/dpos/fork)
[![GoDoc](https://img.shields.io/badge/GoDoc-Reference-green)](https://godoc.org/github.com/csunny/dpos)
## Architecture Design
- Create a P2P Conn-pool
- BlockChain Generate
- Node Manage And Vote
- Pick Node
- Write Block On Blockchain

## Build 
👏
go build -o build/dpos  main/dpos.go

## RUN 
```
git clone git@github.com:csunny/dpos.git

cd dpos    // 切换到源码路径下
go build main/dpos.go
```

connect multi peer 
```
./dpos new --port 3000 --secio
```
## Vote
```
./dpos vote -name QmaxEdbKW4x9mP2vX15zL9fyEsp9b9yV48zwtdrpYddfxe -v 30
```

# Document
[Doc](https://xiaozhuanlan.com/topic/3245810967)

# Licence 
[![](https://img.shields.io/github/license/csunny/dpos)](https://github.com/csunny/dpos/blob/master/LICENSE)

MIT

