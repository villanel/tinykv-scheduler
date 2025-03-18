// Copyright 2015 The etcd Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package raft

import (
	"log"

	pb "github.com/villanel/tinykv-scheduler/proto/pkg/eraftpb"
)

// RaftLog manage the log entries, its struct look like:
//
//	snapshot/first.....applied....committed....stabled.....last
//	--------|------------------------------------------------|
//	                          log entries
//
// for simplify the RaftLog implement should manage all log entries
// that not truncated
type RaftLog struct {
	// storage contains all stable entries since the last snapshot.
	storage Storage

	firstIdx uint64
	// committed is the highest log position that is known to be in
	// stable storage on a quorum of nodes.
	committed uint64

	// applied is the highest log position that the application has
	// been instructed to apply to its state machine.
	// Invariant: applied <= committed
	applied uint64

	// log entries with index <= stabled are persisted to storage.
	// It is used to record the logs that are not persisted by storage yet.
	// Everytime handling `Ready`, the unstabled logs will be included.
	stabled uint64

	// all entries that have not yet compact.
	entries []pb.Entry

	// the incoming unstable snapshot, if any.
	// (Used in 2C)
	pendingSnapshot *pb.Snapshot

	// Your Data Here (2A).
	offset uint64
}

// newLog returns log using the given storage. It recovers the log
// to the state that it just commits and applies the latest snapshot.
func newLog(storage Storage) *RaftLog {
	if storage == nil {
		log.Panic("storage must not be nil")
	}
	log := &RaftLog{
		storage: storage,
	}
	firstIndex, err := storage.FirstIndex()
	if err != nil {
		panic(err) // TODO(bdarnell)
	}
	lastIndex, err := storage.LastIndex()
	if err != nil {
		panic(err) // TODO(bdarnell)
	}
	log.committed = firstIndex - 1
	log.applied = firstIndex - 1
	log.stabled = lastIndex
	log.firstIdx = firstIndex
	//storage not empty
	if lastIndex >= firstIndex {
		ents, err := storage.Entries(firstIndex, lastIndex+1)
		if err != nil {
			panic(err)
		}
		log.entries = make([]pb.Entry, len(ents))
		copy(log.entries, ents)
	}

	return log
}

// We need to compact the log entries in some point of time like
// storage compact stabled log entries prevent the log entries
// grow unlimitedly in memory
func (l *RaftLog) mayeCompact() {
	// Your Code Here (2C).
	//
	//first, _ := l.storage.firstIdx()
	//if first > l.firstIdx {
	//	if len(l.entries) > 0 {
	//		entries := l.entries[l.toSliceIndex(first):]
	//		l.entries = make([]pb.Entry, len(entries))
	//		copy(l.entries, entries)
	//	}
	//	l.firstIdx = first
	//}
	//

	if len(l.entries) == 0 {
		return
	}
	sfirst, _ := l.storage.FirstIndex()
	if sfirst > l.firstIdx {
		if len(l.entries) > 0 {
			l.entries = l.entries[sfirst-l.firstIdx:]
		}
		l.firstIdx = sfirst
	}

}

// unstableEntries return all the unstable entries
func (l *RaftLog) unstableEntries() []pb.Entry {
	//if l.stabled>=l.offset{
	//	l.offset = l.stabled + 1
	//	return  l.entries[l.stabled+1-l.offset:]
	//}
	// Your Code Here (2A).
	if len(l.entries) == 0 {
		return nil
	}

	//entries := l.entries[l.offset-1:]
	entries := l.entries[l.stabled-l.entries[0].Index+1:]

	return entries
}

// nextEnts returns all the committed but not applied entries
func (l *RaftLog) nextEnts() (ents []pb.Entry) {
	if len(l.entries) > 0 {
		return l.entries[l.applied-l.entries[0].Index+1 : l.committed-l.entries[0].Index+1]
	}
	return nil

	//off := max(l.applied+1, l.firstIdx())
	//if l.committed+1 > off {
	//	hi := l.committed + 1
	//	lo := off
	//	var ents []pb.Entry
	//	//if lo < l.offset {
	//	//	storedEnts, _ := l.storage.Entries(lo, min(hi, l.offset))
	//	//	for _, ent := range storedEnts {
	//	//		ents = append(ents, ent)
	//	//	}
	//	//}
	//	//if hi >l.offset{
	//	//	entries := l.entries[max(lo, l.offset)-l.offset:hi-l.offset]
	//	//	for _, entry := range entries {
	//	//		ents = append(ents, entry)
	//	//	}
	//	//for _, entry := range l.entries[lo-1 : hi-1] {
	//	for _, entry := range l.entries[lo-l.entries[0].Index : hi-l.entries[0].Index] {
	//		ents = append(ents, entry)
	//
	//	}
	//	return ents
	//}
}

// LastIndex return the last index of the log entries
func (l *RaftLog) LastIndex() uint64 {
	//lastIndex, _ := l.storage.LastIndex()
	//if lastIndex >=l.offset{
	//	l.offset=lastIndex+1
	//}
	var idx uint64
	// Your Code Here (2A).
	if !IsEmptySnap(l.pendingSnapshot) {
		idx = max(idx, l.pendingSnapshot.Metadata.Index)
	}
	if len := len(l.entries); len != 0 {
		//return l.offset + uint64(len) - 1
		idx = max(l.entries[len-1].Index, idx)
		return idx
	}
	index, err := l.storage.LastIndex()
	if err != nil {
		panic(err)
	}
	// Your Code Here (2A).
	return max(index, idx)
}
func (l *RaftLog) firstIndex() uint64 {
	index, err := l.storage.FirstIndex()
	if err != nil {
		panic(err)
	}
	// Your Code Here (2A).
	return index
}

// Term return the term of the entry in the given index
func (l *RaftLog) Term(i uint64) (uint64, error) {

	//dummyIndex := l.firstIndex() - 1
	//if i < dummyIndex ||
	if i > l.LastIndex() {
		return 0, nil
	}
	if t, ok := l.unstableTerm(i); ok {
		return t, nil
	}

	t, err := l.storage.Term(i)
	if err == nil {
		return t, nil
	}
	snap := l.pendingSnapshot
	if err == ErrUnavailable && !IsEmptySnap(snap) {
		if i == snap.Metadata.Index {
			t = snap.Metadata.Term
			err = nil
		} else if i < snap.Metadata.Index {
			err = ErrCompacted
		}
	}
	return t, err
}

// get entries
func (l *RaftLog) entry(lo uint64) ([]*pb.Entry, error) {
	if lo > l.LastIndex() {
		return nil, nil
	}
	var ents []*pb.Entry

	entries := l.entries[lo-l.entries[0].Index:]
	for i := 0; i < len(entries); i++ {
		ents = append(ents, &entries[i])
	}
	return ents, nil
}

func (l *RaftLog) unstableTerm(i uint64) (uint64, bool) {
	if len(l.entries) > 0 && i >= l.entries[0].Index {
		if i-l.entries[0].Index >= uint64(len(l.entries)) {
			return 0, false
		}

		return l.entries[i-l.entries[0].Index].Term, true
	}
	return 0, false
}

// 返回日志不匹配的index
func (l *RaftLog) findConflictByTerm(index uint64, term uint64) uint64 {
	if li := l.LastIndex(); index > li {
		// NB: such calls should not exist, but since there is a straightfoward
		// way to recover, do it.
		//
		// It is tempting to also check something about the first index, but
		// there is odd behavior with peers that have no log, in which case
		// lastIndex will return zero and firstIndex will return one, which
		// leads to calls with an index of zero into this method.
		return index
	}
	for {
		logTerm, err := l.Term(index)
		if logTerm <= term || err != nil {
			break
		}
		index--
	}
	return index
}

func (l *RaftLog) truncateAndAppend(ents []pb.Entry) {
	after := ents[0].Index
	switch {
	case uint64(len(l.entries))+1 == after:
		l.entries = append(l.entries, ents...)
	case after <= l.stabled:
		l.stabled = after
		if len(l.entries) != 0 {
			l.entries = l.entries[:l.stabled-l.entries[0].Index]
		}
		l.entries = append(l.entries, ents...)
	default:
		if after > l.LastIndex() {
			l.entries = append(l.entries, ents...)
		} else {
			l.entries = append([]pb.Entry{}, l.entries[0:after-l.entries[0].Index]...)
			l.entries = append(l.entries, ents...)
		}
	}
}

// 比较日志，能否投票
func (l *RaftLog) isUpToDate(lasti, term uint64) bool {
	lastTerm, err := l.Term(l.LastIndex())
	if err != nil {
		panic(err)
	}

	return term > lastTerm || (term == lastTerm && lasti >= l.LastIndex())
}
