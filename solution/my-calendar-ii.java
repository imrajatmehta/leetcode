class MyCalendarTwo {
	TreeMap<Integer,Integer> map;
    public MyCalendarTwo() {
		map=new TreeMap();
    }
    
    public boolean book(int start, int end) {
        map.put(start,map.getOrDefault(start,0)+1);
		map.put(end,map.getOrDefault(end,0)-1);

		int count=0;
		for(int d: map.values()){
			count+=d;
			if (count>2){
				map.put(start,map.get(start)-1);
				map.put(end,map.get(end)+1);
				return false;
			}
		}
		return true;
		

    }
}

