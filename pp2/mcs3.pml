// MCS lock implementation in Promela

#define N 3

typedef LNode {
	int next;
	bit waiting
};

LNode queue[N+1];
int lock = 0, index = 1, ncrit = 0, counter = 0;

active[N] proctype MCS(){

	int pred;
	int myNode;

	atomic {
			myNode = index;
			index++;
		}

	// Lock
	again:	

		queue[myNode].next = 0;

		// swap(tail, pred)
		atomic {
			pred = lock;
			lock = myNode;
		}

		if
		::pred ->
			queue[myNode].waiting = 1;
			queue[pred].next = myNode;
			queue[myNode].waiting == 0		// wait for lock
		::else ->
			skip;
		fi;

		// Critical Section
		ncrit++;
		assert(ncrit == 1);
		counter++;
		ncrit--;

		// Unlock
		int ret;
		int succ = queue[myNode].next;

		if
		:: (succ == 0) ->
			// Compare and swap
			atomic {
				ret = lock;
				lock = (lock == myNode -> 0 : lock);
			}

			if
			:: ret == myNode -> goto again;
			:: else -> skip;
			fi;

			succ = queue[myNode].next;

			do
			:: succ == 0 -> succ = queue[myNode].next;
			:: else -> break;
			od;
	::	else ->
			skip;
	fi;
	queue[succ].waiting = 0;

	if 
	:: counter > N -> goto done;
	fi;

	goto again;

	done:
		printf("%d\n", counter);

}
