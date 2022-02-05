package main

import (
	"fmt"
	"strings"
)

//소문자는 대문자로, 대문자는 그대로 
func ToUpper1(str string) string {
	var rst string
	for _, c := range str {
		if c >= 'a' && c <= 'z' {
			rst += string('A' + (c - 'a'))
			fmt.Println(rst)
		} else {
			rst += string(c)
			fmt.Println(rst)
		}
	}
	return rst
}

func ToUpper2(str string) string {
	var builder strings.Builder
	for _, c := range str {
		if c >= 'a' && c <= 'z' {
			builder.WriteRune('A' + (c - 'a'))
		} else {
			builder.WriteRune(c)
		}
	}
	return builder.String()
}

func main() {
	var str string = "Hello World"
	fmt.Println(ToUpper1(str))
	fmt.Println(ToUpper2(str))
}


{
	"Sid": "VisualEditor0",
	"Effect": "Allow",
	"Action": [
		"rds:CreateDBInstance",
		"rds:CreateDBCluster",
		"rds:DescribeDBInstances"
	],
	"Resource": [
		"*"
	]
},
{
	"Sid": "VisualEditor1",
	"Effect": "Deny",
	"Action": [
		"rds:CreateDBInstance",
		"rds:CreateDBCluster",
		"rds:DescribeDBInstances"
	],
	"NotResource": [
		"arn:aws:rds:*:*:db:test*",
		"arn:aws:rds:*:*:og:*",
		"arn:aws:rds:*:*:pg:*",
		"arn:aws:rds:*:*:secgrp:*",
		"arn:aws:rds:*:*:subgrp:*",
		"arn:aws:rds:*:*:cluster:test*",
		"arn:aws:rds:*:*:cluster-pg:*"
	]
}