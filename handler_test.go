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

package main

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestHandler(t *testing.T) {
	tests := []struct {
		name string
		path string

		statusCode int
		goImport   string
	}{
		{
			name:     "infer from project/repo",
			path:     "/rakyll/portmidi",
			goImport: "example.com/rakyll/portmidi git https://source.developers.google.com/p/rakyll/r/portmidi",
		},
		{
			name:     "subpath",
			path:     "/rakyll/portmidi/foo",
			goImport: "example.com/rakyll/portmidi git https://source.developers.google.com/p/rakyll/r/portmidi",
		},
		{
			name:       "no repo",
			path:       "/rakyll",
			statusCode: http.StatusNotFound,
		},
		{
			name:       "no repo trailing slash",
			path:       "/rakyll/",
			statusCode: http.StatusNotFound,
		},
		{
			name:       "empty project",
			path:       "//portmidi",
			statusCode: http.StatusNotFound,
		},
		{
			name:       "no repo or project",
			path:       "/",
			statusCode: http.StatusNotFound,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			h, err := newHandler(0)
			if err != nil {
				t.Fatalf("newHandler: %v", err)
			}
			s := httptest.NewServer(h)
			resp, err := http.Get(s.URL + test.path)
			if err != nil {
				s.Close()
				t.Fatalf("http.Get: %v", err)
			}
			data, err := ioutil.ReadAll(resp.Body)
			resp.Body.Close()
			s.Close()
			if test.statusCode == 0 && resp.StatusCode != http.StatusOK {
				t.Fatalf("status code = %s; want 200 OK", resp.Status)
			}

			if test.statusCode != 0 && test.statusCode != http.StatusOK {
				if resp.StatusCode != test.statusCode {
					t.Fatalf("status code = %s; want %d", resp.Status, test.statusCode)
				}
				// We don't care about the rest of the test
				return
			}

			if err != nil {
				t.Fatalf("ioutil.ReadAll: %v", err)
			}

			test.goImport = strings.Replace(
				test.goImport,
				"example.com",
				strings.Replace(s.URL, "http://", "", 1),
				1,
			)
			if got := findMeta(data, "go-import"); got != test.goImport {
				t.Fatalf("meta go-import = %q; want %q", got, test.goImport)
			}
		})
	}
}

func TestBadCacheAge(t *testing.T) {
	_, err := newHandler(-1 * time.Second)
	if err == nil {
		t.Errorf("expected config to produce an error, but did not")
	}
}

func findMeta(data []byte, name string) string {
	var sep []byte
	sep = append(sep, `<meta name="`...)
	sep = append(sep, name...)
	sep = append(sep, `" content="`...)
	i := bytes.Index(data, sep)
	if i == -1 {
		return ""
	}
	content := data[i+len(sep):]
	j := bytes.IndexByte(content, '"')
	if j == -1 {
		return ""
	}
	return string(content[:j])
}

func TestCacheHeader(t *testing.T) {
	tests := []struct {
		name         string
		config       string
		cacheControl string
		cacheAge     time.Duration
	}{
		{
			name:         "default",
			cacheControl: "public, max-age=86400",
			cacheAge:     86400 * time.Second,
		},
		{
			name:         "specify time",
			cacheControl: "public, max-age=60",
			cacheAge:     60 * time.Second,
		},
		{
			name:         "zero",
			cacheControl: "public, max-age=0",
			cacheAge:     0 * time.Second,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			h, err := newHandler(test.cacheAge)
			if err != nil {
				t.Fatalf("newHandler: %v", err)
			}
			s := httptest.NewServer(h)
			resp, err := http.Get(s.URL + "/rakyll/portmidi")
			if err != nil {
				t.Fatalf("http.Get: %v", err)
			}
			resp.Body.Close()
			got := resp.Header.Get("Cache-Control")
			if got != test.cacheControl {
				t.Fatalf("Cache-Control header = %q; want %q", got, test.cacheControl)
			}
		})
	}
}
