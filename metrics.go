package metrics

import (
	"go.k6.io/k6/js/modules"
	"go.k6.io/k6/metrics"
)

func init() {
	modules.Register("k6/x/metrics", New())
}

type (
	RootModule struct{}

	ModuleInstance struct {
		vu modules.VU
	}
)

var (
	_ modules.Instance = &ModuleInstance{}
	_ modules.Module   = &RootModule{}
)

func New() *RootModule {
	return &RootModule{}
}

func (*RootModule) NewModuleInstance(vu modules.VU) modules.Instance {
	return &ModuleInstance{
		vu: vu,
	}
}

func (mi *ModuleInstance) Exports() modules.Exports {
	return modules.Exports{
		Named: map[string]interface{}{
			"getDropPercentage": func() interface{} {
				droppedIterations := mi.vu.State().BuiltinMetrics.DroppedIterations.Sink.(*metrics.CounterSink).Value
				iterations := mi.vu.State().BuiltinMetrics.Iterations.Sink.(*metrics.CounterSink).Value
				return droppedIterations / (iterations + droppedIterations)
			},
		},
	}
}
