
# Sample Go app with Autometrics
This is a sample backend service written in Go with several API handlers that are erroring or responding slowly, intended to showcase the Autometrics framework with Service Level Objectives.

### Resources

- see the [documentation site](https://docs.autometrics.dev) for more in-depth information on all Autometrics features.
- ask for help and share feedback on [our Discord](https://discord.com/invite/MJr7pYzZQ4)!

## Getting started

To get the sample up and running you can follow these steps:

1. Clone the repository

```bash
git clone autometrics-dev/gettingstarted-am-go

cd gettingstarted-am-go
```

2. Download the Autometrics CLI

If you're on macOS you can use Homebrew:

```bash
brew install autometrics-dev/tap/am
```

or you can grab the binaries directly from the [GitHub release](https://github.com/autometrics-dev/am/releases/).

3. Set up and call the generator

The generator is the binary in cmd/autometrics, so the easiest way to get it is to install it through go:

```bash
go install github.com/autometrics-dev/autometrics-go/cmd/autometrics@latest
```

<details>
<summary> Make sure your `$PATH` is set up</summary>

In order to have `autometrics` visible then, make sure that the directory `$GOBIN` (or the default `$GOPATH/bin`) is in your `$PATH`:

```bash
$ echo "$PATH" | grep -q "${GOBIN:-$GOPATH/bin}" && echo "GOBIN in PATH" || echo "GOBIN not in PATH, please add it"
GOBIN in PATH
```
</details>

You can now call `go generate`:

```bash
go generate ./main.go
```

The generator will augment the doc comment to add quick links to metrics (using the Prometheus URL as base URL), and add a unique defer statement that will take care of instrumenting your code.

4. Build and run the application

Run `go build` to compile the Go application

```bash
go build
```

...and run the generated binary.

```bash
./gettingstarted-am-go
```

The application will start on a port 8080 by default and expose a metrics endpoint.

4. Start the Autometrics CLI and Explorer

Start the Autometrics CLI and point it to the endpoint it can scrape metrics from.

```bash
am start :8080
```

Autometrics CLI will download and run a Prometheus binary under the hood and start scraping metrics.

5. Preview the metrics in Autometrics Explorer

Autometrics CLI will also start a server with the Explorer available on `localhost:6789`. You can browse it and start exploring your sample app metrics! (You might need to ping the endpoints a few times to see the data reflected).

That's all!


