module Stubs
  module Config
    # 環境変数を定義
    name = ENV['GRPC_SERVICE_NAME']
    port = ENV['GRPC_SERVICE_PORT']
    GRPC_HOSTNAME = name + ":" + port

    # lib配下のファイルを読み込むように
    Dir[File.expand_path("#{Rails.root}/lib/pb", __FILE__) << '/*.rb'].each do |file|
      require file
    end
  end
end