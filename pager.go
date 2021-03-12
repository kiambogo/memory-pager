package pager

const DEFAULT_PAGE_SIZE = 1024

type Page struct {
	offset int
	buffer []byte
}

type Pager struct {
	pageSize int
	pages    []Page
}

func NewPager(pageSize int) *Pager {
	if pageSize == 0 {
		pageSize = DEFAULT_PAGE_SIZE
	}
	return &Pager{
		pageSize: pageSize,
	}
}

// Get will return the page at the specified index
func (p *Pager) Get(i int) *Page {
	return nil
}

// GetOrAlloc will return the page at the specified index, allocating it if not already allocated
func (p *Pager) GetOrAlloc(i int) *Page {
	return nil
}

// PageSize will return the page size of the pager
func (p *Pager) PageSize() int {
	return p.pageSize
}

// Len will return the size of the pager (number of pages)
func (p *Pager) Len() int {
	return len(p.pages)
}

// IsEmpty will check if the memory page is empty (has zero pages)
func (p *Pager) IsEmpty() bool {
	return len(p.pages) == 0
}

func (p *Pager) Set(index int, data []byte) {
	page := p.Get(index)
	page.buffer = p.truncate(data)
}

func (p *Pager) resize(index int) {
	// newPageListSize := len(p.pages) * 2
	//	for newPageListSize <= index
}

func (p Pager) truncate(buf []byte) []byte {
	if p.pageSize >= len(buf) {
		return buf
	}
	return buf[:p.pageSize]
}
