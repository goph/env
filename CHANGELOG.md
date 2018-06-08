# Change Log


All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](http://keepachangelog.com/en/1.0.0/)
and this project adheres to [Semantic Versioning](http://semver.org/spec/v2.0.0.html).


## [Unreleased]

### Added

- `NormalizedName` for normalized environment variable names
- Variable name normalization with a sane default


## [0.2.0] - 2018-06-06

### Added

- `EnvVar` struct

### Changes

- `Var` panics when redeclaring an existing variable
- Error output writer
- Store usage


## 0.1.0 - 2018-06-06

- Initial, preview release


[Unreleased]: https://github.com/goph/emperror/compare/v0.2.0...HEAD
[0.2.0]: https://github.com/goph/emperror/compare/v0.1.0...v0.2.0
