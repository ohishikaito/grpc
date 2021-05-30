class ApplicationController < ActionController::API
  Dir[File.expand_path("#{Rails.root}/lib/stubs", __FILE__) << '/*.rb'].each do |file|
    require file
  end
end