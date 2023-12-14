package proxy

// #include "proxy.h"
import "C"

// Temporarily add some drop methods that aren't being generated.

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