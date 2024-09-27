package main

import "fmt"

func httpPort(p int) {
	if p == 80 {
		fmt.Println("You can't use port 80! It's not safe!")
	} else {
		if p == 443 {
			fmt.Println("Perfect Your webserver is safe")
		} else {
			fmt.Println("You are running web application on unofficial port!")
		}
	}
}
