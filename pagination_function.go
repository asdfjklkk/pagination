package pagination

import (
	"math"
	"sort"
)

func NewPaginationData(dataCount int, pageSize int, pageIndex int) (returnValue PaginaionData) {
	if pageIndex <= 0 {
		pageIndex = 1
	}
	//前端總共顯示幾個數字頁籤
	visiblePages := 10
	pages := make([]int, 0)
	isShowPrevious := false
	isShowNext := false
	isShowFirst := false
	isShowLast := false
	beginDataIndex := 0
	endDataIndex := 0
	//共幾頁
	pageCount := math.Ceil(float64(dataCount) / float64(pageSize))
	if pageIndex > int(pageCount) {
		pageIndex = 1
	}
	beginDataIndex = 1 + ((pageIndex - 1) * pageSize)
	endDataIndex = beginDataIndex + (pageSize - 1)
	if endDataIndex > dataCount {
		endDataIndex = dataCount
	}
	pages = append(pages, pageIndex)
	minPageIndex := pageIndex
	//5 = 往前顯示 5 個頁籤，例如目前索引值若為 6，就是顯示 1 ~ 6
	for i := 1; i <= 5; i++ {
		tempI := pageIndex - i
		if tempI <= 0 {
			break
		}
		if len(pages) >= visiblePages {
			break
		}
		pages = append(pages, tempI)
		if tempI < minPageIndex {
			minPageIndex = tempI
		}
	}
	//50 = 往後顯示 50 個頁籤，例如目前索引值若為 6，就是顯示 7 ~ XX，但實際上不會顯示這麼多頁籤
	for i := 1; i <= 50; i++ {
		tempI := pageIndex + i
		if tempI > int(pageCount) {
			break
		}
		pages = append(pages, tempI)
		if len(pages) >= visiblePages {
			break
		}
	}
	//若總頁籤不足，再往前追加
	if len(pages) < visiblePages {
		for i := 1; i <= 5; i++ {
			tempI := minPageIndex - i
			if tempI <= 0 {
				break
			}
			if len(pages) >= visiblePages {
				break
			}
			pages = append(pages, tempI)
		}
	}
	sort.Ints(pages)
	for _, v := range pages {
		if v < pageIndex {
			isShowPrevious = true
			isShowFirst = true
		}
		if v > pageIndex {
			isShowNext = true
			isShowLast = true
		}
	}
	returnValue.DataCount = dataCount
	returnValue.BeginDataIndex = beginDataIndex
	returnValue.EndDataIndex = endDataIndex
	returnValue.PageIndex = pageIndex
	returnValue.PreviousPageIndex = pageIndex - 1
	returnValue.NextPageIndex = pageIndex + 1
	returnValue.FirstPageIndex = 1
	returnValue.LastPageIndex = int(pageCount)
	returnValue.Pages = pages
	returnValue.IsShowPrevious = isShowPrevious
	returnValue.IsShowNext = isShowNext
	returnValue.IsShowFirst = isShowFirst
	returnValue.IsShowLast = isShowLast
	return
}
