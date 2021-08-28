package queue

type Data interface{}

type Queue struct {
	size uint
	data []Data
}

func NewQueue() *Queue {
	return &Queue{}
}

func (q *Queue) Size() uint {
	return q.size
}

func (q *Queue) Enqueue(d Data) {
	if len(q.data) == 0 {
		q.data = append(q.data, d)
	} else {
		q.data = append(q.data[:1], q.data[0:]...)
		q.data[0] = d
	}

	q.size++
}

func (q *Queue) Dequeue() Data {
	result := q.data[q.size-1]

	q.data = q.data[:q.size-1]
	q.size--

	return result
}
