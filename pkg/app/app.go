package app

import (
	nf_context "github.com/andy89923/nf-example/internal/context"
	"github.com/andy89923/nf-example/pkg/factory"
)

type App interface {
	SetLogEnable(enable bool)

	SetLogLevel(level string)

	SetReportCaller(reportCaller bool)

	Start()

	Terminate()

	Context() *nf_context.NFContext

	Config() *factory.Config
}
