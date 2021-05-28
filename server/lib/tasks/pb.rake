namespace :pb do
  desc 'proto'
  task protoc: :environment do
    RUBY_OUT_DIR = "app/pb"
    Dir["#{Rails.root}/proto/**/*.proto"].each do |file_path|
      proto_file = file_path.gsub(Rails.root.to_s, '.')
      system("grpc_tools_ruby_protoc --ruby_out=#{Rails.root}/#{RUBY_OUT_DIR} --proto_path=./proto #{proto_file}")
    end
  end
end