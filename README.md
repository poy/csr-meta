# Go Vanity URLs

Go Vanity URLs is a simple App Engine Go app that allows you to set custom
import paths for your Go packages.

## Quickstart

Install [gcloud][gcloud] and install Go App Engine component:

```
$ gcloud components install app-engine-go
```

Setup a [custom domain][custom-domain] for your app.

Get the application:
```
go get -u -d github.com/poy/govanityurls
cd $(go env GOPATH)/src/github.com/poy/govanityurls
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
  </tbody>
</table>


[gcloud]:        https://cloud.google.com/sdk/downloads
[custom-domain]: https://cloud.google.com/appengine/docs/standard/python/using-custom-domains-and-ssl
