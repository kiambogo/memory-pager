# mempgr

[![build status][1]][2]

an implementation of memory paging in go.

adapted from https://github.com/datrs/memory-pager.

archived as code has been moved over to https://github.com/kiambogo/go-hypercore.

## Installation

``` sh
go get github.com/kiambogo/mempgr
```

## Usage

``` go
import mempgr

pager := mempgr.NewPager(0) // Use default page size of 1024 bytes

page := pager.Get(5)        // get page #5
log.Println(page == nil)    // page is nil, as no page exists at index 5

page = pager.GetOrAlloc(5)  // get page #5
log.Println(page.Offset())  // 5120

pager.Set(10, []byte("hello, world!")) // set the buffer of page at index 10. truncates if buffer > page size
page = pager.Get(10)
log.Println(page.Offset())  // 10240
```

## License
[MIT](./LICENSE)

[1]: https://github.com/kiambogo/memory-pager/actions/workflows/test.yml/badge.svg
[2]: https://github.com/kiambogo/memory-pager/actions/workflows/test.yml
