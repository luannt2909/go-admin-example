package util

import (
	cronDesc "github.com/lnquy/cron"
	"github.com/robfig/cron/v3"
	"time"
)

var exprDesc *cronDesc.ExpressionDescriptor

func init() {
	exprDesc, _ = cronDesc.NewDescriptor()
}

func Parse(spec string) (next time.Time, err error) {
	schedule, err := cron.ParseStandard(spec)
	if err != nil {
		return
	}
	next = schedule.Next(time.Now())
	return
}

func CronHumanReadable(spec string) string {
	if exprDesc == nil {
		expr, err := cronDesc.NewDescriptor()
		if err != nil {
			return spec
		}
		exprDesc = expr
	}
	desc, _ := exprDesc.ToDescription(spec, cronDesc.Locale_en)
	return desc
}
