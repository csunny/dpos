# Implement a simple blockchain use dpos algorithm

## Architecture Design
- Create a P2P Conn-pool
- BlockChain Generate
- Note Manage And Vote
- Pick Node
- Write Block On Blockchain


## RUN 
```
git clone git@github.com:csunny/dpos.git
make dep
go run main/dpos.go -l 3000 -secio
```

## Vote
```
go run main/vote.go -name aaa -v 30
```

# Document
[Doc](https://xiaozhuanlan.com/topic/3245810967)

# Licence 
MIT
