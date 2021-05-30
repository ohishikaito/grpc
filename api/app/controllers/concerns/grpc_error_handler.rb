module GrpcErrorHandler
  # NOTE: errorを返す用のメソッド
  # GRPC::BadStatusはStandardErrorを継承してる
  # case分でerror codeごとにresponseを出し分ける

  # ↓記事を参考に実装
  # https://qiita.com/yyh-gl/items/30bd91c2b33fdfbe49b5

  extend ActiveSupport::Concern

  included do
    rescue_from GRPC::BadStatus do |err|
      error_response(err)
    end
  end

  def error_response(err)
    case err.code
    when 5 then
      render_json_error(err.details, :not_found)
    else # NOTE: どのエラーにも該当しない場合
      render_json_error(JSON.parse(err.debug_error_string), :internal_server_error)
    end
  end

  def render_json_error(response, code)
    render json: { response: response }, status: code
  end
end

# ----------------------------------------------
# error code 簡易まとめ
# 詳しくはdocumentを参照
# https://pkg.go.dev/google.golang.org/grpc/codes

# gRPC	HTTP
# OK = 0	                200 OK
# CANCELLED = 1	          499 Client Closed Request
# UNKNOWN = 2	            500 Internal Server Error
# INVALID_ARGUMENT = 3	  400 Bad Request
# DEADLINE_EXCEEDED = 4	  504 Gateway Timeout
# NOT_FOUND = 5	          404 Not Found
# ALREADY_EXISTS = 6	    409 Conflict
# PERMISSION_DENIED = 7	  403 Forbidden
# UNAUTHENTICATED = 16	  401 Unauthorized
# RESOURCE_EXHAUSTED = 8	429 Too Many Requests
# FAILED_PRECONDITION = 9	400 Bad Request
# ABORTED = 10	          409 Conflict
# OUT_OF_RANGE = 11	      400 Bad Request
# UNIMPLEMENTED = 12	    501 Not Implemented
# INTERNAL = 13	          500 Internal Server Error
# UNAVAILABLE = 14	      503 Service Unavailable
# DATA_LOSS = 15	        500 Internal Server Error