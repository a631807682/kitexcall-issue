#!/bin/bash
kitexcall -idl-path  idl/echo/echo.proto -m Echo/ConvertSint64 -d '{"message":"test", "num":2}' -e 127.0.0.1:10001

kitexcall -idl-path  idl/echo/echo.proto -m Echo/ConvertSint64 -d '{"message":"test"}' -e 127.0.0.1:10001

kitexcall -idl-path  idl/multiple/multiple.proto -m Multiple1/Method1 -d '{}' -e 127.0.0.1:10002

kitexcall -idl-path  idl/multiple/multiple.proto -m Multiple2/Method2 -d '{}' -e 127.0.0.1:10002