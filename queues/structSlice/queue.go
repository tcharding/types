// Copyright ©2013 The gonum Authors. All rights reserved.

// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND
// ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
// WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
// DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
// FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
// DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
// SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
// CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
// OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
package queue

// Modification of Queue implementation found at
// https://github.com/gonum/graph/blob/master/internal/linear/linear.go

// Queue implements a FIFO queue.
type Queue struct {
	head int
	data []item
}

// Item: real implementation would have something useful here.
type item struct {
	data string
}

// Len returns the number of items in the queue.
func (q *Queue) Len() int { return len(q.data) - q.head }

// Enqueue adds the node n to the back of the queue.
func (q *Queue) Enqueue(n item) {
	if len(q.data) == cap(q.data) && q.head > 0 {
		ln := q.Len()
		copy(q.data, q.data[q.head:])
		q.head = 0
		q.data = append(q.data[:ln], n)
	} else {
		q.data = append(q.data, n)
	}
}

// Dequeue returns the item at the front of the queue and
// removes it from the queue.
func (q *Queue) Dequeue() item {
	if q.Len() == 0 {
		panic("queue: empty queue")
	}

	var n item
	n, q.data[q.head] = q.data[q.head], nil
	q.head++

	if q.Len() == 0 {
		q.head = 0
		q.data = q.data[:0]
	}

	return n
}

// Reset clears the queue for reuse.
func (q *Queue) Reset() {
	q.head = 0
	q.data = q.data[:0]
}
