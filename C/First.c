#include <stdio.h>
#include <stdlib.h>
#include <stdio.h>
#include <math.h>

int main()
{

    //------------------------------------------------------
    // ------ these is to define if a number is prime or Not
    // ------------------------------------------------------

    //     int n ,i=2;

    //     printf("what is the number : ");
    //     scanf("%d",&n);

    // if(n<=1){
    //     printf("The number is not a prime number\n",n);
    // }

    //     if(n%i==0){
    //         printf("The number is not a prime number\n",n);
    //     }else {
    //         printf("The number is a prime number\n",n);
    //         }

    //------------------------------------------------------
    // ------ these is to display all prime number from 1 to a number given by the user
    // ------------------------------------------------------

    // int n, i;

    // printf("what is the number : ");
    // scanf("%d", &n);

    // for (i = 1; i <= n; i++)
    // {

    //     if (i % 2 == 0)
    //     {
    //         printf("The number %d is not a prime number\n", i);
    //     }
    //     else
    //     {
    //         printf("The number %d is a prime number\n", i);
    //     }
    // }

    //------------------------------------------------------
    // ------ these is an algo that display all orange money code (0000 to 9999)
    // ------------------------------------------------------

    // int n=9999, i=0000;

    // for(i;i<=n;i++){
    //     printf("%d is an orange numbers code: \n",i);
    // }

    //------------------------------------------------------
    // ------ these is an algo that collect the mark of 10 student and display the average min  mark (i did not use array)
    // ------------------------------------------------------

    // int i, marks, total, average;

    // for (i = 1; i <= 10; i++)
    // {
    //     printf("what is the mark of %d student\n:", i);

    //     scanf("%d", &marks);

    //     printf("the mark of student number %d  is %d\n", i, marks);

    //     total = total + marks;
    // }

    // printf("the sum or total mark is %d\n", total);

    // average = total / 10;

    // printf("the average of the marks is %d\n", average);

    //------------------------------------------------------
    // ------ these is an algo that solve x'n where X and n are given by the user
    // ------------------------------------------------------

    // int i, n, x, solution = 1;
    // printf("what is the value of X: ");
    // scanf("%d", &x);
    // printf("what is the value of n: ");
    // scanf("%d", &n);

    // for (i = 0; i < n; i++)
    // {
    //     solution = solution * x;
    // }

    // printf("the result is %d\n", solution);

    //------------------------------------------------------
    // ------ these is an algo that display the multiplication table of all the num from 1 to given by user (up to 12)
    // ------------------------------------------------------

    // int i,num,mult,j;
    // printf("what is the number:");
    // scanf("%d",&num);

    // for (i=1; i<=num;i++)
    // {
    //     for(j=1,i; j<=12; j++){
    //         mult= i * j ;
    //         printf("%d x %d = %d\n",i,j,mult);
    //     }
    // }

    //------------------------------------------------------
    // ------ these is an algo that display the pascal triangle with number of a line n given buy user
    // -----------------------------------------------------

    // int n, i, j;
    // printf("Enter the number of lines: ");
    // scanf("%d", &n);

    // for (i = 1; i <= n; i++) {
    //     int k = 1;
    //     for (j = 2; j <= i; j++) {
    //         printf("%d ", k);
    //         k = k * (i - j + 1) / (j - 1);
    //     }
    //     printf("1\n");
    // }

    

    return 0;
}