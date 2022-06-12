import 'dart:ffi';
import 'dart:io';

import 'package:ffi/ffi.dart';
import 'package:flutter/foundation.dart';
import 'package:flutter_call_go/dart_lib/go_string.dart';

DynamicLibrary lib = DynamicLibrary.open(getGoLibPath());

// 注册dart api 用于异步回调
initializeDartApi() {
  final initializeApi = lib.lookupFunction<IntPtr Function(Pointer<Void>),
      int Function(Pointer<Void>)>("InitializeDartApi");

  if (initializeApi(NativeApi.initializeApiDLData) != 0) {
    throw "Failed to initialize Dart API";
  }
}

String getGoLibPath() {
  if (Platform.isMacOS || Platform.isIOS) {
    return 'flutter_call_go.dylib';
  }
  if (Platform.isWindows) {
    return 'flutter_call_go.dll';
  }
  if (Platform.isAndroid || Platform.isLinux) {
    return 'flutter_call_go.so';
  }
  return 'flutter_call_go.so';
}

// 同步调用
typedef SynchronousCallFunc = Pointer<Int8> Function(Pointer<GoString>, int);
SynchronousCallFunc goSynchronousCall = lib
    .lookup<NativeFunction<Pointer<Int8> Function(Pointer<GoString>, Int32)>>(
        'SynchronousCall')
    .asFunction();

// 异步调用
typedef AsynchronousCallFunc = Pointer<Int8> Function(Pointer<GoString>, int);
AsynchronousCallFunc goAsynchronousCall = lib
    .lookup<NativeFunction<Pointer<Int8> Function(Pointer<GoString>, Int32)>>(
        'AsynchronousCall')
    .asFunction();

class CallGo {
  // 同步调用
  static String synchronousCall(String param1, int param2) {
    Pointer<Int8> val = goSynchronousCall(GoString.fromString(param1), param2);
    String valStr = val.cast<Utf8>().toDartString();
    return valStr;
  }

  // 异步调用
  static Future<String> asynchronousCall(String param1, int param2) async {
    String result = await compute<Map<String, dynamic>, String>((data) {
      Pointer<Int8> val = goAsynchronousCall(
          GoString.fromString(data['param1'].toString()), data['param2']);
      String valStr = val.cast<Utf8>().toDartString();
      return valStr;
    }, {
      'param1': param1,
      'param2': param2,
    });
    return result;
  }
}
