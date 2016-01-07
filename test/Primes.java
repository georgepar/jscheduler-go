package test;

import java.util.concurrent.ExecutorService;
import java.util.concurrent.LinkedBlockingQueue;
import java.util.concurrent.ThreadPoolExecutor;
import java.util.concurrent.TimeUnit;

public class Primes {
    public static void main(String[] args) {

        ExecutorService[] executors = {
                new ThreadPoolExecutor(
                        5, // core pool size
                        5, // max pool size
                        0, TimeUnit.MILLISECONDS, // keep alive
                        new LinkedBlockingQueue<>(1), // queue
                        new ThreadPoolExecutor.CallerRunsPolicy()),
                new ThreadPoolExecutor(
                        5, // core pool size
                        5, // max pool size
                        0, TimeUnit.MILLISECONDS, // keep alive
                        new LinkedBlockingQueue<>(1), // queue
                        new ThreadPoolExecutor.CallerRunsPolicy())
        }; // handler

        System.out.println("Main thread name " + Thread.currentThread().getName());
        for (int i = 0; i<2; i++) {
            for (int thread = 0; thread < 5; thread++) {
            executors[i].submit(() -> {
                System.out.println("Entering thread " + Thread.currentThread().getName() + " #" + Thread.currentThread().getId());
/*
                final int NUM_TESTS = 1000;
                long start = System.nanoTime();
                for (int j = 0; j < NUM_TESTS; j++) {
                    long sleepTime = 500 * 1000000L; // convert to nanoseconds
                    long startTime = System.nanoTime();
                    while ((System.nanoTime() - startTime) < sleepTime);
                }
                System.out.println("Thread " + Thread.currentThread().getName() +
                        " took " + (System.nanoTime() - start) / 1000000 +
                        "ms (expected " + (NUM_TESTS * 500) + ")");
 */                       

                for (long job = 1; job <= 1000000000; job++) {
                    long count = 0;
                    for (long k=3; k<=10000L; k++) {
                        boolean isPrime = true;
                        for (long j=2; j<=k/2 && isPrime; j++) {
                            isPrime = k % j > 0;
                        }
                        if (isPrime) {
                            count++;
                        }
                    }
                    if(job % 100 == 0) {
                        System.out.println(Thread.currentThread().getName() + ": " + job);
                    } 
                }
            });
            }
            executors[i].shutdown();
        }

    }

}
