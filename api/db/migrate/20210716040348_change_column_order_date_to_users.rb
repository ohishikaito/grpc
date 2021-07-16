class ChangeColumnOrderDateToUsers < ActiveRecord::Migration[6.0]
  def change
    change_column :users, :order_date, :datetime, null: false
  end
end
