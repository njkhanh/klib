package widget

import (
	"math"
)

const (
	percentage = 2
)

type Paginator struct {
	Current    int
	Total      int
	Percentage int
	Next       int
	Prev       int
	First      int
	Middle     int
	Last       int
	Pages      []int
	Url        string
	Param      string
}

type PageLength struct {
	Start int
	End   int
	Total int
}

func RenderPaginator(current int, totalResults int, per int, param string, url string) Paginator {
	total := 0

	if current < 1 {
		current = 1
	}

	if totalResults > 0 && per > 0 {
		total = int(math.Floor(float64(totalResults+per-1) / float64(per)))
	}

	paginator := Paginator{
		Current:    current,
		Total:      total,
		Percentage: percentage,
		Prev:       1,
		Next:       total,
		First:      1,
		Last:       total,
		Param:      param,
		Url:        url,
	}

	if current > 2 {
		paginator.Prev = current - 1
	}

	if current < total {
		paginator.Next = current + 1
	}

	for i := 1; i <= total; i++ {
		if total < percentage*2 || (i >= current-percentage && i <= current+percentage) || (current > percentage*2 && total-i < percentage*2 && total-current < percentage*2) {
			paginator.Pages = append(paginator.Pages, i)
		}
	}

	if total > percentage*2 {
		paginator.Middle = (total + current) / 2
	}

	return paginator

}

func GetPageLength(total int, totalResults int, per int, current int) PageLength {
	start := 0
	if total > 0 {
		start = (current-1)*per + 1
	}

	return PageLength{
		Start: start,
		End:   per*(current-1) + total,
		Total: totalResults,
	}

}
