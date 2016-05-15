/* 6 Frogs Puzzle
Reference: http://www.weizmann.ac.il/sci-tea/benari/keynote/invisible-slides.pdf
*/

#define STONES 7

#define success (\
	(stones[0]==female) && (stones[1]==female) && \
	(stones[2]==female) && (stones[4]==male) && \
	(stones[5]==male)   && (stones[6]==male))

mtype = { none, male, female }
mtype stones[STONES];

// Specify correctness
ltl { []!success }

// Male Process
proctype maleFrog(byte at) {
end:do
	:: 	atomic {
			(at < STONES-1) && 
			(stones[at+1] == none) -> 
			stones[at] = none; 
			stones[at+1] = male;
			at = at + 1;
		}
	:: atomic {
			(at < STONES-2) && 
	   		(stones[at+1] != none) && 
			(stones[at+2] == none) -> 
			stones[at] = none; 
			stones[at+2] = male;
			at = at + 2;
		}
	od
}

// Female Process
proctype femaleFrog(byte at) {
end:do
	:: atomic {
			(at > 0) && 
			(stones[at-1] == none) -> 
			stones[at] = none; 
			stones[at-1] = female;
			at = at - 1;
		}
	:: atomic {
			(at > 1) && 
	   		(stones[at-1] != none) && 
			(stones[at-2] == none) -> 
			stones[at] = none; 
			stones[at-2] = female;
			at = at - 2;
		}
	od
}

// Main
init {
	atomic {
		stones[STONES/2] = none;
		byte I = 0;
        do
        :: I == STONES/2 -> break;
   		:: else -> 
             stones[I] = male;
             run maleFrog(I);
    		 stones[STONES-I-1] = female;
			 run femaleFrog(STONES-I-1);
             I++
        od
	}
}
