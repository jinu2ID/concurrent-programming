#define N  3

//
// need to simulate a linked list of QNodes
// Promela does not support pointers - so use array indices in lieu of pointers
// qnode is array of N+1 QNodes
// the QNode belonging to process _pid in qnode[_pid+1]
// this allows index 0 to represent a NULL pointer
//
typedef QNode {
    byte next;
    byte waiting;
};

QNode qnode[N+1];                                       // see comment above
byte lock;                                              // index into qnode

byte ncs;                                               // number of processes in critical section
byte cs[N];                                             // cs[n] true if process n in critcal section

int counter;

active[N] proctype p()
{
    byte pred;
    byte succ;
    byte r;
    
    again:
    
			if
			:: counter >= N ->
					goto done;
			:: else -> skip;
			fi;

        //
        // aquire lock
        //
        qnode[_pid + 1].next = 0;                       // zero next field of QNode belonging to process
        
        atomic {
            pred = lock;                                // simulate an ...
            lock = _pid + 1;                            // InterlockedExchange
        }

        if
        :: pred ->                                      // if pred...
            qnode[_pid + 1].waiting = 1;                //    add QNode belonging to process...
            qnode[pred].next = _pid + 1;                //    to tail of Q
            qnode[_pid + 1].waiting == 0;               //    wait until true (ie. have lock)
        :: else ->                                      // otherwise...
            skip;                                       // have...
        fi;                                             // lock

        //
        // critical section
        //
        ncs++;
        assert(ncs == 1);
		  counter++;
        cs[_pid] = 1;
        cs[_pid] = 0;
        ncs--;

        //
        // release lock
        //
        succ = qnode[_pid + 1].next;                    // get next process waiting for lock
        if
        :: succ == 0 ->
            atomic {                                    // simulate
                r = lock;                               // a
                lock = (lock == _pid + 1 -> 0 : lock);  // CAS
            }                                           //
            if
            :: r == _pid + 1 ->
                goto again;
            :: else ->
                skip;   
            fi;       
            succ = qnode[_pid + 1].next;
            do
            :: succ == 0 ->
                succ = qnode[_pid + 1].next;
            :: else ->
                break;
            od;
        :: else ->                                      // otherwise...
            skip;                                       // exit
        fi;
        qnode[succ].waiting = 0;                        // pass lock to next process
    
        goto again;

	done:
		printf("%d\n", counter);
        
}

