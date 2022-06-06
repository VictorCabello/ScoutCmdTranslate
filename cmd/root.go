/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"strconv"
	"strings"
)

func Execute(msg string, reverse bool) {
	var key string = "murcielago"
	if reverse {
		var msgTranslated string = msg
		for i := range key {
			msgTranslated = strings.ReplaceAll(msgTranslated, strconv.Itoa(i), string(key[i]))
		}
		fmt.Println(msgTranslated)
	} else {
		var msgTranslated string = msg
		for i := range key {
			msgTranslated = strings.ReplaceAll(msgTranslated, string(key[i]), strconv.Itoa(i))
		}
		fmt.Println(msgTranslated)
	}
}
