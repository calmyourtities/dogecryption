//============================================================================
// Name              : dogecrypt.cpp
// Author            : calmyourtities
// Version           :
// Copyright         : COPYRIGHT HAHAHA
// Description       : Hello World in C++, Ansi-style
// Partial Credit to : http://www.sanfoundry.com/cpp-program-implement-rsa-algorithm/
//============================================================================

#include <iostream>
#include <math.h>
#include <string.h>
#include <stdlib.h>
#include <random>
#include <climits>
#include <ctime>

/* #define ULLONG_MAX 18446744073709551615 */

unsigned long long p, q, n, t, flag, e[100], d[100], temp[100], j, m[100], en[100], i;

std::string msg;

int prime(unsigned long long);

void ce();

long int cd(long int);

void encrypt();

void decrypt();

unsigned long long getBigPrime();
unsigned long long randULL();

int prime(unsigned long long pr) {
    unsigned long long i;

    j = sqrt(pr);

    for (i = 2; i <= j; i++) {

        if (pr % i == 0) {
            return 0;
        }

    }

    return 1;

}

int main() {

	srand(time(NULL));

	std::cout << "BIG PRIME: " << std::endl;

	/*clock_t start = clock();

	p = 0;

	int primesFound = 0;
	for(int tries = 0; primesFound < 2; tries++) {
		unsigned long long cnum = randULL();
		std::cout << "checking: " << cnum << std::endl;
		std::cout << cnum;
		if(prime(cnum)) {
			std::cout << " is prime" << std::endl;
			primesFound++;
			if(p == 0)
				p = cnum;
			else
				q = cnum;
		}
		else
			std::cout << " is not prime" << std::endl;

		std::cout << "tested " << (tries+1) << " numbers " << std::endl << "found " << primesFound << " prime(s)" << std::endl;
	}

	std::cout << "took " << (clock() - start)/1000 << " seconds";*/

	p = 7;
	q = 13;

	std::cout << std::endl << "YOUR PRIMES ARE: " << std::endl << p << " AND " << q;

    std::cout << "\nENTER MESSAGE\n";

    //std::cin.clear();
    //std::cin.ignore(INT_MAX, '\n');
    getline(std::cin, msg);

    std::cout << "done" << std::endl;

    for(long int i = 0; i < msg.size(); i++)
    	m[i] = msg.at(i);

    n = p * q;

    t = (p - 1) * (q - 1);

    std::cout << "done" << std::endl;

    ce();

    std::cout << "\nPOSSIBLE VALUES OF e AND d ARE\n";

    for (i = 0; i < j - 1; i++)
        std::cout << e[i] << "\t" << d[i] << "\n";

    std::cout << "done" << std::endl;

    encrypt();

    decrypt();

    return 0;

}

void ce() {

    int k = 0;

    for (i = 2; i < t; i++) {

        if (t % i == 0)
            continue;

        flag = prime(i);

        if (flag == 1 && i != p && i != q)

        {

            e[k] = i;

            flag = cd(e[k]);

            if (flag > 0) {
                d[k] = flag;
                k++;
            }

            if (k == 99)
                break;
        }

    }

}

long int cd(long int x) {

    long int k = 1;

    while (1)

    {

        k = k + t;

        if (k % x == 0)

            return (k / x);

    }

}

void encrypt()

{

    unsigned long long pt, ct, key = e[0], k, len;

    i = 0;

    len = msg.size();

    while (i != len)

    {

        pt = m[i];

        pt = pt - 96;

        k = 1;

        for (j = 0; j < key; j++)

        {

            k = k * pt;

            k = k % n;

        }

        temp[i] = k;

        ct = k + 96;

        en[i] = ct;

        i++;

    }

    en[i] = -1;

    std::cout << "\nTHE ENCRYPTED MESSAGE IS\n";

    for (i = 0; en[i] != -1; i++)
        std::cout << en[i] << std::endl;

}

void decrypt() {

    long int pt, ct, key = d[0], k;
    i = 0;
    while (en[i] != -1) {
        ct = temp[i];
        k = 1;

        for (j = 0; j < key; j++) {

            k = k * ct;

            k = k % n;

        }

        pt = k + 96;

        m[i] = pt;

        i++;

    }

    m[i] = -1;

    std::cout << "\nTHE DECRYPTED MESSAGE IS\n";

    for (i = 0; m[i] != -1; i++)
        std::cout << (char) m[i];

}

unsigned long long getBigPrime() {
	while(true) {
		unsigned long long prime = randULL();
		unsigned long long primec = 13;
		unsigned long long limit = sqrt(prime);
		std::cout << "TESTING: " << primec << std::endl;
		if(prime % 2 == 0) continue;
		if(prime % 3 == 0) continue;
		if(prime % 5 == 0) continue;
		if(prime % 7 == 0) continue;
		if(prime % 11 == 0) continue;
		if(prime % 13 == 0) continue;
		while(primec++ < limit)
			if(prime % primec == 0)
				continue;
		return prime;
	}
}

unsigned long long randULL() {
    unsigned long long r = 0;

    for (int i = 0; i < 5; ++i) {
        r = (r << 15) | (rand() & 0x7FFF);
    }

    return r & 0xFFFFFFFFFFFFFFFFULL;
}
