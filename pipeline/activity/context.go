package activity

import "github.com/naveego/api/types/pipeline"

type Context struct {
	Pipeline        pipeline.Pipeline // The Pipeline being executed
	Activity        pipeline.Activity // The currently executing Activity
	OutputCollector OutputCollector   // The output collector for the activity
}
