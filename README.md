# UUID Wrong Import Path Demo

## Clean run

1. Clear your module cache:
   `go clean -modcache`
2. Generate ent code:
   `go generate -x ./...`
3. Compile (it should be OK):
   `go build -v ./...`

## Failure run

1. Clear your module cache again
2. Seed `github.com/gofrs/uuid` into your module cache:
   `( cd \$(mktemp -d) && go mod init test && go get github.com/gofrs/uuid )`
   - Note that this is not making _any_ changes to the test code, it is _only_ changing the shared go cache!
3. Generate ent code again
4. Try to compile again, this time will fail with:
   `ent/runtime/runtime.go:22:20: cannot use exampleDescID.Default.(func() "github.com/gofrs/uuid".UUID) (type func() "github.com/gofrs/uuid".UUID) as type func() "github.com/google/uuid".UUID in assignment`
5. Check `go.mod` and `go.sum` -- notice that `github.com/gofrs/uuid` have been added
