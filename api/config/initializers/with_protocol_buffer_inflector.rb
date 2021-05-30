# Zeitwerkをオンにする場合、設定する必要あり

class WithProtocolBufferInflector < Zeitwerk::Inflector
  def camelize(basename, abspath)
    if basename =~ /\A.*_pb$/
      basename.gsub("_pb", '').camelize #<- `_pb` をここで除去した形でRails側でロードするように対応
    else
      super
    end
  end
end

Rails.autoloaders.each do |autoloader|
  autoloader.inflector = WithProtocolBufferInflector.new
end

# ■ゴール
# pb_user_user_service.rbで読み込ませたい
# ■そのために
# Pb::User::UserServiceって認識できるようにZeitwerkを調教する必要あり
# ↓
# 前提として、user_pbって読み込める？