// Ticket Lock

#define N 6 

int ticket =0, serving = 0, counter = 0, ncrit = 0;

active [N] proctype ticketLock(){
	
	int myTick;

	do
	::(counter >= N) -> break;
	:: else ->
		// Fetch-and-increment
		atomic {
			myTick = ticket;
			ticket = ticket + 1;
		}

		(myTick == serving); // await mytick = serving
		ncrit++;					// critical section
		assert(ncrit == 1);	
		counter++;			 		
		ncrit--;
		serving++;
	od;

}
