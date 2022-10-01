package queue

import "errors"

// Queue Fake redis
type Queue interface {
	Get(string) (string, error)
	Set(string, string)
	Pop(string) (string, error)
}

type FakeRedisQueue struct {
	queue map[string]string
}

func (q *FakeRedisQueue) Get(key string) (string, error) {
	if val, ok := q.queue[key]; ok {
		return val, nil
	}

	return "", errors.New("key doesn't exists")
}

func (q *FakeRedisQueue) Set(key string, value string) {
	q.queue[key] = value
}

func (q *FakeRedisQueue) Pop(key string) (string, error) {
	if val, ok := q.queue[key]; ok {
		delete(q.queue, key)
		return val, nil
	}

	return "", errors.New("key doesn't exists")
}

var queue Queue = &FakeRedisQueue{
	queue: map[string]string{},
}
