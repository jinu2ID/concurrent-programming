#define _GNU_SOURCE
#include <stdio.h>
#include <pthread.h>
#include <sched.h>

void *increment();

int counter = 0;

int main(){

	// Set processor affinity
	pthread_attr_t myattr;
   cpu_set_t cpuset;

   pthread_attr_init(&myattr);
   CPU_ZERO(&cpuset);
   CPU_SET(0, &cpuset);
   pthread_attr_setaffinity_np(&myattr, sizeof(cpu_set_t), &cpuset);


	int rc1, rc2;
	pthread_t t1, t2;

	if ((rc1 = pthread_create(&t1, NULL, &increment, NULL)))
	{
		printf("Error creating thread1: %d\n", rc1);
	}

	if ((rc2 = pthread_create(&t2, NULL, &increment, NULL)))
	{
		printf("Error creating thread2: %d\n", rc2);
	}

	pthread_join(t1, NULL);
	pthread_join(t2, NULL);

	printf("%d\n", counter);

	return 0;

}


void *increment()
{
	int i;
//	int *counter = (int *)counter_ptr;
	for (i = 0; i < 100000000; i++){
		counter++;
	}
}

