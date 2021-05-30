class ApplicationController < ActionController::API
  Dir[File.expand_path("#{Rails.root}/lib/pb", __FILE__) << '/*.rb'].each do |file|
    require file
  end
end
# Dir[File.expand_path('../commands', __FILE__) << '/*.rb'].each do |file|
#   require file
# end