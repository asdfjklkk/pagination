package pagination

type PaginaionData struct {
	DataCount         int
	BeginDataIndex    int
	EndDataIndex      int
	Pages             []int
	PageIndex         int
	PreviousPageIndex int
	NextPageIndex     int
	FirstPageIndex    int
	LastPageIndex     int
	IsShowPrevious    bool
	IsShowNext        bool
	IsShowFirst       bool
	IsShowLast        bool
}
