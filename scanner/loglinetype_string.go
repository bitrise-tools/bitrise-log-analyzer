// Code generated by "stringer -type=LogLineType"; DO NOT EDIT

package scanner

import "fmt"

const _LogLineType_name = "BeforeFirstStepStepInfoHeaderStepLogStepInfoFooterBetweenStepsBuildSummaryStepInfoHeaderOrBuildSummarySectionStarterAfterBuildSummary"

var _LogLineType_index = [...]uint8{0, 15, 29, 36, 50, 62, 74, 116, 133}

func (i LogLineType) String() string {
	if i < 0 || i >= LogLineType(len(_LogLineType_index)-1) {
		return fmt.Sprintf("LogLineType(%d)", i)
	}
	return _LogLineType_name[_LogLineType_index[i]:_LogLineType_index[i+1]]
}
