package helper

import (
	"fmt"
	"math"
	"path/filepath"
	"strings"
)

type Page struct {
	total     int
	curPage   int
	prePage   int
	urlPrefix string
	totalPage int
}

func NewPage(urlPrefix string, curPage, prePage, total int) *Page {
	return &Page{
		total:     total,
		curPage:   curPage,
		prePage:   prePage,
		urlPrefix: urlPrefix,
		totalPage: int(math.Ceil(float64(total) / float64(prePage))),
	}
}

func (p *Page) GetFirst() string {
	return p.getUrl(1)
}

func (p *Page) HasPrev() bool {
	return p.curPage > 1
}

func (p *Page) HasNext() bool {
	return p.curPage < p.totalPage
}

func (p *Page) getUrl(page int) string {
	format := "index.html"
	if page > 1 {
		format = fmt.Sprintf("index_%d.html", page)
	}
	return filepath.Join(p.urlPrefix, format)
}

func (p *Page) GetPrev() string {
	return p.getUrl(p.curPage-1)
}

func (p *Page) GetNext() string {
	return p.getUrl(p.curPage+1)
}

func (p *Page) GetLast() string {
	return p.getUrl(p.totalPage)
}

func (p *Page) String() string {
	if p.totalPage == 0 {
		return ""
	}
	var page strings.Builder
	fmt.Fprintf(&page, "<a href='%s'>首页</a>", p.GetFirst())
	if p.HasPrev() {
		fmt.Fprintf(&page, "<a href='%s'>上一页</a>", p.GetPrev())
	}
	start := p.curPage - 3
	endStep := 3
	if start <= 1 {
		endStep -= start
		start = 1
	}
	end := p.curPage + endStep
	for i := start; i <= end; i++ {
		if i > p.totalPage {
			break
		}
		fmt.Fprintf(&page, "<a href='%s'>%d</a>", p.getUrl(i), i)
	}
	if p.HasNext() {
		fmt.Fprintf(&page, "<a href='%s'>下一页</a>", p.GetNext())
	}
	fmt.Fprintf(&page, "<a href='%s'>尾页</a>", p.GetLast())
	return page.String()
}
