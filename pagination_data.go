package pagination

type PaginaionData struct {
	DataCount         int `form:"DataCount" json:"DataCount" xml:"DataCount"`
	BeginDataIndex    int `form:"BeginDataIndex" json:"BeginDataIndex" xml:"BeginDataIndex"`
	EndDataIndex      int `form:"EndDataIndex" json:"EndDataIndex" xml:"EndDataIndex"`
	Pages             []int `form:"Pages" json:"Pages" xml:"Pages"`
	PageIndex         int `form:"PageIndex" json:"PageIndex" xml:"PageIndex"`
	PreviousPageIndex int `form:"PreviousPageIndex" json:"PreviousPageIndex" xml:"PreviousPageIndex"`
	NextPageIndex     int `form:"NextPageIndex" json:"NextPageIndex" xml:"NextPageIndex"`
	FirstPageIndex    int `form:"FirstPageIndex" json:"FirstPageIndex" xml:"FirstPageIndex"`
	LastPageIndex     int `form:"LastPageIndex" json:"LastPageIndex" xml:"LastPageIndex"`
	IsShowPrevious    bool `form:"IsShowPrevious" json:"IsShowPrevious" xml:"IsShowPrevious"`
	IsShowNext        bool `form:"IsShowNext" json:"IsShowNext" xml:"IsShowNext"`
	IsShowFirst       bool `form:"IsShowFirst" json:"IsShowFirst" xml:"IsShowFirst"`
	IsShowLast        bool `form:"IsShowLast" json:"IsShowLast" xml:"IsShowLast"`
}
