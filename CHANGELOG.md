<p align="center"><img src="https://raw.githubusercontent.com/svengreb/golib/main/assets/images/repository-hero.svg?sanitize=true"/></p>

<p align="center"><a href="https://github.com/svengreb/golib/releases/latest" target="_blank"><img src="https://img.shields.io/github/release/svengreb/golib.svg?style=flat-square&label=Release&logo=github&logoColor=eceff4&colorA=4c566a&colorB=88c0d0"/></a> <a href="https://pkg.go.dev/github.com/svengreb/golib" target="_blank"><img src="https://img.shields.io/github/release/svengreb/golib.svg?style=flat-square&label=GoDoc&logo=go&logoColor=eceff4&colorA=4c566a&colorB=88c0d0"/></a></p>

<p align="center">Changelog for <a href="https://github.com/svengreb/golib" target="_blank">golib</a> — A core library of <a href="https://go.dev" target="_blank">Go</a> packages.</p>

<!--lint disable no-duplicate-headings no-duplicate-headings-in-section-->

# 0.1.0

![Release Date: 2020-11-21](https://img.shields.io/static/v1?style=flat-square&label=Release%20Date&message=2020-11-21&colorA=4c566a&colorB=88c0d0) [![Project Board](https://img.shields.io/static/v1?style=flat-square&label=Project%20Board&message=0.1.0&logo=github&logoColor=eceff4&colorA=4c566a&colorB=88c0d0)](https://github.com/svengreb/golib/projects/4) [![Milestone](https://img.shields.io/static/v1?style=flat-square&label=Milestone&message=0.1.0&logo=github&logoColor=eceff4&colorA=4c566a&colorB=88c0d0)](https://github.com/svengreb/golib/milestone/1)

⇅ [Show all commits][repo-compare-tag-init_v0.1.0]

This is the initial release version of _golib_.
The basic project setup, structure and development workflow has been bootstrapped by [the base _tmpl-go_ template repository][svengreb/tmpl-go].
Any additional change is covered in the following sections of this version changelog to introduce features, used technologies and explain why several decisions have been made.

## Features

<details>
<summary><strong>Bootstrapped project based on <em>tmpl-go</em> template repository</strong> — #1 (⊶ eb0b5be9)</summary>

<p align="center"><img src="https://github.com/svengreb/tmpl-go/blob/main/assets/images/repository-hero.svg?raw=true"/></p>

↠ The basic project setup, structure and development workflow has been bootstrapped [from version 0.5.0][svengreb/tmpl-go-rl-v0.5.0] of the [_tmpl-go_ template repository][svengreb/tmpl-go].
Additionally specific assets like the repository hero image have been replaced and documentations like the _README_ and GitHub issue/PR templates are adjusted.

</details>

<details>
<summary><strong>"io/fs" package: I/O utility functions for filesystem interactions</strong> — #3 ⇄ #4 (⊶ d1c7cc2b)</summary>

↠ The standard library provides the [io/ioutil][go-docs-pkg-io/ioutil] package that simplifies some common tasks like [reading all data from a file][go-docs-pkg-io/ioutil#readfile] and many more useful functions.

To simplify common I/O tasks related to interacting with filesystems, _golib_ provides the [io/fs][go-pkg-pkg/io/fs] package:

- `DirExists(string) (bool, error)` — checks if a directory exists.
- `FileExists(string) (bool, error)` — checks if a regular file or directory exists.
- `IsFileWritable(string) (bool, error)` — checks if a file is writable.
- `IsSymlink(string) (bool, error)` — checks if a file is a symbolic link.
- `RegularFileExists(string) (bool, error)` — checks if a regular file exists.

The package makes use of the underlying filesystem and only serves as additional utility for the [os][go-docs-pkg-os] Go standard library package. For more advanced and extended features see packages like [github.com/spf13/afero][] instead.

</details>

<details>
<summary><strong>"io/fs/filepath" package: utility functions for filename path</strong> — #5 ⇄ #6 (⊶ 008c9976)</summary>

↠ The standard library provides the [path/filepath][go-docs-pkg-path/filepath] package that simplifies some common filesystem path handling tasks like [_glob_ pattern matching][go-docs-pkg-path/filepath#glob] and many more useful functions.

_golib_ provides the [io/fs/filepath][go-pkg-pkg/io/fs/filepath] package that extends the [filepath][go-docs-pkg-path/filepath] package with more utilities.
For more advanced and extended features see packages like [github.com/spf13/afero][] instead.
Please note that some functions will interact with the underlying filesystem through on-disk operations!

- `IsSubDir(parentPath, subPath string, evalSymlinks bool) (bool, error)` — checks if a path is a subdirectory of another path.

</details>

<details>
<summary><strong>"vcs/git" package: VCS utility functions to interact with Git repositories</strong> — #8 ⇄ #9 (⊶ d21a6996)</summary>

↠ In order to simplify interactions with various [version control systems][wikip-vcs], the [vcs/git][go-pkg-pkg/vcs/git] package provides utility functions for [Git][] repositories.

It makes use of the [github.com/go-git/go-git/v5][] module, a highly extensible Git implementation in pure Go, instead of just trying to use the Git binary executable that may have been installed on the current system.
The [github.com/Masterminds/semver/v3][] module is used to provide a stable and flexible way to work with [“Semantic Versions“][semver].

</details>

<details>
<summary><strong>Initial project documentation</strong> — #7 ⇄ #10 (⊶ 41126f14)</summary>

↠ Wrote the initial [project documentation][gh-blob-readme] that includes…

1. …an project introduction and motivation.
2. …an overview of the module packages.
3. …information about how to contribute to this project.

</details>

<!--
+------------------+
+ Formatting Notes +
+------------------+

The `<summary />` tag must be separated with a blank line from the actual item content paragraph,
otherwise Markdown elements are not parsed and rendered!

+------------------+
+ Symbol Reference +
+------------------+
↠ (U+21A0): Start of a log section description
— (U+2014): Separator between a log section title and the metadata
⇄ (U+21C4): Separator between a issue ID and pull request ID in a log metadata
⊶ (U+22B6): Icon prefix for the short commit SHA checksum in a log metadata
⇅ (U+21C5): Icon prefix for the link of the Git commit history comparison on GitHub
-->

<!--lint disable final-definition-->

<!-- Base Links -->

<!-- v0.1.0 -->

[gh-blob-readme]: https://github.com/svengreb/golib/blob/main/README.md
[git]: https://git-scm.com
[github.com/go-git/go-git/v5]: https://pkg.go.dev/github.com/go-git/go-git/v5
[github.com/masterminds/semver/v3]: https://pkg.go.dev/github.com/Masterminds/semver/v3
[github.com/spf13/afero]: https://github.com/spf13/afero
[go-docs-pkg-io/ioutil]: https://golang.org/pkg/io/ioutil
[go-docs-pkg-io/ioutil#readfile]: https://golang.org/pkg/io/ioutil/#ReadFile
[go-docs-pkg-os]: https://golang.org/pkg/os
[go-docs-pkg-path/filepath]: https://golang.org/pkg/path/filepath
[go-docs-pkg-path/filepath#glob]: https://golang.org/pkg/path/filepath/#Glob
[go-pkg-pkg/io/fs]: https://pkg.go.dev/github.com/svengreb/golib/pkg/io/fs
[go-pkg-pkg/io/fs/filepath]: https://pkg.go.dev/github.com/svengreb/golib/pkg/io/fs/filepath
[go-pkg-pkg/vcs/git]: https://pkg.go.dev/github.com/svengreb/golib/pkg/vcs/git
[repo-compare-tag-init_v0.1.0]: https://github.com/svengreb/golib/compare/eb0b5be9...v0.1.0
[semver]: https://semver.org
[svengreb/tmpl-go-rl-v0.5.0]: https://github.com/svengreb/tmpl-go/releases/tag/v0.5.0
[svengreb/tmpl-go]: https://github.com/svengreb/tmpl-go
[wikip-vcs]: https://en.wikipedia.org/wiki/Version_control
