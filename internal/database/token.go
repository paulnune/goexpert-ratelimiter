package database

import (
	"strconv"
	"strings"
)

type TokenLimit struct {
	Name  string `json:"name"`
	Limit int    `json:"limit"`
}

type TokenLimitList struct {
	List map[string]TokenLimit
}

func NewTokenLimitList(limitsParam string) (limitList TokenLimitList) {
	limitList.List = make(map[string]TokenLimit)
	arr := strings.Split(limitsParam, ",")
	for _, v := range arr {
		limite, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
		token := "Token" + v
		limitList.List[token] = TokenLimit{Name: token, Limit: limite}
	}
	return limitList
}

func (tll *TokenLimitList) GetLimit(token string) int {
	limite := tll.List[token].Limit
	return limite
}
