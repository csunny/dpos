# 实现一个简易的DPOS算法


## 架构设计
- 创建一个P2P连接池
- BlockChain生成
- 节点管理与投票
- 选取节点生成区块
- 区块上链


## 代码运行
```
git clone git@github.com:csunny/dpos.git
make dep
go run main/dpos.go -l 3000 -secio
```

## 投票
```
go run main/vote.go -name aaa -v 30
```

# 说明文档
[文档](https://xiaozhuanlan.com/topic/3245810967)

# Licence 
MIT