// automatically generated by stateify.

package pipe

import (
	"context"

	"gvisor.dev/gvisor/pkg/state"
)

func (p *pipe) StateTypeName() string {
	return "pkg/tcpip/link/sharedmem/pipe.pipe"
}

func (p *pipe) StateFields() []string {
	return []string{
		"buffer",
	}
}

func (p *pipe) beforeSave() {}

// +checklocksignore
func (p *pipe) StateSave(stateSinkObject state.Sink) {
	p.beforeSave()
	stateSinkObject.Save(0, &p.buffer)
}

func (p *pipe) afterLoad(context.Context) {}

// +checklocksignore
func (p *pipe) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &p.buffer)
}

func (r *Rx) StateTypeName() string {
	return "pkg/tcpip/link/sharedmem/pipe.Rx"
}

func (r *Rx) StateFields() []string {
	return []string{
		"p",
		"tail",
		"head",
	}
}

func (r *Rx) beforeSave() {}

// +checklocksignore
func (r *Rx) StateSave(stateSinkObject state.Sink) {
	r.beforeSave()
	stateSinkObject.Save(0, &r.p)
	stateSinkObject.Save(1, &r.tail)
	stateSinkObject.Save(2, &r.head)
}

func (r *Rx) afterLoad(context.Context) {}

// +checklocksignore
func (r *Rx) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &r.p)
	stateSourceObject.Load(1, &r.tail)
	stateSourceObject.Load(2, &r.head)
}

func (t *Tx) StateTypeName() string {
	return "pkg/tcpip/link/sharedmem/pipe.Tx"
}

func (t *Tx) StateFields() []string {
	return []string{
		"p",
		"maxPayloadSize",
		"head",
		"tail",
		"next",
		"tailHeader",
	}
}

func (t *Tx) beforeSave() {}

// +checklocksignore
func (t *Tx) StateSave(stateSinkObject state.Sink) {
	t.beforeSave()
	stateSinkObject.Save(0, &t.p)
	stateSinkObject.Save(1, &t.maxPayloadSize)
	stateSinkObject.Save(2, &t.head)
	stateSinkObject.Save(3, &t.tail)
	stateSinkObject.Save(4, &t.next)
	stateSinkObject.Save(5, &t.tailHeader)
}

func (t *Tx) afterLoad(context.Context) {}

// +checklocksignore
func (t *Tx) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &t.p)
	stateSourceObject.Load(1, &t.maxPayloadSize)
	stateSourceObject.Load(2, &t.head)
	stateSourceObject.Load(3, &t.tail)
	stateSourceObject.Load(4, &t.next)
	stateSourceObject.Load(5, &t.tailHeader)
}

func init() {
	state.Register((*pipe)(nil))
	state.Register((*Rx)(nil))
	state.Register((*Tx)(nil))
}