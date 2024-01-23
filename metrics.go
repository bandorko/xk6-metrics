package metrics

import (
	"sync"
	"time"

	"go.k6.io/k6/event"
	"go.k6.io/k6/js/modules"
	"go.k6.io/k6/metrics"
)

func init() {
	modules.Register("k6/x/metrics", New())
}

type extraMetrics struct {
	DropPercentage *metrics.Metric
}

type (
	RootModule struct {
		initOnce sync.Once
	}

	ModuleInstance struct {
		vu modules.VU
	}
)

var (
	_ modules.Instance = &ModuleInstance{}
	_ modules.Module   = &RootModule{}
)

func New() *RootModule {
	return &RootModule{
		initOnce: sync.Once{},
	}
}

func (rm *RootModule) NewModuleInstance(vu modules.VU) modules.Instance {
	rm.initOnce.Do(
		func() {
			registry := vu.InitEnv().Registry
			m := &extraMetrics{
				DropPercentage: registry.MustNewMetric("drop_percentage", metrics.Counter),
			}
			sid, evtCh := vu.Events().Global.Subscribe(event.TestEnd)
			go func() {
				for {
					select {
					case evt, ok := <-evtCh:
						if !ok {
							return
						}
						evt.Done()
						if evt.Type == event.TestEnd {
							droppedIterations := registry.Get("dropped_iterations").Sink.(*metrics.CounterSink).Value
							iterations := registry.Get("iterations").Sink.(*metrics.CounterSink).Value
							dropPercentage := droppedIterations / (iterations + droppedIterations)
							m.DropPercentage.Sink.Add(
								metrics.Sample{
									Time: time.Now(),
									TimeSeries: metrics.TimeSeries{
										Metric: m.DropPercentage,
										Tags:   &metrics.TagSet{},
									},
									Value: float64(dropPercentage),
								})
							vu.Events().Global.Unsubscribe(sid)
						}
					case <-vu.Context().Done():
						return
					}
				}
			}()
		})

	return &ModuleInstance{
		vu: vu,
	}
}

func (mi *ModuleInstance) Exports() modules.Exports {
	return modules.Exports{
		Named: map[string]interface{}{
			"registry": mi.vu.InitEnv().Registry,
		},
	}
}
