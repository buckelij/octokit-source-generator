# source-generator

This repository is a prototype of code generation from GitHub's OpenAPI specification.

## Usage

### For Kiota

1. Install Kiota: First install .NET 7, then run `dotnet tool install --global --prerelease Microsoft.OpenApi.Kiota`
1. Run generation: `kiota generate -l go --ll trace -o generated/go -n kiota -d schemas/updated/api.github.com.json --co > output.txt 2>&1`
	1. Alternately, you may debug using the VSCode launch.json in the microsoft/kiota repo.
1. Run `go mod init github.com/octokit/kiota`
1. There's a command to print each dependency that's required; run it and then install dependencies: `kiota info -d /home/kfcampbell/github/dev/source-generator/schemas/updated/api.github.com.json -l Go`
1. Run `go build ./...` from the directory in which you've output the generated code to see errors.

To run pre-fixed/compiling Kiota demo app:
1. Export your PAT as GITHUB_TOKEN
1. `cd` to `compiling-kiota-go/go`
1. Run `go run cmd/main.go`
