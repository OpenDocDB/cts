module github.com/OpenDocDB/cts/opendocdb-cts

go 1.25

toolchain go1.25.0

// replace github.com/FerretDB/wire => ../../wire

require (
	github.com/FerretDB/wire v0.1.7
	github.com/alecthomas/kong v1.12.1
	github.com/neilotoole/slogt v1.1.0
	github.com/pmezard/go-difflib v1.0.0
	github.com/sethvargo/go-githubactions v1.3.1
	github.com/stretchr/testify v1.11.1
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/xdg-go/pbkdf2 v1.0.0 // indirect
	github.com/xdg-go/scram v1.1.2 // indirect
	github.com/xdg-go/stringprep v1.0.4 // indirect
	go.mongodb.org/mongo-driver v1.17.4 // indirect
	go.mongodb.org/mongo-driver/v2 v2.2.2 // indirect
	golang.org/x/text v0.25.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
