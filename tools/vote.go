/**
Tools package. It's contain some useful tools, just like vote and so on.
This file is created by magic at 2018-9-3
**/
package tools

import (
	"fmt"
	"os"
	"io/ioutil"
	"strings"
	"strconv"
	"github.com/urfave/cli"
	"github.com/outbrain/golib/log"
	"github.com/csunny/dpos"

)

// NodeVote 节点投票命令
var NodeVote = cli.Command{
	Name: "vote",
	Usage: "vote for node",
	Flags: []cli.Flag{
		cli.StringFlag{
			Name: "name",
			Value: "",
			Usage: "节点名称",
		},
		cli.IntFlag{
			Name: "v",
			Value: 0,
			Usage: "投票数量",
		},
	}, 
	Action: func(context *cli.Context) error{
		if err := Vote(context); err != nil{
			return err
		} 
		return nil
	},
}
// Vote for node. The votes of node is origin vote plus new vote.
// votes = originVote + vote 
func Vote(context *cli.Context) error {
	name := context.String("name")
	vote := context.Int("v")

	if name == "" {
		log.Errorf("节点名称不能为空")
	}

	if vote < 1 {
		log.Errorf("最小投票数目为1")
	}

	f, err := ioutil.ReadFile(dpos.FileName)
	if err != nil {
		log.Errorf(err.Error())
		return err
	}
	res := strings.Split(string(f), "\n")

	voteMap := make(map[string]string)
	for _, node := range res {
		nodeSplit := strings.Split(node, ":")
		if len(nodeSplit) > 1 {
			voteMap[nodeSplit[0]] = fmt.Sprintf("%s", nodeSplit[1])
		}
	}

	originVote, err := strconv.Atoi(voteMap[name])
	if err != nil {
		log.Errorf(err.Error())
		return err
	}
	votes := originVote + vote
	voteMap[name] = fmt.Sprintf("%d", votes)

	log.Infof("节点%s新增票数%d", name, vote)
	str := ""
	for k, v := range voteMap {
		str += k + ":" + v + "\n"
	}

	file, err := os.OpenFile(dpos.FileName, os.O_RDWR, 0666)
	if err != nil{
		return err
	}

	defer file.Close()
	file.WriteString(str)

	return nil
}