package proxy

// #include "proxy.h"
import "C"

import "unsafe"

import "fmt"

type WasiPollPollPollable = uint32
type WasiClocksWallClockDatetime struct {
  Seconds uint64
  Nanoseconds uint32
}

type WasiClocksTimezoneDatetime = WasiClocksWallClockDatetime
type WasiClocksTimezoneTimezoneDisplay struct {
  UtcOffset int32
  Name string
  InDaylightSavingTime bool
}

type WasiClocksTimezoneTimezone = uint32
type WasiClocksMonotonicClockPollable = uint32
type WasiClocksMonotonicClockInstant = uint64
type WasiIoStreamsPollable = uint32
type WasiIoStreamsStreamError struct {
}

type WasiIoStreamsOutputStream = uint32
type WasiIoStreamsInputStream = uint32
type WasiIoStreamsTuple2ListU8TBoolT struct {
  F0 []uint8
  F1 bool
}

type WasiIoStreamsTuple2U64BoolT struct {
  F0 uint64
  F1 bool
}

type WasiCliStdoutOutputStream = uint32
type WasiCliStdinInputStream = uint32
type WasiCliStderrOutputStream = uint32
type WasiHttpTypesInputStream = uint32
type WasiHttpTypesOutputStream = uint32
type WasiHttpTypesPollable = uint32
type WasiHttpTypesStatusCode = uint16
type WasiHttpTypesSchemeKind int

const (
WasiHttpTypesSchemeKindHttp WasiHttpTypesSchemeKind = iota
WasiHttpTypesSchemeKindHttps
WasiHttpTypesSchemeKindOther
)

type WasiHttpTypesScheme struct {
  kind WasiHttpTypesSchemeKind
  val any
}

func (n WasiHttpTypesScheme) Kind() WasiHttpTypesSchemeKind {
  return n.kind
}

func WasiHttpTypesSchemeHttp() WasiHttpTypesScheme{
  return WasiHttpTypesScheme{kind: WasiHttpTypesSchemeKindHttp}
}

func WasiHttpTypesSchemeHttps() WasiHttpTypesScheme{
  return WasiHttpTypesScheme{kind: WasiHttpTypesSchemeKindHttps}
}

func WasiHttpTypesSchemeOther(v string) WasiHttpTypesScheme{
  return WasiHttpTypesScheme{kind: WasiHttpTypesSchemeKindOther, val: v}
}

func (n WasiHttpTypesScheme) GetOther() string{
  if g, w := n.Kind(), WasiHttpTypesSchemeKindOther; g != w {
    panic(fmt.Sprintf("Attr kind is %v, not %v", g, w))
  }
  return n.val.(string)
}

func (n *WasiHttpTypesScheme) SetOther(v string) {
  n.val = v
  n.kind = WasiHttpTypesSchemeKindOther
}

type WasiHttpTypesResponseOutparam = uint32
type WasiHttpTypesRequestOptions struct {
  ConnectTimeoutMs Option[uint32]
  FirstByteTimeoutMs Option[uint32]
  BetweenBytesTimeoutMs Option[uint32]
}

type WasiHttpTypesOutgoingStream = uint32
type WasiHttpTypesOutgoingResponse = uint32
type WasiHttpTypesOutgoingRequest = uint32
type WasiHttpTypesMethodKind int

const (
WasiHttpTypesMethodKindGet WasiHttpTypesMethodKind = iota
WasiHttpTypesMethodKindHead
WasiHttpTypesMethodKindPost
WasiHttpTypesMethodKindPut
WasiHttpTypesMethodKindDelete
WasiHttpTypesMethodKindConnect
WasiHttpTypesMethodKindOptions
WasiHttpTypesMethodKindTrace
WasiHttpTypesMethodKindPatch
WasiHttpTypesMethodKindOther
)

type WasiHttpTypesMethod struct {
  kind WasiHttpTypesMethodKind
  val any
}

func (n WasiHttpTypesMethod) Kind() WasiHttpTypesMethodKind {
  return n.kind
}

func WasiHttpTypesMethodGet() WasiHttpTypesMethod{
  return WasiHttpTypesMethod{kind: WasiHttpTypesMethodKindGet}
}

func WasiHttpTypesMethodHead() WasiHttpTypesMethod{
  return WasiHttpTypesMethod{kind: WasiHttpTypesMethodKindHead}
}

func WasiHttpTypesMethodPost() WasiHttpTypesMethod{
  return WasiHttpTypesMethod{kind: WasiHttpTypesMethodKindPost}
}

func WasiHttpTypesMethodPut() WasiHttpTypesMethod{
  return WasiHttpTypesMethod{kind: WasiHttpTypesMethodKindPut}
}

func WasiHttpTypesMethodDelete() WasiHttpTypesMethod{
  return WasiHttpTypesMethod{kind: WasiHttpTypesMethodKindDelete}
}

func WasiHttpTypesMethodConnect() WasiHttpTypesMethod{
  return WasiHttpTypesMethod{kind: WasiHttpTypesMethodKindConnect}
}

func WasiHttpTypesMethodOptions() WasiHttpTypesMethod{
  return WasiHttpTypesMethod{kind: WasiHttpTypesMethodKindOptions}
}

func WasiHttpTypesMethodTrace() WasiHttpTypesMethod{
  return WasiHttpTypesMethod{kind: WasiHttpTypesMethodKindTrace}
}

func WasiHttpTypesMethodPatch() WasiHttpTypesMethod{
  return WasiHttpTypesMethod{kind: WasiHttpTypesMethodKindPatch}
}

func WasiHttpTypesMethodOther(v string) WasiHttpTypesMethod{
  return WasiHttpTypesMethod{kind: WasiHttpTypesMethodKindOther, val: v}
}

func (n WasiHttpTypesMethod) GetOther() string{
  if g, w := n.Kind(), WasiHttpTypesMethodKindOther; g != w {
    panic(fmt.Sprintf("Attr kind is %v, not %v", g, w))
  }
  return n.val.(string)
}

func (n *WasiHttpTypesMethod) SetOther(v string) {
  n.val = v
  n.kind = WasiHttpTypesMethodKindOther
}

type WasiHttpTypesIncomingStream = uint32
type WasiHttpTypesIncomingResponse = uint32
type WasiHttpTypesIncomingRequest = uint32
type WasiHttpTypesFutureIncomingResponse = uint32
type WasiHttpTypesFields = uint32
type WasiHttpTypesTrailers = uint32
type WasiHttpTypesHeaders = uint32
type WasiHttpTypesErrorKind int

const (
WasiHttpTypesErrorKindInvalidUrl WasiHttpTypesErrorKind = iota
WasiHttpTypesErrorKindTimeoutError
WasiHttpTypesErrorKindProtocolError
WasiHttpTypesErrorKindUnexpectedError
)

type WasiHttpTypesError struct {
  kind WasiHttpTypesErrorKind
  val any
}

func (n WasiHttpTypesError) Kind() WasiHttpTypesErrorKind {
  return n.kind
}

func WasiHttpTypesErrorInvalidUrl(v string) WasiHttpTypesError{
  return WasiHttpTypesError{kind: WasiHttpTypesErrorKindInvalidUrl, val: v}
}

func (n WasiHttpTypesError) GetInvalidUrl() string{
  if g, w := n.Kind(), WasiHttpTypesErrorKindInvalidUrl; g != w {
    panic(fmt.Sprintf("Attr kind is %v, not %v", g, w))
  }
  return n.val.(string)
}

func (n *WasiHttpTypesError) SetInvalidUrl(v string) {
  n.val = v
  n.kind = WasiHttpTypesErrorKindInvalidUrl
}

func WasiHttpTypesErrorTimeoutError(v string) WasiHttpTypesError{
  return WasiHttpTypesError{kind: WasiHttpTypesErrorKindTimeoutError, val: v}
}

func (n WasiHttpTypesError) GetTimeoutError() string{
  if g, w := n.Kind(), WasiHttpTypesErrorKindTimeoutError; g != w {
    panic(fmt.Sprintf("Attr kind is %v, not %v", g, w))
  }
  return n.val.(string)
}

func (n *WasiHttpTypesError) SetTimeoutError(v string) {
  n.val = v
  n.kind = WasiHttpTypesErrorKindTimeoutError
}

func WasiHttpTypesErrorProtocolError(v string) WasiHttpTypesError{
  return WasiHttpTypesError{kind: WasiHttpTypesErrorKindProtocolError, val: v}
}

func (n WasiHttpTypesError) GetProtocolError() string{
  if g, w := n.Kind(), WasiHttpTypesErrorKindProtocolError; g != w {
    panic(fmt.Sprintf("Attr kind is %v, not %v", g, w))
  }
  return n.val.(string)
}

func (n *WasiHttpTypesError) SetProtocolError(v string) {
  n.val = v
  n.kind = WasiHttpTypesErrorKindProtocolError
}

func WasiHttpTypesErrorUnexpectedError(v string) WasiHttpTypesError{
  return WasiHttpTypesError{kind: WasiHttpTypesErrorKindUnexpectedError, val: v}
}

func (n WasiHttpTypesError) GetUnexpectedError() string{
  if g, w := n.Kind(), WasiHttpTypesErrorKindUnexpectedError; g != w {
    panic(fmt.Sprintf("Attr kind is %v, not %v", g, w))
  }
  return n.val.(string)
}

func (n *WasiHttpTypesError) SetUnexpectedError(v string) {
  n.val = v
  n.kind = WasiHttpTypesErrorKindUnexpectedError
}

type WasiHttpTypesTuple2StringStringT struct {
  F0 string
  F1 string
}

type WasiHttpTypesTuple2StringListU8TT struct {
  F0 string
  F1 []uint8
}

type WasiHttpOutgoingHandlerOutgoingRequest = uint32
type WasiHttpOutgoingHandlerRequestOptions = WasiHttpTypesRequestOptions
type WasiHttpOutgoingHandlerFutureIncomingResponse = uint32
type WasiHttpIncomingHandlerIncomingRequest = uint32
type WasiHttpIncomingHandlerResponseOutparam = uint32
// Import functions from wasi:clocks/wall-clock
func WasiClocksWallClockNow() WasiClocksWallClockDatetime {
  var ret C.wasi_clocks_wall_clock_datetime_t
  C.wasi_clocks_wall_clock_now(&ret)
  var lift_ret WasiClocksWallClockDatetime
  var lift_ret_Seconds uint64
  lift_ret_Seconds = uint64(ret.seconds)
  lift_ret.Seconds = lift_ret_Seconds
  var lift_ret_Nanoseconds uint32
  lift_ret_Nanoseconds = uint32(ret.nanoseconds)
  lift_ret.Nanoseconds = lift_ret_Nanoseconds
  return lift_ret
}

func WasiClocksWallClockResolution() WasiClocksWallClockDatetime {
  var ret C.wasi_clocks_wall_clock_datetime_t
  C.wasi_clocks_wall_clock_resolution(&ret)
  var lift_ret WasiClocksWallClockDatetime
  var lift_ret_Seconds uint64
  lift_ret_Seconds = uint64(ret.seconds)
  lift_ret.Seconds = lift_ret_Seconds
  var lift_ret_Nanoseconds uint32
  lift_ret_Nanoseconds = uint32(ret.nanoseconds)
  lift_ret.Nanoseconds = lift_ret_Nanoseconds
  return lift_ret
}

// Import functions from wasi:poll/poll
func WasiPollPollDropPollable(this uint32) {
  var lower_this C.uint32_t
  lower_this_val := C.uint32_t(this)
  lower_this = lower_this_val
  C.wasi_poll_poll_drop_pollable(lower_this)
}

func WasiPollPollPollOneoff(in []uint32) []uint8 {
  var lower_in C.proxy_list_pollable_t
  if len(in) == 0 {
    lower_in.ptr = nil
    lower_in.len = 0
  } else {
    var empty_lower_in C.wasi_poll_poll_pollable_t
    lower_in.ptr = (*C.wasi_poll_poll_pollable_t)(C.malloc(C.size_t(len(in)) * C.size_t(unsafe.Sizeof(empty_lower_in))))
    lower_in.len = C.size_t(len(in))
    for lower_in_i := range in {
      lower_in_ptr := (*C.wasi_poll_poll_pollable_t)(unsafe.Pointer(uintptr(unsafe.Pointer(lower_in.ptr)) +
      uintptr(lower_in_i)*unsafe.Sizeof(empty_lower_in)))
      var lower_in_ptr_value C.uint32_t
      lower_in_ptr_value_val := C.uint32_t(in[lower_in_i])
      lower_in_ptr_value = lower_in_ptr_value_val
      *lower_in_ptr = lower_in_ptr_value
    }
  }
  defer C.proxy_list_pollable_free(&lower_in)
  var ret C.proxy_list_u8_t
  C.wasi_poll_poll_poll_oneoff(&lower_in, &ret)
  var lift_ret []uint8
  lift_ret = make([]uint8, ret.len)
  if ret.len > 0 {
    for lift_ret_i := 0; lift_ret_i < int(ret.len); lift_ret_i++ {
      var empty_lift_ret C.uint8_t
      lift_ret_ptr := *(*C.uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(ret.ptr)) +
      uintptr(lift_ret_i)*unsafe.Sizeof(empty_lift_ret)))
      var list_lift_ret uint8
      list_lift_ret = uint8(lift_ret_ptr)
      lift_ret[lift_ret_i] = list_lift_ret
    }
  }
  return lift_ret
}

// Import functions from wasi:clocks/monotonic-clock
func WasiClocksMonotonicClockNow() uint64 {
  ret := C.wasi_clocks_monotonic_clock_now()
  var lift_ret uint64
  var lift_ret_val uint64
  lift_ret_val = uint64(ret)
  lift_ret = lift_ret_val
  return lift_ret
}

func WasiClocksMonotonicClockResolution() uint64 {
  ret := C.wasi_clocks_monotonic_clock_resolution()
  var lift_ret uint64
  var lift_ret_val uint64
  lift_ret_val = uint64(ret)
  lift_ret = lift_ret_val
  return lift_ret
}

func WasiClocksMonotonicClockSubscribe(when uint64, absolute bool) uint32 {
  var lower_when C.uint64_t
  lower_when_val := C.uint64_t(when)
  lower_when = lower_when_val
  lower_absolute := absolute
  ret := C.wasi_clocks_monotonic_clock_subscribe(lower_when, lower_absolute)
  var lift_ret uint32
  var lift_ret_val uint32
  var lift_ret_val_val uint32
  lift_ret_val_val = uint32(ret)
  lift_ret_val = lift_ret_val_val
  lift_ret = lift_ret_val
  return lift_ret
}

// Import functions from wasi:clocks/timezone
func WasiClocksTimezoneDisplay(this uint32, when WasiClocksWallClockDatetime) WasiClocksTimezoneTimezoneDisplay {
  var lower_this C.uint32_t
  lower_this_val := C.uint32_t(this)
  lower_this = lower_this_val
  var lower_when C.wasi_clocks_wall_clock_datetime_t
  var lower_when_val C.wasi_clocks_wall_clock_datetime_t
  lower_when_val_seconds := C.uint64_t(when.Seconds)
  lower_when_val.seconds = lower_when_val_seconds
  lower_when_val_nanoseconds := C.uint32_t(when.Nanoseconds)
  lower_when_val.nanoseconds = lower_when_val_nanoseconds
  lower_when = lower_when_val
  var ret C.wasi_clocks_timezone_timezone_display_t
  C.wasi_clocks_timezone_display(lower_this, &lower_when, &ret)
  var lift_ret WasiClocksTimezoneTimezoneDisplay
  var lift_ret_UtcOffset int32
  lift_ret_UtcOffset = int32(ret.utc_offset)
  lift_ret.UtcOffset = lift_ret_UtcOffset
  var lift_ret_Name string
  lift_ret_Name = C.GoStringN(ret.name.ptr, C.int(ret.name.len))
  lift_ret.Name = lift_ret_Name
  lift_ret_InDaylightSavingTime := ret.in_daylight_saving_time
  lift_ret.InDaylightSavingTime = lift_ret_InDaylightSavingTime
  return lift_ret
}

func WasiClocksTimezoneUtcOffset(this uint32, when WasiClocksWallClockDatetime) int32 {
  var lower_this C.uint32_t
  lower_this_val := C.uint32_t(this)
  lower_this = lower_this_val
  var lower_when C.wasi_clocks_wall_clock_datetime_t
  var lower_when_val C.wasi_clocks_wall_clock_datetime_t
  lower_when_val_seconds := C.uint64_t(when.Seconds)
  lower_when_val.seconds = lower_when_val_seconds
  lower_when_val_nanoseconds := C.uint32_t(when.Nanoseconds)
  lower_when_val.nanoseconds = lower_when_val_nanoseconds
  lower_when = lower_when_val
  ret := C.wasi_clocks_timezone_utc_offset(lower_this, &lower_when)
  var lift_ret int32
  lift_ret = int32(ret)
  return lift_ret
}

func WasiClocksTimezoneDropTimezone(this uint32) {
  var lower_this C.uint32_t
  lower_this_val := C.uint32_t(this)
  lower_this = lower_this_val
  C.wasi_clocks_timezone_drop_timezone(lower_this)
}

// Import functions from wasi:random/random
func WasiRandomRandomGetRandomBytes(len uint64) []uint8 {
  lower_len := C.uint64_t(len)
  var ret C.proxy_list_u8_t
  C.wasi_random_random_get_random_bytes(lower_len, &ret)
  var lift_ret []uint8
  lift_ret = make([]uint8, ret.len)
  if ret.len > 0 {
    for lift_ret_i := 0; lift_ret_i < int(ret.len); lift_ret_i++ {
      var empty_lift_ret C.uint8_t
      lift_ret_ptr := *(*C.uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(ret.ptr)) +
      uintptr(lift_ret_i)*unsafe.Sizeof(empty_lift_ret)))
      var list_lift_ret uint8
      list_lift_ret = uint8(lift_ret_ptr)
      lift_ret[lift_ret_i] = list_lift_ret
    }
  }
  return lift_ret
}

func WasiRandomRandomGetRandomU64() uint64 {
  ret := C.wasi_random_random_get_random_u64()
  var lift_ret uint64
  lift_ret = uint64(ret)
  return lift_ret
}

// Import functions from wasi:io/streams
func WasiIoStreamsRead(this uint32, len uint64) Result[WasiIoStreamsTuple2ListU8TBoolT, WasiIoStreamsStreamError] {
  var lower_this C.uint32_t
  lower_this_val := C.uint32_t(this)
  lower_this = lower_this_val
  lower_len := C.uint64_t(len)
  var ret C.proxy_result_tuple2_list_u8_bool_stream_error_t
  C.wasi_io_streams_read(lower_this, lower_len, &ret)
  var lift_ret Result[WasiIoStreamsTuple2ListU8TBoolT, WasiIoStreamsStreamError]
  if ret.is_err {
    var lift_ret_val WasiIoStreamsStreamError
    lift_ret.SetErr(lift_ret_val)
  } else {
    lift_ret_ptr := *(*C.proxy_tuple2_list_u8_bool_t)(unsafe.Pointer(&ret.val))
    var lift_ret_val WasiIoStreamsTuple2ListU8TBoolT
    var lift_ret_val_F0 []uint8
    lift_ret_val_F0 = make([]uint8, lift_ret_ptr.f0.len)
    if lift_ret_ptr.f0.len > 0 {
      for lift_ret_val_F0_i := 0; lift_ret_val_F0_i < int(lift_ret_ptr.f0.len); lift_ret_val_F0_i++ {
        var empty_lift_ret_val_F0 C.uint8_t
        lift_ret_val_F0_ptr := *(*C.uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(lift_ret_ptr.f0.ptr)) +
        uintptr(lift_ret_val_F0_i)*unsafe.Sizeof(empty_lift_ret_val_F0)))
        var list_lift_ret_val_F0 uint8
        list_lift_ret_val_F0 = uint8(lift_ret_val_F0_ptr)
        lift_ret_val_F0[lift_ret_val_F0_i] = list_lift_ret_val_F0
      }
    }
    lift_ret_val.F0 = lift_ret_val_F0
    lift_ret_val_F1 := lift_ret_ptr.f1
    lift_ret_val.F1 = lift_ret_val_F1
    lift_ret.Set(lift_ret_val)
  }
  return lift_ret
}

func WasiIoStreamsBlockingRead(this uint32, len uint64) Result[WasiIoStreamsTuple2ListU8TBoolT, WasiIoStreamsStreamError] {
  var lower_this C.uint32_t
  lower_this_val := C.uint32_t(this)
  lower_this = lower_this_val
  lower_len := C.uint64_t(len)
  var ret C.proxy_result_tuple2_list_u8_bool_stream_error_t
  C.wasi_io_streams_blocking_read(lower_this, lower_len, &ret)
  var lift_ret Result[WasiIoStreamsTuple2ListU8TBoolT, WasiIoStreamsStreamError]
  if ret.is_err {
    var lift_ret_val WasiIoStreamsStreamError
    lift_ret.SetErr(lift_ret_val)
  } else {
    lift_ret_ptr := *(*C.proxy_tuple2_list_u8_bool_t)(unsafe.Pointer(&ret.val))
    var lift_ret_val WasiIoStreamsTuple2ListU8TBoolT
    var lift_ret_val_F0 []uint8
    lift_ret_val_F0 = make([]uint8, lift_ret_ptr.f0.len)
    if lift_ret_ptr.f0.len > 0 {
      for lift_ret_val_F0_i := 0; lift_ret_val_F0_i < int(lift_ret_ptr.f0.len); lift_ret_val_F0_i++ {
        var empty_lift_ret_val_F0 C.uint8_t
        lift_ret_val_F0_ptr := *(*C.uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(lift_ret_ptr.f0.ptr)) +
        uintptr(lift_ret_val_F0_i)*unsafe.Sizeof(empty_lift_ret_val_F0)))
        var list_lift_ret_val_F0 uint8
        list_lift_ret_val_F0 = uint8(lift_ret_val_F0_ptr)
        lift_ret_val_F0[lift_ret_val_F0_i] = list_lift_ret_val_F0
      }
    }
    lift_ret_val.F0 = lift_ret_val_F0
    lift_ret_val_F1 := lift_ret_ptr.f1
    lift_ret_val.F1 = lift_ret_val_F1
    lift_ret.Set(lift_ret_val)
  }
  return lift_ret
}

func WasiIoStreamsSkip(this uint32, len uint64) Result[WasiIoStreamsTuple2U64BoolT, WasiIoStreamsStreamError] {
  var lower_this C.uint32_t
  lower_this_val := C.uint32_t(this)
  lower_this = lower_this_val
  lower_len := C.uint64_t(len)
  var ret C.proxy_result_tuple2_u64_bool_stream_error_t
  C.wasi_io_streams_skip(lower_this, lower_len, &ret)
  var lift_ret Result[WasiIoStreamsTuple2U64BoolT, WasiIoStreamsStreamError]
  if ret.is_err {
    var lift_ret_val WasiIoStreamsStreamError
    lift_ret.SetErr(lift_ret_val)
  } else {
    lift_ret_ptr := *(*C.proxy_tuple2_u64_bool_t)(unsafe.Pointer(&ret.val))
    var lift_ret_val WasiIoStreamsTuple2U64BoolT
    var lift_ret_val_F0 uint64
    lift_ret_val_F0 = uint64(lift_ret_ptr.f0)
    lift_ret_val.F0 = lift_ret_val_F0
    lift_ret_val_F1 := lift_ret_ptr.f1
    lift_ret_val.F1 = lift_ret_val_F1
    lift_ret.Set(lift_ret_val)
  }
  return lift_ret
}

func WasiIoStreamsBlockingSkip(this uint32, len uint64) Result[WasiIoStreamsTuple2U64BoolT, WasiIoStreamsStreamError] {
  var lower_this C.uint32_t
  lower_this_val := C.uint32_t(this)
  lower_this = lower_this_val
  lower_len := C.uint64_t(len)
  var ret C.proxy_result_tuple2_u64_bool_stream_error_t
  C.wasi_io_streams_blocking_skip(lower_this, lower_len, &ret)
  var lift_ret Result[WasiIoStreamsTuple2U64BoolT, WasiIoStreamsStreamError]
  if ret.is_err {
    var lift_ret_val WasiIoStreamsStreamError
    lift_ret.SetErr(lift_ret_val)
  } else {
    lift_ret_ptr := *(*C.proxy_tuple2_u64_bool_t)(unsafe.Pointer(&ret.val))
    var lift_ret_val WasiIoStreamsTuple2U64BoolT
    var lift_ret_val_F0 uint64
    lift_ret_val_F0 = uint64(lift_ret_ptr.f0)
    lift_ret_val.F0 = lift_ret_val_F0
    lift_ret_val_F1 := lift_ret_ptr.f1
    lift_ret_val.F1 = lift_ret_val_F1
    lift_ret.Set(lift_ret_val)
  }
  return lift_ret
}

func WasiIoStreamsSubscribeToInputStream(this uint32) uint32 {
  var lower_this C.uint32_t
  lower_this_val := C.uint32_t(this)
  lower_this = lower_this_val
  ret := C.wasi_io_streams_subscribe_to_input_stream(lower_this)
  var lift_ret uint32
  var lift_ret_val uint32
  var lift_ret_val_val uint32
  lift_ret_val_val = uint32(ret)
  lift_ret_val = lift_ret_val_val
  lift_ret = lift_ret_val
  return lift_ret
}

func WasiIoStreamsDropInputStream(this uint32) {
  var lower_this C.uint32_t
  lower_this_val := C.uint32_t(this)
  lower_this = lower_this_val
  C.wasi_io_streams_drop_input_stream(lower_this)
}

func WasiIoStreamsWrite(this uint32, buf []uint8) Result[uint64, WasiIoStreamsStreamError] {
  var lower_this C.uint32_t
  lower_this_val := C.uint32_t(this)
  lower_this = lower_this_val
  var lower_buf C.proxy_list_u8_t
  if len(buf) == 0 {
    lower_buf.ptr = nil
    lower_buf.len = 0
  } else {
    var empty_lower_buf C.uint8_t
    lower_buf.ptr = (*C.uint8_t)(C.malloc(C.size_t(len(buf)) * C.size_t(unsafe.Sizeof(empty_lower_buf))))
    lower_buf.len = C.size_t(len(buf))
    for lower_buf_i := range buf {
      lower_buf_ptr := (*C.uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(lower_buf.ptr)) +
      uintptr(lower_buf_i)*unsafe.Sizeof(empty_lower_buf)))
      lower_buf_ptr_value := C.uint8_t(buf[lower_buf_i])
      *lower_buf_ptr = lower_buf_ptr_value
    }
  }
  defer C.proxy_list_u8_free(&lower_buf)
  var ret C.proxy_result_u64_stream_error_t
  C.wasi_io_streams_write(lower_this, &lower_buf, &ret)
  var lift_ret Result[uint64, WasiIoStreamsStreamError]
  if ret.is_err {
    var lift_ret_val WasiIoStreamsStreamError
    lift_ret.SetErr(lift_ret_val)
  } else {
    lift_ret_ptr := *(*C.uint64_t)(unsafe.Pointer(&ret.val))
    var lift_ret_val uint64
    lift_ret_val = uint64(lift_ret_ptr)
    lift_ret.Set(lift_ret_val)
  }
  return lift_ret
}

func WasiIoStreamsBlockingWrite(this uint32, buf []uint8) Result[uint64, WasiIoStreamsStreamError] {
  var lower_this C.uint32_t
  lower_this_val := C.uint32_t(this)
  lower_this = lower_this_val
  var lower_buf C.proxy_list_u8_t
  if len(buf) == 0 {
    lower_buf.ptr = nil
    lower_buf.len = 0
  } else {
    var empty_lower_buf C.uint8_t
    lower_buf.ptr = (*C.uint8_t)(C.malloc(C.size_t(len(buf)) * C.size_t(unsafe.Sizeof(empty_lower_buf))))
    lower_buf.len = C.size_t(len(buf))
    for lower_buf_i := range buf {
      lower_buf_ptr := (*C.uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(lower_buf.ptr)) +
      uintptr(lower_buf_i)*unsafe.Sizeof(empty_lower_buf)))
      lower_buf_ptr_value := C.uint8_t(buf[lower_buf_i])
      *lower_buf_ptr = lower_buf_ptr_value
    }
  }
  defer C.proxy_list_u8_free(&lower_buf)
  var ret C.proxy_result_u64_stream_error_t
  C.wasi_io_streams_blocking_write(lower_this, &lower_buf, &ret)
  var lift_ret Result[uint64, WasiIoStreamsStreamError]
  if ret.is_err {
    var lift_ret_val WasiIoStreamsStreamError
    lift_ret.SetErr(lift_ret_val)
  } else {
    lift_ret_ptr := *(*C.uint64_t)(unsafe.Pointer(&ret.val))
    var lift_ret_val uint64
    lift_ret_val = uint64(lift_ret_ptr)
    lift_ret.Set(lift_ret_val)
  }
  return lift_ret
}

func WasiIoStreamsWriteZeroes(this uint32, len uint64) Result[uint64, WasiIoStreamsStreamError] {
  var lower_this C.uint32_t
  lower_this_val := C.uint32_t(this)
  lower_this = lower_this_val
  lower_len := C.uint64_t(len)
  var ret C.proxy_result_u64_stream_error_t
  C.wasi_io_streams_write_zeroes(lower_this, lower_len, &ret)
  var lift_ret Result[uint64, WasiIoStreamsStreamError]
  if ret.is_err {
    var lift_ret_val WasiIoStreamsStreamError
    lift_ret.SetErr(lift_ret_val)
  } else {
    lift_ret_ptr := *(*C.uint64_t)(unsafe.Pointer(&ret.val))
    var lift_ret_val uint64
    lift_ret_val = uint64(lift_ret_ptr)
    lift_ret.Set(lift_ret_val)
  }
  return lift_ret
}

func WasiIoStreamsBlockingWriteZeroes(this uint32, len uint64) Result[uint64, WasiIoStreamsStreamError] {
  var lower_this C.uint32_t
  lower_this_val := C.uint32_t(this)
  lower_this = lower_this_val
  lower_len := C.uint64_t(len)
  var ret C.proxy_result_u64_stream_error_t
  C.wasi_io_streams_blocking_write_zeroes(lower_this, lower_len, &ret)
  var lift_ret Result[uint64, WasiIoStreamsStreamError]
  if ret.is_err {
    var lift_ret_val WasiIoStreamsStreamError
    lift_ret.SetErr(lift_ret_val)
  } else {
    lift_ret_ptr := *(*C.uint64_t)(unsafe.Pointer(&ret.val))
    var lift_ret_val uint64
    lift_ret_val = uint64(lift_ret_ptr)
    lift_ret.Set(lift_ret_val)
  }
  return lift_ret
}

func WasiIoStreamsSplice(this uint32, src uint32, len uint64) Result[WasiIoStreamsTuple2U64BoolT, WasiIoStreamsStreamError] {
  var lower_this C.uint32_t
  lower_this_val := C.uint32_t(this)
  lower_this = lower_this_val
  var lower_src C.uint32_t
  lower_src_val := C.uint32_t(src)
  lower_src = lower_src_val
  lower_len := C.uint64_t(len)
  var ret C.proxy_result_tuple2_u64_bool_stream_error_t
  C.wasi_io_streams_splice(lower_this, lower_src, lower_len, &ret)
  var lift_ret Result[WasiIoStreamsTuple2U64BoolT, WasiIoStreamsStreamError]
  if ret.is_err {
    var lift_ret_val WasiIoStreamsStreamError
    lift_ret.SetErr(lift_ret_val)
  } else {
    lift_ret_ptr := *(*C.proxy_tuple2_u64_bool_t)(unsafe.Pointer(&ret.val))
    var lift_ret_val WasiIoStreamsTuple2U64BoolT
    var lift_ret_val_F0 uint64
    lift_ret_val_F0 = uint64(lift_ret_ptr.f0)
    lift_ret_val.F0 = lift_ret_val_F0
    lift_ret_val_F1 := lift_ret_ptr.f1
    lift_ret_val.F1 = lift_ret_val_F1
    lift_ret.Set(lift_ret_val)
  }
  return lift_ret
}

func WasiIoStreamsBlockingSplice(this uint32, src uint32, len uint64) Result[WasiIoStreamsTuple2U64BoolT, WasiIoStreamsStreamError] {
  var lower_this C.uint32_t
  lower_this_val := C.uint32_t(this)
  lower_this = lower_this_val
  var lower_src C.uint32_t
  lower_src_val := C.uint32_t(src)
  lower_src = lower_src_val
  lower_len := C.uint64_t(len)
  var ret C.proxy_result_tuple2_u64_bool_stream_error_t
  C.wasi_io_streams_blocking_splice(lower_this, lower_src, lower_len, &ret)
  var lift_ret Result[WasiIoStreamsTuple2U64BoolT, WasiIoStreamsStreamError]
  if ret.is_err {
    var lift_ret_val WasiIoStreamsStreamError
    lift_ret.SetErr(lift_ret_val)
  } else {
    lift_ret_ptr := *(*C.proxy_tuple2_u64_bool_t)(unsafe.Pointer(&ret.val))
    var lift_ret_val WasiIoStreamsTuple2U64BoolT
    var lift_ret_val_F0 uint64
    lift_ret_val_F0 = uint64(lift_ret_ptr.f0)
    lift_ret_val.F0 = lift_ret_val_F0
    lift_ret_val_F1 := lift_ret_ptr.f1
    lift_ret_val.F1 = lift_ret_val_F1
    lift_ret.Set(lift_ret_val)
  }
  return lift_ret
}

func WasiIoStreamsForward(this uint32, src uint32) Result[uint64, WasiIoStreamsStreamError] {
  var lower_this C.uint32_t
  lower_this_val := C.uint32_t(this)
  lower_this = lower_this_val
  var lower_src C.uint32_t
  lower_src_val := C.uint32_t(src)
  lower_src = lower_src_val
  var ret C.proxy_result_u64_stream_error_t
  C.wasi_io_streams_forward(lower_this, lower_src, &ret)
  var lift_ret Result[uint64, WasiIoStreamsStreamError]
  if ret.is_err {
    var lift_ret_val WasiIoStreamsStreamError
    lift_ret.SetErr(lift_ret_val)
  } else {
    lift_ret_ptr := *(*C.uint64_t)(unsafe.Pointer(&ret.val))
    var lift_ret_val uint64
    lift_ret_val = uint64(lift_ret_ptr)
    lift_ret.Set(lift_ret_val)
  }
  return lift_ret
}

func WasiIoStreamsSubscribeToOutputStream(this uint32) uint32 {
  var lower_this C.uint32_t
  lower_this_val := C.uint32_t(this)
  lower_this = lower_this_val
  ret := C.wasi_io_streams_subscribe_to_output_stream(lower_this)
  var lift_ret uint32
  var lift_ret_val uint32
  var lift_ret_val_val uint32
  lift_ret_val_val = uint32(ret)
  lift_ret_val = lift_ret_val_val
  lift_ret = lift_ret_val
  return lift_ret
}

func WasiIoStreamsDropOutputStream(this uint32) {
  var lower_this C.uint32_t
  lower_this_val := C.uint32_t(this)
  lower_this = lower_this_val
  C.wasi_io_streams_drop_output_stream(lower_this)
}

// Import functions from wasi:cli/stdout
func WasiCliStdoutGetStdout() uint32 {
  ret := C.wasi_cli_stdout_get_stdout()
  var lift_ret uint32
  var lift_ret_val uint32
  var lift_ret_val_val uint32
  lift_ret_val_val = uint32(ret)
  lift_ret_val = lift_ret_val_val
  lift_ret = lift_ret_val
  return lift_ret
}

// Import functions from wasi:cli/stderr
func WasiCliStderrGetStderr() uint32 {
  ret := C.wasi_cli_stderr_get_stderr()
  var lift_ret uint32
  var lift_ret_val uint32
  var lift_ret_val_val uint32
  lift_ret_val_val = uint32(ret)
  lift_ret_val = lift_ret_val_val
  lift_ret = lift_ret_val
  return lift_ret
}

// Import functions from wasi:cli/stdin
func WasiCliStdinGetStdin() uint32 {
  ret := C.wasi_cli_stdin_get_stdin()
  var lift_ret uint32
  var lift_ret_val uint32
  var lift_ret_val_val uint32
  lift_ret_val_val = uint32(ret)
  lift_ret_val = lift_ret_val_val
  lift_ret = lift_ret_val
  return lift_ret
}

// Import functions from wasi:http/types
func WasiHttpTypesDropFields(fields uint32) {
  var lower_fields C.uint32_t
  lower_fields_val := C.uint32_t(fields)
  lower_fields = lower_fields_val
  C.wasi_http_types_drop_fields(lower_fields)
}

func WasiHttpTypesNewFields(entries []WasiHttpTypesTuple2StringStringT) uint32 {
  var lower_entries C.proxy_list_tuple2_string_string_t
  if len(entries) == 0 {
    lower_entries.ptr = nil
    lower_entries.len = 0
  } else {
    var empty_lower_entries C.proxy_tuple2_string_string_t
    lower_entries.ptr = (*C.proxy_tuple2_string_string_t)(C.malloc(C.size_t(len(entries)) * C.size_t(unsafe.Sizeof(empty_lower_entries))))
    lower_entries.len = C.size_t(len(entries))
    for lower_entries_i := range entries {
      lower_entries_ptr := (*C.proxy_tuple2_string_string_t)(unsafe.Pointer(uintptr(unsafe.Pointer(lower_entries.ptr)) +
      uintptr(lower_entries_i)*unsafe.Sizeof(empty_lower_entries)))
      var lower_entries_ptr_value C.proxy_tuple2_string_string_t
      var lower_entries_ptr_value_f0 C.proxy_string_t
      
      lower_entries_ptr_value_f0.ptr = C.CString(entries[lower_entries_i].F0)
      lower_entries_ptr_value_f0.len = C.size_t(len(entries[lower_entries_i].F0))
      lower_entries_ptr_value.f0 = lower_entries_ptr_value_f0
      var lower_entries_ptr_value_f1 C.proxy_string_t
      
      lower_entries_ptr_value_f1.ptr = C.CString(entries[lower_entries_i].F1)
      lower_entries_ptr_value_f1.len = C.size_t(len(entries[lower_entries_i].F1))
      lower_entries_ptr_value.f1 = lower_entries_ptr_value_f1
      *lower_entries_ptr = lower_entries_ptr_value
    }
  }
  defer C.proxy_list_tuple2_string_string_free(&lower_entries)
  ret := C.wasi_http_types_new_fields(&lower_entries)
  var lift_ret uint32
  var lift_ret_val uint32
  lift_ret_val = uint32(ret)
  lift_ret = lift_ret_val
  return lift_ret
}

func WasiHttpTypesFieldsGet(fields uint32, name string) [][]uint8 {
  var lower_fields C.uint32_t
  lower_fields_val := C.uint32_t(fields)
  lower_fields = lower_fields_val
  var lower_name C.proxy_string_t
  
  lower_name.ptr = C.CString(name)
  lower_name.len = C.size_t(len(name))
  defer C.proxy_string_free(&lower_name)
  var ret C.proxy_list_list_u8_t
  C.wasi_http_types_fields_get(lower_fields, &lower_name, &ret)
  var lift_ret [][]uint8
  lift_ret = make([][]uint8, ret.len)
  if ret.len > 0 {
    for lift_ret_i := 0; lift_ret_i < int(ret.len); lift_ret_i++ {
      var empty_lift_ret C.proxy_list_u8_t
      lift_ret_ptr := *(*C.proxy_list_u8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(ret.ptr)) +
      uintptr(lift_ret_i)*unsafe.Sizeof(empty_lift_ret)))
      var list_lift_ret []uint8
      list_lift_ret = make([]uint8, lift_ret_ptr.len)
      if lift_ret_ptr.len > 0 {
        for list_lift_ret_i := 0; list_lift_ret_i < int(lift_ret_ptr.len); list_lift_ret_i++ {
          var empty_list_lift_ret C.uint8_t
          list_lift_ret_ptr := *(*C.uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(lift_ret_ptr.ptr)) +
          uintptr(list_lift_ret_i)*unsafe.Sizeof(empty_list_lift_ret)))
          var list_list_lift_ret uint8
          list_list_lift_ret = uint8(list_lift_ret_ptr)
          list_lift_ret[list_lift_ret_i] = list_list_lift_ret
        }
      }
      lift_ret[lift_ret_i] = list_lift_ret
    }
  }
  return lift_ret
}

func WasiHttpTypesFieldsSet(fields uint32, name string, value [][]uint8) {
  var lower_fields C.uint32_t
  lower_fields_val := C.uint32_t(fields)
  lower_fields = lower_fields_val
  var lower_name C.proxy_string_t
  
  lower_name.ptr = C.CString(name)
  lower_name.len = C.size_t(len(name))
  defer C.proxy_string_free(&lower_name)
  var lower_value C.proxy_list_list_u8_t
  if len(value) == 0 {
    lower_value.ptr = nil
    lower_value.len = 0
  } else {
    var empty_lower_value C.proxy_list_u8_t
    lower_value.ptr = (*C.proxy_list_u8_t)(C.malloc(C.size_t(len(value)) * C.size_t(unsafe.Sizeof(empty_lower_value))))
    lower_value.len = C.size_t(len(value))
    for lower_value_i := range value {
      lower_value_ptr := (*C.proxy_list_u8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(lower_value.ptr)) +
      uintptr(lower_value_i)*unsafe.Sizeof(empty_lower_value)))
      if len(value[lower_value_i]) == 0 {
        lower_value_ptr.ptr = nil
        lower_value_ptr.len = 0
      } else {
        var empty_lower_value_ptr C.uint8_t
        lower_value_ptr.ptr = (*C.uint8_t)(C.malloc(C.size_t(len(value[lower_value_i])) * C.size_t(unsafe.Sizeof(empty_lower_value_ptr))))
        lower_value_ptr.len = C.size_t(len(value[lower_value_i]))
        for lower_value_ptr_i := range value[lower_value_i] {
          lower_value_ptr_ptr := (*C.uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(lower_value_ptr.ptr)) +
          uintptr(lower_value_ptr_i)*unsafe.Sizeof(empty_lower_value_ptr)))
          lower_value_ptr_ptr_value := C.uint8_t(value[lower_value_i][lower_value_ptr_i])
          *lower_value_ptr_ptr = lower_value_ptr_ptr_value
        }
      }
    }
  }
  defer C.proxy_list_list_u8_free(&lower_value)
  C.wasi_http_types_fields_set(lower_fields, &lower_name, &lower_value)
}

func WasiHttpTypesFieldsDelete(fields uint32, name string) {
  var lower_fields C.uint32_t
  lower_fields_val := C.uint32_t(fields)
  lower_fields = lower_fields_val
  var lower_name C.proxy_string_t
  
  lower_name.ptr = C.CString(name)
  lower_name.len = C.size_t(len(name))
  defer C.proxy_string_free(&lower_name)
  C.wasi_http_types_fields_delete(lower_fields, &lower_name)
}

func WasiHttpTypesFieldsAppend(fields uint32, name string, value []uint8) {
  var lower_fields C.uint32_t
  lower_fields_val := C.uint32_t(fields)
  lower_fields = lower_fields_val
  var lower_name C.proxy_string_t
  
  lower_name.ptr = C.CString(name)
  lower_name.len = C.size_t(len(name))
  defer C.proxy_string_free(&lower_name)
  var lower_value C.proxy_list_u8_t
  if len(value) == 0 {
    lower_value.ptr = nil
    lower_value.len = 0
  } else {
    var empty_lower_value C.uint8_t
    lower_value.ptr = (*C.uint8_t)(C.malloc(C.size_t(len(value)) * C.size_t(unsafe.Sizeof(empty_lower_value))))
    lower_value.len = C.size_t(len(value))
    for lower_value_i := range value {
      lower_value_ptr := (*C.uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(lower_value.ptr)) +
      uintptr(lower_value_i)*unsafe.Sizeof(empty_lower_value)))
      lower_value_ptr_value := C.uint8_t(value[lower_value_i])
      *lower_value_ptr = lower_value_ptr_value
    }
  }
  defer C.proxy_list_u8_free(&lower_value)
  C.wasi_http_types_fields_append(lower_fields, &lower_name, &lower_value)
}

func WasiHttpTypesFieldsEntries(fields uint32) []WasiHttpTypesTuple2StringListU8TT {
  var lower_fields C.uint32_t
  lower_fields_val := C.uint32_t(fields)
  lower_fields = lower_fields_val
  var ret C.proxy_list_tuple2_string_list_u8_t
  C.wasi_http_types_fields_entries(lower_fields, &ret)
  var lift_ret []WasiHttpTypesTuple2StringListU8TT
  lift_ret = make([]WasiHttpTypesTuple2StringListU8TT, ret.len)
  if ret.len > 0 {
    for lift_ret_i := 0; lift_ret_i < int(ret.len); lift_ret_i++ {
      var empty_lift_ret C.proxy_tuple2_string_list_u8_t
      lift_ret_ptr := *(*C.proxy_tuple2_string_list_u8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(ret.ptr)) +
      uintptr(lift_ret_i)*unsafe.Sizeof(empty_lift_ret)))
      var list_lift_ret WasiHttpTypesTuple2StringListU8TT
      var list_lift_ret_F0 string
      list_lift_ret_F0 = C.GoStringN(lift_ret_ptr.f0.ptr, C.int(lift_ret_ptr.f0.len))
      list_lift_ret.F0 = list_lift_ret_F0
      var list_lift_ret_F1 []uint8
      list_lift_ret_F1 = make([]uint8, lift_ret_ptr.f1.len)
      if lift_ret_ptr.f1.len > 0 {
        for list_lift_ret_F1_i := 0; list_lift_ret_F1_i < int(lift_ret_ptr.f1.len); list_lift_ret_F1_i++ {
          var empty_list_lift_ret_F1 C.uint8_t
          list_lift_ret_F1_ptr := *(*C.uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(lift_ret_ptr.f1.ptr)) +
          uintptr(list_lift_ret_F1_i)*unsafe.Sizeof(empty_list_lift_ret_F1)))
          var list_list_lift_ret_F1 uint8
          list_list_lift_ret_F1 = uint8(list_lift_ret_F1_ptr)
          list_lift_ret_F1[list_lift_ret_F1_i] = list_list_lift_ret_F1
        }
      }
      list_lift_ret.F1 = list_lift_ret_F1
      lift_ret[lift_ret_i] = list_lift_ret
    }
  }
  return lift_ret
}

func WasiHttpTypesFieldsClone(fields uint32) uint32 {
  var lower_fields C.uint32_t
  lower_fields_val := C.uint32_t(fields)
  lower_fields = lower_fields_val
  ret := C.wasi_http_types_fields_clone(lower_fields)
  var lift_ret uint32
  var lift_ret_val uint32
  lift_ret_val = uint32(ret)
  lift_ret = lift_ret_val
  return lift_ret
}

func WasiHttpTypesFinishIncomingStream(s uint32) Option[uint32] {
  var lower_s C.wasi_http_types_input_stream_t
  var lower_s_val C.wasi_io_streams_input_stream_t
  var lower_s_val_val C.uint32_t
  lower_s_val_val_val := C.uint32_t(s)
  lower_s_val_val = lower_s_val_val_val
  lower_s_val = lower_s_val_val
  lower_s = lower_s_val
  var ret C.proxy_option_trailers_t
  C.wasi_http_types_finish_incoming_stream(lower_s, &ret)
  var lift_ret Option[uint32]
  if ret.is_some {
    var lift_ret_val uint32
    var lift_ret_val_val uint32
    var lift_ret_val_val_val uint32
    lift_ret_val_val_val = uint32(ret.val)
    lift_ret_val_val = lift_ret_val_val_val
    lift_ret_val = lift_ret_val_val
    lift_ret.Set(lift_ret_val)
  } else {
    lift_ret.Unset()
  }
  return lift_ret
}

func WasiHttpTypesFinishOutgoingStream(s uint32, trailers Option[uint32]) {
  var lower_s C.wasi_http_types_output_stream_t
  var lower_s_val C.wasi_io_streams_output_stream_t
  var lower_s_val_val C.uint32_t
  lower_s_val_val_val := C.uint32_t(s)
  lower_s_val_val = lower_s_val_val_val
  lower_s_val = lower_s_val_val
  lower_s = lower_s_val
  var lower_trailers C.proxy_option_trailers_t
  if trailers.IsSome() {
    var lower_trailers_val C.wasi_http_types_fields_t
    var lower_trailers_val_val C.uint32_t
    lower_trailers_val_val_val := C.uint32_t(trailers.Unwrap())
    lower_trailers_val_val = lower_trailers_val_val_val
    lower_trailers_val = lower_trailers_val_val
    lower_trailers.val = lower_trailers_val
    lower_trailers.is_some = true
  }
  C.wasi_http_types_finish_outgoing_stream(lower_s, &lower_trailers)
}

func WasiHttpTypesDropIncomingRequest(request uint32) {
  var lower_request C.uint32_t
  lower_request_val := C.uint32_t(request)
  lower_request = lower_request_val
  C.wasi_http_types_drop_incoming_request(lower_request)
}

func WasiHttpTypesDropOutgoingRequest(request uint32) {
  var lower_request C.uint32_t
  lower_request_val := C.uint32_t(request)
  lower_request = lower_request_val
  C.wasi_http_types_drop_outgoing_request(lower_request)
}

func WasiHttpTypesIncomingRequestMethod(request uint32) WasiHttpTypesMethod {
  var lower_request C.uint32_t
  lower_request_val := C.uint32_t(request)
  lower_request = lower_request_val
  var ret C.wasi_http_types_method_t
  C.wasi_http_types_incoming_request_method(lower_request, &ret)
  var lift_ret WasiHttpTypesMethod
  if ret.tag == 0 {
    lift_ret = WasiHttpTypesMethodGet()
  }
  if ret.tag == 1 {
    lift_ret = WasiHttpTypesMethodHead()
  }
  if ret.tag == 2 {
    lift_ret = WasiHttpTypesMethodPost()
  }
  if ret.tag == 3 {
    lift_ret = WasiHttpTypesMethodPut()
  }
  if ret.tag == 4 {
    lift_ret = WasiHttpTypesMethodDelete()
  }
  if ret.tag == 5 {
    lift_ret = WasiHttpTypesMethodConnect()
  }
  if ret.tag == 6 {
    lift_ret = WasiHttpTypesMethodOptions()
  }
  if ret.tag == 7 {
    lift_ret = WasiHttpTypesMethodTrace()
  }
  if ret.tag == 8 {
    lift_ret = WasiHttpTypesMethodPatch()
  }
  if ret.tag == 9 {
    lift_ret_ptr := *(*C.proxy_string_t)(unsafe.Pointer(&ret.val))
    var lift_ret_val string
    lift_ret_val = C.GoStringN(lift_ret_ptr.ptr, C.int(lift_ret_ptr.len))
    lift_ret = WasiHttpTypesMethodOther(lift_ret_val)
  }
  return lift_ret
}

func WasiHttpTypesIncomingRequestPathWithQuery(request uint32) Option[string] {
  var lower_request C.uint32_t
  lower_request_val := C.uint32_t(request)
  lower_request = lower_request_val
  var ret C.proxy_option_string_t
  C.wasi_http_types_incoming_request_path_with_query(lower_request, &ret)
  var lift_ret Option[string]
  if ret.is_some {
    var lift_ret_val string
    lift_ret_val = C.GoStringN(ret.val.ptr, C.int(ret.val.len))
    lift_ret.Set(lift_ret_val)
  } else {
    lift_ret.Unset()
  }
  return lift_ret
}

func WasiHttpTypesIncomingRequestScheme(request uint32) Option[WasiHttpTypesScheme] {
  var lower_request C.uint32_t
  lower_request_val := C.uint32_t(request)
  lower_request = lower_request_val
  var ret C.proxy_option_scheme_t
  C.wasi_http_types_incoming_request_scheme(lower_request, &ret)
  var lift_ret Option[WasiHttpTypesScheme]
  if ret.is_some {
    var lift_ret_val WasiHttpTypesScheme
    if ret.val.tag == 0 {
      lift_ret_val = WasiHttpTypesSchemeHttp()
    }
    if ret.val.tag == 1 {
      lift_ret_val = WasiHttpTypesSchemeHttps()
    }
    if ret.val.tag == 2 {
      lift_ret_val_ptr := *(*C.proxy_string_t)(unsafe.Pointer(&ret.val.val))
      var lift_ret_val_val string
      lift_ret_val_val = C.GoStringN(lift_ret_val_ptr.ptr, C.int(lift_ret_val_ptr.len))
      lift_ret_val = WasiHttpTypesSchemeOther(lift_ret_val_val)
    }
    lift_ret.Set(lift_ret_val)
  } else {
    lift_ret.Unset()
  }
  return lift_ret
}

func WasiHttpTypesIncomingRequestAuthority(request uint32) Option[string] {
  var lower_request C.uint32_t
  lower_request_val := C.uint32_t(request)
  lower_request = lower_request_val
  var ret C.proxy_option_string_t
  C.wasi_http_types_incoming_request_authority(lower_request, &ret)
  var lift_ret Option[string]
  if ret.is_some {
    var lift_ret_val string
    lift_ret_val = C.GoStringN(ret.val.ptr, C.int(ret.val.len))
    lift_ret.Set(lift_ret_val)
  } else {
    lift_ret.Unset()
  }
  return lift_ret
}

func WasiHttpTypesIncomingRequestHeaders(request uint32) uint32 {
  var lower_request C.uint32_t
  lower_request_val := C.uint32_t(request)
  lower_request = lower_request_val
  ret := C.wasi_http_types_incoming_request_headers(lower_request)
  var lift_ret uint32
  var lift_ret_val uint32
  var lift_ret_val_val uint32
  lift_ret_val_val = uint32(ret)
  lift_ret_val = lift_ret_val_val
  lift_ret = lift_ret_val
  return lift_ret
}

func WasiHttpTypesIncomingRequestConsume(request uint32) Result[uint32, struct{}] {
  var lower_request C.uint32_t
  lower_request_val := C.uint32_t(request)
  lower_request = lower_request_val
  var ret C.proxy_result_incoming_stream_void_t
  C.wasi_http_types_incoming_request_consume(lower_request, &ret)
  var lift_ret Result[uint32, struct{}]
  if ret.is_err {
    lift_ret.SetErr(struct{}{})
  } else {
    lift_ret_ptr := *(*C.wasi_http_types_incoming_stream_t)(unsafe.Pointer(&ret.val))
    var lift_ret_val uint32
    var lift_ret_val_val uint32
    var lift_ret_val_val_val uint32
    var lift_ret_val_val_val_val uint32
    lift_ret_val_val_val_val = uint32(lift_ret_ptr)
    lift_ret_val_val_val = lift_ret_val_val_val_val
    lift_ret_val_val = lift_ret_val_val_val
    lift_ret_val = lift_ret_val_val
    lift_ret.Set(lift_ret_val)
  }
  return lift_ret
}

func WasiHttpTypesNewOutgoingRequest(method WasiHttpTypesMethod, path_with_query Option[string], scheme Option[WasiHttpTypesScheme], authority Option[string], headers uint32) uint32 {
  var lower_method C.wasi_http_types_method_t
  if method.Kind() == WasiHttpTypesMethodKindGet {
    lower_method.tag = 0
  }
  if method.Kind() == WasiHttpTypesMethodKindHead {
    lower_method.tag = 1
  }
  if method.Kind() == WasiHttpTypesMethodKindPost {
    lower_method.tag = 2
  }
  if method.Kind() == WasiHttpTypesMethodKindPut {
    lower_method.tag = 3
  }
  if method.Kind() == WasiHttpTypesMethodKindDelete {
    lower_method.tag = 4
  }
  if method.Kind() == WasiHttpTypesMethodKindConnect {
    lower_method.tag = 5
  }
  if method.Kind() == WasiHttpTypesMethodKindOptions {
    lower_method.tag = 6
  }
  if method.Kind() == WasiHttpTypesMethodKindTrace {
    lower_method.tag = 7
  }
  if method.Kind() == WasiHttpTypesMethodKindPatch {
    lower_method.tag = 8
  }
  if method.Kind() == WasiHttpTypesMethodKindOther {
    
    lower_method.tag = 9
    lower_method_ptr := (*C.proxy_string_t)(unsafe.Pointer(&lower_method.val))
    var lower_method_val C.proxy_string_t
    
    lower_method_val.ptr = C.CString(method.GetOther())
    lower_method_val.len = C.size_t(len(method.GetOther()))
    *lower_method_ptr = lower_method_val
  }
  defer C.wasi_http_types_method_free(&lower_method)
  var lower_path_with_query C.proxy_option_string_t
  if path_with_query.IsSome() {
    var lower_path_with_query_val C.proxy_string_t
    
    lower_path_with_query_val.ptr = C.CString(path_with_query.Unwrap())
    lower_path_with_query_val.len = C.size_t(len(path_with_query.Unwrap()))
    lower_path_with_query.val = lower_path_with_query_val
    lower_path_with_query.is_some = true
  }
  defer C.proxy_option_string_free(&lower_path_with_query)
  var lower_scheme C.proxy_option_scheme_t
  if scheme.IsSome() {
    var lower_scheme_val C.wasi_http_types_scheme_t
    if scheme.Unwrap().Kind() == WasiHttpTypesSchemeKindHttp {
      lower_scheme_val.tag = 0
    }
    if scheme.Unwrap().Kind() == WasiHttpTypesSchemeKindHttps {
      lower_scheme_val.tag = 1
    }
    if scheme.Unwrap().Kind() == WasiHttpTypesSchemeKindOther {
      
      lower_scheme_val.tag = 2
      lower_scheme_val_ptr := (*C.proxy_string_t)(unsafe.Pointer(&lower_scheme_val.val))
      var lower_scheme_val_val C.proxy_string_t
      
      lower_scheme_val_val.ptr = C.CString(scheme.Unwrap().GetOther())
      lower_scheme_val_val.len = C.size_t(len(scheme.Unwrap().GetOther()))
      *lower_scheme_val_ptr = lower_scheme_val_val
    }
    lower_scheme.val = lower_scheme_val
    lower_scheme.is_some = true
  }
  defer C.proxy_option_scheme_free(&lower_scheme)
  var lower_authority C.proxy_option_string_t
  if authority.IsSome() {
    var lower_authority_val C.proxy_string_t
    
    lower_authority_val.ptr = C.CString(authority.Unwrap())
    lower_authority_val.len = C.size_t(len(authority.Unwrap()))
    lower_authority.val = lower_authority_val
    lower_authority.is_some = true
  }
  defer C.proxy_option_string_free(&lower_authority)
  var lower_headers C.wasi_http_types_fields_t
  var lower_headers_val C.uint32_t
  lower_headers_val_val := C.uint32_t(headers)
  lower_headers_val = lower_headers_val_val
  lower_headers = lower_headers_val
  ret := C.wasi_http_types_new_outgoing_request(&lower_method, &lower_path_with_query, &lower_scheme, &lower_authority, lower_headers)
  var lift_ret uint32
  var lift_ret_val uint32
  lift_ret_val = uint32(ret)
  lift_ret = lift_ret_val
  return lift_ret
}

func WasiHttpTypesOutgoingRequestWrite(request uint32) Result[uint32, struct{}] {
  var lower_request C.uint32_t
  lower_request_val := C.uint32_t(request)
  lower_request = lower_request_val
  var ret C.proxy_result_outgoing_stream_void_t
  C.wasi_http_types_outgoing_request_write(lower_request, &ret)
  var lift_ret Result[uint32, struct{}]
  if ret.is_err {
    lift_ret.SetErr(struct{}{})
  } else {
    lift_ret_ptr := *(*C.wasi_http_types_outgoing_stream_t)(unsafe.Pointer(&ret.val))
    var lift_ret_val uint32
    var lift_ret_val_val uint32
    var lift_ret_val_val_val uint32
    var lift_ret_val_val_val_val uint32
    lift_ret_val_val_val_val = uint32(lift_ret_ptr)
    lift_ret_val_val_val = lift_ret_val_val_val_val
    lift_ret_val_val = lift_ret_val_val_val
    lift_ret_val = lift_ret_val_val
    lift_ret.Set(lift_ret_val)
  }
  return lift_ret
}

func WasiHttpTypesDropResponseOutparam(response uint32) {
  var lower_response C.uint32_t
  lower_response_val := C.uint32_t(response)
  lower_response = lower_response_val
  C.wasi_http_types_drop_response_outparam(lower_response)
}

func WasiHttpTypesSetResponseOutparam(param uint32, response Result[uint32, WasiHttpTypesError]) Result[struct{}, struct{}] {
  var lower_param C.uint32_t
  lower_param_val := C.uint32_t(param)
  lower_param = lower_param_val
  var lower_response C.proxy_result_outgoing_response_error_t
  lower_response.is_err = response.IsErr()
  if response.IsOk() {
    lower_response_ptr := (*C.wasi_http_types_outgoing_response_t)(unsafe.Pointer(&lower_response.val))
    var lower_response_val C.uint32_t
    lower_response_val_val := C.uint32_t(response.Unwrap())
    lower_response_val = lower_response_val_val
    *lower_response_ptr = lower_response_val
  } else {
    lower_response_ptr := (*C.wasi_http_types_error_t)(unsafe.Pointer(&lower_response.val))
    var lower_response_val C.wasi_http_types_error_t
    if response.UnwrapErr().Kind() == WasiHttpTypesErrorKindInvalidUrl {
      
      lower_response_val.tag = 0
      lower_response_val_ptr := (*C.proxy_string_t)(unsafe.Pointer(&lower_response_val.val))
      var lower_response_val_val C.proxy_string_t
      
      lower_response_val_val.ptr = C.CString(response.UnwrapErr().GetInvalidUrl())
      lower_response_val_val.len = C.size_t(len(response.UnwrapErr().GetInvalidUrl()))
      *lower_response_val_ptr = lower_response_val_val
    }
    if response.UnwrapErr().Kind() == WasiHttpTypesErrorKindTimeoutError {
      
      lower_response_val.tag = 1
      lower_response_val_ptr := (*C.proxy_string_t)(unsafe.Pointer(&lower_response_val.val))
      var lower_response_val_val C.proxy_string_t
      
      lower_response_val_val.ptr = C.CString(response.UnwrapErr().GetTimeoutError())
      lower_response_val_val.len = C.size_t(len(response.UnwrapErr().GetTimeoutError()))
      *lower_response_val_ptr = lower_response_val_val
    }
    if response.UnwrapErr().Kind() == WasiHttpTypesErrorKindProtocolError {
      
      lower_response_val.tag = 2
      lower_response_val_ptr := (*C.proxy_string_t)(unsafe.Pointer(&lower_response_val.val))
      var lower_response_val_val C.proxy_string_t
      
      lower_response_val_val.ptr = C.CString(response.UnwrapErr().GetProtocolError())
      lower_response_val_val.len = C.size_t(len(response.UnwrapErr().GetProtocolError()))
      *lower_response_val_ptr = lower_response_val_val
    }
    if response.UnwrapErr().Kind() == WasiHttpTypesErrorKindUnexpectedError {
      
      lower_response_val.tag = 3
      lower_response_val_ptr := (*C.proxy_string_t)(unsafe.Pointer(&lower_response_val.val))
      var lower_response_val_val C.proxy_string_t
      
      lower_response_val_val.ptr = C.CString(response.UnwrapErr().GetUnexpectedError())
      lower_response_val_val.len = C.size_t(len(response.UnwrapErr().GetUnexpectedError()))
      *lower_response_val_ptr = lower_response_val_val
    }
    *lower_response_ptr = lower_response_val
  }
  defer C.proxy_result_outgoing_response_error_free(&lower_response)
  var ret C.proxy_result_void_void_t
  C.wasi_http_types_set_response_outparam(lower_param, &lower_response, &ret)
  var lift_ret Result[struct{}, struct{}]
  if ret.is_err {
    lift_ret.SetErr(struct{}{})
  } else {
  }
  return lift_ret
}

func WasiHttpTypesDropIncomingResponse(response uint32) {
  var lower_response C.uint32_t
  lower_response_val := C.uint32_t(response)
  lower_response = lower_response_val
  C.wasi_http_types_drop_incoming_response(lower_response)
}

func WasiHttpTypesDropOutgoingResponse(response uint32) {
  var lower_response C.uint32_t
  lower_response_val := C.uint32_t(response)
  lower_response = lower_response_val
  C.wasi_http_types_drop_outgoing_response(lower_response)
}

func WasiHttpTypesIncomingResponseStatus(response uint32) uint16 {
  var lower_response C.uint32_t
  lower_response_val := C.uint32_t(response)
  lower_response = lower_response_val
  ret := C.wasi_http_types_incoming_response_status(lower_response)
  var lift_ret uint16
  var lift_ret_val uint16
  lift_ret_val = uint16(ret)
  lift_ret = lift_ret_val
  return lift_ret
}

func WasiHttpTypesIncomingResponseHeaders(response uint32) uint32 {
  var lower_response C.uint32_t
  lower_response_val := C.uint32_t(response)
  lower_response = lower_response_val
  ret := C.wasi_http_types_incoming_response_headers(lower_response)
  var lift_ret uint32
  var lift_ret_val uint32
  var lift_ret_val_val uint32
  lift_ret_val_val = uint32(ret)
  lift_ret_val = lift_ret_val_val
  lift_ret = lift_ret_val
  return lift_ret
}

func WasiHttpTypesIncomingResponseConsume(response uint32) Result[uint32, struct{}] {
  var lower_response C.uint32_t
  lower_response_val := C.uint32_t(response)
  lower_response = lower_response_val
  var ret C.proxy_result_incoming_stream_void_t
  C.wasi_http_types_incoming_response_consume(lower_response, &ret)
  var lift_ret Result[uint32, struct{}]
  if ret.is_err {
    lift_ret.SetErr(struct{}{})
  } else {
    lift_ret_ptr := *(*C.wasi_http_types_incoming_stream_t)(unsafe.Pointer(&ret.val))
    var lift_ret_val uint32
    var lift_ret_val_val uint32
    var lift_ret_val_val_val uint32
    var lift_ret_val_val_val_val uint32
    lift_ret_val_val_val_val = uint32(lift_ret_ptr)
    lift_ret_val_val_val = lift_ret_val_val_val_val
    lift_ret_val_val = lift_ret_val_val_val
    lift_ret_val = lift_ret_val_val
    lift_ret.Set(lift_ret_val)
  }
  return lift_ret
}

func WasiHttpTypesNewOutgoingResponse(status_code uint16, headers uint32) uint32 {
  var lower_status_code C.uint16_t
  lower_status_code_val := C.uint16_t(status_code)
  lower_status_code = lower_status_code_val
  var lower_headers C.wasi_http_types_fields_t
  var lower_headers_val C.uint32_t
  lower_headers_val_val := C.uint32_t(headers)
  lower_headers_val = lower_headers_val_val
  lower_headers = lower_headers_val
  ret := C.wasi_http_types_new_outgoing_response(lower_status_code, lower_headers)
  var lift_ret uint32
  var lift_ret_val uint32
  lift_ret_val = uint32(ret)
  lift_ret = lift_ret_val
  return lift_ret
}

func WasiHttpTypesOutgoingResponseWrite(response uint32) Result[uint32, struct{}] {
  var lower_response C.uint32_t
  lower_response_val := C.uint32_t(response)
  lower_response = lower_response_val
  var ret C.proxy_result_outgoing_stream_void_t
  C.wasi_http_types_outgoing_response_write(lower_response, &ret)
  var lift_ret Result[uint32, struct{}]
  if ret.is_err {
    lift_ret.SetErr(struct{}{})
  } else {
    lift_ret_ptr := *(*C.wasi_http_types_outgoing_stream_t)(unsafe.Pointer(&ret.val))
    var lift_ret_val uint32
    var lift_ret_val_val uint32
    var lift_ret_val_val_val uint32
    var lift_ret_val_val_val_val uint32
    lift_ret_val_val_val_val = uint32(lift_ret_ptr)
    lift_ret_val_val_val = lift_ret_val_val_val_val
    lift_ret_val_val = lift_ret_val_val_val
    lift_ret_val = lift_ret_val_val
    lift_ret.Set(lift_ret_val)
  }
  return lift_ret
}

func WasiHttpTypesDropFutureIncomingResponse(f uint32) {
  var lower_f C.uint32_t
  lower_f_val := C.uint32_t(f)
  lower_f = lower_f_val
  C.wasi_http_types_drop_future_incoming_response(lower_f)
}

func WasiHttpTypesFutureIncomingResponseGet(f uint32) Option[Result[uint32, WasiHttpTypesError]] {
  var lower_f C.uint32_t
  lower_f_val := C.uint32_t(f)
  lower_f = lower_f_val
  var ret C.proxy_option_result_incoming_response_error_t
  C.wasi_http_types_future_incoming_response_get(lower_f, &ret)
  var lift_ret Option[Result[uint32, WasiHttpTypesError]]
  if ret.is_some {
    var lift_ret_val Result[uint32, WasiHttpTypesError]
    if ret.val.is_err {
      lift_ret_val_ptr := *(*C.wasi_http_types_error_t)(unsafe.Pointer(&ret.val.val))
      var lift_ret_val_val WasiHttpTypesError
      if lift_ret_val_ptr.tag == 0 {
        lift_ret_val_val_ptr := *(*C.proxy_string_t)(unsafe.Pointer(&lift_ret_val_ptr.val))
        var lift_ret_val_val_val string
        lift_ret_val_val_val = C.GoStringN(lift_ret_val_val_ptr.ptr, C.int(lift_ret_val_val_ptr.len))
        lift_ret_val_val = WasiHttpTypesErrorInvalidUrl(lift_ret_val_val_val)
      }
      if lift_ret_val_ptr.tag == 1 {
        lift_ret_val_val_ptr := *(*C.proxy_string_t)(unsafe.Pointer(&lift_ret_val_ptr.val))
        var lift_ret_val_val_val string
        lift_ret_val_val_val = C.GoStringN(lift_ret_val_val_ptr.ptr, C.int(lift_ret_val_val_ptr.len))
        lift_ret_val_val = WasiHttpTypesErrorTimeoutError(lift_ret_val_val_val)
      }
      if lift_ret_val_ptr.tag == 2 {
        lift_ret_val_val_ptr := *(*C.proxy_string_t)(unsafe.Pointer(&lift_ret_val_ptr.val))
        var lift_ret_val_val_val string
        lift_ret_val_val_val = C.GoStringN(lift_ret_val_val_ptr.ptr, C.int(lift_ret_val_val_ptr.len))
        lift_ret_val_val = WasiHttpTypesErrorProtocolError(lift_ret_val_val_val)
      }
      if lift_ret_val_ptr.tag == 3 {
        lift_ret_val_val_ptr := *(*C.proxy_string_t)(unsafe.Pointer(&lift_ret_val_ptr.val))
        var lift_ret_val_val_val string
        lift_ret_val_val_val = C.GoStringN(lift_ret_val_val_ptr.ptr, C.int(lift_ret_val_val_ptr.len))
        lift_ret_val_val = WasiHttpTypesErrorUnexpectedError(lift_ret_val_val_val)
      }
      lift_ret_val.SetErr(lift_ret_val_val)
    } else {
      lift_ret_val_ptr := *(*C.wasi_http_types_incoming_response_t)(unsafe.Pointer(&ret.val.val))
      var lift_ret_val_val uint32
      var lift_ret_val_val_val uint32
      lift_ret_val_val_val = uint32(lift_ret_val_ptr)
      lift_ret_val_val = lift_ret_val_val_val
      lift_ret_val.Set(lift_ret_val_val)
    }
    lift_ret.Set(lift_ret_val)
  } else {
    lift_ret.Unset()
  }
  return lift_ret
}

func WasiHttpTypesListenToFutureIncomingResponse(f uint32) uint32 {
  var lower_f C.uint32_t
  lower_f_val := C.uint32_t(f)
  lower_f = lower_f_val
  ret := C.wasi_http_types_listen_to_future_incoming_response(lower_f)
  var lift_ret uint32
  var lift_ret_val uint32
  var lift_ret_val_val uint32
  lift_ret_val_val = uint32(ret)
  lift_ret_val = lift_ret_val_val
  lift_ret = lift_ret_val
  return lift_ret
}

// Import functions from wasi:http/outgoing-handler
func WasiHttpOutgoingHandlerHandle(request uint32, options Option[WasiHttpTypesRequestOptions]) uint32 {
  var lower_request C.wasi_http_types_outgoing_request_t
  var lower_request_val C.uint32_t
  lower_request_val_val := C.uint32_t(request)
  lower_request_val = lower_request_val_val
  lower_request = lower_request_val
  var lower_options C.proxy_option_request_options_t
  if options.IsSome() {
    var lower_options_val C.wasi_http_types_request_options_t
    var lower_options_val_val C.wasi_http_types_request_options_t
    var lower_options_val_val_connect_timeout_ms C.proxy_option_u32_t
    if options.Unwrap().ConnectTimeoutMs.IsSome() {
      lower_options_val_val_connect_timeout_ms_val := C.uint32_t(options.Unwrap().ConnectTimeoutMs.Unwrap())
      lower_options_val_val_connect_timeout_ms.val = lower_options_val_val_connect_timeout_ms_val
      lower_options_val_val_connect_timeout_ms.is_some = true
    }
    lower_options_val_val.connect_timeout_ms = lower_options_val_val_connect_timeout_ms
    var lower_options_val_val_first_byte_timeout_ms C.proxy_option_u32_t
    if options.Unwrap().FirstByteTimeoutMs.IsSome() {
      lower_options_val_val_first_byte_timeout_ms_val := C.uint32_t(options.Unwrap().FirstByteTimeoutMs.Unwrap())
      lower_options_val_val_first_byte_timeout_ms.val = lower_options_val_val_first_byte_timeout_ms_val
      lower_options_val_val_first_byte_timeout_ms.is_some = true
    }
    lower_options_val_val.first_byte_timeout_ms = lower_options_val_val_first_byte_timeout_ms
    var lower_options_val_val_between_bytes_timeout_ms C.proxy_option_u32_t
    if options.Unwrap().BetweenBytesTimeoutMs.IsSome() {
      lower_options_val_val_between_bytes_timeout_ms_val := C.uint32_t(options.Unwrap().BetweenBytesTimeoutMs.Unwrap())
      lower_options_val_val_between_bytes_timeout_ms.val = lower_options_val_val_between_bytes_timeout_ms_val
      lower_options_val_val_between_bytes_timeout_ms.is_some = true
    }
    lower_options_val_val.between_bytes_timeout_ms = lower_options_val_val_between_bytes_timeout_ms
    lower_options_val = lower_options_val_val
    lower_options.val = lower_options_val
    lower_options.is_some = true
  }
  ret := C.wasi_http_outgoing_handler_handle(lower_request, &lower_options)
  var lift_ret uint32
  var lift_ret_val uint32
  var lift_ret_val_val uint32
  lift_ret_val_val = uint32(ret)
  lift_ret_val = lift_ret_val_val
  lift_ret = lift_ret_val
  return lift_ret
}

// Export functions from wasi:http/incoming-handler
var wasi_http_incoming_handler ExportsWasiHttpIncomingHandler = nil
func SetExportsWasiHttpIncomingHandler(i ExportsWasiHttpIncomingHandler) {
  wasi_http_incoming_handler = i
}
type ExportsWasiHttpIncomingHandler interface {
  Handle(request uint32, response_out uint32) 
}
//export exports_wasi_http_incoming_handler_handle
func ExportsWasiHttpIncomingHandlerHandle(request C.wasi_http_incoming_handler_incoming_request_t, response_out C.wasi_http_incoming_handler_response_outparam_t) {
  var lift_request uint32
  var lift_request_val uint32
  var lift_request_val_val uint32
  lift_request_val_val = uint32(request)
  lift_request_val = lift_request_val_val
  lift_request = lift_request_val
  var lift_response_out uint32
  var lift_response_out_val uint32
  var lift_response_out_val_val uint32
  lift_response_out_val_val = uint32(response_out)
  lift_response_out_val = lift_response_out_val_val
  lift_response_out = lift_response_out_val
  wasi_http_incoming_handler.Handle(lift_request, lift_response_out)
  
}
