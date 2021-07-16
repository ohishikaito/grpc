class AddColumnOrderDateToUsers < ActiveRecord::Migration[6.0]
  def change
    add_column :users, :order_date, :datetime
  end
end
