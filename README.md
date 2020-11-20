<p align="center"><img src="https://raw.githubusercontent.com/svengreb/golib/main/assets/images/repository-hero.svg?sanitize=true"/></p>

<p align="center"><a href="https://github.com/svengreb/golib/releases/latest" target="_blank"><img src="https://img.shields.io/github/release/svengreb/golib.svg?style=flat-square&label=Release&logo=github&logoColor=eceff4&colorA=4c566a&colorB=88c0d0"/></a> <a href="https://pkg.go.dev/github.com/svengreb/golib" target="_blank"><img src="https://img.shields.io/github/release/svengreb/golib.svg?style=flat-square&label=GoDoc&logo=go&logoColor=eceff4&colorA=4c566a&colorB=88c0d0"/></a> <a href="https://github.com/svengreb/golib/blob/main/CHANGELOG.md" target="_blank"><img src="https://img.shields.io/github/release/svengreb/golib.svg?style=flat-square&label=Changelog&logo=github&logoColor=eceff4&colorA=4c566a&colorB=88c0d0"/></a></p>

<p align="center"><a href="https://github.com/svengreb/golib/actions?query=workflow%3Aci" target="_blank"><img src="https://img.shields.io/github/workflow/status/svengreb/golib/ci.svg?style=flat-square&label=CI&logo=github&logoColor=eceff4&colorA=4c566a"/></a></p>

<p align="center"><a href="https://golang.org/doc/effective_go.html#formatting" target="_blank"><img src="https://img.shields.io/static/v1?style=flat-square&label=Go%20Style%20Guide&message=gofmt&logo=go&logoColor=eceff4&colorA=4c566a&colorB=88c0d0"/></a> <a href="https://github.com/arcticicestudio/styleguide-markdown/releases/latest" target="_blank"><img src="https://img.shields.io/github/release/arcticicestudio/styleguide-markdown.svg?style=flat-square&label=Markdown%20Style%20Guide&logoColor=eceff4&colorA=4c566a&colorB=88c0d0&logo=data%3Aimage%2Fsvg%2Bxml%3Bbase64%2CPHN2ZyB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIHdpZHRoPSIzOSIgaGVpZ2h0PSIzOSIgdmlld0JveD0iMCAwIDM5IDM5Ij48cGF0aCBmaWxsPSJub25lIiBzdHJva2U9IiNEOERFRTkiIHN0cm9rZS13aWR0aD0iMyIgc3Ryb2tlLW1pdGVybGltaXQ9IjEwIiBkPSJNMS41IDEuNWgzNnYzNmgtMzZ6Ii8%2BPHBhdGggZmlsbD0iI0Q4REVFOSIgZD0iTTIwLjY4MyAyNS42NTVsNS44NzItMTMuNDhoLjU2Nmw1Ljg3MyAxMy40OGgtMS45OTZsLTQuMTU5LTEwLjA1Ni00LjE2MSAxMC4wNTZoLTEuOTk1em0tMi42OTYgMGwtMTMuNDgtNS44NzJ2LS41NjZsMTMuNDgtNS44NzJ2MS45OTVMNy45MzEgMTkuNWwxMC4wNTYgNC4xNnoiLz48L3N2Zz4%3D"/></a> <a href="https://github.com/arcticicestudio/styleguide-git/releases/latest" target="_blank"><img src="https://img.shields.io/github/release/arcticicestudio/styleguide-git.svg?style=flat-square&label=Git%20Style%20Guide&logoColor=eceff4&colorA=4c566a&colorB=88c0d0&logo=git"/></a></p>

<p align="center">A core library of <a href="https://go.dev" target="_blank">Go</a> packages</p>

_golib_ is a [collection of Go packages][go-doc-mod] that provide various helpful functions and types for many common use cases. It is designed with and for projects that follow the [DRY][wikip-dry], [KISS][wikip-kiss_prin] and/or [Unix][wikip-unix_phil] principles to improve the [reuse of code][wikip-code_reuse] while reducing [“Copy-and-paste programming“][wikip-cp_prog] and [duplicated code][wikip-dup_code] as much as possible. The packages are created to be flexible for universal usage and composition for different use cases and environments.

Please always keep in mind that _golib_ was mainly created for usage in my open source and personal projects!
Even though this module follows the [_SemVer_ 2.0.0 specification][semver-spec-v2.0.0] and development of the provided packages take stability and backwards compatibility into account, there might be (possibly breaking) changes, at least while this module in [_SemVer_ major version zero][semver#major_ver_zero].
Also note that packages might be extracted into own modules if the scope exceeds the goal of _golib_. Such changes will be communicated by raising the [Go module major version][go-docs-mod#versions].
Most packages are independent from each other and strive to depend on as few external packages as possible, but ensure that you're aware of the internal structure of the packages you're using.

## Overview

This is an rough overview of the packages included in this module. Please refer to the individual package documentations for more details.

- [pkg/io/fs][go-pkg-pkg/io/fs] — provides I/O utility functions for filesystem interactions. Please note that it makes use of the underlying filesystem and only serves as additional utility for the [os][go-docs-pkg-os] Go standard library package.
  For more advanced and extended features see packages like [github.com/spf13/afero][go-pkg-github.com/spf13/afero] instead.
- [pkg/io/fs/filepath][go-pkg-pkg/io/fs/filepath] — provides utility functions for manipulating filename paths for the target operating system-defined file paths, using either forward slashes or backslashes. It extends the [path/filepath][go-docs-pkg-path/filepath] Go standard library package with more utilities.
  Please note that some functions interact with the underlying filesystem through on-disk operations.
- [pkg/vcs][go-pkg-pkg/vcs] — provides packages and utility functions to interact with [version control systems][wikip-vcs].
  - [pkg/vcs/git][go-pkg-pkg/vcs/git] — provides VCS utility functions to interact with [Git][] repositories.

## Contributing

_golib_ is an open source project and contributions are always welcome!

There are many ways to contribute, from [writing- and improving documentation and tutorials][contrib-guide-docs], [reporting bugs][contrib-guide-bugs], [submitting enhancement suggestions][contrib-guide-enhance] that can be added to _golib_ by [submitting pull requests][contrib-guide-pr].

Please take a moment to read the [contributing guide][contrib-guide] to learn about the development process, the [styleguides][contrib-guide-styles] to which this project adheres as well as the [branch organization][contrib-guide-branching] and [versioning][contrib-guide-versioning] model.

The guide also includes information about [minimal, complete, and verifiable examples][contrib-guide-mcve] and other ways to contribute to the project like [improving existing issues][contrib-guide-impr-issues] and [giving feedback on issues and pull requests][contrib-guide-feedback].

<p align="center">Copyright &copy; 2020-present <a href="https://www.svengreb.de" target="_blank">Sven Greb</a></p>

<p align="center"><a href="https://github.com/svengreb/golib/blob/main/LICENSE"><img src="https://img.shields.io/static/v1.svg?style=flat-square&label=License&message=MIT&logoColor=eceff4&logo=github&colorA=4c566a&colorB=88c0d0"/></a></p>

[contrib-guide-branching]: https://github.com/svengreb/golib/blob/main/CONTRIBUTING.md#branch-organization
[contrib-guide-bugs]: https://github.com/svengreb/golib/blob/main/CONTRIBUTING.md#bug-reports
[contrib-guide-docs]: https://github.com/svengreb/golib/blob/main/CONTRIBUTING.md#documentations
[contrib-guide-enhance]: https://github.com/svengreb/golib/blob/main/CONTRIBUTING.md#enhancement-suggestions
[contrib-guide-feedback]: https://github.com/svengreb/golib/blob/main/CONTRIBUTING.md#give-feedback-on-issues-and-pull-requests
[contrib-guide-impr-issues]: https://github.com/svengreb/golib/blob/main/CONTRIBUTING.md#improve-issues
[contrib-guide-mcve]: https://github.com/svengreb/golib/blob/main/CONTRIBUTING.md#mcve
[contrib-guide-pr]: https://github.com/svengreb/golib/blob/main/CONTRIBUTING.md#pull-requests
[contrib-guide-styles]: https://github.com/svengreb/golib/blob/main/CONTRIBUTING.md#style-guides
[contrib-guide-versioning]: https://github.com/svengreb/golib/blob/main/CONTRIBUTING.md#versioning
[contrib-guide]: https://github.com/svengreb/golib/blob/main/CONTRIBUTING.md
[git]: https://git-scm.com
[go-doc-mod]: https://golang.org/ref/mod
[go-docs-mod#versions]: https://golang.org/ref/mod#versions
[go-docs-pkg-os]: https://golang.org/pkg/os
[go-docs-pkg-path/filepath]: https://golang.org/pkg/path/filepath
[go-pkg-github.com/spf13/afero]: https://pkg.go.dev/github.com/spf13/afero
[go-pkg-pkg/io/fs]: https://pkg.go.dev/github.com/svengreb/golib/pkg/io/fs
[go-pkg-pkg/io/fs/filepath]: https://pkg.go.dev/github.com/svengreb/golib/pkg/io/fs/filepath
[go-pkg-pkg/vcs]: https://pkg.go.dev/github.com/svengreb/golib/pkg/vcs
[go-pkg-pkg/vcs/git]: https://pkg.go.dev/github.com/svengreb/golib/pkg/vcs/git
[semver-spec-v2.0.0]: https://semver.org/spec/v2.0.0.html
[semver#major_ver_zero]: https://semver.org/#spec-item-4
[wikip-code_reuse]: https://en.wikipedia.org/wiki/Code_reuse
[wikip-cp_prog]: https://en.wikipedia.org/wiki/Copy-and-paste_programming
[wikip-dry]: https://en.wikipedia.org/wiki/Don%27t_repeat_yourself
[wikip-dup_code]: https://en.wikipedia.org/wiki/Duplicate_code
[wikip-kiss_prin]: https://en.wikipedia.org/wiki/KISS_principle
[wikip-unix_phil]: https://en.wikipedia.org/wiki/Unix_philosophy
[wikip-vcs]: https://en.wikipedia.org/wiki/Version_control
