package segmenttree_sum

/*
This Segment Tree implementation is memory optimised, by only creating a node which is provided in update and contains only parents. not sibling which are not requred i.e range which are not required.
Update() is called with {1,3,9}  say arr so it will create tree which is shown below,
Not all the nodes like child (2,2),(4,4),(5,5),(6,6),(7,7),(8,8),(10,10), and there parents, but parent will be there for given arr.

				(1-10)
				/    \
			(1-5)    (6-10)
			/   \    /   \
		(1,3)             (9,10)
		/   \             /    \
	(1,2)   (3,3)       (9,9)
	/   \

(1,1)
*/
type SegmentTreeSum struct {
	left, right     *SegmentTreeSum
	sum, start, end int
}

func Build(s, e int) *SegmentTreeSum {
	if s > e {
		return nil
	}
	return &SegmentTreeSum{start: s, end: e}
}

func (st *SegmentTreeSum) Update(val, place int) {
	if st == nil {
		return
	}
	if st.start == place && st.end == place {
		st.sum += val
		return
	}
	mid := st.start + (st.end-st.start)/2
	if place <= mid {
		if st.left == nil {
			st.left = &SegmentTreeSum{start: st.start, end: mid}
		}
		st.left.Update(val, place)
	} else {
		if st.right == nil {
			st.right = &SegmentTreeSum{start: mid + 1, end: st.end}
		}
		st.right.Update(val, place)
	}
	left, right := 0, 0
	if st.left != nil {
		left = st.left.sum
	}
	if st.right != nil {
		right = st.right.sum
	}
	st.sum = right + left

}
func (st *SegmentTreeSum) GetSum(s, e int) int {
	if s > e || st == nil {
		return 0
	}
	if st.start == s && st.end == e {
		return st.sum
	}
	mid := st.start + (st.end-st.start)/2
	if e <= mid {
		return st.left.GetSum(s, e)
	} else if mid < s {
		return st.right.GetSum(s, e)
	}
	return st.left.GetSum(s, mid) + st.right.GetSum(mid+1, e)
}
