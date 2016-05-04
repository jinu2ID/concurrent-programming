.PHONY : all
all : p1 p2 p3

CC = gcc
CFLAGS = -pthread -o
TARGET1 = p1
TARGET2 = p2
TARGET3 = p3

p1 : p1.c
	$(CC) $(CFLAGS) $(TARGET1).o $(TARGET1).c
	./$(TARGET1).o

p2 : p2.c 
	$(CC) $(CFLAGS) $(TARGET2).o $(TARGET2).c
	./$(TARGET2).o

p3 : p3.c 
	$(CC) $(CFLAGS) $(TARGET3).o $(TARGET3).c
	./$(TARGET3).o

.PHONY : clean
clean :
	-\rm *.o
