#ifndef GOCODE_WRAPPER_H
#define GOCODE_WRAPPER_H

unsigned long long int read_cpu_tick();
unsigned long read_time_from_pid(int pid);
unsigned int num_cores();

#endif // GOCODE_WRAPPER_H