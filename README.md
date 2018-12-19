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

Edit `vanity.yaml` to add any number of git repos. E.g.,
`customdomain.com/portmidi` will serve the
[https://github.com/rakyll/portmidi][portmidi] repo.

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

## Configuration File

```
host: example.com
paths:
  /foo:
    repo: https://github.com/example/foo
    display: "https://github.com/example/foo https://github.com/example/foo/tree/master{/dir} https://github.com/example/foo/blob/master{/dir}/{file}#L{line}"
```

<table>
  <thead>
    <tr>
      <th scope="col">Key</th>
      <th scope="col">Required</th>
      <th scope="col">Description</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <th scope="row"><code>host</code></th>
      <td>optional</td>
      <td>Host name to use in meta tags.  If omitted, uses the App Engine default version host or the Host header on non-App Engine Standard environments.  You can use this option to fix the host when using this service behind a reverse proxy or a <a href="https://cloud.google.com/appengine/docs/standard/go/how-requests-are-routed#routing_with_a_dispatch_file">custom dispatch file</a>.</td>
    </tr>
    <tr>
      <th scope="row"><code>paths</code></th>
      <td>required</td>
      <td>Map of paths to path configurations. Each key is a path that will point to the root of a repository hosted elsewhere. The fields are documented in the Path Configuration section below.</td>
    </tr>
  </tbody>
</table>

### Path Configuration

<table>
  <thead>
    <tr>
      <th scope="col">Key</th>
      <th scope="col">Required</th>
      <th scope="col">Description</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <th scope="row"><code>display</code></th>
      <td>optional</td>
      <td>The last three fields of the <a href="https://github.com/golang/gddo/wiki/Source-Code-Links"><code>go-source</code> meta tag</a>.  If omitted, it is inferred from the code hosting service if possible.</td>
    </tr>
    <tr>
      <th scope="row"><code>repo</code></th>
      <td>required</td>
      <td>Root URL of the repository as it would appear in <a href="https://golang.org/cmd/go/#hdr-Remote_import_paths"><code>go-import</code> meta tag</a>.</td>
    </tr>
  </tbody>
</table>


[gcloud]:        https://cloud.google.com/sdk/downloads
[custom-domain]: https://cloud.google.com/appengine/docs/standard/python/using-custom-domains-and-ssl
[portmidi]:      https://github.com/rakyll/portmidi
