#include <stdio.h>
#include <stdbool.h>
#include <unistd.h>

#define MAX_NAME 128


unsigned long long int read_cpu_tick() {
  unsigned long long int usertime, nicetime, systemtime, idletime;
  unsigned long long int ioWait, irq, softIrq, steal, guest, guestnice;
  usertime = nicetime = systemtime = idletime = 0;
  ioWait = irq = softIrq = steal = guest = guestnice = 0;

  FILE *fp;
  fp = fopen("/proc/stat", "r");
  if (fp != NULL) {
    if (fscanf(fp,   "cpu  %16llu %16llu %16llu %16llu %16llu %16llu %16llu %16llu %16llu %16llu",
               &usertime, &nicetime, &systemtime, &idletime,
               &ioWait, &irq, &softIrq, &steal, &guest, &guestnice) == EOF) {
      fclose(fp);
      return 0;
    } else {
      fclose(fp);
      return usertime + nicetime + systemtime + idletime + ioWait + irq + softIrq + steal + guest + guestnice;
    }
  }else{
    return 0;
  }
}

unsigned long read_time_from_pid(int pid) {
  char fn[MAX_NAME + 1];
  snprintf(fn, sizeof fn, "/proc/%i/stat", pid);

  unsigned long utime = 0;
  unsigned long stime = 0;

  FILE *fp;
  fp = fopen(fn, "r");
  if (fp != NULL) {
    bool ans = fscanf(fp, "%*d %*s %*c %*d %*d %*d %*d %*d %*u %*u %*u %*u %*u %lu"
                      "%lu %*ld %*ld %*d %*d %*d %*d %*u %*lu %*ld",
                      &utime, &stime) != EOF;
    if (!ans) {
      fclose(fp);
      return 0;
    } else {
      fclose(fp);
      return utime + stime;
    }
  } else {
    return 0;
  }
}

unsigned int num_cores(){
  return sysconf(_SC_NPROCESSORS_ONLN);
}

