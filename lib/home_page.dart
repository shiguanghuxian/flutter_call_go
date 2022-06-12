import 'dart:developer';

import 'package:flutter/material.dart';
import 'package:flutter_call_go/dart_lib/dart_lib.dart';
import 'package:ftoast/ftoast.dart';

class HomePage extends StatefulWidget {
  const HomePage({Key? key}) : super(key: key);

  @override
  State<HomePage> createState() => _HomePageState();
}

class _HomePageState extends State<HomePage> {
  // 同步调用
  void synchronousCall() {
    log('同步调用');
    String valStr = CallGo.synchronousCall('synchronousCall', 9);
    FToast.toast(context, msg: "同步调用结果：$valStr");
  }

  // 异步调用
  void asynchronousCall() async {
    log('异步调用');
    String valStr = await CallGo.asynchronousCall('asynchronousCall', 9);
    FToast.toast(context, msg: "同步调用结果：$valStr");
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text('flutter call go'),
      ),
      body: Center(
        child: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          children: <Widget>[
            TextButton(
              onPressed: synchronousCall,
              child: SizedBox(
                width: MediaQuery.of(context).size.width * 80 / 100,
                child: const Center(
                  child: Text('同步调用'),
                ),
              ),
            ),
            TextButton(
              onPressed: asynchronousCall,
              child: SizedBox(
                width: MediaQuery.of(context).size.width * 80 / 100,
                child: const Center(
                  child: Text('异步调用'),
                ),
              ),
            ),
          ],
        ),
      ),
    );
  }
}
