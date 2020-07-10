package problem04

import (
	"math"
	"math/rand"
	"sync"
	"time"
)

type IDGenerator interface {
	GenerateID() int32
}

func NEWIDGenerator() *idGenerator {
	idGen := new(idGenerator)
	idGen.startTime, idGen.endTime = idGen.getNextIDRange()
	return idGen
}

type idGenerator struct {
	sync.Mutex
	startTime int
	endTime int
}

func (i *idGenerator) GenerateID() int32 {
	i.Lock()
	defer i.Unlock()
	for i.startTime >= i.endTime {
		i.startTime, i.endTime = i.getNextIDRange()
	}

	i.startTime++
	id := boundedPseudoEncrypt2(int32(i.startTime))

	return id
}

func (i *idGenerator) getNextIDRange() (start int, end int) {
	return 0, 100
}

var a, b, c int
func init() {
	rand.Seed(time.Now().UnixNano())
	a = 1000 + rand.Intn(500)
	b = 150000 + rand.Intn(1000)
	c = 700000 + rand.Intn(1000)
}

const maxValue = int32(99999999)
func boundedPseudoEncrypt2(value int32) int32 {
	for {
		value = pseudoEncrypt2(value)
		if value <= maxValue {
			return value
		}
	}
}

func pseudoEncrypt2(value int32) int32 {
	var l1, r1, l2, r2 int32
	l1 = (value >> 13) & (8192 - 1)
	r1 = value & (8192 - 1)
	for i:=0; i<3; i++ {
		/* round() is used to produce the same values as the
		   plpgsql implementation that does an SQL cast to INT */
		l2, r2 = r1, l1 ^ int32(math.Round((float64((a*int(r1) + b) % c) / float64(c)) * (8192 - 1)))
		l1, r1 = l2, r2
	}
	return (l1 << 13) + r1
}



