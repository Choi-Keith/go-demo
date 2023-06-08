package util

import (
	"fmt"
	"strconv"

	"github.com/bwmarrin/snowflake"
)

func GenUUID(prefix string) string {
	node, err := snowflake.NewNode(1)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	if prefix != "" {
		return prefix + strconv.Itoa(int(node.Generate().Int64()))
	}
	return strconv.Itoa(int(node.Generate().Int64())) + ""
}
