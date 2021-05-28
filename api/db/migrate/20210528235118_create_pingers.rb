class CreatePingers < ActiveRecord::Migration[6.0]
  def change
    create_table :pingers do |t|

      t.timestamps
    end
  end
end
