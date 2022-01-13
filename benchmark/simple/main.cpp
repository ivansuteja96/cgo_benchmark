#include <stdlib.h>
#include <unistd.h>

int Process1(int prev,int curr) {
	return prev+curr;
}

int* Process2(int length) {
    int *x = (int *) malloc(length*2);
	int prev = 0;
	for (int i = 0; i < length; ++i)
	{	
		x[i] = Process1(prev,i);
		prev = x[i];
	}
	return x;
}