# Change Log


All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](http://keepachangelog.com/en/1.0.0/)
and this project adheres to [Semantic Versioning](http://semver.org/spec/v2.0.0.html).


## [Unreleased]

### Added

- `Init` function


## [0.7.0] - 2018-07-30

### Added

- VisitAll function
- Preserve order of variables (and also sort them if necessary)
- Variable lookup

### Changed

- Variable usage lists variables in lexicographical or primordial order
- Empty values are not filtered out from query string maps


## [0.6.0] - 2018-06-22

### Added

- String slice value type


## [0.5.0] - 2018-06-14

### Added

- Query string value type
- `HasEnvVars` to determine if the `EnvVarSet` has variables defined


## [0.4.0] - 2018-06-12

### Added

- Simple usage string


## [0.3.0] - 2018-06-08

### Added

- `NormalizedName` for normalized environment variable names
- Variable name normalization with a sane default
- `NormalizeFunc` to make variable normalization configurable
- `VarE` function to return the created environment variable

### Changed

- Global `Parse` does not return an error anymore


## [0.2.0] - 2018-06-06

### Added

- `EnvVar` struct

### Changes

- `Var` panics when redeclaring an existing variable
- Error output writer
- Store usage


## 0.1.0 - 2018-06-06

- Initial, preview release


[Unreleased]: https://github.com/goph/env/compare/v0.7.0...HEAD
[0.7.0]: https://github.com/goph/env/compare/v0.6.0...v0.7.0
[0.6.0]: https://github.com/goph/env/compare/v0.5.0...v0.6.0
[0.5.0]: https://github.com/goph/env/compare/v0.4.0...v0.5.0
[0.4.0]: https://github.com/goph/env/compare/v0.3.0...v0.4.0
[0.3.0]: https://github.com/goph/env/compare/v0.2.0...v0.3.0
[0.2.0]: https://github.com/goph/env/compare/v0.1.0...v0.2.0
