# poeitemstore [![CircleCI](https://circleci.com/gh/Everlag/poeitemstore.svg?style=svg)](https://circleci.com/gh/Everlag/poeitemstore)

This is a [Path of Exile stash tab](https://www.pathofexile.com/developer/docs/api-resource-public-stash-tabs) indexer based on [boltdb](https://github.com/boltdb/bolt) aimed at maximum performance with minimal disk space.

All tests are in `dbTest` and test only functionality exposed by `db`.

## Optimizations

Crossed out indicates didn't work out.

### Indexes

Bucketing IDs into temporally and value-wise similar

~~Compression of index values~~ overhead was too high for our workload, may revist in future with added metadata and optional compression based on workload in IndexEntry.

~~Set pooling~~ clearing maps costs too much between IndexQueries. Switching to bitsets, both [dense](https://github.com/willf/bitset) and [sparse](https://github.com/js-ojus/sparsebitset) end up with significantly poorer performance. Did not try roaring bitmaps.

## License

poeitemstore is licensed under either of

 * Apache License, Version 2.0, ([LICENSE-APACHE](LICENSE-APACHE) or
   http://www.apache.org/licenses/LICENSE-2.0)
 * MIT license ([LICENSE-MIT](LICENSE-MIT) or
   http://opensource.org/licenses/MIT)

at your option.
