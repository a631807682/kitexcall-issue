#!/bin/bash
kitexcall -idl-path  idl/echo/echo.proto -m Echo/ConvertSint64 -d '{"message":"test", "num":2}' -e 127.0.0.1:10001