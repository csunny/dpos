// implement a simple p2p blockchain which use dpos algorithm.
// this file just for a simple block generate and validate.
// This file is created by magic at 2018-9-2

package dpos

import (
	"crypto/sha256"
	"encoding/hex"
	"time"
)

// BlockChain slice to storage Block
var BlockChain []Block

// Block struct, A block contain 以下信息:
// Index 索引、Timestamp(时间戳)、BPM、Hash(自己的hash值)、PreHash(上一个块的Hash值)、validator(此区块的生产者信息)
type Block struct {
	Index     int
	Timestamp string
	BPM       int
	Hash      string
	PrevHash  string
	validator string
}

// CaculateHash 计算string的hash值
func CaculateHash(s string) string {
	h := sha256.New()
	h.Write([]byte(s))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

// CaculateBlockHash 计算Block的hash值
func CaculateBlockHash(block Block) string {
	record := string(block.Index) + block.Timestamp + string(block.BPM) + block.PrevHash
	return CaculateHash(record)
}

// GenerateBlock 根据上一个区块信息，生成新的区块
func GenerateBlock(oldBlock Block, BPM int, address string) (Block, error) {
	var newBlock Block
	t := time.Now()

	newBlock.Index = oldBlock.Index + 1
	newBlock.BPM = BPM
	newBlock.Timestamp = t.String()
	newBlock.PrevHash = oldBlock.Hash
	newBlock.Hash = CaculateBlockHash(newBlock)
	newBlock.validator = address

	return newBlock, nil
}

// IsBlockValid 校验区块是否合法
func IsBlockValid(newBlock, oldBlock Block) bool {
	if oldBlock.Index+1 != newBlock.Index {
		return false
	}

	if oldBlock.Hash != newBlock.PrevHash {
		return false
	}

	if CaculateBlockHash(newBlock) != newBlock.Hash {
		return false
	}
	return true
}
