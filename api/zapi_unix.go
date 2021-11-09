// mksyscall_unix.pl api.go
// MACHINE GENERATED BY THE COMMAND ABOVE; DO NOT EDIT

// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build (darwin || linux) && cgo
// +build darwin linux
// +build cgo

package api

import "unsafe"

// #cgo darwin LDFLAGS: -L/etc/altibase-server-7.1.0/lib  -lodbccli -ldl -lpthread -lcrypt -lrt -lstdc++ -lm
// #cgo darwin CFLAGS: -I/etc/altibase-server-7.1.0/include
// #include <sqlcli.h>
import "C"

func SQLAllocHandle(handleType SQLSMALLINT, inputHandle SQLHANDLE, outputHandle *SQLHANDLE) (ret SQLRETURN) {
	r := C.SQLAllocHandle(C.SQLSMALLINT(handleType), C.SQLHANDLE(inputHandle), (*C.SQLHANDLE)(outputHandle))
	return SQLRETURN(r)
}

func SQLBindCol(statementHandle SQLHSTMT, columnNumber SQLUSMALLINT, targetType SQLSMALLINT, targetValuePtr SQLPOINTER, bufferLength SQLLEN, vallen *SQLLEN) (ret SQLRETURN) {
	r := C.SQLBindCol(C.SQLHSTMT(statementHandle), C.SQLUSMALLINT(columnNumber), C.SQLSMALLINT(targetType), C.SQLPOINTER(targetValuePtr), C.SQLLEN(bufferLength), (*C.SQLLEN)(vallen))
	return SQLRETURN(r)
}

func SQLBindParameter(statementHandle SQLHSTMT, parameterNumber SQLUSMALLINT, inputOutputType SQLSMALLINT, valueType SQLSMALLINT, parameterType SQLSMALLINT, columnSize SQLULEN, decimalDigits SQLSMALLINT, parameterValue SQLPOINTER, bufferLength SQLLEN, ind *SQLLEN) (ret SQLRETURN) {
	r := C.SQLBindParameter(C.SQLHSTMT(statementHandle), C.SQLUSMALLINT(parameterNumber), C.SQLSMALLINT(inputOutputType), C.SQLSMALLINT(valueType), C.SQLSMALLINT(parameterType), C.SQLULEN(columnSize), C.SQLSMALLINT(decimalDigits), C.SQLPOINTER(parameterValue), C.SQLLEN(bufferLength), (*C.SQLLEN)(ind))
	return SQLRETURN(r)
}

func SQLCloseCursor(statementHandle SQLHSTMT) (ret SQLRETURN) {
	r := C.SQLCloseCursor(C.SQLHSTMT(statementHandle))
	return SQLRETURN(r)
}

func SQLDescribeCol(statementHandle SQLHSTMT, columnNumber SQLUSMALLINT, columnName *SQLWCHAR, bufferLength SQLSMALLINT, nameLengthPtr *SQLSMALLINT, dataTypePtr *SQLSMALLINT, columnSizePtr *SQLULEN, decimalDigitsPtr *SQLSMALLINT, nullablePtr *SQLSMALLINT) (ret SQLRETURN) {
	r := C.SQLDescribeColW(C.SQLHSTMT(statementHandle), C.SQLUSMALLINT(columnNumber), (*C.SQLWCHAR)(unsafe.Pointer(columnName)), C.SQLSMALLINT(bufferLength), (*C.SQLSMALLINT)(nameLengthPtr), (*C.SQLSMALLINT)(dataTypePtr), (*C.SQLULEN)(columnSizePtr), (*C.SQLSMALLINT)(decimalDigitsPtr), (*C.SQLSMALLINT)(nullablePtr))
	return SQLRETURN(r)
}

func SQLDescribeParam(statementHandle SQLHSTMT, parameterNumber SQLUSMALLINT, dataTypePtr *SQLSMALLINT, parameterSizePtr *SQLULEN, decimalDigitsPtr *SQLSMALLINT, nullablePtr *SQLSMALLINT) (ret SQLRETURN) {
	r := C.SQLDescribeParam(C.SQLHSTMT(statementHandle), C.SQLUSMALLINT(parameterNumber), (*C.SQLSMALLINT)(dataTypePtr), (*C.SQLULEN)(parameterSizePtr), (*C.SQLSMALLINT)(decimalDigitsPtr), (*C.SQLSMALLINT)(nullablePtr))
	return SQLRETURN(r)
}

func SQLDisconnect(connectionHandle SQLHDBC) (ret SQLRETURN) {
	r := C.SQLDisconnect(C.SQLHDBC(connectionHandle))
	return SQLRETURN(r)
}

func SQLDriverConnect(connectionHandle SQLHDBC, windowHandle SQLHWND, inConnectionString *SQLWCHAR, stringLength1 SQLSMALLINT, outConnectionString *SQLWCHAR, bufferLength SQLSMALLINT, stringLength2Ptr *SQLSMALLINT, driverCompletion SQLUSMALLINT) (ret SQLRETURN) {
	r := C.SQLDriverConnectW(C.SQLHDBC(connectionHandle), C.SQLHWND(windowHandle), (*C.SQLWCHAR)(unsafe.Pointer(inConnectionString)), C.SQLSMALLINT(stringLength1), (*C.SQLWCHAR)(unsafe.Pointer(outConnectionString)), C.SQLSMALLINT(bufferLength), (*C.SQLSMALLINT)(stringLength2Ptr), C.SQLUSMALLINT(driverCompletion))
	return SQLRETURN(r)
}

func SQLEndTran(handleType SQLSMALLINT, handle SQLHANDLE, completionType SQLSMALLINT) (ret SQLRETURN) {
	r := C.SQLEndTran(C.SQLSMALLINT(handleType), C.SQLHANDLE(handle), C.SQLSMALLINT(completionType))
	return SQLRETURN(r)
}

func SQLExecute(statementHandle SQLHSTMT) (ret SQLRETURN) {
	r := C.SQLExecute(C.SQLHSTMT(statementHandle))
	return SQLRETURN(r)
}

func SQLFetch(statementHandle SQLHSTMT) (ret SQLRETURN) {
	r := C.SQLFetch(C.SQLHSTMT(statementHandle))
	return SQLRETURN(r)
}

func SQLFreeHandle(handleType SQLSMALLINT, handle SQLHANDLE) (ret SQLRETURN) {
	r := C.SQLFreeHandle(C.SQLSMALLINT(handleType), C.SQLHANDLE(handle))
	return SQLRETURN(r)
}

func SQLGetData(statementHandle SQLHSTMT, colOrParamNum SQLUSMALLINT, targetType SQLSMALLINT, targetValuePtr SQLPOINTER, bufferLength SQLLEN, vallen *SQLLEN) (ret SQLRETURN) {
	r := C.SQLGetData(C.SQLHSTMT(statementHandle), C.SQLUSMALLINT(colOrParamNum), C.SQLSMALLINT(targetType), C.SQLPOINTER(targetValuePtr), C.SQLLEN(bufferLength), (*C.SQLLEN)(vallen))
	return SQLRETURN(r)
}

func SQLGetDiagRec(handleType SQLSMALLINT, handle SQLHANDLE, recNumber SQLSMALLINT, sqlState *SQLWCHAR, nativeErrorPtr *SQLINTEGER, messageText *SQLWCHAR, bufferLength SQLSMALLINT, textLengthPtr *SQLSMALLINT) (ret SQLRETURN) {
	r := C.SQLGetDiagRecW(C.SQLSMALLINT(handleType), C.SQLHANDLE(handle), C.SQLSMALLINT(recNumber), (*C.SQLWCHAR)(unsafe.Pointer(sqlState)), (*C.SQLINTEGER)(nativeErrorPtr), (*C.SQLWCHAR)(unsafe.Pointer(messageText)), C.SQLSMALLINT(bufferLength), (*C.SQLSMALLINT)(textLengthPtr))
	return SQLRETURN(r)
}

func SQLNumParams(statementHandle SQLHSTMT, parameterCountPtr *SQLSMALLINT) (ret SQLRETURN) {
	r := C.SQLNumParams(C.SQLHSTMT(statementHandle), (*C.SQLSMALLINT)(parameterCountPtr))
	return SQLRETURN(r)
}

func SQLNumResultCols(statementHandle SQLHSTMT, columnCountPtr *SQLSMALLINT) (ret SQLRETURN) {
	r := C.SQLNumResultCols(C.SQLHSTMT(statementHandle), (*C.SQLSMALLINT)(columnCountPtr))
	return SQLRETURN(r)
}

func SQLPrepare(statementHandle SQLHSTMT, statementText *SQLWCHAR, textLength SQLINTEGER) (ret SQLRETURN) {
	r := C.SQLPrepareW(C.SQLHSTMT(statementHandle), (*C.SQLWCHAR)(unsafe.Pointer(statementText)), C.SQLINTEGER(textLength))
	return SQLRETURN(r)
}

func SQLRowCount(statementHandle SQLHSTMT, rowCountPtr *SQLLEN) (ret SQLRETURN) {
	r := C.SQLRowCount(C.SQLHSTMT(statementHandle), (*C.SQLLEN)(rowCountPtr))
	return SQLRETURN(r)
}

func SQLSetEnvAttr(environmentHandle SQLHENV, attribute SQLINTEGER, valuePtr SQLPOINTER, stringLength SQLINTEGER) (ret SQLRETURN) {
	r := C.SQLSetEnvAttr(C.SQLHENV(environmentHandle), C.SQLINTEGER(attribute), C.SQLPOINTER(valuePtr), C.SQLINTEGER(stringLength))
	return SQLRETURN(r)
}

func SQLSetConnectAttr(connectionHandle SQLHDBC, attribute SQLINTEGER, valuePtr SQLPOINTER, stringLength SQLINTEGER) (ret SQLRETURN) {
	r := C.SQLSetConnectAttrW(C.SQLHDBC(connectionHandle), C.SQLINTEGER(attribute), C.SQLPOINTER(valuePtr), C.SQLINTEGER(stringLength))
	return SQLRETURN(r)
}
