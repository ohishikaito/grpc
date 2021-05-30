require_relative 'boot'

require "rails"
# Pick the frameworks you want:
require "active_model/railtie"
require "active_job/railtie"
require "active_record/railtie"
require "active_storage/engine"
require "action_controller/railtie"
require "action_mailer/railtie"
require "action_mailbox/engine"
require "action_text/engine"
require "action_view/railtie"
require "action_cable/engine"
# require "sprockets/railtie"
require "rails/test_unit/railtie"

# Require the gems listed in Gemfile, including any gems
# you've limited to :test, :development, or :production.
Bundler.require(*Rails.groups)

module Api
  class Application < Rails::Application
    # Initialize configuration defaults for originally generated Rails version.
    config.load_defaults 6.0

    # ZeitwerkをONにすると、↓にしてlib配下を読ませる必要あり
    config.paths.add 'lib', eager_load: true
    # autoload_pathsの一覧
    # ActiveSupport::Dependencies.autoload_paths

    # NOTE: Unable to add defs to DescriptorPool: duplicate symbol 'pb.user.Users' errorが発生するため
    # https://github.com/ruby-protobuf/protobuf/issues/381
    grpc_root = File.join(config.root, "/lib/pb")
    $LOAD_PATH.unshift(grpc_root)
    Dir[File.join(grpc_root, "*.rb")].each do |file|
      require file
    end

    # Settings in config/environments/* take precedence over those specified here.
    # Application configuration can go into files in config/initializers
    # -- all .rb files in that directory are automatically loaded after loading
    # the framework and any gems in your application.

    # Only loads a smaller set of middleware suitable for API only apps.
    # Middleware like session, flash, cookies can be added back manually.
    # Skip views, helpers and assets when generating a new resource.
    config.api_only = true
  end
end
