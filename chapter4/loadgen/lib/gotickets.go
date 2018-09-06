package lib

import (
	"errors"
	"fmt"
)

// GoTickets 表示Goroutine票池的接口。
type GoTickets interface {
	// 取走一张票。
	Take()
	// 归还一张票。
	Return()
	// 票池是否已被激活。
	Active() bool
	// 票的总数。
	Total() uint32
	// 剩余的票数。
	Remainder() uint32
}

// myGoTickets 表示Goroutine票池的实现。
type myGoTickets struct {
	total    uint32        // 票的总数。
	ticketCh chan struct{} // 票的容器。
	active   bool          // 票池是否已被激活。
}

// NewGoTickets 会新建一个Goroutine票池。
func NewGoTickets(total uint32) (GoTickets, error) {
	gt := myGoTickets{}
	if !gt.init(total) {
		errMsg := fmt.Sprintf("The goroutine ticket pool can NOT be initialized! (total=%d)\n", total)
		return nil, errors.New(errMsg)
	}
	return &gt, nil
}

// 初始化票池
func (gt *myGoTickets) init(total uint32) bool {
	if gt.active {
		return false
	}
	if total == 0 {
		return false
	}

	ch := make(chan struct{}, total)
	n := int(total)
	for i := 0; i < n; i++ {
		ch <- struct{}{}
	}

	gt.ticketCh = ch
	gt.total = total
	gt.active = true

	return true
}

// 取走一张票。
func (gt *myGoTickets) Take() {
	<-gt.ticketCh
}

// 归还一张票。
func (gt *myGoTickets) Return() {
	gt.ticketCh <- struct{}{}
}

// 票池是否已被激活。
func (gt *myGoTickets) Active() bool {
	return gt.active
}

// 票的总数。
func (gt *myGoTickets) Total() uint32 {
	return gt.total
}

// 剩余的票数。
func (gt *myGoTickets) Remainder() uint32 {
	return uint32(len(gt.ticketCh))
}
