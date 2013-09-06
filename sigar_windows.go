package sigar

import (
	"errors"
	//"fmt"
)

const errNotImplemented = "not implemented"

func (self *LoadAverage) Get() (err error) {
	return errors.New(errNotImplemented)
}

func (self *Uptime) Get() (err error) {
	self.Length, err = PdhGetCounterValueFloat64(PdhCounter_UpTime)
	return
}

func (self *Mem) Get() (err error) {

	var m MemoryStatusEx
	err = GlobalMemoryStatusEx(&m)
	if err != nil {
		return err
	}
	self.Total = m.ullTotalPhys
	self.Free = m.ullAvailPhys
	self.Used = self.Total - self.Free

	self.ActualUsed = self.Used
	self.ActualFree = self.Free

	cache, err := PdhGetCounterValueInt64(PdhCounter_UpTime)
	if err == nil {
		self.ActualUsed -= uint64(cache)
		self.ActualFree += uint64(cache)
	}

	return nil
}

func (self *Swap) Get() (err error) {
	return errors.New(errNotImplemented)
}

func (self *Cpu) Get() (err error) {
	return errors.New(errNotImplemented)
}

func (self *CpuList) Get() (err error) {
	return errors.New(errNotImplemented)
}

func (self *FileSystemList) Get() (err error) {
	return errors.New(errNotImplemented)
}

func (self *ProcList) Get() (err error) {
	return errors.New(errNotImplemented)
}

func (self *ProcState) Get(pid int) (err error) {
	return errors.New(errNotImplemented)
}

func (self *ProcMem) Get(pid int) (err error) {
	return errors.New(errNotImplemented)
}

func (self *ProcTime) Get(pid int) (err error) {
	return errors.New(errNotImplemented)
}

func (self *ProcArgs) Get(pid int) (err error) {
	return errors.New(errNotImplemented)
}

func (self *ProcExe) Get(pid int) (err error) {
	return errors.New(errNotImplemented)
}

func (self *FileSystemUsage) Get(path string) (err error) {
	return errors.New(errNotImplemented)
}
