// automatically generated by stateify.

package arp

import (
	"context"

	"gvisor.dev/gvisor/pkg/state"
)

func (e *endpoint) StateTypeName() string {
	return "pkg/tcpip/network/arp.endpoint"
}

func (e *endpoint) StateFields() []string {
	return []string{
		"protocol",
		"enabled",
		"nic",
		"stats",
		"dad",
	}
}

func (e *endpoint) beforeSave() {}

// +checklocksignore
func (e *endpoint) StateSave(stateSinkObject state.Sink) {
	e.beforeSave()
	stateSinkObject.Save(0, &e.protocol)
	stateSinkObject.Save(1, &e.enabled)
	stateSinkObject.Save(2, &e.nic)
	stateSinkObject.Save(3, &e.stats)
	stateSinkObject.Save(4, &e.dad)
}

func (e *endpoint) afterLoad(context.Context) {}

// +checklocksignore
func (e *endpoint) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &e.protocol)
	stateSourceObject.Load(1, &e.enabled)
	stateSourceObject.Load(2, &e.nic)
	stateSourceObject.Load(3, &e.stats)
	stateSourceObject.Load(4, &e.dad)
}

func (p *protocol) StateTypeName() string {
	return "pkg/tcpip/network/arp.protocol"
}

func (p *protocol) StateFields() []string {
	return []string{
		"stack",
		"options",
	}
}

func (p *protocol) beforeSave() {}

// +checklocksignore
func (p *protocol) StateSave(stateSinkObject state.Sink) {
	p.beforeSave()
	stateSinkObject.Save(0, &p.stack)
	stateSinkObject.Save(1, &p.options)
}

func (p *protocol) afterLoad(context.Context) {}

// +checklocksignore
func (p *protocol) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &p.stack)
	stateSourceObject.Load(1, &p.options)
}

func (o *Options) StateTypeName() string {
	return "pkg/tcpip/network/arp.Options"
}

func (o *Options) StateFields() []string {
	return []string{
		"DADConfigs",
	}
}

func (o *Options) beforeSave() {}

// +checklocksignore
func (o *Options) StateSave(stateSinkObject state.Sink) {
	o.beforeSave()
	stateSinkObject.Save(0, &o.DADConfigs)
}

func (o *Options) afterLoad(context.Context) {}

// +checklocksignore
func (o *Options) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &o.DADConfigs)
}

func (s *Stats) StateTypeName() string {
	return "pkg/tcpip/network/arp.Stats"
}

func (s *Stats) StateFields() []string {
	return []string{
		"ARP",
	}
}

func (s *Stats) beforeSave() {}

// +checklocksignore
func (s *Stats) StateSave(stateSinkObject state.Sink) {
	s.beforeSave()
	stateSinkObject.Save(0, &s.ARP)
}

func (s *Stats) afterLoad(context.Context) {}

// +checklocksignore
func (s *Stats) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &s.ARP)
}

func (s *sharedStats) StateTypeName() string {
	return "pkg/tcpip/network/arp.sharedStats"
}

func (s *sharedStats) StateFields() []string {
	return []string{
		"localStats",
		"arp",
	}
}

func (s *sharedStats) beforeSave() {}

// +checklocksignore
func (s *sharedStats) StateSave(stateSinkObject state.Sink) {
	s.beforeSave()
	stateSinkObject.Save(0, &s.localStats)
	stateSinkObject.Save(1, &s.arp)
}

func (s *sharedStats) afterLoad(context.Context) {}

// +checklocksignore
func (s *sharedStats) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &s.localStats)
	stateSourceObject.Load(1, &s.arp)
}

func (m *multiCounterARPStats) StateTypeName() string {
	return "pkg/tcpip/network/arp.multiCounterARPStats"
}

func (m *multiCounterARPStats) StateFields() []string {
	return []string{
		"packetsReceived",
		"disabledPacketsReceived",
		"malformedPacketsReceived",
		"requestsReceived",
		"requestsReceivedUnknownTargetAddress",
		"outgoingRequestInterfaceHasNoLocalAddressErrors",
		"outgoingRequestBadLocalAddressErrors",
		"outgoingRequestsDropped",
		"outgoingRequestsSent",
		"repliesReceived",
		"outgoingRepliesDropped",
		"outgoingRepliesSent",
	}
}

func (m *multiCounterARPStats) beforeSave() {}

// +checklocksignore
func (m *multiCounterARPStats) StateSave(stateSinkObject state.Sink) {
	m.beforeSave()
	stateSinkObject.Save(0, &m.packetsReceived)
	stateSinkObject.Save(1, &m.disabledPacketsReceived)
	stateSinkObject.Save(2, &m.malformedPacketsReceived)
	stateSinkObject.Save(3, &m.requestsReceived)
	stateSinkObject.Save(4, &m.requestsReceivedUnknownTargetAddress)
	stateSinkObject.Save(5, &m.outgoingRequestInterfaceHasNoLocalAddressErrors)
	stateSinkObject.Save(6, &m.outgoingRequestBadLocalAddressErrors)
	stateSinkObject.Save(7, &m.outgoingRequestsDropped)
	stateSinkObject.Save(8, &m.outgoingRequestsSent)
	stateSinkObject.Save(9, &m.repliesReceived)
	stateSinkObject.Save(10, &m.outgoingRepliesDropped)
	stateSinkObject.Save(11, &m.outgoingRepliesSent)
}

func (m *multiCounterARPStats) afterLoad(context.Context) {}

// +checklocksignore
func (m *multiCounterARPStats) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &m.packetsReceived)
	stateSourceObject.Load(1, &m.disabledPacketsReceived)
	stateSourceObject.Load(2, &m.malformedPacketsReceived)
	stateSourceObject.Load(3, &m.requestsReceived)
	stateSourceObject.Load(4, &m.requestsReceivedUnknownTargetAddress)
	stateSourceObject.Load(5, &m.outgoingRequestInterfaceHasNoLocalAddressErrors)
	stateSourceObject.Load(6, &m.outgoingRequestBadLocalAddressErrors)
	stateSourceObject.Load(7, &m.outgoingRequestsDropped)
	stateSourceObject.Load(8, &m.outgoingRequestsSent)
	stateSourceObject.Load(9, &m.repliesReceived)
	stateSourceObject.Load(10, &m.outgoingRepliesDropped)
	stateSourceObject.Load(11, &m.outgoingRepliesSent)
}

func init() {
	state.Register((*endpoint)(nil))
	state.Register((*protocol)(nil))
	state.Register((*Options)(nil))
	state.Register((*Stats)(nil))
	state.Register((*sharedStats)(nil))
	state.Register((*multiCounterARPStats)(nil))
}