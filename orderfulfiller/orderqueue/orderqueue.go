// This package defines a queue for order which dedupes orders that are in the queue. It also enforces a requeue delay
// after an order is popped from the queue wherein the same order cannot be requeued.

package orderqueue

import (
	"github.com/skip-mev/go-fast-solver/db/gen/db"
	"golang.org/x/net/context"
	"sync"
	"time"
)

const (
	cleanupInterval = 1 * time.Minute
)

// OrderQueue contains pending orders and fulfills them FIFO
type OrderQueue struct {
	orderRequeueTime map[int64]time.Time
	ordersInQueue    map[int64]bool
	orderChan        chan db.Order
	requeueDelay     time.Duration
	cleanupTicker    *time.Ticker
	lock             sync.Mutex
	stopCleanup      chan struct{}
}

func NewOrderQueue(ctx context.Context, requeueDelay time.Duration, queueCapacity int) *OrderQueue {
	orderQueue := &OrderQueue{
		orderRequeueTime: make(map[int64]time.Time),
		ordersInQueue:    make(map[int64]bool),
		orderChan:        make(chan db.Order, queueCapacity),
		requeueDelay:     requeueDelay,
		cleanupTicker:    time.NewTicker(cleanupInterval),
		stopCleanup:      make(chan struct{}),
	}
	go orderQueue.startCleanup(ctx)
	return orderQueue
}

func (d *OrderQueue) QueueOrder(order db.Order) bool {
	d.lock.Lock()
	defer d.lock.Unlock()
	if _, ok := d.ordersInQueue[order.ID]; ok {
		return false
	}
	if _, ok := d.orderRequeueTime[order.ID]; ok && time.Now().Before(d.orderRequeueTime[order.ID]) {
		return false
	}
	select {
	case d.orderChan <- order:
		d.ordersInQueue[order.ID] = true
		return true
	default:
		return false
	}
}

func (d *OrderQueue) PopOrder() <-chan db.Order {
	out := make(chan db.Order)
	go func() {
		defer close(out)
		order := <-d.orderChan
		out <- order
		d.lock.Lock()
		d.orderRequeueTime[order.ID] = time.Now().Add(d.requeueDelay)
		delete(d.ordersInQueue, order.ID)
		d.lock.Unlock()
	}()
	return out
}

func (d *OrderQueue) startCleanup(ctx context.Context) {
	for {
		select {
		case <-d.cleanupTicker.C:
			d.lock.Lock()
			for id, delay := range d.orderRequeueTime {
				if time.Now().After(delay) {
					delete(d.orderRequeueTime, id)
				}
			}
			d.lock.Unlock()
		case <-ctx.Done():
			d.cleanupTicker.Stop()
			return
		}
	}
}
