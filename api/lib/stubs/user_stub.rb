module Stubs
  module UserStub
    extend Stubs::Config

    Stub = create_stub(Pb::UserService::Stub)
  end
end