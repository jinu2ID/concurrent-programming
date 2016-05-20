// C implementation of ticket lock

#include <stdio.h>
#include <pthread.h>

int ticket = 0;
int serving = 0;
int counter = 0;

pthread_mutex_t lock;

void *increment();

int main(){
	
	int rc1, rc2, rc3, rc4, rc5, rc6;
	pthread_t t1, t2, t3, t4, t5, t6;

	if ((rc1 = pthread_create(&t1, NULL, &increment, NULL)))
	{
		printf("Error creating thread1: %d\n", rc1);
	}
	if ((rc2 = pthread_create(&t2, NULL, &increment, NULL)))
	{
		printf("Error creating thread2: %d\n", rc2);
	}
	if ((rc3 = pthread_create(&t3, NULL, &increment, NULL)))
	{
		printf("Error creating thread3: %d\n", rc3);
	}
	if ((rc4 = pthread_create(&t4, NULL, &increment, NULL)))
	{
		printf("Error creating thread4: %d\n", rc4);
	}
	if ((rc5 = pthread_create(&t5, NULL, &increment, NULL)))
	{
		printf("Error creating thread5: %d\n", rc5);
	}
	if ((rc6 = pthread_create(&t6, NULL, &increment, NULL)))
	{
		printf("Error creating thread6: %d\n", rc6);
	}


	pthread_join(t1, NULL);
	pthread_join(t2, NULL);
	pthread_join(t3, NULL);
	pthread_join(t4, NULL);
	pthread_join(t5, NULL);
	pthread_join(t6, NULL);

	printf("counter = %d\n", counter);

 	return 0;

}

// atomic fetch and add
int fetch_add(){
	pthread_mutex_lock(&lock);
	int value = ticket;
	ticket++;
	pthread_mutex_unlock(&lock);
	return value;
}

void ticketLock_acquire(){
	int myTick = fetch_add();
	while(serving != myTick){
		pthread_yield();	
	}
}

void ticketLock_release(){
	serving++;
}

void *increment(){
	int i;
	for (i = 0; i < 1000000; i++){
		ticketLock_acquire();
		counter++;
		ticketLock_release();
	}
}

