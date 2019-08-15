// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: google/rpc/code.proto

package rpc

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import strconv "strconv"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// The canonical error codes for Google APIs.
//
//
// Sometimes multiple error codes may apply.  Services should return
// the most specific error code that applies.  For example, prefer
// `OUT_OF_RANGE` over `FAILED_PRECONDITION` if both codes apply.
// Similarly prefer `NOT_FOUND` or `ALREADY_EXISTS` over `FAILED_PRECONDITION`.
type Code int32

const (
	// Not an error; returned on success
	//
	// HTTP Mapping: 200 OK
	OK Code = 0
	// The operation was cancelled, typically by the caller.
	//
	// HTTP Mapping: 499 Client Closed Request
	CANCELLED Code = 1
	// Unknown error.  For example, this error may be returned when
	// a `Status` value received from another address space belongs to
	// an error space that is not known in this address space.  Also
	// errors raised by APIs that do not return enough error information
	// may be converted to this error.
	//
	// HTTP Mapping: 500 Internal Server Error
	UNKNOWN Code = 2
	// The client specified an invalid argument.  Note that this differs
	// from `FAILED_PRECONDITION`.  `INVALID_ARGUMENT` indicates arguments
	// that are problematic regardless of the state of the system
	// (e.g., a malformed file name).
	//
	// HTTP Mapping: 400 Bad Request
	INVALID_ARGUMENT Code = 3
	// The deadline expired before the operation could complete. For operations
	// that change the state of the system, this error may be returned
	// even if the operation has completed successfully.  For example, a
	// successful response from a server could have been delayed long
	// enough for the deadline to expire.
	//
	// HTTP Mapping: 504 Gateway Timeout
	DEADLINE_EXCEEDED Code = 4
	// Some requested entity (e.g., file or directory) was not found.
	//
	// Note to server developers: if a request is denied for an entire class
	// of users, such as gradual feature rollout or undocumented whitelist,
	// `NOT_FOUND` may be used. If a request is denied for some users within
	// a class of users, such as user-based access control, `PERMISSION_DENIED`
	// must be used.
	//
	// HTTP Mapping: 404 Not Found
	NOT_FOUND Code = 5
	// The entity that a client attempted to create (e.g., file or directory)
	// already exists.
	//
	// HTTP Mapping: 409 Conflict
	ALREADY_EXISTS Code = 6
	// The caller does not have permission to execute the specified
	// operation. `PERMISSION_DENIED` must not be used for rejections
	// caused by exhausting some resource (use `RESOURCE_EXHAUSTED`
	// instead for those errors). `PERMISSION_DENIED` must not be
	// used if the caller can not be identified (use `UNAUTHENTICATED`
	// instead for those errors). This error code does not imply the
	// request is valid or the requested entity exists or satisfies
	// other pre-conditions.
	//
	// HTTP Mapping: 403 Forbidden
	PERMISSION_DENIED Code = 7
	// The request does not have valid authentication credentials for the
	// operation.
	//
	// HTTP Mapping: 401 Unauthorized
	UNAUTHENTICATED Code = 16
	// Some resource has been exhausted, perhaps a per-user quota, or
	// perhaps the entire file system is out of space.
	//
	// HTTP Mapping: 429 Too Many Requests
	RESOURCE_EXHAUSTED Code = 8
	// The operation was rejected because the system is not in a state
	// required for the operation's execution.  For example, the directory
	// to be deleted is non-empty, an rmdir operation is applied to
	// a non-directory, etc.
	//
	// Service implementors can use the following guidelines to decide
	// between `FAILED_PRECONDITION`, `ABORTED`, and `UNAVAILABLE`:
	//  (a) Use `UNAVAILABLE` if the client can retry just the failing call.
	//  (b) Use `ABORTED` if the client should retry at a higher level
	//      (e.g., when a client-specified test-and-set fails, indicating the
	//      client should restart a read-modify-write sequence).
	//  (c) Use `FAILED_PRECONDITION` if the client should not retry until
	//      the system state has been explicitly fixed.  E.g., if an "rmdir"
	//      fails because the directory is non-empty, `FAILED_PRECONDITION`
	//      should be returned since the client should not retry unless
	//      the files are deleted from the directory.
	//
	// HTTP Mapping: 400 Bad Request
	FAILED_PRECONDITION Code = 9
	// The operation was aborted, typically due to a concurrency issue such as
	// a sequencer check failure or transaction abort.
	//
	// See the guidelines above for deciding between `FAILED_PRECONDITION`,
	// `ABORTED`, and `UNAVAILABLE`.
	//
	// HTTP Mapping: 409 Conflict
	ABORTED Code = 10
	// The operation was attempted past the valid range.  E.g., seeking or
	// reading past end-of-file.
	//
	// Unlike `INVALID_ARGUMENT`, this error indicates a problem that may
	// be fixed if the system state changes. For example, a 32-bit file
	// system will generate `INVALID_ARGUMENT` if asked to read at an
	// offset that is not in the range [0,2^32-1], but it will generate
	// `OUT_OF_RANGE` if asked to read from an offset past the current
	// file size.
	//
	// There is a fair bit of overlap between `FAILED_PRECONDITION` and
	// `OUT_OF_RANGE`.  We recommend using `OUT_OF_RANGE` (the more specific
	// error) when it applies so that callers who are iterating through
	// a space can easily look for an `OUT_OF_RANGE` error to detect when
	// they are done.
	//
	// HTTP Mapping: 400 Bad Request
	OUT_OF_RANGE Code = 11
	// The operation is not implemented or is not supported/enabled in this
	// service.
	//
	// HTTP Mapping: 501 Not Implemented
	UNIMPLEMENTED Code = 12
	// Internal errors.  This means that some invariants expected by the
	// underlying system have been broken.  This error code is reserved
	// for serious errors.
	//
	// HTTP Mapping: 500 Internal Server Error
	INTERNAL Code = 13
	// The service is currently unavailable.  This is most likely a
	// transient condition, which can be corrected by retrying with
	// a backoff.
	//
	// See the guidelines above for deciding between `FAILED_PRECONDITION`,
	// `ABORTED`, and `UNAVAILABLE`.
	//
	// HTTP Mapping: 503 Service Unavailable
	UNAVAILABLE Code = 14
	// Unrecoverable data loss or corruption.
	//
	// HTTP Mapping: 500 Internal Server Error
	DATA_LOSS Code = 15
)

var Code_name = map[int32]string{
	0:  "OK",
	1:  "CANCELLED",
	2:  "UNKNOWN",
	3:  "INVALID_ARGUMENT",
	4:  "DEADLINE_EXCEEDED",
	5:  "NOT_FOUND",
	6:  "ALREADY_EXISTS",
	7:  "PERMISSION_DENIED",
	16: "UNAUTHENTICATED",
	8:  "RESOURCE_EXHAUSTED",
	9:  "FAILED_PRECONDITION",
	10: "ABORTED",
	11: "OUT_OF_RANGE",
	12: "UNIMPLEMENTED",
	13: "INTERNAL",
	14: "UNAVAILABLE",
	15: "DATA_LOSS",
}
var Code_value = map[string]int32{
	"OK":                  0,
	"CANCELLED":           1,
	"UNKNOWN":             2,
	"INVALID_ARGUMENT":    3,
	"DEADLINE_EXCEEDED":   4,
	"NOT_FOUND":           5,
	"ALREADY_EXISTS":      6,
	"PERMISSION_DENIED":   7,
	"UNAUTHENTICATED":     16,
	"RESOURCE_EXHAUSTED":  8,
	"FAILED_PRECONDITION": 9,
	"ABORTED":             10,
	"OUT_OF_RANGE":        11,
	"UNIMPLEMENTED":       12,
	"INTERNAL":            13,
	"UNAVAILABLE":         14,
	"DATA_LOSS":           15,
}

func (Code) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_code_29c527ae4b4f0e5d, []int{0}
}

func init() {
	proto.RegisterEnum("google.rpc.Code", Code_name, Code_value)
}
func (x Code) String() string {
	s, ok := Code_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}

func init() { proto.RegisterFile("google/rpc/code.proto", fileDescriptor_code_29c527ae4b4f0e5d) }

var fileDescriptor_code_29c527ae4b4f0e5d = []byte{
	// 393 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x44, 0x91, 0x3d, 0x6e, 0x13, 0x41,
	0x14, 0xc7, 0x3d, 0x76, 0x70, 0xe2, 0xf1, 0xd7, 0xcb, 0x84, 0x40, 0x37, 0x07, 0xa0, 0x70, 0x0a,
	0x4e, 0xf0, 0xbc, 0xf3, 0x9c, 0x8c, 0x32, 0x7e, 0xb3, 0x9a, 0x9d, 0x09, 0x01, 0x21, 0xad, 0xc4,
	0xc6, 0x4a, 0x03, 0x5a, 0xcb, 0xe2, 0x00, 0x9c, 0x85, 0x8a, 0x1b, 0x70, 0x85, 0x94, 0x29, 0x29,
	0xf1, 0xa6, 0xa1, 0x74, 0x49, 0x89, 0x06, 0x0a, 0xda, 0x9f, 0xde, 0xc7, 0xff, 0x43, 0x9e, 0xdf,
	0xb7, 0xed, 0xfd, 0xc7, 0xcd, 0xc5, 0x6e, 0xdb, 0x5c, 0x34, 0xed, 0xdd, 0x66, 0xb1, 0xdd, 0xb5,
	0x9f, 0x5b, 0x25, 0xff, 0xe1, 0xc5, 0x6e, 0xdb, 0xbc, 0xfa, 0xde, 0x97, 0x47, 0x45, 0x7b, 0xb7,
	0x51, 0x43, 0xd9, 0xf7, 0xd7, 0xd0, 0x53, 0x53, 0x39, 0x2a, 0x90, 0x0b, 0x72, 0x8e, 0x0c, 0x08,
	0x35, 0x96, 0xc7, 0x89, 0xaf, 0xd9, 0xbf, 0x61, 0xe8, 0xab, 0xe7, 0x12, 0x2c, 0xdf, 0xa0, 0xb3,
	0xa6, 0xc6, 0x70, 0x99, 0xd6, 0xc4, 0x11, 0x06, 0xea, 0x5c, 0x9e, 0x1a, 0x42, 0xe3, 0x2c, 0x53,
	0x4d, 0xb7, 0x05, 0x91, 0x21, 0x03, 0x47, 0xf9, 0x10, 0xfb, 0x58, 0xaf, 0x7c, 0x62, 0x03, 0xcf,
	0x94, 0x92, 0x33, 0x74, 0x81, 0xd0, 0xbc, 0xad, 0xe9, 0xd6, 0x56, 0xb1, 0x82, 0x61, 0xde, 0x2c,
	0x29, 0xac, 0x6d, 0x55, 0x59, 0xcf, 0xb5, 0x21, 0xb6, 0x64, 0xe0, 0x58, 0x9d, 0xc9, 0x79, 0x62,
	0x4c, 0xf1, 0x8a, 0x38, 0xda, 0x02, 0x23, 0x19, 0x00, 0xf5, 0x42, 0xaa, 0x40, 0x95, 0x4f, 0xa1,
	0xc8, 0x5f, 0xae, 0x30, 0x55, 0x99, 0x9f, 0xa8, 0x97, 0xf2, 0x6c, 0x85, 0xd6, 0x91, 0xa9, 0xcb,
	0x40, 0x85, 0x67, 0x63, 0xa3, 0xf5, 0x0c, 0xa3, 0xac, 0x1c, 0x97, 0x3e, 0xe4, 0x29, 0xa9, 0x40,
	0x4e, 0x7c, 0x8a, 0xb5, 0x5f, 0xd5, 0x01, 0xf9, 0x92, 0x60, 0xac, 0x4e, 0xe5, 0x34, 0xb1, 0x5d,
	0x97, 0x8e, 0xb2, 0x0d, 0x32, 0x30, 0x51, 0x13, 0x79, 0x62, 0x39, 0x52, 0x60, 0x74, 0x30, 0x55,
	0x73, 0x39, 0x4e, 0x8c, 0x37, 0x68, 0x1d, 0x2e, 0x1d, 0xc1, 0x2c, 0x1b, 0x32, 0x18, 0xb1, 0x76,
	0xbe, 0xaa, 0x60, 0xbe, 0x7c, 0xff, 0xb8, 0xd7, 0xbd, 0x1f, 0x7b, 0xdd, 0x3b, 0xec, 0xb5, 0xf8,
	0xbd, 0xd7, 0xe2, 0x4b, 0xa7, 0xc5, 0xb7, 0x4e, 0x8b, 0x87, 0x4e, 0x8b, 0xc7, 0x4e, 0x8b, 0x9f,
	0x9d, 0x16, 0xbf, 0x3a, 0xdd, 0x3b, 0x64, 0xfe, 0xa4, 0xc5, 0xc3, 0x93, 0x16, 0x72, 0xd6, 0xb4,
	0x9f, 0x16, 0xff, 0xf3, 0x5f, 0x8e, 0x72, 0xf8, 0x65, 0xae, 0xa5, 0x14, 0xef, 0x06, 0xbb, 0x6d,
	0xf3, 0xb5, 0x3f, 0x08, 0x65, 0xf1, 0x61, 0xf8, 0xb7, 0xaa, 0xd7, 0x7f, 0x02, 0x00, 0x00, 0xff,
	0xff, 0x03, 0xd4, 0x27, 0xff, 0xc3, 0x01, 0x00, 0x00,
}
