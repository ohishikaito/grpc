class Hoge < ActiveRecord::Migration[6.0]
  def change
    add_column :posts, :hoge, :string
  end
end
