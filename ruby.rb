#!/usr/bin/env ruby

require 'ffi'

module Goroutines
  extend FFI::Library
  ffi_lib './go.so'
  attach_function :Goroutines, [], :void
end

Goroutines.Goroutines
