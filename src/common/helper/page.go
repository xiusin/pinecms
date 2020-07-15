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
	params    map[string][]string
	dynamic   bool
}

func NewPage(urlPrefix string, curPage, prePage, total int, params map[string][]string, dynamic bool) *Page {
	if dynamic == false {
		params = nil // 静态地址放弃参数传递
	} else if params != nil {
		delete(params, "page")
	}
	return &Page{
		total:     total,
		curPage:   curPage,
		prePage:   prePage,
		urlPrefix: urlPrefix,
		dynamic:   dynamic,
		params:    params,
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

func (p *Page) buildParams(page int) string {
	var q []string
	q = append(q, fmt.Sprintf("page=%d", page))
	if p.params != nil && p.dynamic {
		for k, v := range p.params {
			if len(v) == 0 {
				v[0] = ""
			}
			q = append(q, fmt.Sprintf("%s=%s", k, v[0]))
		}
	}
	return strings.Join(q, "&")
}

func (p *Page) getUrl(page int) string {
	if p.dynamic {
		return fmt.Sprintf("%s?%s", p.urlPrefix, p.buildParams(page))
	} else {
		format := "index.html"
		if page > 1 {
			format = fmt.Sprintf("index_%d.html", page)
		}
		return strings.TrimSuffix(filepath.Join(p.urlPrefix, format), "index.html")
	}
}

func (p *Page) GetPrev() string {
	return p.getUrl(p.curPage - 1)
}

func (p *Page) GetNext() string {
	return p.getUrl(p.curPage + 1)
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
		var cur string
		var url = p.getUrl(i)
		if p.curPage == i {
			cur = "class='curPage'"
			url = "javascript:void(0);"
		}
		fmt.Fprintf(&page, "<a href='%s' %s>%d</a>", url, cur, i)
	}
	if p.HasNext() {
		fmt.Fprintf(&page, "<a href='%s'>下一页</a>", p.GetNext())
	}
	fmt.Fprintf(&page, "<a href='%s'>尾页</a>", p.GetLast())
	return page.String()
}

