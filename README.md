# Website Learn Backend

A backend written in go lang

We use the `GIN` Package to create the APIs

To have a look at api specifications go here: [`/apiContracts`](/apiContracts)

To run the backend use the command

```sh
go run ./src/main.go
```

To run the tests use the command

```sh
go test ./src/tests
```

After cloning the repository run the following command to install all the dependencies:

```sh
go mod tidy
```

Create a document named local.toml in `/src/config` folder and add a variable named `FIRESTORE_CREDENTIALS` and paste the service account json downloaded.
Follow the formatting provided in `/src/config/default.toml`

To validate the setup run the following command:

```sh
go run ./src/main.go validate
```
