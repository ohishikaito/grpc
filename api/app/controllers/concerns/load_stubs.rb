module LoadStubs
  # NOTE: 使い回さないからConcernにはしない
  # extend ActiveSupport::Concern

  def load_stubs
    Dir[File.expand_path("#{Rails.root}/lib/stubs", __FILE__) << '/*.rb'].each do |file|
      require file
    end
  end
end