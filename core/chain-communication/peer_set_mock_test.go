// Copyright 2019, Keychain Foundation Ltd.
// This file is part of the dipperin-core library.
//
// The dipperin-core library is free software: you can redistribute
// it and/or modify it under the terms of the GNU Lesser General Public License
// as published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.
//
// The dipperin-core library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

// Automatically generated by MockGen. DO NOT EDIT!
// Source: github.com/caiqingfeng/dipperin-core/core/chain-communication (interfaces: AbstractPeerSet)

package chain_communication

import (
	p2p "github.com/dipperin/dipperin-core/third-party/p2p"
	gomock "github.com/golang/mock/gomock"
)

// Mock of AbstractPeerSet interface
type MockAbstractPeerSet struct {
	ctrl     *gomock.Controller
	recorder *_MockAbstractPeerSetRecorder
}

// Recorder for MockAbstractPeerSet (not exported)
type _MockAbstractPeerSetRecorder struct {
	mock *MockAbstractPeerSet
}

func NewMockAbstractPeerSet(ctrl *gomock.Controller) *MockAbstractPeerSet {
	mock := &MockAbstractPeerSet{ctrl: ctrl}
	mock.recorder = &_MockAbstractPeerSetRecorder{mock}
	return mock
}

func (_m *MockAbstractPeerSet) EXPECT() *_MockAbstractPeerSetRecorder {
	return _m.recorder
}

func (_m *MockAbstractPeerSet) AddPeer(_param0 PmAbstractPeer) error {
	ret := _m.ctrl.Call(_m, "AddPeer", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockAbstractPeerSetRecorder) AddPeer(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "AddPeer", arg0)
}

func (_m *MockAbstractPeerSet) BestPeer() PmAbstractPeer {
	ret := _m.ctrl.Call(_m, "BestPeer")
	ret0, _ := ret[0].(PmAbstractPeer)
	return ret0
}

func (_mr *_MockAbstractPeerSetRecorder) BestPeer() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "BestPeer")
}

func (_m *MockAbstractPeerSet) Close() {
	_m.ctrl.Call(_m, "Close")
}

func (_mr *_MockAbstractPeerSetRecorder) Close() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Close")
}

func (_m *MockAbstractPeerSet) GetPeers() map[string]PmAbstractPeer {
	ret := _m.ctrl.Call(_m, "GetPeers")
	ret0, _ := ret[0].(map[string]PmAbstractPeer)
	return ret0
}

func (_mr *_MockAbstractPeerSetRecorder) GetPeers() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetPeers")
}

func (_m *MockAbstractPeerSet) GetPeersInfo() []*p2p.CsPeerInfo {
	ret := _m.ctrl.Call(_m, "GetPeersInfo")
	ret0, _ := ret[0].([]*p2p.CsPeerInfo)
	return ret0
}

func (_mr *_MockAbstractPeerSetRecorder) GetPeersInfo() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetPeersInfo")
}

func (_m *MockAbstractPeerSet) Len() int {
	ret := _m.ctrl.Call(_m, "Len")
	ret0, _ := ret[0].(int)
	return ret0
}

func (_mr *_MockAbstractPeerSetRecorder) Len() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Len")
}

func (_m *MockAbstractPeerSet) Peer(_param0 string) PmAbstractPeer {
	ret := _m.ctrl.Call(_m, "Peer", _param0)
	ret0, _ := ret[0].(PmAbstractPeer)
	return ret0
}

func (_mr *_MockAbstractPeerSetRecorder) Peer(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Peer", arg0)
}

func (_m *MockAbstractPeerSet) RemovePeer(_param0 string) error {
	ret := _m.ctrl.Call(_m, "RemovePeer", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockAbstractPeerSetRecorder) RemovePeer(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "RemovePeer", arg0)
}

func (_m *MockAbstractPeerSet) ReplacePeers(_param0 map[string]PmAbstractPeer) {
	_m.ctrl.Call(_m, "ReplacePeers", _param0)
}

func (_mr *_MockAbstractPeerSetRecorder) ReplacePeers(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ReplacePeers", arg0)
}