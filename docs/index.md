# csr-meta

[Cloud Source Repositories][csr] (CSR) is great for hosting proprietary and
private source code. However if you have ever tried to use CSR to host Go
code, you have more than likely ran into the following issue:

```
$ go get source.developers.google.com/p/<PROJECT-ID>/r/<REPO-NAME>

go get source.developers.google.com/p/<PROJECT-ID>/r/<REPO-NAME>: unrecognized import path "source.developers.google.com/p/<PROJECT-ID>/r/<REPO-NAME>" (parse https://source.developers.google.com/p/<PROJECT-ID>/r/<REPO-NAME>?go-get=1: no go-import meta tags ())
```

The real clue as to why this doesn't work is:

```
no go-import meta tags ()
```

This is where **csr-meta** comes in. It will read the request and infer the
tags and redirect the `go` build tool to the correct place on CSR.

## How to use

Its simple! All you have to do is update your import paths:

```go
package bar_test

import (
  // Project ID => foo
  // Repo Name  => bar
  "go-csr.appspot.com/foo/bar"
)
```

Then when you download it via `go get`:

```
$ go get go-csr.appspot.com/foo/bar
$ cd $GOPATH/src/go-csr.appspot.com/foo/bar
```

**NOTE**: Repository names with a `/` are not supported.

[csr]: https://source.cloud.google.com/

## Security

**csr-meta** does **not** proxy the code. It simply redirects the build tool.
This means that the hosted solution does not have access to your code. Your
private source code is kept private.
