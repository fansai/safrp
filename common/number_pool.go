package common

import (
	"sync/atomic"
)

type NumberPool struct {
	numberArr []uint64
	number uint64
	maxVal uint64
	add uint64
}

/**
 * 创建一个编号池
 * @param		maxVal, add uint64		最大编号, 每次增加值
 * @return		*NumberPool				编号池对象的指针
 * func NewNumberPool(maxVal, add uint64) *NumberPool;
 */
func NewNumberPool(maxVal, add uint64) *NumberPool {
	return &NumberPool{
		numberArr:make([]uint64, maxVal+1),
		number: 1,
		maxVal: maxVal,
		add:    add,
	}
}

/**
 * 从编号池中取出一个未使用的编号
 * @param		nil
 * @return		uint64, bool	编号, 是否可取
 * func (n *NumberPool)Get() (uint64, bool);
 */
func (n *NumberPool)Get() (uint64, bool) {
	num := 0
	for i := atomic.LoadUint64(&n.number);;i = atomic.AddUint64(&n.number, n.add) {
		atomic.CompareAndSwapUint64(&n.number, n.maxVal, 1)
		num++
		if num / int(n.maxVal) >= 3 {
			return 0, false
		}
		if i > n.maxVal {
			i = 1
		}
		if atomic.CompareAndSwapUint64(&n.numberArr[i], 0, 1) {
			return i, true
		}
	}
	return 0, false
}

/**
 * 将编号放入编号池中
 * @param		number int		编号
 * @return		nil
 * func (n *NumberPool)Put(number int);
 */
func (n *NumberPool)Put(number int) {
	atomic.CompareAndSwapUint64(&n.numberArr[number], 1, 0)
}
