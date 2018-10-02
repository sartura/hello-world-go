//
// Copyright (C) 2018 Sartura Ltd.
//
// Permission is granted to copy, distribute and/or modify this document
// under the terms of the GNU Free Documentation License, Version 1.2
// or any later version published by the Free Software Foundation;
// with no Invariant Sections, no Front-Cover Texts, and no Back-Cover
// Texts.  A copy of the license is included in the section entitled "GNU
// Free Documentation License".

package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(w, "Hello world!")
	})
	log.Fatal(http.ListenAndServe(":8007", nil))
}
