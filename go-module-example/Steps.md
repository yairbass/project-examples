## Setting up local environment

1. Clone this repository in a location outside your $HOME/go folder or $GOPATH folder if you have it set
2. If you want to resolve dependencies from a private repo, set up your GOPROXY environment variable properly
3. If you want to resolve all dependencies again you need to remove the contents at your go local cache by deleting your $HOME/go folder or $GOPATH folder if you have it set
4. Go to $CLONE_LOCATION/go-module-example

## Building and testing the project

1. Run `go test -v ./demo`

## Running the project

1. Run `go run . --text="Go is awesome!"`

## Creating the go executable binary

1. Run `go build .` 

## Publishing the project as a go module to Artifactory

1. Make sure you have the latest version of the [JFrog CLI](https://jfrog.com/getcli/)
2. Setup your JFrog CLI Artifactory connection: 

```
jfrog rt config --interactive=false --url=<ARTIFACTORY_URL> --user=<USER> --password=<PASSWORD> <ARTIFACTORY_ID>
```

Example:

```
jfrog rt config --interactive=false --url=http://localhost:8081/artifactory --user=user --password=password my-artifactory
```

3. Publish the module to Artifactory:

```
jfrog rt go-publish --server-id=<ARTIFACTORY_ID> --self=true <GO_REPO> <VERSION> 
```

Example:

```
jfrog rt go-publish --server-id=my-artifactory --self=true go-virtual v1.0.0
```

4. If you want to publish all dependencies to Artifactory as well add the option `--deps=ALL` to the command above
