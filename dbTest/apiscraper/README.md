# apiscraper

This packages provdes a `apiscraper` binary which will fetch from the GGG stash tab api until the provided maximum size is reached. Individual responses are stored in `stash.CompressedResponse` while a chronologically ordered list of responses are in `stash.ChangeSet`.

For up to date info: `apiscraper -help`