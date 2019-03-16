# Implement a simple blockchain use dpos algorithm
 <a href="https://travis-ci.org/csunny/dpos"><img src="https://travis-ci.org/csunny/dpos.svg?branch=master" /></a>

## Architecture Design
- Create a P2P Conn-pool
- BlockChain Generate
- Node Manage And Vote
- Pick Node
- Write Block On Blockchain

## RUN 
```
git clone git@github.com:csunny/dpos.git

cd dpos    // 切换到源码路径下
go build -o build/dpos  main/dpos.go
```

connect multi peer 
```
./build/dpos new --port 3000 --secio
```
## Vote
```
./build/dpos vote -name QmaxEdbKW4x9mP2vX15zL9fyEsp9b9yV48zwtdrpYddfxe -v 30
```

# Document
[Doc](https://xiaozhuanlan.com/topic/3245810967)

# Licence 
MIT
