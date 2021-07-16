class User < ApplicationRecord
  attribute :liked
  attribute :epoc_created_at
  attribute :bazirisuku_time

  validates :order_date, presence: true
end
