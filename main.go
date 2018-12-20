// Copyright 2017 Google Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//+build !appengine

package main

import (
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	log := log.New(os.Stderr, "", log.LstdFlags)
	h, err := newHandler(host(), cacheAge(log))
	if err != nil {
		log.Fatal(err)
	}
	http.Handle("/", h)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func host() string {
	h := os.Getenv("HOST")
	if h == "" {
		return "code.gopher.run"
	}

	return h
}

func cacheAge(log *log.Logger) time.Duration {
	age := os.Getenv("CACHE_AGE")
	if age == "" {
		return 24 * time.Hour
	}

	d, err := time.ParseDuration(age)
	if err != nil {
		log.Fatalf("failed to parse CACHE_AGE %q: %s", age, err)
	}
	return d
}
