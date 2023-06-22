package proxy

// #include "proxy.h"
import "C"

import "unsafe"

import "fmt"

// random
func RandomGetRandomBytes(len uint64) []uint8 {
  lower_len := C.uint64_t(len)
  var ret C.random_list_u8_t
  C.random_get_random_bytes(lower_len, &ret)
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

func RandomGetRandomU64() uint64 {
  ret := C.random_get_random_u64()
  var lift_ret uint64
  lift_ret = uint64(ret)
  return lift_ret
}

func RandomInsecureRandom() RandomTuple2U64U64T {
  var ret C.random_tuple2_u64_u64_t
  C.random_insecure_random(&ret)
  var lift_ret RandomTuple2U64U64T
  var lift_ret_F0 uint64
  lift_ret_F0 = uint64(ret.f0)
  lift_ret.F0 = lift_ret_F0
  var lift_ret_F1 uint64
  lift_ret_F1 = uint64(ret.f1)
  lift_ret.F1 = lift_ret_F1
  return lift_ret
}

type RandomTuple2U64U64T struct {
  F0 uint64
  F1 uint64
}

// console
type ConsoleLevelKind int

const (
ConsoleLevelKindTrace ConsoleLevelKind = iota
ConsoleLevelKindDebug
ConsoleLevelKindInfo
ConsoleLevelKindWarn
ConsoleLevelKindError
)

type ConsoleLevel struct {
  kind ConsoleLevelKind
}

func (n ConsoleLevel) Kind() ConsoleLevelKind {
  return n.kind
}

func ConsoleLevelTrace() ConsoleLevel{
  return ConsoleLevel{kind: ConsoleLevelKindTrace}
}

func ConsoleLevelDebug() ConsoleLevel{
  return ConsoleLevel{kind: ConsoleLevelKindDebug}
}

func ConsoleLevelInfo() ConsoleLevel{
  return ConsoleLevel{kind: ConsoleLevelKindInfo}
}

func ConsoleLevelWarn() ConsoleLevel{
  return ConsoleLevel{kind: ConsoleLevelKindWarn}
}

func ConsoleLevelError() ConsoleLevel{
  return ConsoleLevel{kind: ConsoleLevelKindError}
}

func ConsoleLog(level ConsoleLevel, context string, message string) {
  var lower_level C.console_level_t
  if level.Kind() == ConsoleLevelKindTrace {
    lower_level = 0
  }
  if level.Kind() == ConsoleLevelKindDebug {
    lower_level = 1
  }
  if level.Kind() == ConsoleLevelKindInfo {
    lower_level = 2
  }
  if level.Kind() == ConsoleLevelKindWarn {
    lower_level = 3
  }
  if level.Kind() == ConsoleLevelKindError {
    lower_level = 4
  }
  var lower_context C.proxy_string_t
  
  lower_context.ptr = C.CString(context)
  lower_context.len = C.size_t(len(context))
  defer C.proxy_string_free(&lower_context)
  var lower_message C.proxy_string_t
  
  lower_message.ptr = C.CString(message)
  lower_message.len = C.size_t(len(message))
  defer C.proxy_string_free(&lower_message)
  C.console_log(lower_level, &lower_context, &lower_message)
}

// poll
type PollPollable = uint32
func PollDropPollable(this uint32) {
  var lower_this C.uint32_t
  lower_this_val := C.uint32_t(this)
  lower_this = lower_this_val
  C.poll_drop_pollable(lower_this)
}

func PollPollOneoff(in []uint32) []uint8 {
  var lower_in C.poll_list_pollable_t
  if len(in) == 0 {
    lower_in.ptr = nil
    lower_in.len = 0
  } else {
    var empty_lower_in C.poll_pollable_t
    lower_in.ptr = (*C.poll_pollable_t)(C.malloc(C.size_t(len(in)) * C.size_t(unsafe.Sizeof(empty_lower_in))))
    lower_in.len = C.size_t(len(in))
    for lower_in_i := range in {
      lower_in_ptr := (*C.poll_pollable_t)(unsafe.Pointer(uintptr(unsafe.Pointer(lower_in.ptr)) +
      uintptr(lower_in_i)*unsafe.Sizeof(empty_lower_in)))
      var lower_in_ptr_value C.uint32_t
      lower_in_ptr_value_val := C.uint32_t(in[lower_in_i])
      lower_in_ptr_value = lower_in_ptr_value_val
      *lower_in_ptr = lower_in_ptr_value
    }
  }
  defer C.poll_list_pollable_free(&lower_in)
  var ret C.poll_list_u8_t
  C.poll_poll_oneoff(&lower_in, &ret)
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

// streams
type StreamsPollable = uint32
type StreamsStreamError struct {
}

type StreamsOutputStream = uint32
type StreamsInputStream = uint32
func StreamsRead(this uint32, len uint64) Result[StreamsTuple2ListU8TBoolT, StreamsStreamError] {
  var lower_this C.uint32_t
  lower_this_val := C.uint32_t(this)
  lower_this = lower_this_val
  lower_len := C.uint64_t(len)
  var err C.streams_stream_error_t
  var ret C.streams_tuple2_list_u8_bool_t
  is_ret_ok := C.streams_read(lower_this, lower_len, &ret, &err)
  
  var lift_ret Result[StreamsTuple2ListU8TBoolT, StreamsStreamError]
  if !is_ret_ok {
    var lift_ret_val StreamsStreamError
    lift_ret.SetErr(lift_ret_val)
  } else {
    var lift_ret_val StreamsTuple2ListU8TBoolT
    var lift_ret_val_F0 []uint8
    lift_ret_val_F0 = make([]uint8, ret.f0.len)
    if ret.f0.len > 0 {
      for lift_ret_val_F0_i := 0; lift_ret_val_F0_i < int(ret.f0.len); lift_ret_val_F0_i++ {
        var empty_lift_ret_val_F0 C.uint8_t
        lift_ret_val_F0_ptr := *(*C.uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(ret.f0.ptr)) +
        uintptr(lift_ret_val_F0_i)*unsafe.Sizeof(empty_lift_ret_val_F0)))
        var list_lift_ret_val_F0 uint8
        list_lift_ret_val_F0 = uint8(lift_ret_val_F0_ptr)
        lift_ret_val_F0[lift_ret_val_F0_i] = list_lift_ret_val_F0
      }
    }
    lift_ret_val.F0 = lift_ret_val_F0
    lift_ret_val_F1 := ret.f1
    lift_ret_val.F1 = lift_ret_val_F1
    lift_ret.Set(lift_ret_val)
  }
  return lift_ret
}

func StreamsBlockingRead(this uint32, len uint64) Result[StreamsTuple2ListU8TBoolT, StreamsStreamError] {
  var lower_this C.uint32_t
  lower_this_val := C.uint32_t(this)
  lower_this = lower_this_val
  lower_len := C.uint64_t(len)
  var err C.streams_stream_error_t
  var ret C.streams_tuple2_list_u8_bool_t
  is_ret_ok := C.streams_blocking_read(lower_this, lower_len, &ret, &err)
  
  var lift_ret Result[StreamsTuple2ListU8TBoolT, StreamsStreamError]
  if !is_ret_ok {
    var lift_ret_val StreamsStreamError
    lift_ret.SetErr(lift_ret_val)
  } else {
    var lift_ret_val StreamsTuple2ListU8TBoolT
    var lift_ret_val_F0 []uint8
    lift_ret_val_F0 = make([]uint8, ret.f0.len)
    if ret.f0.len > 0 {
      for lift_ret_val_F0_i := 0; lift_ret_val_F0_i < int(ret.f0.len); lift_ret_val_F0_i++ {
        var empty_lift_ret_val_F0 C.uint8_t
        lift_ret_val_F0_ptr := *(*C.uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(ret.f0.ptr)) +
        uintptr(lift_ret_val_F0_i)*unsafe.Sizeof(empty_lift_ret_val_F0)))
        var list_lift_ret_val_F0 uint8
        list_lift_ret_val_F0 = uint8(lift_ret_val_F0_ptr)
        lift_ret_val_F0[lift_ret_val_F0_i] = list_lift_ret_val_F0
      }
    }
    lift_ret_val.F0 = lift_ret_val_F0
    lift_ret_val_F1 := ret.f1
    lift_ret_val.F1 = lift_ret_val_F1
    lift_ret.Set(lift_ret_val)
  }
  return lift_ret
}

func StreamsSkip(this uint32, len uint64) Result[StreamsTuple2U64BoolT, StreamsStreamError] {
  var lower_this C.uint32_t
  lower_this_val := C.uint32_t(this)
  lower_this = lower_this_val
  lower_len := C.uint64_t(len)
  var err C.streams_stream_error_t
  var ret C.streams_tuple2_u64_bool_t
  is_ret_ok := C.streams_skip(lower_this, lower_len, &ret, &err)
  
  var lift_ret Result[StreamsTuple2U64BoolT, StreamsStreamError]
  if !is_ret_ok {
    var lift_ret_val StreamsStreamError
    lift_ret.SetErr(lift_ret_val)
  } else {
    var lift_ret_val StreamsTuple2U64BoolT
    var lift_ret_val_F0 uint64
    lift_ret_val_F0 = uint64(ret.f0)
    lift_ret_val.F0 = lift_ret_val_F0
    lift_ret_val_F1 := ret.f1
    lift_ret_val.F1 = lift_ret_val_F1
    lift_ret.Set(lift_ret_val)
  }
  return lift_ret
}

func StreamsBlockingSkip(this uint32, len uint64) Result[StreamsTuple2U64BoolT, StreamsStreamError] {
  var lower_this C.uint32_t
  lower_this_val := C.uint32_t(this)
  lower_this = lower_this_val
  lower_len := C.uint64_t(len)
  var err C.streams_stream_error_t
  var ret C.streams_tuple2_u64_bool_t
  is_ret_ok := C.streams_blocking_skip(lower_this, lower_len, &ret, &err)
  
  var lift_ret Result[StreamsTuple2U64BoolT, StreamsStreamError]
  if !is_ret_ok {
    var lift_ret_val StreamsStreamError
    lift_ret.SetErr(lift_ret_val)
  } else {
    var lift_ret_val StreamsTuple2U64BoolT
    var lift_ret_val_F0 uint64
    lift_ret_val_F0 = uint64(ret.f0)
    lift_ret_val.F0 = lift_ret_val_F0
    lift_ret_val_F1 := ret.f1
    lift_ret_val.F1 = lift_ret_val_F1
    lift_ret.Set(lift_ret_val)
  }
  return lift_ret
}

func StreamsSubscribeToInputStream(this uint32) uint32 {
  var lower_this C.uint32_t
  lower_this_val := C.uint32_t(this)
  lower_this = lower_this_val
  ret := C.streams_subscribe_to_input_stream(lower_this)
  var lift_ret uint32
  var lift_ret_val uint32
  var lift_ret_val_val uint32
  lift_ret_val_val = uint32(ret)
  lift_ret_val = lift_ret_val_val
  lift_ret = lift_ret_val
  return lift_ret
}

func StreamsDropInputStream(this uint32) {
  var lower_this C.uint32_t
  lower_this_val := C.uint32_t(this)
  lower_this = lower_this_val
  C.streams_drop_input_stream(lower_this)
}

func StreamsWrite(this uint32, buf []uint8) Result[uint64, StreamsStreamError] {
  var lower_this C.uint32_t
  lower_this_val := C.uint32_t(this)
  lower_this = lower_this_val
  var lower_buf C.streams_list_u8_t
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
  defer C.streams_list_u8_free(&lower_buf)
  var err C.streams_stream_error_t
  var ret C.uint64_t
  is_ret_ok := C.streams_write(lower_this, &lower_buf, &ret, &err)
  
  var lift_ret Result[uint64, StreamsStreamError]
  if !is_ret_ok {
    var lift_ret_val StreamsStreamError
    lift_ret.SetErr(lift_ret_val)
  } else {
    var lift_ret_val uint64
    lift_ret_val = uint64(ret)
    lift_ret.Set(lift_ret_val)
  }
  return lift_ret
}

func StreamsBlockingWrite(this uint32, buf []uint8) Result[uint64, StreamsStreamError] {
  var lower_this C.uint32_t
  lower_this_val := C.uint32_t(this)
  lower_this = lower_this_val
  var lower_buf C.streams_list_u8_t
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
  defer C.streams_list_u8_free(&lower_buf)
  var err C.streams_stream_error_t
  var ret C.uint64_t
  is_ret_ok := C.streams_blocking_write(lower_this, &lower_buf, &ret, &err)
  
  var lift_ret Result[uint64, StreamsStreamError]
  if !is_ret_ok {
    var lift_ret_val StreamsStreamError
    lift_ret.SetErr(lift_ret_val)
  } else {
    var lift_ret_val uint64
    lift_ret_val = uint64(ret)
    lift_ret.Set(lift_ret_val)
  }
  return lift_ret
}

func StreamsWriteZeroes(this uint32, len uint64) Result[uint64, StreamsStreamError] {
  var lower_this C.uint32_t
  lower_this_val := C.uint32_t(this)
  lower_this = lower_this_val
  lower_len := C.uint64_t(len)
  var err C.streams_stream_error_t
  var ret C.uint64_t
  is_ret_ok := C.streams_write_zeroes(lower_this, lower_len, &ret, &err)
  
  var lift_ret Result[uint64, StreamsStreamError]
  if !is_ret_ok {
    var lift_ret_val StreamsStreamError
    lift_ret.SetErr(lift_ret_val)
  } else {
    var lift_ret_val uint64
    lift_ret_val = uint64(ret)
    lift_ret.Set(lift_ret_val)
  }
  return lift_ret
}

func StreamsBlockingWriteZeroes(this uint32, len uint64) Result[uint64, StreamsStreamError] {
  var lower_this C.uint32_t
  lower_this_val := C.uint32_t(this)
  lower_this = lower_this_val
  lower_len := C.uint64_t(len)
  var err C.streams_stream_error_t
  var ret C.uint64_t
  is_ret_ok := C.streams_blocking_write_zeroes(lower_this, lower_len, &ret, &err)
  
  var lift_ret Result[uint64, StreamsStreamError]
  if !is_ret_ok {
    var lift_ret_val StreamsStreamError
    lift_ret.SetErr(lift_ret_val)
  } else {
    var lift_ret_val uint64
    lift_ret_val = uint64(ret)
    lift_ret.Set(lift_ret_val)
  }
  return lift_ret
}

func StreamsSplice(this uint32, src uint32, len uint64) Result[StreamsTuple2U64BoolT, StreamsStreamError] {
  var lower_this C.uint32_t
  lower_this_val := C.uint32_t(this)
  lower_this = lower_this_val
  var lower_src C.uint32_t
  lower_src_val := C.uint32_t(src)
  lower_src = lower_src_val
  lower_len := C.uint64_t(len)
  var err C.streams_stream_error_t
  var ret C.streams_tuple2_u64_bool_t
  is_ret_ok := C.streams_splice(lower_this, lower_src, lower_len, &ret, &err)
  
  var lift_ret Result[StreamsTuple2U64BoolT, StreamsStreamError]
  if !is_ret_ok {
    var lift_ret_val StreamsStreamError
    lift_ret.SetErr(lift_ret_val)
  } else {
    var lift_ret_val StreamsTuple2U64BoolT
    var lift_ret_val_F0 uint64
    lift_ret_val_F0 = uint64(ret.f0)
    lift_ret_val.F0 = lift_ret_val_F0
    lift_ret_val_F1 := ret.f1
    lift_ret_val.F1 = lift_ret_val_F1
    lift_ret.Set(lift_ret_val)
  }
  return lift_ret
}

func StreamsBlockingSplice(this uint32, src uint32, len uint64) Result[StreamsTuple2U64BoolT, StreamsStreamError] {
  var lower_this C.uint32_t
  lower_this_val := C.uint32_t(this)
  lower_this = lower_this_val
  var lower_src C.uint32_t
  lower_src_val := C.uint32_t(src)
  lower_src = lower_src_val
  lower_len := C.uint64_t(len)
  var err C.streams_stream_error_t
  var ret C.streams_tuple2_u64_bool_t
  is_ret_ok := C.streams_blocking_splice(lower_this, lower_src, lower_len, &ret, &err)
  
  var lift_ret Result[StreamsTuple2U64BoolT, StreamsStreamError]
  if !is_ret_ok {
    var lift_ret_val StreamsStreamError
    lift_ret.SetErr(lift_ret_val)
  } else {
    var lift_ret_val StreamsTuple2U64BoolT
    var lift_ret_val_F0 uint64
    lift_ret_val_F0 = uint64(ret.f0)
    lift_ret_val.F0 = lift_ret_val_F0
    lift_ret_val_F1 := ret.f1
    lift_ret_val.F1 = lift_ret_val_F1
    lift_ret.Set(lift_ret_val)
  }
  return lift_ret
}

func StreamsForward(this uint32, src uint32) Result[uint64, StreamsStreamError] {
  var lower_this C.uint32_t
  lower_this_val := C.uint32_t(this)
  lower_this = lower_this_val
  var lower_src C.uint32_t
  lower_src_val := C.uint32_t(src)
  lower_src = lower_src_val
  var err C.streams_stream_error_t
  var ret C.uint64_t
  is_ret_ok := C.streams_forward(lower_this, lower_src, &ret, &err)
  
  var lift_ret Result[uint64, StreamsStreamError]
  if !is_ret_ok {
    var lift_ret_val StreamsStreamError
    lift_ret.SetErr(lift_ret_val)
  } else {
    var lift_ret_val uint64
    lift_ret_val = uint64(ret)
    lift_ret.Set(lift_ret_val)
  }
  return lift_ret
}

func StreamsSubscribeToOutputStream(this uint32) uint32 {
  var lower_this C.uint32_t
  lower_this_val := C.uint32_t(this)
  lower_this = lower_this_val
  ret := C.streams_subscribe_to_output_stream(lower_this)
  var lift_ret uint32
  var lift_ret_val uint32
  var lift_ret_val_val uint32
  lift_ret_val_val = uint32(ret)
  lift_ret_val = lift_ret_val_val
  lift_ret = lift_ret_val
  return lift_ret
}

func StreamsDropOutputStream(this uint32) {
  var lower_this C.uint32_t
  lower_this_val := C.uint32_t(this)
  lower_this = lower_this_val
  C.streams_drop_output_stream(lower_this)
}

type StreamsTuple2U64BoolT struct {
  F0 uint64
  F1 bool
}

type StreamsTuple2ListU8TBoolT struct {
  F0 []uint8
  F1 bool
}

// types
type TypesInputStream = uint32
type TypesOutputStream = uint32
type TypesPollable = uint32
type TypesStatusCode = uint16
type TypesSchemeKind int

const (
TypesSchemeKindHttp TypesSchemeKind = iota
TypesSchemeKindHttps
TypesSchemeKindOther
)

type TypesScheme struct {
  kind TypesSchemeKind
  val any
}

func (n TypesScheme) Kind() TypesSchemeKind {
  return n.kind
}

func TypesSchemeHttp() TypesScheme{
  return TypesScheme{kind: TypesSchemeKindHttp}
}

func TypesSchemeHttps() TypesScheme{
  return TypesScheme{kind: TypesSchemeKindHttps}
}

func TypesSchemeOther(v string) TypesScheme{
  return TypesScheme{kind: TypesSchemeKindOther, val: v}
}

func (n TypesScheme) GetOther() string{
  if g, w := n.Kind(), TypesSchemeKindOther; g != w {
    panic(fmt.Sprintf("Attr kind is %v, not %v", g, w))
  }
  return n.val.(string)
}

func (n *TypesScheme) SetOther(v string) {
  n.val = v
  n.kind = TypesSchemeKindOther
}

type TypesResponseOutparam = uint32
type TypesRequestOptions struct {
  ConnectTimeoutMs Option[uint32]
  FirstByteTimeoutMs Option[uint32]
  BetweenBytesTimeoutMs Option[uint32]
}

type TypesOutgoingStream = uint32
type TypesOutgoingResponse = uint32
type TypesOutgoingRequest = uint32
type TypesMethodKind int

const (
TypesMethodKindGet TypesMethodKind = iota
TypesMethodKindHead
TypesMethodKindPost
TypesMethodKindPut
TypesMethodKindDelete
TypesMethodKindConnect
TypesMethodKindOptions
TypesMethodKindTrace
TypesMethodKindPatch
TypesMethodKindOther
)

type TypesMethod struct {
  kind TypesMethodKind
  val any
}

func (n TypesMethod) Kind() TypesMethodKind {
  return n.kind
}

func TypesMethodGet() TypesMethod{
  return TypesMethod{kind: TypesMethodKindGet}
}

func TypesMethodHead() TypesMethod{
  return TypesMethod{kind: TypesMethodKindHead}
}

func TypesMethodPost() TypesMethod{
  return TypesMethod{kind: TypesMethodKindPost}
}

func TypesMethodPut() TypesMethod{
  return TypesMethod{kind: TypesMethodKindPut}
}

func TypesMethodDelete() TypesMethod{
  return TypesMethod{kind: TypesMethodKindDelete}
}

func TypesMethodConnect() TypesMethod{
  return TypesMethod{kind: TypesMethodKindConnect}
}

func TypesMethodOptions() TypesMethod{
  return TypesMethod{kind: TypesMethodKindOptions}
}

func TypesMethodTrace() TypesMethod{
  return TypesMethod{kind: TypesMethodKindTrace}
}

func TypesMethodPatch() TypesMethod{
  return TypesMethod{kind: TypesMethodKindPatch}
}

func TypesMethodOther(v string) TypesMethod{
  return TypesMethod{kind: TypesMethodKindOther, val: v}
}

func (n TypesMethod) GetOther() string{
  if g, w := n.Kind(), TypesMethodKindOther; g != w {
    panic(fmt.Sprintf("Attr kind is %v, not %v", g, w))
  }
  return n.val.(string)
}

func (n *TypesMethod) SetOther(v string) {
  n.val = v
  n.kind = TypesMethodKindOther
}

type TypesIncomingStream = uint32
type TypesIncomingResponse = uint32
type TypesIncomingRequest = uint32
type TypesFutureIncomingResponse = uint32
type TypesFields = uint32
type TypesTrailers = uint32
type TypesHeaders = uint32
type TypesErrorKind int

const (
TypesErrorKindInvalidUrl TypesErrorKind = iota
TypesErrorKindTimeoutError
TypesErrorKindProtocolError
TypesErrorKindUnexpectedError
)

type TypesError struct {
  kind TypesErrorKind
  val any
}

func (n TypesError) Kind() TypesErrorKind {
  return n.kind
}

func TypesErrorInvalidUrl(v string) TypesError{
  return TypesError{kind: TypesErrorKindInvalidUrl, val: v}
}

func (n TypesError) GetInvalidUrl() string{
  if g, w := n.Kind(), TypesErrorKindInvalidUrl; g != w {
    panic(fmt.Sprintf("Attr kind is %v, not %v", g, w))
  }
  return n.val.(string)
}

func (n *TypesError) SetInvalidUrl(v string) {
  n.val = v
  n.kind = TypesErrorKindInvalidUrl
}

func TypesErrorTimeoutError(v string) TypesError{
  return TypesError{kind: TypesErrorKindTimeoutError, val: v}
}

func (n TypesError) GetTimeoutError() string{
  if g, w := n.Kind(), TypesErrorKindTimeoutError; g != w {
    panic(fmt.Sprintf("Attr kind is %v, not %v", g, w))
  }
  return n.val.(string)
}

func (n *TypesError) SetTimeoutError(v string) {
  n.val = v
  n.kind = TypesErrorKindTimeoutError
}

func TypesErrorProtocolError(v string) TypesError{
  return TypesError{kind: TypesErrorKindProtocolError, val: v}
}

func (n TypesError) GetProtocolError() string{
  if g, w := n.Kind(), TypesErrorKindProtocolError; g != w {
    panic(fmt.Sprintf("Attr kind is %v, not %v", g, w))
  }
  return n.val.(string)
}

func (n *TypesError) SetProtocolError(v string) {
  n.val = v
  n.kind = TypesErrorKindProtocolError
}

func TypesErrorUnexpectedError(v string) TypesError{
  return TypesError{kind: TypesErrorKindUnexpectedError, val: v}
}

func (n TypesError) GetUnexpectedError() string{
  if g, w := n.Kind(), TypesErrorKindUnexpectedError; g != w {
    panic(fmt.Sprintf("Attr kind is %v, not %v", g, w))
  }
  return n.val.(string)
}

func (n *TypesError) SetUnexpectedError(v string) {
  n.val = v
  n.kind = TypesErrorKindUnexpectedError
}

func TypesDropFields(fields uint32) {
  var lower_fields C.uint32_t
  lower_fields_val := C.uint32_t(fields)
  lower_fields = lower_fields_val
  C.types_drop_fields(lower_fields)
}

func TypesNewFields(entries []TypesTuple2StringStringT) uint32 {
  var lower_entries C.types_list_tuple2_string_string_t
  if len(entries) == 0 {
    lower_entries.ptr = nil
    lower_entries.len = 0
  } else {
    var empty_lower_entries C.types_tuple2_string_string_t
    lower_entries.ptr = (*C.types_tuple2_string_string_t)(C.malloc(C.size_t(len(entries)) * C.size_t(unsafe.Sizeof(empty_lower_entries))))
    lower_entries.len = C.size_t(len(entries))
    for lower_entries_i := range entries {
      lower_entries_ptr := (*C.types_tuple2_string_string_t)(unsafe.Pointer(uintptr(unsafe.Pointer(lower_entries.ptr)) +
      uintptr(lower_entries_i)*unsafe.Sizeof(empty_lower_entries)))
      var lower_entries_ptr_value C.types_tuple2_string_string_t
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
  defer C.types_list_tuple2_string_string_free(&lower_entries)
  ret := C.types_new_fields(&lower_entries)
  var lift_ret uint32
  var lift_ret_val uint32
  lift_ret_val = uint32(ret)
  lift_ret = lift_ret_val
  return lift_ret
}

func TypesFieldsGet(fields uint32, name string) []string {
  var lower_fields C.uint32_t
  lower_fields_val := C.uint32_t(fields)
  lower_fields = lower_fields_val
  var lower_name C.proxy_string_t
  
  lower_name.ptr = C.CString(name)
  lower_name.len = C.size_t(len(name))
  defer C.proxy_string_free(&lower_name)
  var ret C.types_list_string_t
  C.types_fields_get(lower_fields, &lower_name, &ret)
  var lift_ret []string
  lift_ret = make([]string, ret.len)
  if ret.len > 0 {
    for lift_ret_i := 0; lift_ret_i < int(ret.len); lift_ret_i++ {
      var empty_lift_ret C.proxy_string_t
      lift_ret_ptr := *(*C.proxy_string_t)(unsafe.Pointer(uintptr(unsafe.Pointer(ret.ptr)) +
      uintptr(lift_ret_i)*unsafe.Sizeof(empty_lift_ret)))
      var list_lift_ret string
      list_lift_ret = C.GoStringN(lift_ret_ptr.ptr, C.int(lift_ret_ptr.len))
      lift_ret[lift_ret_i] = list_lift_ret
    }
  }
  return lift_ret
}

func TypesFieldsSet(fields uint32, name string, value []string) {
  var lower_fields C.uint32_t
  lower_fields_val := C.uint32_t(fields)
  lower_fields = lower_fields_val
  var lower_name C.proxy_string_t
  
  lower_name.ptr = C.CString(name)
  lower_name.len = C.size_t(len(name))
  defer C.proxy_string_free(&lower_name)
  var lower_value C.types_list_string_t
  if len(value) == 0 {
    lower_value.ptr = nil
    lower_value.len = 0
  } else {
    var empty_lower_value C.proxy_string_t
    lower_value.ptr = (*C.proxy_string_t)(C.malloc(C.size_t(len(value)) * C.size_t(unsafe.Sizeof(empty_lower_value))))
    lower_value.len = C.size_t(len(value))
    for lower_value_i := range value {
      lower_value_ptr := (*C.proxy_string_t)(unsafe.Pointer(uintptr(unsafe.Pointer(lower_value.ptr)) +
      uintptr(lower_value_i)*unsafe.Sizeof(empty_lower_value)))
      var lower_value_ptr_value C.proxy_string_t
      
      lower_value_ptr_value.ptr = C.CString(value[lower_value_i])
      lower_value_ptr_value.len = C.size_t(len(value[lower_value_i]))
      *lower_value_ptr = lower_value_ptr_value
    }
  }
  defer C.types_list_string_free(&lower_value)
  C.types_fields_set(lower_fields, &lower_name, &lower_value)
}

func TypesFieldsDelete(fields uint32, name string) {
  var lower_fields C.uint32_t
  lower_fields_val := C.uint32_t(fields)
  lower_fields = lower_fields_val
  var lower_name C.proxy_string_t
  
  lower_name.ptr = C.CString(name)
  lower_name.len = C.size_t(len(name))
  defer C.proxy_string_free(&lower_name)
  C.types_fields_delete(lower_fields, &lower_name)
}

func TypesFieldsAppend(fields uint32, name string, value string) {
  var lower_fields C.uint32_t
  lower_fields_val := C.uint32_t(fields)
  lower_fields = lower_fields_val
  var lower_name C.proxy_string_t
  
  lower_name.ptr = C.CString(name)
  lower_name.len = C.size_t(len(name))
  defer C.proxy_string_free(&lower_name)
  var lower_value C.proxy_string_t
  
  lower_value.ptr = C.CString(value)
  lower_value.len = C.size_t(len(value))
  defer C.proxy_string_free(&lower_value)
  C.types_fields_append(lower_fields, &lower_name, &lower_value)
}

func TypesFieldsEntries(fields uint32) []TypesTuple2StringStringT {
  var lower_fields C.uint32_t
  lower_fields_val := C.uint32_t(fields)
  lower_fields = lower_fields_val
  var ret C.types_list_tuple2_string_string_t
  C.types_fields_entries(lower_fields, &ret)
  var lift_ret []TypesTuple2StringStringT
  lift_ret = make([]TypesTuple2StringStringT, ret.len)
  if ret.len > 0 {
    for lift_ret_i := 0; lift_ret_i < int(ret.len); lift_ret_i++ {
      var empty_lift_ret C.types_tuple2_string_string_t
      lift_ret_ptr := *(*C.types_tuple2_string_string_t)(unsafe.Pointer(uintptr(unsafe.Pointer(ret.ptr)) +
      uintptr(lift_ret_i)*unsafe.Sizeof(empty_lift_ret)))
      var list_lift_ret TypesTuple2StringStringT
      var list_lift_ret_F0 string
      list_lift_ret_F0 = C.GoStringN(lift_ret_ptr.f0.ptr, C.int(lift_ret_ptr.f0.len))
      list_lift_ret.F0 = list_lift_ret_F0
      var list_lift_ret_F1 string
      list_lift_ret_F1 = C.GoStringN(lift_ret_ptr.f1.ptr, C.int(lift_ret_ptr.f1.len))
      list_lift_ret.F1 = list_lift_ret_F1
      lift_ret[lift_ret_i] = list_lift_ret
    }
  }
  return lift_ret
}

func TypesFieldsClone(fields uint32) uint32 {
  var lower_fields C.uint32_t
  lower_fields_val := C.uint32_t(fields)
  lower_fields = lower_fields_val
  ret := C.types_fields_clone(lower_fields)
  var lift_ret uint32
  var lift_ret_val uint32
  lift_ret_val = uint32(ret)
  lift_ret = lift_ret_val
  return lift_ret
}

func TypesFinishIncomingStream(s uint32) Option[uint32] {
  var lower_s C.types_input_stream_t
  var lower_s_val C.streams_input_stream_t
  var lower_s_val_val C.uint32_t
  lower_s_val_val_val := C.uint32_t(s)
  lower_s_val_val = lower_s_val_val_val
  lower_s_val = lower_s_val_val
  lower_s = lower_s_val
  var ret C.types_trailers_t
  C.types_finish_incoming_stream(lower_s, &ret)
  var lift_ret Option[uint32]
  if ret == 0 {
    lift_ret.Unset()
  } else {
    var lift_ret_val uint32
    var lift_ret_val_val uint32
    var lift_ret_val_val_val uint32
    lift_ret_val_val_val = uint32(ret)
    lift_ret_val_val = lift_ret_val_val_val
    lift_ret_val = lift_ret_val_val
    lift_ret.Set(lift_ret_val)
  }
  return lift_ret
}

func TypesFinishOutgoingStream(s uint32, trailers Option[uint32]) {
  var lower_s C.types_output_stream_t
  var lower_s_val C.streams_output_stream_t
  var lower_s_val_val C.uint32_t
  lower_s_val_val_val := C.uint32_t(s)
  lower_s_val_val = lower_s_val_val_val
  lower_s_val = lower_s_val_val
  lower_s = lower_s_val
  var lower_trailers C.types_trailers_t
  if trailers.IsSome() {
    var lower_trailers_val C.types_fields_t
    var lower_trailers_val_val C.uint32_t
    lower_trailers_val_val_val := C.uint32_t(trailers.Unwrap())
    lower_trailers_val_val = lower_trailers_val_val_val
    lower_trailers_val = lower_trailers_val_val
    lower_trailers = lower_trailers_val
  }
  C.types_finish_outgoing_stream(lower_s, &lower_trailers)
}

func TypesDropIncomingRequest(request uint32) {
  var lower_request C.uint32_t
  lower_request_val := C.uint32_t(request)
  lower_request = lower_request_val
  C.types_drop_incoming_request(lower_request)
}

func TypesDropOutgoingRequest(request uint32) {
  var lower_request C.uint32_t
  lower_request_val := C.uint32_t(request)
  lower_request = lower_request_val
  C.types_drop_outgoing_request(lower_request)
}

func TypesIncomingRequestMethod(request uint32) TypesMethod {
  var lower_request C.uint32_t
  lower_request_val := C.uint32_t(request)
  lower_request = lower_request_val
  var ret C.types_method_t
  C.types_incoming_request_method(lower_request, &ret)
  var lift_ret TypesMethod
  if ret.tag == 0 {
    lift_ret = TypesMethodGet()
  }
  if ret.tag == 1 {
    lift_ret = TypesMethodHead()
  }
  if ret.tag == 2 {
    lift_ret = TypesMethodPost()
  }
  if ret.tag == 3 {
    lift_ret = TypesMethodPut()
  }
  if ret.tag == 4 {
    lift_ret = TypesMethodDelete()
  }
  if ret.tag == 5 {
    lift_ret = TypesMethodConnect()
  }
  if ret.tag == 6 {
    lift_ret = TypesMethodOptions()
  }
  if ret.tag == 7 {
    lift_ret = TypesMethodTrace()
  }
  if ret.tag == 8 {
    lift_ret = TypesMethodPatch()
  }
  if ret.tag == 9 {
    lift_ret_ptr := *(*C.proxy_string_t)(unsafe.Pointer(&ret.val))
    var lift_ret_val string
    lift_ret_val = C.GoStringN(lift_ret_ptr.ptr, C.int(lift_ret_ptr.len))
    lift_ret = TypesMethodOther(lift_ret_val)
  }
  return lift_ret
}

func TypesIncomingRequestPath(request uint32) string {
  var lower_request C.uint32_t
  lower_request_val := C.uint32_t(request)
  lower_request = lower_request_val
  var ret C.proxy_string_t
  C.types_incoming_request_path(lower_request, &ret)
  var lift_ret string
  lift_ret = C.GoStringN(ret.ptr, C.int(ret.len))
  return lift_ret
}

func TypesIncomingRequestQuery(request uint32) string {
  var lower_request C.uint32_t
  lower_request_val := C.uint32_t(request)
  lower_request = lower_request_val
  var ret C.proxy_string_t
  C.types_incoming_request_query(lower_request, &ret)
  var lift_ret string
  lift_ret = C.GoStringN(ret.ptr, C.int(ret.len))
  return lift_ret
}

func TypesIncomingRequestScheme(request uint32) Option[TypesScheme] {
  var lower_request C.uint32_t
  lower_request_val := C.uint32_t(request)
  lower_request = lower_request_val
  var ret C.types_scheme_t
  C.types_incoming_request_scheme(lower_request, &ret)
  var lift_ret Option[TypesScheme]
  var lift_ret_c C.types_scheme_t
  defer C.types_scheme_free(&lift_ret_c)
  if ret == lift_ret_c {
    lift_ret.Unset()
  } else {
    var lift_ret_val TypesScheme
    if ret.tag == 0 {
      lift_ret_val = TypesSchemeHttp()
    }
    if ret.tag == 1 {
      lift_ret_val = TypesSchemeHttps()
    }
    if ret.tag == 2 {
      lift_ret_val_ptr := *(*C.proxy_string_t)(unsafe.Pointer(&ret.val))
      var lift_ret_val_val string
      lift_ret_val_val = C.GoStringN(lift_ret_val_ptr.ptr, C.int(lift_ret_val_ptr.len))
      lift_ret_val = TypesSchemeOther(lift_ret_val_val)
    }
    lift_ret.Set(lift_ret_val)
  }
  return lift_ret
}

func TypesIncomingRequestAuthority(request uint32) string {
  var lower_request C.uint32_t
  lower_request_val := C.uint32_t(request)
  lower_request = lower_request_val
  var ret C.proxy_string_t
  C.types_incoming_request_authority(lower_request, &ret)
  var lift_ret string
  lift_ret = C.GoStringN(ret.ptr, C.int(ret.len))
  return lift_ret
}

func TypesIncomingRequestHeaders(request uint32) uint32 {
  var lower_request C.uint32_t
  lower_request_val := C.uint32_t(request)
  lower_request = lower_request_val
  ret := C.types_incoming_request_headers(lower_request)
  var lift_ret uint32
  var lift_ret_val uint32
  var lift_ret_val_val uint32
  lift_ret_val_val = uint32(ret)
  lift_ret_val = lift_ret_val_val
  lift_ret = lift_ret_val
  return lift_ret
}

func TypesIncomingRequestConsume(request uint32) Result[uint32, struct{}] {
  var lower_request C.uint32_t
  lower_request_val := C.uint32_t(request)
  lower_request = lower_request_val
  var ret C.types_incoming_stream_t
  is_ret_ok := C.types_incoming_request_consume(lower_request, &ret)
  
  var lift_ret Result[uint32, struct{}]
  if !is_ret_ok {
  } else {
    var lift_ret_val uint32
    var lift_ret_val_val uint32
    var lift_ret_val_val_val uint32
    var lift_ret_val_val_val_val uint32
    lift_ret_val_val_val_val = uint32(ret)
    lift_ret_val_val_val = lift_ret_val_val_val_val
    lift_ret_val_val = lift_ret_val_val_val
    lift_ret_val = lift_ret_val_val
    lift_ret.Set(lift_ret_val)
  }
  return lift_ret
}

func TypesNewOutgoingRequest(method TypesMethod, path string, query string, scheme Option[TypesScheme], authority string, headers uint32) uint32 {
  var lower_method C.types_method_t
  if method.Kind() == TypesMethodKindGet {
    lower_method.tag = 0
  }
  if method.Kind() == TypesMethodKindHead {
    lower_method.tag = 1
  }
  if method.Kind() == TypesMethodKindPost {
    lower_method.tag = 2
  }
  if method.Kind() == TypesMethodKindPut {
    lower_method.tag = 3
  }
  if method.Kind() == TypesMethodKindDelete {
    lower_method.tag = 4
  }
  if method.Kind() == TypesMethodKindConnect {
    lower_method.tag = 5
  }
  if method.Kind() == TypesMethodKindOptions {
    lower_method.tag = 6
  }
  if method.Kind() == TypesMethodKindTrace {
    lower_method.tag = 7
  }
  if method.Kind() == TypesMethodKindPatch {
    lower_method.tag = 8
  }
  if method.Kind() == TypesMethodKindOther {
    
    lower_method.tag = 9
    lower_method_ptr := (*C.proxy_string_t)(unsafe.Pointer(&lower_method.val))
    var lower_method_val C.proxy_string_t
    
    lower_method_val.ptr = C.CString(method.GetOther())
    lower_method_val.len = C.size_t(len(method.GetOther()))
    *lower_method_ptr = lower_method_val
  }
  defer C.types_method_free(&lower_method)
  var lower_path C.proxy_string_t
  
  lower_path.ptr = C.CString(path)
  lower_path.len = C.size_t(len(path))
  defer C.proxy_string_free(&lower_path)
  var lower_query C.proxy_string_t
  
  lower_query.ptr = C.CString(query)
  lower_query.len = C.size_t(len(query))
  defer C.proxy_string_free(&lower_query)
  var lower_scheme C.types_scheme_t
  if scheme.IsSome() {
    var lower_scheme_val C.types_scheme_t
    if scheme.Unwrap().Kind() == TypesSchemeKindHttp {
      lower_scheme_val.tag = 0
    }
    if scheme.Unwrap().Kind() == TypesSchemeKindHttps {
      lower_scheme_val.tag = 1
    }
    if scheme.Unwrap().Kind() == TypesSchemeKindOther {
      
      lower_scheme_val.tag = 2
      lower_scheme_val_ptr := (*C.proxy_string_t)(unsafe.Pointer(&lower_scheme_val.val))
      var lower_scheme_val_val C.proxy_string_t
      
      lower_scheme_val_val.ptr = C.CString(scheme.Unwrap().GetOther())
      lower_scheme_val_val.len = C.size_t(len(scheme.Unwrap().GetOther()))
      *lower_scheme_val_ptr = lower_scheme_val_val
    }
    lower_scheme = lower_scheme_val
  }
  defer C.types_scheme_free(&lower_scheme)
  var lower_authority C.proxy_string_t
  
  lower_authority.ptr = C.CString(authority)
  lower_authority.len = C.size_t(len(authority))
  defer C.proxy_string_free(&lower_authority)
  var lower_headers C.types_fields_t
  var lower_headers_val C.uint32_t
  lower_headers_val_val := C.uint32_t(headers)
  lower_headers_val = lower_headers_val_val
  lower_headers = lower_headers_val
  ret := C.types_new_outgoing_request(&lower_method, &lower_path, &lower_query, &lower_scheme, &lower_authority, lower_headers)
  var lift_ret uint32
  var lift_ret_val uint32
  lift_ret_val = uint32(ret)
  lift_ret = lift_ret_val
  return lift_ret
}

func TypesOutgoingRequestWrite(request uint32) Result[uint32, struct{}] {
  var lower_request C.uint32_t
  lower_request_val := C.uint32_t(request)
  lower_request = lower_request_val
  var ret C.types_outgoing_stream_t
  is_ret_ok := C.types_outgoing_request_write(lower_request, &ret)
  
  var lift_ret Result[uint32, struct{}]
  if !is_ret_ok {
  } else {
    var lift_ret_val uint32
    var lift_ret_val_val uint32
    var lift_ret_val_val_val uint32
    var lift_ret_val_val_val_val uint32
    lift_ret_val_val_val_val = uint32(ret)
    lift_ret_val_val_val = lift_ret_val_val_val_val
    lift_ret_val_val = lift_ret_val_val_val
    lift_ret_val = lift_ret_val_val
    lift_ret.Set(lift_ret_val)
  }
  return lift_ret
}

func TypesDropResponseOutparam(response uint32) {
  var lower_response C.uint32_t
  lower_response_val := C.uint32_t(response)
  lower_response = lower_response_val
  C.types_drop_response_outparam(lower_response)
}

func TypesSetResponseOutparam(param uint32, response Result[uint32, TypesError]) Result[struct{}, struct{}] {
  var lower_param C.uint32_t
  lower_param_val := C.uint32_t(param)
  lower_param = lower_param_val
  var lower_response C.types_result_outgoing_response_error_t
  lower_response.is_err = response.IsErr()
  if response.IsOk() {
    lower_response_ptr := (*C.types_outgoing_response_t)(unsafe.Pointer(&lower_response.val))
    var lower_response_val C.uint32_t
    lower_response_val_val := C.uint32_t(response.Unwrap())
    lower_response_val = lower_response_val_val
    *lower_response_ptr = lower_response_val
  } else {
    lower_response_ptr := (*C.types_error_t)(unsafe.Pointer(&lower_response.val))
    var lower_response_val C.types_error_t
    if response.UnwrapErr().Kind() == TypesErrorKindInvalidUrl {
      
      lower_response_val.tag = 0
      lower_response_val_ptr := (*C.proxy_string_t)(unsafe.Pointer(&lower_response_val.val))
      var lower_response_val_val C.proxy_string_t
      
      lower_response_val_val.ptr = C.CString(response.UnwrapErr().GetInvalidUrl())
      lower_response_val_val.len = C.size_t(len(response.UnwrapErr().GetInvalidUrl()))
      *lower_response_val_ptr = lower_response_val_val
    }
    if response.UnwrapErr().Kind() == TypesErrorKindTimeoutError {
      
      lower_response_val.tag = 1
      lower_response_val_ptr := (*C.proxy_string_t)(unsafe.Pointer(&lower_response_val.val))
      var lower_response_val_val C.proxy_string_t
      
      lower_response_val_val.ptr = C.CString(response.UnwrapErr().GetTimeoutError())
      lower_response_val_val.len = C.size_t(len(response.UnwrapErr().GetTimeoutError()))
      *lower_response_val_ptr = lower_response_val_val
    }
    if response.UnwrapErr().Kind() == TypesErrorKindProtocolError {
      
      lower_response_val.tag = 2
      lower_response_val_ptr := (*C.proxy_string_t)(unsafe.Pointer(&lower_response_val.val))
      var lower_response_val_val C.proxy_string_t
      
      lower_response_val_val.ptr = C.CString(response.UnwrapErr().GetProtocolError())
      lower_response_val_val.len = C.size_t(len(response.UnwrapErr().GetProtocolError()))
      *lower_response_val_ptr = lower_response_val_val
    }
    if response.UnwrapErr().Kind() == TypesErrorKindUnexpectedError {
      
      lower_response_val.tag = 3
      lower_response_val_ptr := (*C.proxy_string_t)(unsafe.Pointer(&lower_response_val.val))
      var lower_response_val_val C.proxy_string_t
      
      lower_response_val_val.ptr = C.CString(response.UnwrapErr().GetUnexpectedError())
      lower_response_val_val.len = C.size_t(len(response.UnwrapErr().GetUnexpectedError()))
      *lower_response_val_ptr = lower_response_val_val
    }
    *lower_response_ptr = lower_response_val
  }
  defer C.types_result_outgoing_response_error_free(&lower_response)
  is_ret_ok := C.types_set_response_outparam(lower_param, &lower_response, )
  
  var lift_ret Result[struct{}, struct{}]
  if !is_ret_ok {
  } else {
  }
  return lift_ret
}

func TypesDropIncomingResponse(response uint32) {
  var lower_response C.uint32_t
  lower_response_val := C.uint32_t(response)
  lower_response = lower_response_val
  C.types_drop_incoming_response(lower_response)
}

func TypesDropOutgoingResponse(response uint32) {
  var lower_response C.uint32_t
  lower_response_val := C.uint32_t(response)
  lower_response = lower_response_val
  C.types_drop_outgoing_response(lower_response)
}

func TypesIncomingResponseStatus(response uint32) uint16 {
  var lower_response C.uint32_t
  lower_response_val := C.uint32_t(response)
  lower_response = lower_response_val
  ret := C.types_incoming_response_status(lower_response)
  var lift_ret uint16
  var lift_ret_val uint16
  lift_ret_val = uint16(ret)
  lift_ret = lift_ret_val
  return lift_ret
}

func TypesIncomingResponseHeaders(response uint32) uint32 {
  var lower_response C.uint32_t
  lower_response_val := C.uint32_t(response)
  lower_response = lower_response_val
  ret := C.types_incoming_response_headers(lower_response)
  var lift_ret uint32
  var lift_ret_val uint32
  var lift_ret_val_val uint32
  lift_ret_val_val = uint32(ret)
  lift_ret_val = lift_ret_val_val
  lift_ret = lift_ret_val
  return lift_ret
}

func TypesIncomingResponseConsume(response uint32) Result[uint32, struct{}] {
  var lower_response C.uint32_t
  lower_response_val := C.uint32_t(response)
  lower_response = lower_response_val
  var ret C.types_incoming_stream_t
  is_ret_ok := C.types_incoming_response_consume(lower_response, &ret)
  
  var lift_ret Result[uint32, struct{}]
  if !is_ret_ok {
  } else {
    var lift_ret_val uint32
    var lift_ret_val_val uint32
    var lift_ret_val_val_val uint32
    var lift_ret_val_val_val_val uint32
    lift_ret_val_val_val_val = uint32(ret)
    lift_ret_val_val_val = lift_ret_val_val_val_val
    lift_ret_val_val = lift_ret_val_val_val
    lift_ret_val = lift_ret_val_val
    lift_ret.Set(lift_ret_val)
  }
  return lift_ret
}

func TypesNewOutgoingResponse(status_code uint16, headers uint32) uint32 {
  var lower_status_code C.uint16_t
  lower_status_code_val := C.uint16_t(status_code)
  lower_status_code = lower_status_code_val
  var lower_headers C.types_fields_t
  var lower_headers_val C.uint32_t
  lower_headers_val_val := C.uint32_t(headers)
  lower_headers_val = lower_headers_val_val
  lower_headers = lower_headers_val
  ret := C.types_new_outgoing_response(lower_status_code, lower_headers)
  var lift_ret uint32
  var lift_ret_val uint32
  lift_ret_val = uint32(ret)
  lift_ret = lift_ret_val
  return lift_ret
}

func TypesOutgoingResponseWrite(response uint32) Result[uint32, struct{}] {
  var lower_response C.uint32_t
  lower_response_val := C.uint32_t(response)
  lower_response = lower_response_val
  var ret C.types_outgoing_stream_t
  is_ret_ok := C.types_outgoing_response_write(lower_response, &ret)
  
  var lift_ret Result[uint32, struct{}]
  if !is_ret_ok {
  } else {
    var lift_ret_val uint32
    var lift_ret_val_val uint32
    var lift_ret_val_val_val uint32
    var lift_ret_val_val_val_val uint32
    lift_ret_val_val_val_val = uint32(ret)
    lift_ret_val_val_val = lift_ret_val_val_val_val
    lift_ret_val_val = lift_ret_val_val_val
    lift_ret_val = lift_ret_val_val
    lift_ret.Set(lift_ret_val)
  }
  return lift_ret
}

func TypesDropFutureIncomingResponse(f uint32) {
  var lower_f C.uint32_t
  lower_f_val := C.uint32_t(f)
  lower_f = lower_f_val
  C.types_drop_future_incoming_response(lower_f)
}

func TypesFutureIncomingResponseGet(f uint32) Option[Result[uint32, TypesError]] {
  var lower_f C.uint32_t
  lower_f_val := C.uint32_t(f)
  lower_f = lower_f_val
  var ret C.types_result_incoming_response_error_t
  C.types_future_incoming_response_get(lower_f, &ret)
  var lift_ret Option[Result[uint32, TypesError]]
  var lift_ret_c C.types_result_incoming_response_error_t
  defer C.types_result_incoming_response_error_free(&lift_ret_c)
  if ret == lift_ret_c {
    lift_ret.Unset()
  } else {
    var lift_ret_val Result[uint32, TypesError]
    if ret.is_err {
      lift_ret_val_ptr := *(*C.types_error_t)(unsafe.Pointer(&ret.val))
      var lift_ret_val_val TypesError
      if lift_ret_val_ptr.tag == 0 {
        lift_ret_val_val_ptr := *(*C.proxy_string_t)(unsafe.Pointer(&lift_ret_val_ptr.val))
        var lift_ret_val_val_val string
        lift_ret_val_val_val = C.GoStringN(lift_ret_val_val_ptr.ptr, C.int(lift_ret_val_val_ptr.len))
        lift_ret_val_val = TypesErrorInvalidUrl(lift_ret_val_val_val)
      }
      if lift_ret_val_ptr.tag == 1 {
        lift_ret_val_val_ptr := *(*C.proxy_string_t)(unsafe.Pointer(&lift_ret_val_ptr.val))
        var lift_ret_val_val_val string
        lift_ret_val_val_val = C.GoStringN(lift_ret_val_val_ptr.ptr, C.int(lift_ret_val_val_ptr.len))
        lift_ret_val_val = TypesErrorTimeoutError(lift_ret_val_val_val)
      }
      if lift_ret_val_ptr.tag == 2 {
        lift_ret_val_val_ptr := *(*C.proxy_string_t)(unsafe.Pointer(&lift_ret_val_ptr.val))
        var lift_ret_val_val_val string
        lift_ret_val_val_val = C.GoStringN(lift_ret_val_val_ptr.ptr, C.int(lift_ret_val_val_ptr.len))
        lift_ret_val_val = TypesErrorProtocolError(lift_ret_val_val_val)
      }
      if lift_ret_val_ptr.tag == 3 {
        lift_ret_val_val_ptr := *(*C.proxy_string_t)(unsafe.Pointer(&lift_ret_val_ptr.val))
        var lift_ret_val_val_val string
        lift_ret_val_val_val = C.GoStringN(lift_ret_val_val_ptr.ptr, C.int(lift_ret_val_val_ptr.len))
        lift_ret_val_val = TypesErrorUnexpectedError(lift_ret_val_val_val)
      }
      lift_ret_val.SetErr(lift_ret_val_val)
    } else {
      lift_ret_val_ptr := *(*C.types_incoming_response_t)(unsafe.Pointer(&ret.val))
      var lift_ret_val_val uint32
      var lift_ret_val_val_val uint32
      lift_ret_val_val_val = uint32(lift_ret_val_ptr)
      lift_ret_val_val = lift_ret_val_val_val
      lift_ret_val.Set(lift_ret_val_val)
    }
    lift_ret.Set(lift_ret_val)
  }
  return lift_ret
}

func TypesListenToFutureIncomingResponse(f uint32) uint32 {
  var lower_f C.uint32_t
  lower_f_val := C.uint32_t(f)
  lower_f = lower_f_val
  ret := C.types_listen_to_future_incoming_response(lower_f)
  var lift_ret uint32
  var lift_ret_val uint32
  var lift_ret_val_val uint32
  lift_ret_val_val = uint32(ret)
  lift_ret_val = lift_ret_val_val
  lift_ret = lift_ret_val
  return lift_ret
}

type TypesTuple2StringStringT struct {
  F0 string
  F1 string
}

// default-outgoing-HTTP
type DefaultOutgoingHttpOutgoingRequest = uint32
type DefaultOutgoingHttpRequestOptions = TypesRequestOptions
type DefaultOutgoingHttpFutureIncomingResponse = uint32
func DefaultOutgoingHttpHandle(request uint32, options Option[TypesRequestOptions]) uint32 {
  var lower_request C.types_outgoing_request_t
  var lower_request_val C.uint32_t
  lower_request_val_val := C.uint32_t(request)
  lower_request_val = lower_request_val_val
  lower_request = lower_request_val
  var lower_options C.default_outgoing_http_request_options_t
  if options.IsSome() {
    var lower_options_val C.types_request_options_t
    var lower_options_val_val C.types_request_options_t
    var lower_options_val_val_connect_timeout_ms C.default_outgoing_http_option_u32_t
    if options.Unwrap().ConnectTimeoutMs.IsSome() {
      lower_options_val_val_connect_timeout_ms_val := C.uint32_t(options.Unwrap().ConnectTimeoutMs.Unwrap())
      lower_options_val_val_connect_timeout_ms.val = lower_options_val_val_connect_timeout_ms_val
      lower_options_val_val_connect_timeout_ms.is_some = true
    }
    lower_options_val_val.connect_timeout_ms = lower_options_val_val_connect_timeout_ms
    var lower_options_val_val_first_byte_timeout_ms C.default_outgoing_http_option_u32_t
    if options.Unwrap().FirstByteTimeoutMs.IsSome() {
      lower_options_val_val_first_byte_timeout_ms_val := C.uint32_t(options.Unwrap().FirstByteTimeoutMs.Unwrap())
      lower_options_val_val_first_byte_timeout_ms.val = lower_options_val_val_first_byte_timeout_ms_val
      lower_options_val_val_first_byte_timeout_ms.is_some = true
    }
    lower_options_val_val.first_byte_timeout_ms = lower_options_val_val_first_byte_timeout_ms
    var lower_options_val_val_between_bytes_timeout_ms C.default_outgoing_http_option_u32_t
    if options.Unwrap().BetweenBytesTimeoutMs.IsSome() {
      lower_options_val_val_between_bytes_timeout_ms_val := C.uint32_t(options.Unwrap().BetweenBytesTimeoutMs.Unwrap())
      lower_options_val_val_between_bytes_timeout_ms.val = lower_options_val_val_between_bytes_timeout_ms_val
      lower_options_val_val_between_bytes_timeout_ms.is_some = true
    }
    lower_options_val_val.between_bytes_timeout_ms = lower_options_val_val_between_bytes_timeout_ms
    lower_options_val = lower_options_val_val
    lower_options = lower_options_val
  }
  ret := C.default_outgoing_http_handle(lower_request, &lower_options)
  var lift_ret uint32
  var lift_ret_val uint32
  var lift_ret_val_val uint32
  lift_ret_val_val = uint32(ret)
  lift_ret_val = lift_ret_val_val
  lift_ret = lift_ret_val
  return lift_ret
}

// HTTP
type HttpIncomingRequest = uint32
type HttpResponseOutparam = uint32
var http Http = nil
func SetHttp(i Http) {
  http = i
}
type Http interface {
  Handle(request uint32, response_out uint32) 
}
//export http_handle
func HttpHandle(request C.http_incoming_request_t, response_out C.http_response_outparam_t) {
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
  http.Handle(lift_request, lift_response_out)
  
}
