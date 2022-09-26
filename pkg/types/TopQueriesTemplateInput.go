package types

import (
	"bytes"
	"fmt"
	"github.com/ChistaDATA/ChistaDATA-Profiler-for-ClickHouse.git/pkg/formatters"
	log "github.com/sirupsen/logrus"
	"html/template"
	"strconv"
	"strings"
)

type TopQueriesTemplateInputRecord struct {
	Query                   string
	Pos                     string
	TotalDuration           string
	Count                   string
	TotalDurationPercentage string
	ResponseTimePerCall     string
}

func InitTopQueriesTemplateInputRecord(info *QueryInfo, queryInfoTemplateInput *QueryInfoTemplateInput, totalExecutionTime float64) TopQueriesTemplateInputRecord {
	queryLengthLimit := 80
	if len(info.Query) < queryLengthLimit {
		queryLengthLimit = len(info.Query)
	}
	return TopQueriesTemplateInputRecord{
		Query:                   info.Query[:queryLengthLimit],
		Pos:                     formatters.PrefixSpace(strings.TrimSpace(queryInfoTemplateInput.Pos), 4),
		TotalDuration:           formatters.PrefixSpace(formatters.Float64SecondsToString(info.GetTotalDuration(), 7), 7),
		Count:                   formatters.PrefixSpace(strconv.Itoa(info.Count), 5),
		TotalDurationPercentage: formatters.PrefixSpace(fmt.Sprintf("%.2f%s", info.GetTotalDuration()/totalExecutionTime, "%"), 7),
		ResponseTimePerCall:     formatters.PrefixSpace(formatters.Float64SecondsToString(info.GetTotalDuration()/float64(info.Count), 6), 6),
	}
}

type TopQueriesTemplateInput struct {
	Records string
}

func InitTopQueriesTemplateInput(records []TopQueriesTemplateInputRecord, recordTemplateString string) TopQueriesTemplateInput {
	temp, err := template.New("tmp").Parse(recordTemplateString)
	if err != nil {
		log.Fatalln(err.Error())
		panic(err)
	}
	recordString := ""
	for _, record := range records {
		var bf bytes.Buffer
		err := temp.Execute(&bf, record)
		if err == nil {
			recordString = recordString + bf.String() + "\n"
		}
	}
	return TopQueriesTemplateInput{Records: recordString}
}
