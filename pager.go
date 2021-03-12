package pager

const DEFAULT_PAGE_SIZE = 1024

type Page struct {
	offset int
	buffer *[]byte
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

func (p *Pager) Get(i int) *Page {
	return nil
}

func (p *Pager) Set(index int, data []byte) {
	page := p.Get(index)
	page.buffer = p.truncate(data)
}

func (p *Pager) ToBuffer() {}

func (p *Pager) grow(index int) {
	newPageListSize := len(p.pages) * 2
	for newPageListSize <= index
}

func (p Pager) truncate(buf []byte) []byte {
	if p.pageSize >= len(buf) {
		return buf
	}
	return buf[:p.pageSize]
}
