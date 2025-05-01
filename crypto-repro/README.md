# crypto-repro
This is a sample project to demonstrate the issues we have encountered with the fips image running some crypto libraries. In this example I am demonstrating the issues encountered with some functions in github.com/opencontainers/go-digest which is used across the container ecosystem. From debugging, it seems that the issue stems from a check to isMarshallable in the golang-fips library: https://github.com/golang-fips/openssl/blob/c494c216979153a2fc07afab57c0858f055ccdc0/hash.go#L227. This has the following function:

```go
// isHashMarshallable returns true if the memory layout of md
// is known by this library and can therefore be marshalled.
func isHashMarshallable(md ossl.EVP_MD_PTR) bool {
	if vMajor == 1 {
		return true
	}
	prov := ossl.EVP_MD_get0_provider(md)
	if prov == nil {
		return false
	}
	cname := ossl.OSSL_PROVIDER_get0_name(prov)
	if cname == nil {
		return false
	}
	name := C.GoString((*C.char)(unsafe.Pointer(cname)))
	// We only know the memory layout of the built-in providers.
	// See evpHash.hashState for more details.
	marshallable := name == "default" || name == "fips"
	return marshallable
}
```

in the fips image this check seems to fail because variable **name**=**symcryptprovider**. We are not entirely sure why this is happening as I understood that we would not end up invoking the fips library from go in the msft image.

## Repro

I have setup a couple of docker images to demonstrate the issue which can be built and ran from the make file as follows:

### Failure scenario
```bash
make run-fips-scenario
```

### Working in a non fips image
```bash
make run-normal-scenario
```

### Debugging 
For convenience I also added a devcontainer to this folder so the code can be ran with the debugger in the fips environment. It can be ran with the devcontainer extension and then repro.go can be launched using the default run (It will ask to install dlv at the bottom right of the vscode screen before running). Note that this folder has to be opened as the root of the vscode workspace.