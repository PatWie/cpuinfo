package main

import (
  "github.com/patwie/cpuinfo/proc"
  "fmt"
  "time"
  "sort"
)

// hash map of all processes
var m map[int]*proc.Process

func DisplayProcessList(procs map[int]*proc.Process, factor float32, max int){

  var m_display []proc.Process

  for _, v := range procs{
    if v.Dirty == false && v.Active == true{
      if v.TimeCur - v.TimePrev > 0{
        usage := 1. / factor * float32(v.TimeCur - v.TimePrev)
        copy_proc := proc.Process{v.Pid, 0,0,false, false, usage}
        m_display = append(m_display, copy_proc)
      }
    }
  }

  sort.Sort(proc.ByUsage(m_display))

  if len(m_display) > 0 {
    fmt.Printf("\033[2J")
    for i := 0; i < max; i++ {
        fmt.Printf("pid: %d \t\t usage: %.2f\n", m_display[i].Pid, m_display[i].Usage)
    }
    
  }

}

func MarkDirtyProcessList(procs map[int]*proc.Process){
  for k, _ := range procs{
    procs[k].Dirty = true
    procs[k].Active = false
  }
}

func main(){

  m = make(map[int]*proc.Process)

  cpu_tick_prev := int64(0)
  cpu_tick_cur := int64(0)
  cores := proc.NumCores()

  for{
    MarkDirtyProcessList(m)

    cpu_tick_cur = proc.CpuTick()
    proc.UpdateProcessList(m)

    factor := float32(cpu_tick_cur - cpu_tick_prev) / float32(cores) / 100.
    DisplayProcessList(m, factor, 10)

    cpu_tick_prev = cpu_tick_cur
    time.Sleep(3 * time.Second)


  }

/**/
  // t_p := proc.TimeFromPid(20251)
  // for{
  //   time.Sleep(3 * time.Second)
  //   t_c := proc.TimeFromPid(20251)
  //   fmt.Println(1. / 3. * float32(t_c - t_p))
  //   t_p = t_c
  // }

/**/
  // fmt.Println(proc.CpuTick())
  // fmt.Println(proc.NumCores())
  // fmt.Println(proc.TimeFromPid(20251))

  // UpdateProcessList(m)
  // DisplayProcessList(m)
  
  // fmt.Println(m[20251])
}