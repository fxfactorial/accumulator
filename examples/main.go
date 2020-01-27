package main

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"math/big"
	"time"

	"github.com/fxfactorial/accumulator"
	"github.com/fxfactorial/accumulator/group"
	h "github.com/fxfactorial/accumulator/hash"
	"github.com/fxfactorial/accumulator/hash/primality"
	"github.com/fxfactorial/accumulator/proof"
	"github.com/google/uuid"
)

const (
	totalNumMiners  = 5
	totalNumBridges = 5
	totalNumUsers   = 5
	blockTime       = time.Millisecond * 5000
)

type utxo struct {
	id     uuid.UUID
	userID uint64
}

type set []utxo

func toStringInt(b *big.Int) string {
	return hex.EncodeToString(b.Bytes())
}

func toStringBuffer(b *bytes.Buffer) string {
	return hex.EncodeToString(b.Bytes())
}

func (utxo set) hashToPrimeBlake2b(record utxo) uint64 {
	payload, _ := record.id.MarshalBinary()
	utxoBytes := bytes.NewBuffer(payload)
	binary.PutUvarint(utxoBytes.Bytes(), record.userID)
	jobChan := make(chan struct{}, 32)
	counterChan := make(chan uint64)
	foundChan := make(chan *big.Int)

	go func() {
		jobChan <- struct{}{}
		counterChan <- 0
	}()

	// Keep going till you find it
	for {
		_, ok := <-jobChan

		if !ok {
			break
		}

		go func() {
			counterNow := <-counterChan
			b := make([]byte, 8)
			binary.LittleEndian.PutUint64(b, counterNow)
			blakeHasher := h.New(h.Blake2B, nil)
			blakeHasher.Write(payload)
			blakeHasher.Write(b)
			dump := blakeHasher.Sum(nil)
			copy(dump[:1], []byte{1})
			candidate := new(big.Int).SetBytes(dump)

			defer func() {
				if primality.IsProbPrime(candidate) {
					close(jobChan)
					close(counterChan)
					foundChan <- candidate
				} else {
					jobChan <- struct{}{}
					counterChan <- counterNow + 1
				}
			}()

		}()
	}

	found := <-foundChan
	fmt.Println("found at-", found)
	return found.Uint64()
}

func (utxo set) primeHashProduct() uint64 {
	large := make([]uint64, len(utxo))
	var result uint64 = 1
	for i := range utxo {
		large[i] = utxo.hashToPrimeBlake2b(utxo[i])
	}
	for i := range large {
		result *= large[i]
	}
	return result
}

func (utxo set) Compute() (uint64, uint64) {
	product := utxo.primeHashProduct()
	fmt.Println("big it is", product)
	return 0, 0
}

type transaction struct {
	utxosCreated          []interface{}
	utxosSpentWithWitness []interface{}
}

type block struct {
	height       uint64
	transactions []transaction
	accNew       accumulator.Tracker
	proofAdded   proof.Membership
	proofDeleted proof.Membership
}

type handler func()

type broadCastQueue struct {
	listeners []handler
}

func newQueue() *broadCastQueue {
	return &broadCastQueue{[]handler{}}
}

func (q *broadCastQueue) addToQueue(h handler) {
	q.listeners = append(q.listeners, h)
}

func runSimulation() {
	blockSenderReceiverQueue, txnSenderReceiverQueue := newQueue(), newQueue()
	go func() {
		// handle one queue
	}()
	go func() {
		// handle the other queue
	}()
	fmt.Println("start simulation", blockSenderReceiverQueue, txnSenderReceiverQueue)
	userUtxos := make(set, totalNumUsers)
	for i := 0; i < totalNumUsers; i++ {
		id, _ := uuid.NewUUID()
		userUtxos[i] = utxo{id, uint64(i)}
	}

	initAccum := accumulator.New(group.RSA2048)
	initAccum.Add(userUtxos)

}

func main() {
	runSimulation()
	accumSet := []interface{}{}
	accum := accumulator.New(group.RSA2048)
	fmt.Println(accumSet, accum)
}
