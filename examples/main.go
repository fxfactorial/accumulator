package main

import (
	"bytes"
	"encoding/binary"
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

func (utxo set) hashToPrimeBlake2b(record utxo) uint64 {
	var counter uint64 = 0
	payload, _ := record.id.MarshalBinary()

	utxoBytes := bytes.NewBuffer(payload)
	binary.PutUvarint(utxoBytes.Bytes(), record.userID)

	// Keep going till you find it
	for {
		payload := bytes.NewBuffer(utxoBytes.Bytes())
		binary.PutUvarint(payload.Bytes(), counter)
		utxoHash := h.New(h.Blake2B, payload.Bytes())
		writeOut := bytes.Buffer{}
		utxoHash.Write(writeOut.Bytes())
		// Make candidate prime odd
		binary.PutUvarint(writeOut.Bytes()[:1], 1)
		candidatePrime := new(big.Int).SetBytes(writeOut.Bytes())
		if primality.IsProbPrime(candidatePrime) {
			return candidatePrime.Uint64()
		}
		counter++
	}

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
