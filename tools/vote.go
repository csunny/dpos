/**
Tools package. It's contain some useful tools, just like vote and so on.
This file is created by magic at 2018-9-3
**/
package tools

import (
	"flag"
	"log"
	"fmt"
	"os"
	"io/ioutil"
	"strings"
	"strconv"
	"github.com/csunny/argo/src/dpos"

)

// Vote for node. The votes of node is origin vote plus new vote.
// votes = originVote + vote 
func Vote() {
	name := flag.String("name", "", "节点名称")
	vote := flag.Int("v", 0, "投票数量")
	flag.Parse()

	if *name == "" {
		log.Fatal("节点名称不能为空")
	}

	if *vote < 1 {
		log.Fatal("最小投票数目为1")
	}

	f, err := ioutil.ReadFile(dpos.FileName)
	if err != nil {
		log.Fatal(err)
	}
	res := strings.Split(string(f), "\n")

	voteMap := make(map[string]string)
	for _, node := range res {
		nodeSplit := strings.Split(node, ":")
		if len(nodeSplit) > 1 {
			voteMap[nodeSplit[0]] = fmt.Sprintf("%s", nodeSplit[1])
		}
	}

	originVote, err := strconv.Atoi(voteMap[*name])
	if err != nil {
		log.Fatal(err)
	}
	votes := originVote + *vote
	voteMap[*name] = fmt.Sprintf("%d", votes)

	log.Printf("节点%s新增票数%d", *name, vote)
	str := ""
	for k, v := range voteMap {
		str += k + ":" + v + "\n"
	}

	file, err := os.OpenFile(dpos.FileName, os.O_RDWR, 0666)
	file.WriteString(str)
}