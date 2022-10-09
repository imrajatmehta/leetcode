class MyCalendarThree {
	TreeMap<Integer,Integer> map;
    public MyCalendarThree() {
		map=new TreeMap();
    }
    
    public int book(int start, int end) {
        map.put(start,map.getOrDefault(start,0)+1);
		map.put(end,map.getOrDefault(end,0)-1);

		int count=0;
        int max=0;
		for(int d: map.values()){
			count+=d;
			if (count>max){
                max=count;
            }
		}
		return max;
		

    }
}

