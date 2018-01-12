package main

import (
  "github.com/patwie/cpuinfo/proc"
  "fmt"
  "io/ioutil"
  "strconv"
  "time"
)

type Process struct{
  Pid       int
  TimePrev int64
  TimeCur  int64
  Dirty     bool
}


var m map[int]*Process

func UpdateProcessList(procs map[int]*Process){
  files, err := ioutil.ReadDir("/proc")
  if err != nil {
    panic(err)
  }

  for _, file := range files {
    // list all possible directories
    name := file.Name()
    // we are just interested in numerical names
    if (name[0] < '0' || name[0] > '9') {
      continue;
    }
    // get pid
    pid, err := strconv.Atoi(name)
    if err != nil{
      continue
    }

    p := procs[pid]
    if p == nil{
      // is a new process
      p = &Process{pid, 0, 0, true}
    }else{
      // just update
      p.Dirty = false
      p.TimePrev = p.TimeCur
    }

    p.TimeCur = proc.TimeFromPid(pid)
    procs[pid] = p
  }
}

func DisplayProcessList(procs map[int]*Process, factor float32){
  for k, v := range procs{
    if v.Dirty == false{

      fmt.Println(k, v, (1. / factor * float32(v.TimeCur - v . TimePrev)))
    }
  }
}

func MarkDirtyProcessList(procs map[int]*Process){
  for k, _ := range procs{
    procs[k].Dirty = true
  }
}

func main(){

  m = make(map[int]*Process)

  cpu_tick_prev := int64(0)
  cpu_tick_cur := int64(0)
  cores := proc.NumCores()

  for{
    MarkDirtyProcessList(m)
    cpu_tick_cur = proc.CpuTick()
    UpdateProcessList(m)
    time.Sleep(3 * time.Second)

    cpu_tick_prev = cpu_tick_cur
    cpu_tick_cur = proc.CpuTick()
    factor := float32(cpu_tick_cur - cpu_tick_prev) / float32(cores) / 100.

    DisplayProcessList(m, factor)
  }

  // fmt.Println(proc.CpuTick())
  // fmt.Println(proc.NumCores())
  // fmt.Println(proc.TimeFromPid(20251))

  // UpdateProcessList(m)
  // DisplayProcessList(m)
  
  // fmt.Println(m[20251])
}