// MCS lock implementation in Promela

#define N 3

typedef LNode {
	int next;
	bit waiting
};

LNode queue[N+1];
int lock = 0, index = 1, ncrit = 0, counter = 0;

inline acquire_lock(_myNode){

	int pred;
	
	// swap(tail, pred)
	atomic {
		pred = lock;
		lock = _myNode;
	}

	if
	::pred ->
		queue[_myNode].waiting = 1;
		queue[pred].next = _myNode;
		queue[_myNode].waiting == 0		// wait for lock
	::else ->
		skip;
	fi;

} 

inline release_lock(_myNode){

	int ret;
	int succ = queue[_myNode].next;

	if
	:: (succ == 0) ->
		// Compare and swap
		atomic {
			ret = lock;
			lock = (lock == _myNode -> 0 : lock);
		}

		if
		:: ret == _myNode -> goto done;
		:: else -> skip;
		fi;

		succ = queue[_myNode].next;

		do
		:: succ == 0 -> succ = queue[_myNode].next;
		:: else -> break;
		od;
	else ->
		skip;
	fi;
	queue[succ].waiting = 0;

	done:
		skip;

}

active[N] proctype increment(){
	
	int myNode;
	
	atomic{
		queue[index].next = 0;
		myNode = index;
		index++;
	}
	
	do
	::(counter >= N) -> break
	:: else ->
		acquire_lock(myNode);		// Lock
		counter++;						// Critical Section
		release_lock(myNode);		// Unlock	
	od;

	printf("%d\n", counter);
}
