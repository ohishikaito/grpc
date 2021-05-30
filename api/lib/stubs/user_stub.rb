require 'config'

module Stubs
  module UserStub
    extend Stubs::Config

    Stub = create_stub(Pb::User::UserService::Stub)
  end
end