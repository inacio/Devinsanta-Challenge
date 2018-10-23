/*
 * 20181022
 * Gabriel Cesar
 * gabrielcesar2@gmail.com
 */

#include <stdio.h>

int main(void) 
{
	for(int i = 1; i <= 200; i++)
	{
		if(i % 6 == 0 && i % 5 == 0)
			printf("%d Santarem\n", i);
		else if(i % 6 == 0)
			printf("%d Santa\n", i);
		else if(i % 5 == 0)
			printf("%d rem\n", i);
		else					
			printf("%d\n", i);
	}
  return 0;
}