package main

import (
	"strings"
)

type User struct {
	FullName string
}

func NewUser(fullName string) *User {
	return &User{FullName: fullName}
}

var AllNames []string

type User2 struct {
	names []uint8
}

func NewUser2(fullName string) *User2 {
	getOrAdd := func(s string) uint8 {
		for i := 0; i < len(AllNames); i++ {
			if AllNames[i] == s {
				return uint8(i)
			}
		}
		AllNames = append(AllNames, s)
		return uint8(len(AllNames) - 1)
	}
	result := &User2{}
	parts := strings.Split(fullName, " ")
	for _, p := range parts {
		index := getOrAdd(p)
		result.names = append(result.names, index)
	}
	return result
}

func (u *User2) FullName() string {
	var parts []string
	for _, i := range u.names {
		parts = append(parts, AllNames[i])
	}
	return strings.Join(parts, " ")
}
