#!/bin/bash
kitex -no-fast-api \
 -module kitexcall-issue idl/echo/echo.proto

kitex -no-fast-api \
 -module kitexcall-issue idl/multiple/multiple.proto