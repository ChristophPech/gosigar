package sigar

import (
	"errors"
	"fmt"
	"syscall"
	"unsafe"
)

var (
	kernel32             = syscall.NewLazyDLL("kernel32.dll")
	globalMemoryStatusEx = kernel32.NewProc("GlobalMemoryStatusEx")

	pdh                         = syscall.NewLazyDLL("Pdh.dll")
	pdhOpenQueryW               = pdh.NewProc("PdhOpenQueryW")
	pdhCloseQuery               = pdh.NewProc("PdhCloseQuery")
	pdhAddEnglishCounterW       = pdh.NewProc("PdhAddEnglishCounterW")
	pdhRemoveCounter            = pdh.NewProc("PdhRemoveCounter")
	pdhCollectQueryData         = pdh.NewProc("PdhCollectQueryData")
	pdhGetFormattedCounterValue = pdh.NewProc("PdhGetFormattedCounterValue")
)

const errInvArg = "invalid argument"

type MemoryStatusEx struct {
	dwLength                uint32
	dwMemoryLoad            uint32
	ullTotalPhys            uint64
	ullAvailPhys            uint64
	ullTotalPageFile        uint64
	ullAvailPageFile        uint64
	ullTotalVirtual         uint64
	ullAvailVirtual         uint64
	ullAvailExtendedVirtual uint64
}

const PdhCounter_UpTime = "\\System\\System Up Time"
const PdhCounter_SystemCache = "\\Memory\\System Cache Resident Bytes"

const (
	pdh_fmt_unicode = 0x00000040
	pdh_fmt_long    = 0x00000100
	pdh_fmt_double  = 0x00000200
	pdh_fmt_large   = 0x00000400
)

type pdh_fmt_countervalue_int32 struct {
	cStatus uint32
	value   int32
	padding int32
}
type pdh_fmt_countervalue_int64 struct {
	cStatus uint32
	value   int64
}
type pdh_fmt_countervalue_float64 struct {
	cStatus uint32
	value   float64
}
type pdh_fmt_countervalue_unicode struct {
	cStatus uint32
	value   *uint16
	padding int32
}

func UTF16PtrToString(s *uint16) string {
	if s == nil {
		return ""
	}
	return syscall.UTF16ToString((*[1 << 29]uint16)(unsafe.Pointer(s))[0:])
}

func GlobalMemoryStatusEx(s *MemoryStatusEx) (err error) {
	s.dwLength = 64
	r, _, e := syscall.Syscall(globalMemoryStatusEx.Addr(), 1, uintptr(unsafe.Pointer(s)), 0, 0)
	if r == 0 {
		if e != 0 {
			err = error(e)
		} else {
			err = errors.New(errInvArg)
		}
	}
	return
}

func fmtdummy() { fmt.Println(1) }

const (
	pdh_cstatus_no_machine        = 0x800007d0
	pdh_more_data                 = 0x800007d2
	pdh_no_data                   = 0x800007d5
	pdh_cstatus_no_object         = 0xc0000bb8
	pdh_cstatus_no_counter        = 0xc0000bb9
	pdh_memory_allocation_failure = 0xc0000bbb
	pdh_invalid_handle            = 0xc0000bbc
	pdh_invalid_argument          = 0xc0000bbd
	pdh_cstatus_bad_countername   = 0xc0000bc0
	pdh_insufficient_buffer       = 0xc0000bc2
	pdh_invalid_data              = 0xc0000bc6
	pdh_not_implemented           = 0xc0000bd3
	pdh_string_not_found          = 0xc0000bd4
)

func pdhResToError(e uint32) error {
	switch e {
	case 0:
		return nil
	case pdh_cstatus_no_machine:
		return errors.New("pdh_cstatus_no_machine")
	case pdh_more_data:
		return errors.New("pdh_more_data")
	case pdh_no_data:
		return errors.New("pdh_no_data")
	case pdh_cstatus_no_object:
		return errors.New("pdh_cstatus_no_object")
	case pdh_cstatus_no_counter:
		return errors.New("pdh_cstatus_no_counter")
	case pdh_memory_allocation_failure:
		return errors.New("pdh_memory_allocation_failure")
	case pdh_invalid_handle:
		return errors.New("pdh_invalid_handle")
	case pdh_invalid_argument:
		return errors.New("pdh_invalid_argument")
	case pdh_cstatus_bad_countername:
		return errors.New("pdh_cstatus_bad_countername")
	case pdh_insufficient_buffer:
		return errors.New("pdh_insufficient_buffer")
	case pdh_invalid_data:
		return errors.New("pdh_invalid_data")
	case pdh_not_implemented:
		return errors.New("pdh_not_implemented")
	case pdh_string_not_found:
		return errors.New("pdh_string_not_found")
	}
	return nil
}

func PdhOpenQuery(hq *syscall.Handle, userdata uint32) (err error) {
	r, _, e := syscall.Syscall(pdhOpenQueryW.Addr(), 3, 0, uintptr(userdata), uintptr(unsafe.Pointer(hq)))
	err = pdhResToError(uint32(r))
	if err != nil {
		return
	}
	if r != 0 {
		if e != 0 {
			err = error(e)
		} else {
			err = errors.New(errInvArg)
		}
	}
	return
}

func PdhCloseQuery(hq syscall.Handle) (err error) {
	r, _, e := syscall.Syscall(pdhCloseQuery.Addr(), 1, uintptr(hq), 0, 0)
	err = pdhResToError(uint32(r))
	if err != nil {
		return
	}
	if r != 0 {
		if e != 0 {
			err = error(e)
		} else {
			err = errors.New(errInvArg)
		}
	}
	return
}

func PdhRemoveCounter(hq syscall.Handle) (err error) {
	r, _, e := syscall.Syscall(pdhRemoveCounter.Addr(), 1, uintptr(hq), 0, 0)
	err = pdhResToError(uint32(r))
	if err != nil {
		return
	}
	if r != 0 {
		if e != 0 {
			err = error(e)
		} else {
			err = errors.New(errInvArg)
		}
	}
	return
}

func PdhCollectQueryData(hq syscall.Handle) (err error) {
	r, _, e := syscall.Syscall(pdhCollectQueryData.Addr(), 1, uintptr(hq), 0, 0)
	err = pdhResToError(uint32(r))
	if err != nil {
		return
	}
	if r != 0 {
		if e != 0 {
			err = error(e)
		} else {
			err = errors.New(errInvArg)
		}
	}
	return
}

func PdhAddCounter(hq syscall.Handle, path string, userdata uint32, hc *syscall.Handle) (err error) {
	var _p0 *uint16
	_p0, err = syscall.UTF16PtrFromString(path)
	if err != nil {
		return
	}

	r, _, e := syscall.Syscall6(pdhAddEnglishCounterW.Addr(), 4, uintptr(hq), uintptr(unsafe.Pointer(_p0)), uintptr(userdata), uintptr(unsafe.Pointer(hc)), 0, 0)
	err = pdhResToError(uint32(r))
	if err != nil {
		return
	}

	if r != 0 {
		if e != 0 {
			err = error(e)
		} else {
			err = errors.New(errInvArg)
		}
	}
	return
}

func PdhGetFormattedCounterValueString(hc syscall.Handle) (res string, err error) {
	var val pdh_fmt_countervalue_unicode
	r, _, e := syscall.Syscall6(pdhGetFormattedCounterValue.Addr(), 4, uintptr(hc), uintptr(pdh_fmt_unicode), 0, uintptr(unsafe.Pointer(&val)), 0, 0)
	err = pdhResToError(uint32(r))
	if err != nil {
		return
	}
	if r != 0 {
		if e != 0 {
			err = error(e)
		} else {
			err = errors.New(errInvArg)
		}
	}
	res = UTF16PtrToString(val.value)
	return
}

func PdhGetFormattedCounterValueInt32(hc syscall.Handle) (res int32, err error) {
	var val pdh_fmt_countervalue_int32
	r, _, e := syscall.Syscall6(pdhGetFormattedCounterValue.Addr(), 4, uintptr(hc), uintptr(pdh_fmt_long), 0, uintptr(unsafe.Pointer(&val)), 0, 0)
	err = pdhResToError(uint32(r))
	if err != nil {
		return
	}
	if r != 0 {
		if e != 0 {
			err = error(e)
		} else {
			err = errors.New(errInvArg)
		}
	}
	res = val.value
	return
}

func PdhGetFormattedCounterValueInt64(hc syscall.Handle) (res int64, err error) {
	var val pdh_fmt_countervalue_int64
	r, _, e := syscall.Syscall6(pdhGetFormattedCounterValue.Addr(), 4, uintptr(hc), uintptr(pdh_fmt_large), 0, uintptr(unsafe.Pointer(&val)), 0, 0)
	err = pdhResToError(uint32(r))
	if err != nil {
		return
	}
	if r != 0 {
		if e != 0 {
			err = error(e)
		} else {
			err = errors.New(errInvArg)
		}
	}
	res = val.value
	return
}

func PdhGetFormattedCounterValueFloat64(hc syscall.Handle) (res float64, err error) {
	var val pdh_fmt_countervalue_float64
	r, _, e := syscall.Syscall6(pdhGetFormattedCounterValue.Addr(), 4, uintptr(hc), uintptr(pdh_fmt_double), 0, uintptr(unsafe.Pointer(&val)), 0, 0)
	err = pdhResToError(uint32(r))
	if err != nil {
		return
	}
	if r != 0 {
		if e != 0 {
			err = error(e)
		} else {
			err = errors.New(errInvArg)
		}
	}
	res = val.value
	return
}

func PdhGetCounterValueInt32(path string) (res int32, err error) {
	var hq syscall.Handle
	err = PdhOpenQuery(&hq, 0)
	if err != nil {
		return
	}
	defer PdhCloseQuery(hq)
	var hc syscall.Handle
	err = PdhAddCounter(hq, path, 0, &hc)
	if err != nil {
		return
	}
	defer PdhRemoveCounter(hc)
	err = PdhCollectQueryData(hq)
	if err != nil {
		return
	}
	return PdhGetFormattedCounterValueInt32(hc)
}

func PdhGetCounterValueInt64(path string) (res int64, err error) {
	var hq syscall.Handle
	err = PdhOpenQuery(&hq, 0)
	if err != nil {
		return
	}
	defer PdhCloseQuery(hq)
	var hc syscall.Handle
	err = PdhAddCounter(hq, path, 0, &hc)
	if err != nil {
		return
	}
	defer PdhRemoveCounter(hc)
	err = PdhCollectQueryData(hq)
	if err != nil {
		return
	}
	return PdhGetFormattedCounterValueInt64(hc)
}

func PdhGetCounterValueFloat64(path string) (res float64, err error) {
	var hq syscall.Handle
	err = PdhOpenQuery(&hq, 0)
	if err != nil {
		return
	}
	defer PdhCloseQuery(hq)
	var hc syscall.Handle
	err = PdhAddCounter(hq, path, 0, &hc)
	if err != nil {
		return
	}
	defer PdhRemoveCounter(hc)
	err = PdhCollectQueryData(hq)
	if err != nil {
		return
	}
	return PdhGetFormattedCounterValueFloat64(hc)
}

func PdhGetCounterValueString(path string) (res string, err error) {
	var hq syscall.Handle
	err = PdhOpenQuery(&hq, 0)
	if err != nil {
		return
	}
	defer PdhCloseQuery(hq)
	var hc syscall.Handle
	err = PdhAddCounter(hq, path, 0, &hc)
	if err != nil {
		return
	}
	defer PdhRemoveCounter(hc)
	err = PdhCollectQueryData(hq)
	if err != nil {
		return
	}
	return PdhGetFormattedCounterValueString(hc)
}
