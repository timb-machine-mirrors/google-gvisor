// automatically generated by stateify.

package fsutil

import (
	"context"

	"gvisor.dev/gvisor/pkg/state"
)

func (d *DirtyInfo) StateTypeName() string {
	return "pkg/sentry/fsutil.DirtyInfo"
}

func (d *DirtyInfo) StateFields() []string {
	return []string{
		"Keep",
	}
}

func (d *DirtyInfo) beforeSave() {}

// +checklocksignore
func (d *DirtyInfo) StateSave(stateSinkObject state.Sink) {
	d.beforeSave()
	stateSinkObject.Save(0, &d.Keep)
}

func (d *DirtyInfo) afterLoad(context.Context) {}

// +checklocksignore
func (d *DirtyInfo) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &d.Keep)
}

func (f *FrameRefSegInfo) StateTypeName() string {
	return "pkg/sentry/fsutil.FrameRefSegInfo"
}

func (f *FrameRefSegInfo) StateFields() []string {
	return []string{
		"refs",
		"memCgID",
	}
}

func (f *FrameRefSegInfo) beforeSave() {}

// +checklocksignore
func (f *FrameRefSegInfo) StateSave(stateSinkObject state.Sink) {
	f.beforeSave()
	stateSinkObject.Save(0, &f.refs)
	stateSinkObject.Save(1, &f.memCgID)
}

func (f *FrameRefSegInfo) afterLoad(context.Context) {}

// +checklocksignore
func (f *FrameRefSegInfo) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &f.refs)
	stateSourceObject.Load(1, &f.memCgID)
}

func (f *HostFileMapper) StateTypeName() string {
	return "pkg/sentry/fsutil.HostFileMapper"
}

func (f *HostFileMapper) StateFields() []string {
	return []string{
		"refs",
	}
}

func (f *HostFileMapper) beforeSave() {}

// +checklocksignore
func (f *HostFileMapper) StateSave(stateSinkObject state.Sink) {
	f.beforeSave()
	stateSinkObject.Save(0, &f.refs)
}

// +checklocksignore
func (f *HostFileMapper) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &f.refs)
	stateSourceObject.AfterLoad(func() { f.afterLoad(ctx) })
}

func (s *mappingSet) StateTypeName() string {
	return "pkg/sentry/fsutil.mappingSet"
}

func (s *mappingSet) StateFields() []string {
	return []string{
		"root",
	}
}

func (s *mappingSet) beforeSave() {}

// +checklocksignore
func (s *mappingSet) StateSave(stateSinkObject state.Sink) {
	s.beforeSave()
	var rootValue []mappingFlatSegment
	rootValue = s.saveRoot()
	stateSinkObject.SaveValue(0, rootValue)
}

func (s *mappingSet) afterLoad(context.Context) {}

// +checklocksignore
func (s *mappingSet) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.LoadValue(0, new([]mappingFlatSegment), func(y any) { s.loadRoot(ctx, y.([]mappingFlatSegment)) })
}

func (n *mappingnode) StateTypeName() string {
	return "pkg/sentry/fsutil.mappingnode"
}

func (n *mappingnode) StateFields() []string {
	return []string{
		"nrSegments",
		"parent",
		"parentIndex",
		"hasChildren",
		"maxGap",
		"keys",
		"values",
		"children",
	}
}

func (n *mappingnode) beforeSave() {}

// +checklocksignore
func (n *mappingnode) StateSave(stateSinkObject state.Sink) {
	n.beforeSave()
	stateSinkObject.Save(0, &n.nrSegments)
	stateSinkObject.Save(1, &n.parent)
	stateSinkObject.Save(2, &n.parentIndex)
	stateSinkObject.Save(3, &n.hasChildren)
	stateSinkObject.Save(4, &n.maxGap)
	stateSinkObject.Save(5, &n.keys)
	stateSinkObject.Save(6, &n.values)
	stateSinkObject.Save(7, &n.children)
}

func (n *mappingnode) afterLoad(context.Context) {}

// +checklocksignore
func (n *mappingnode) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &n.nrSegments)
	stateSourceObject.Load(1, &n.parent)
	stateSourceObject.Load(2, &n.parentIndex)
	stateSourceObject.Load(3, &n.hasChildren)
	stateSourceObject.Load(4, &n.maxGap)
	stateSourceObject.Load(5, &n.keys)
	stateSourceObject.Load(6, &n.values)
	stateSourceObject.Load(7, &n.children)
}

func (m *mappingFlatSegment) StateTypeName() string {
	return "pkg/sentry/fsutil.mappingFlatSegment"
}

func (m *mappingFlatSegment) StateFields() []string {
	return []string{
		"Start",
		"End",
		"Value",
	}
}

func (m *mappingFlatSegment) beforeSave() {}

// +checklocksignore
func (m *mappingFlatSegment) StateSave(stateSinkObject state.Sink) {
	m.beforeSave()
	stateSinkObject.Save(0, &m.Start)
	stateSinkObject.Save(1, &m.End)
	stateSinkObject.Save(2, &m.Value)
}

func (m *mappingFlatSegment) afterLoad(context.Context) {}

// +checklocksignore
func (m *mappingFlatSegment) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &m.Start)
	stateSourceObject.Load(1, &m.End)
	stateSourceObject.Load(2, &m.Value)
}

func (f *PreciseHostFileMapper) StateTypeName() string {
	return "pkg/sentry/fsutil.PreciseHostFileMapper"
}

func (f *PreciseHostFileMapper) StateFields() []string {
	return []string{
		"refs",
		"mappings",
	}
}

func (f *PreciseHostFileMapper) beforeSave() {}

// +checklocksignore
func (f *PreciseHostFileMapper) StateSave(stateSinkObject state.Sink) {
	f.beforeSave()
	stateSinkObject.Save(0, &f.refs)
	stateSinkObject.Save(1, &f.mappings)
}

func (f *PreciseHostFileMapper) afterLoad(context.Context) {}

// +checklocksignore
func (f *PreciseHostFileMapper) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &f.refs)
	stateSourceObject.Load(1, &f.mappings)
}

func (s *refsSet) StateTypeName() string {
	return "pkg/sentry/fsutil.refsSet"
}

func (s *refsSet) StateFields() []string {
	return []string{
		"root",
	}
}

func (s *refsSet) beforeSave() {}

// +checklocksignore
func (s *refsSet) StateSave(stateSinkObject state.Sink) {
	s.beforeSave()
	var rootValue []refsFlatSegment
	rootValue = s.saveRoot()
	stateSinkObject.SaveValue(0, rootValue)
}

func (s *refsSet) afterLoad(context.Context) {}

// +checklocksignore
func (s *refsSet) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.LoadValue(0, new([]refsFlatSegment), func(y any) { s.loadRoot(ctx, y.([]refsFlatSegment)) })
}

func (n *refsnode) StateTypeName() string {
	return "pkg/sentry/fsutil.refsnode"
}

func (n *refsnode) StateFields() []string {
	return []string{
		"nrSegments",
		"parent",
		"parentIndex",
		"hasChildren",
		"maxGap",
		"keys",
		"values",
		"children",
	}
}

func (n *refsnode) beforeSave() {}

// +checklocksignore
func (n *refsnode) StateSave(stateSinkObject state.Sink) {
	n.beforeSave()
	stateSinkObject.Save(0, &n.nrSegments)
	stateSinkObject.Save(1, &n.parent)
	stateSinkObject.Save(2, &n.parentIndex)
	stateSinkObject.Save(3, &n.hasChildren)
	stateSinkObject.Save(4, &n.maxGap)
	stateSinkObject.Save(5, &n.keys)
	stateSinkObject.Save(6, &n.values)
	stateSinkObject.Save(7, &n.children)
}

func (n *refsnode) afterLoad(context.Context) {}

// +checklocksignore
func (n *refsnode) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &n.nrSegments)
	stateSourceObject.Load(1, &n.parent)
	stateSourceObject.Load(2, &n.parentIndex)
	stateSourceObject.Load(3, &n.hasChildren)
	stateSourceObject.Load(4, &n.maxGap)
	stateSourceObject.Load(5, &n.keys)
	stateSourceObject.Load(6, &n.values)
	stateSourceObject.Load(7, &n.children)
}

func (r *refsFlatSegment) StateTypeName() string {
	return "pkg/sentry/fsutil.refsFlatSegment"
}

func (r *refsFlatSegment) StateFields() []string {
	return []string{
		"Start",
		"End",
		"Value",
	}
}

func (r *refsFlatSegment) beforeSave() {}

// +checklocksignore
func (r *refsFlatSegment) StateSave(stateSinkObject state.Sink) {
	r.beforeSave()
	stateSinkObject.Save(0, &r.Start)
	stateSinkObject.Save(1, &r.End)
	stateSinkObject.Save(2, &r.Value)
}

func (r *refsFlatSegment) afterLoad(context.Context) {}

// +checklocksignore
func (r *refsFlatSegment) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &r.Start)
	stateSourceObject.Load(1, &r.End)
	stateSourceObject.Load(2, &r.Value)
}

func init() {
	state.Register((*DirtyInfo)(nil))
	state.Register((*FrameRefSegInfo)(nil))
	state.Register((*HostFileMapper)(nil))
	state.Register((*mappingSet)(nil))
	state.Register((*mappingnode)(nil))
	state.Register((*mappingFlatSegment)(nil))
	state.Register((*PreciseHostFileMapper)(nil))
	state.Register((*refsSet)(nil))
	state.Register((*refsnode)(nil))
	state.Register((*refsFlatSegment)(nil))
}