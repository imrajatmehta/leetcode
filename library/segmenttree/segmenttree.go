package segmenttree_sum

type SegmentTreeSum struct {
	left, right     *SegmentTreeSum
	sum, start, end int
}

func Build(s, e int) *SegmentTreeSum {
	if s > e {
		return nil
	}
	par := &SegmentTreeSum{start: s, end: e}
	if s == e {
		return par
	}
	mid := s + (e-s)/2
	par.left = Build(s, mid)
	par.right = Build(mid+1, e)
	return par
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
		st.left.Update(val, place)
	} else {
		st.right.Update(val, place)
	}
	st.sum = st.left.sum + st.right.sum
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
