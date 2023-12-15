package proxy

// #include "proxy.h"
import "C"

// Temporarily add some drop methods that aren't being generated.

func StaticPollableDrop(handle  WasiHttp0_2_0_rc_2023_11_10_TypesPollable) {
	var lower_self C.wasi_io_0_2_0_rc_2023_11_10_poll_own_pollable_t
	lower_self.__handle = C.int32_t(handle)
  
	C.wasi_io_0_2_0_rc_2023_11_10_poll_pollable_drop_own(lower_self)
}

func StaticFutureIncomingResponseDrop(handle WasiHttp0_2_0_rc_2023_11_10_TypesFutureIncomingResponse) {
	var lower_self C.wasi_http_0_2_0_rc_2023_11_10_types_own_future_incoming_response_t
	lower_self.__handle = C.int32_t(handle)
  
	C.wasi_http_0_2_0_rc_2023_11_10_types_future_incoming_response_drop_own(lower_self)
}

func StaticFieldsDrop(handle WasiHttp0_2_0_rc_2023_11_10_TypesFields) {
	var lower_self C.wasi_http_0_2_0_rc_2023_11_10_types_own_fields_t
	lower_self.__handle = C.int32_t(handle)
  
	C.wasi_http_0_2_0_rc_2023_11_10_types_fields_drop_own(lower_self)
}


func StaticOutgoingRequestDrop(handle WasiHttp0_2_0_rc_2023_11_10_TypesOutgoingRequest) {
	var lower_self C.wasi_http_0_2_0_rc_2023_11_10_types_own_outgoing_request_t
	lower_self.__handle = C.int32_t(handle)
  
	C.wasi_http_0_2_0_rc_2023_11_10_types_outgoing_request_drop_own(lower_self)
}

func StaticIncomingResponseDrop(handle WasiHttp0_2_0_rc_2023_11_10_TypesIncomingResponse) {
	var lower_self C.wasi_http_0_2_0_rc_2023_11_10_types_own_incoming_response_t
	lower_self.__handle = C.int32_t(handle)
  
	C.wasi_http_0_2_0_rc_2023_11_10_types_incoming_response_drop_own(lower_self)
}

func StaticOutgoingStreamDrop(handle WasiIo0_2_0_rc_2023_11_10_StreamsOutputStream) {
	var lower_self C.wasi_io_0_2_0_rc_2023_11_10_streams_own_output_stream_t
	lower_self.__handle = C.int32_t(handle)
  
	C.wasi_io_0_2_0_rc_2023_11_10_streams_output_stream_drop_own(lower_self)
}

func StaticIncomingStreamDrop(handle WasiIo0_2_0_rc_2023_11_10_StreamsInputStream) {
	var lower_self C.wasi_io_0_2_0_rc_2023_11_10_streams_own_input_stream_t
	lower_self.__handle = C.int32_t(handle)
  
	C.wasi_io_0_2_0_rc_2023_11_10_streams_input_stream_drop_own(lower_self)
}