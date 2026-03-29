package main

import (
	"strconv"
)

func citation(id string, unixTime int64) []byte {
	content := []byte(
		`[citation meta]`                               +"\n"+
		`uuid: ` + id                                   +"\n"+
		`unix-time: ` + strconv.FormatInt(unixTime, 10) +"\n"+
		``                                              +"\n"+
		`[citation]`                                    +"\n"+
		`author:`                                       +"\n"+
		`date:`                                         +"\n"+
		`title book:`                                   +"\n"+
		`title chapter:`                                +"\n"+
		`quotation:`                                    +"\n"+
		``,
	)

	return content
}
