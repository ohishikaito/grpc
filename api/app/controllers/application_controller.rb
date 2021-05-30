class ApplicationController < ActionController::API
  extend LoadStubs

  # NOTE: gRPCのstubを読み込む
  load_stubs
end