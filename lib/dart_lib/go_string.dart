import 'dart:convert';
import 'dart:ffi';
import 'package:ffi/ffi.dart';

class GoString extends Struct {
  external Pointer<Uint8> string;

  @IntPtr()
  external int length;

  String toString() {
    List<int> units = [];
    for (int i = 0; i < length; ++i) {
      units.add(string.elementAt(i).value);
    }
    return const Utf8Decoder().convert(units);
  }

  static Pointer<GoString> fromString(String string) {
    List<int> units = const Utf8Encoder().convert(string);
    final ptr = malloc<Uint8>(units.length);
    for (int i = 0; i < units.length; ++i) {
      ptr.elementAt(i).value = units[i];
    }
    Pointer<GoString> goPtr = malloc<GoString>();
    final GoString str = goPtr.ref;
    str.length = units.length;
    str.string = ptr;
    return goPtr;
  }
}
