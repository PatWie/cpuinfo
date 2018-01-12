package proc

// #include "wrapper.h"
import "C"


func CpuTick() (t int64) {
  return int64(C.read_cpu_tick())
}

func TimeFromPid(pid int) (t int64) {
  return int64(C.read_time_from_pid(C.int(pid)))
}

func NumCores() (n int){
  return int(C.num_cores())
}