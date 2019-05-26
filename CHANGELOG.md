# Changelog
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]
### Changed
- No changes yet.

## [0.1.0] - 2019-05-26
### Added

- Middleware and components for [Gin](https://github.com/gin-gonic/gin) environment (panic recovery, error dispatcher and other).
- Configuration service (based on [spf13/viper](https://github.com/spf13/viper)).
- Request/Response models and the default handler implementation.
- Handler for `/index` action.
- Handler for `/tag/groups` action.
- Worker pool & events system for parallel performance.
- Concurrent rule-based processor for text analysis w/ [kljensen/snowball](https://github.com/kljensen/snowball) stemmer.
- Sanity index (SI) calculation implemented by a simple [weighted average](https://en.wikipedia.org/wiki/Weighted_arithmetic_mean) formula.
- Tests & benchmarks for the most important components (using [stretchr/testify](https://github.com/stretchr/testify)).

[Unreleased]: https://github.com/symfony-doge/ministry-of-truth-cis/compare/0.1.0...0.x
[0.1.0]: https://github.com/symfony-doge/ministry-of-truth-cis/releases/tag/0.1.0