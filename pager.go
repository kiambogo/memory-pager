package mpager

const DEFAULT_PAGE_SIZE = 1024

type Page struct {
	offset int
	buffer []byte
}

// Offset returns the byte offset of the page relvative to the other pages within the pager
func (p Page) Offset() int {
	return p.offset
}

type Pager struct {
	pageSize int
	pages    []*Page
}

func NewPager(pageSize int) Pager {
	if pageSize == 0 {
		pageSize = DEFAULT_PAGE_SIZE
	}
	return Pager{
		pageSize: pageSize,
		pages:    []*Page{},
	}
}

func (p Pager) newPage(index int, buf []byte) *Page {
	return &Page{
		offset: index * p.pageSize,
		buffer: buf,
	}
}

// Get will return the page at the specified index
func (p Pager) Get(pageNum int) *Page {
	if pageNum >= len(p.pages) {
		return nil
	}

	return p.pages[pageNum]
}

// GetOrAlloc will return the page at the specified index, allocating it if not already allocated
func (p *Pager) GetOrAlloc(pageNum int) (page *Page) {
	p.growPages(pageNum)

	if page = p.pages[pageNum]; page == nil {
		p.pages[pageNum] = p.newPage(pageNum, nil)
		page = p.pages[pageNum]
	}

	return page
}

// PageSize will return the page size of the pager
func (p Pager) PageSize() int {
	return p.pageSize
}

// Len will return the size of the pager (number of pages)
func (p Pager) Len() int {
	return len(p.pages)
}

// IsEmpty will check if the memory page is empty (has zero pages)
func (p Pager) IsEmpty() bool {
	return len(p.pages) == 0
}

// Set will set the contents of ajj
func (p *Pager) Set(pageNum int, data []byte) {
	page := p.GetOrAlloc(pageNum)
	page.buffer = p.truncate(data)
}

// growPages will increases the size of the pager's page buffer up till the supplied index
func (p *Pager) growPages(index int) {
	for i := len(p.pages); i <= index; i++ {
		p.pages = append(p.pages, nil)
	}
}

func (p Pager) truncate(buf []byte) []byte {
	if p.pageSize >= len(buf) {
		return buf
	}
	return buf[:p.pageSize]
}
