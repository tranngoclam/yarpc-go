// Copyright (c) 2018 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package yarpctest

import yarpc "go.uber.org/yarpc/v2"

// FakePeer is a fake peer with an identifier.
type FakePeer struct {
	id yarpc.Identifier
	// subscribers needs to be modified under lock in FakeTransport
	subscribers []yarpc.Subscriber
	status      yarpc.Status
}

// Identifier returns the fake peer identifier.
func (p *FakePeer) Identifier() string {
	return p.id.Identifier()
}

// Status returns the fake peer status.
func (p *FakePeer) Status() yarpc.Status {
	return p.status
}

// StartRequest increments pending request count.
func (p *FakePeer) StartRequest() {
	p.status.PendingRequestCount++
}

// EndRequest decrements pending request count.
func (p *FakePeer) EndRequest() {
	p.status.PendingRequestCount--
}

func (p *FakePeer) simulateConnect() {
	p.status.ConnectionStatus = yarpc.Available
	p.broadcast()
}

func (p *FakePeer) simulateDisconnect() {
	p.status.ConnectionStatus = yarpc.Unavailable
	p.broadcast()
}

func (p *FakePeer) broadcast() {
	for _, sub := range p.subscribers {
		sub.NotifyStatusChanged(p.id)
	}
}