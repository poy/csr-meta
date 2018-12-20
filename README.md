# csr-meta

Cloud Source Repositories (CSR) does not serve the required meta tags for `go
get` to work correctly. **csr-meta** attaches the required meta tags. It
assumes that the repo name does not have any slashes and uses each request to
build the meta tags. Therefore

```
<hostname>/myproj/myrepo
```

corresponds to

```
source.developers.google.com/p/myproj/r/myrepo
```

## Quickstart

Install [gcloud][gcloud] and install Go App Engine component:

```
$ gcloud components install app-engine-go
```

Setup a [custom domain][custom-domain] for your app.

Get the application:
```
go get -u -d github.com/poy/csr-meta
cd $(go env GOPATH)/src/github.com/poy/csr-meta
```

```
paths:
  /portmidi:
    repo: https://github.com/rakyll/portmidi
```

You can add as many rules as you wish.

Deploy the app:

```
$ gcloud app deploy
```

That's it! You can use `go get` to get the package from your custom domain.

```
$ go get customdomain.com/portmidi
```

## Environment Variables

<table>
  <thead>
    <tr>
      <th scope="col">Name</th>
      <th scope="col">Required</th>
      <th scope="col">Description</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <th scope="row"><code>CACHE_AGE</code></th>
      <td>optional</td>
      <td>The amount of time to cache package pages as a time.Duration (e.g., <code>24h</code>). It will be rounded to the nearest second. Controls the <code>max-age</code> directive sent in the <a href="https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Cache-Control"><code>Cache-Control</code></a> HTTP header.</td>
    </tr>
    <tr>
      <th scope="row"><code>HOST</code></th>
      <td>optional</td>
      <td>The host that the code is being redirected from. It defaults to <code>code.gopher.run</code>.</td>
    </tr>
  </tbody>
</table>


[gcloud]:        https://cloud.google.com/sdk/downloads
[custom-domain]: https://cloud.google.com/appengine/docs/standard/python/using-custom-domains-and-ssl
