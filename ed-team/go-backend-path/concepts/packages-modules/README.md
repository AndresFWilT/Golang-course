# Packages - Modules

We can create packages (directories) where we can have multiple .go files, each file must have the same package.
We can only use the functions, constants, types (stuff) of the files from another package if those stuff is an exportable member. We can infer that a function, structure, type is exportable if it is written as Camel Case.

```Go
func ExportableFunction(){
	// some stuff
}
```

To use from other packages our module we can use the `packageName`.`ExportableFunction`.

You can also push to a remote repository and import the package as a module in your app. After pushed the repository sample **GitHub** we can:

```sh
go mod init user/github/package-url
```

and we can import it in the required modules:

```Go
package main

import (
	"fmt"
	"pushed-package/package"
)

func main() {
	package.ExportableFunction()
}
```

To download all dependencies required and clean up the unused ones, use:

```shell
go mod tidy
```

That command will download the latest versions, unless a dependency that depends on others, specify a version

The packages in the `go.mod` has a hash which allows Go to identify the dependencies in a correct manner.

If we want to list all the modules our project has, use:

```shell
go list -m all
```
The dependencies direct and indirect require sometimes more dependencies to work, so it's probably that the go list will show us all the modules used in the package, **ALL**

If we want to know the dependencies that depends on other modules use:

```shell
go mod graph
```

## Versioning

All the packages has a version, each of those have a notation: `V1.0.0`, but if a module has git tags, Go will also use the git tag to the versioning: `V1.0.0-git-tag` if Go don't found the git tag, Go will create one (`v1.0.0-date-commithash`.). [Semver](https://semver.org)

```text
VMayor.Minor.patch-identifier
V1.2.3-pre-release
```

non stable versions:

```text
v0.0.0
```

stable versions:

```text
v1.0.0 with git tag release
```

- Patch: bugs related to minor version
- Minor: Updates related to major version
- Major: Break compatibility with earlier versions

## Downgrade Versions

to show the versions of a package use:

```shell
go list -m -versions package
```

to upgrade to latest versions:

```shell
go get package
```

To downgrade to a desired version:

```shell
go get package@v.1.0.0
```

## Upgrade Versions

By default, when we import a package:

```Go
package main

import "package/Module"
```

it will be imported as latest version of V1. And the command list will show only the versions from v0 to v1. If we want to use mayor versions then:

```Go
package main

import "package/Module/V2"
```

That was a sample, but we must check the tag releases.

To create aliases to packages:

```Go
package main

import (
	"package/module"
	moduleV2 "package/module/V2"
)
```

### Clean Unused Versions

To clean all packages that are unused, use:

```shell
go mod tidy
```

## Import a package with secondary effects

```Go
package main

import (
	"fmt"
	_ "github.com/lib/pq"
)
```

this will import the package for secondary effects: it will call the `init` func and will block to not use the exported identifiers inside the package.
 