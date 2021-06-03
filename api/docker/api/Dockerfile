# localでdocker-compose upすると
# Could not load the 'listen' gem. Add `gem 'listen'` to the development group of your Gemfile (LoadError)
# って言われるから一旦保留

FROM ruby:2.6.3

WORKDIR /app

COPY Gemfile .
COPY Gemfile.lock .
# NOTE: 本番用はdevelopmentとtest groupのgemをinstallしないように
RUN bundle install -j4 --without=development test

COPY . /app

COPY entrypoint_prod.sh /usr/bin/
RUN chmod +x /usr/bin/entrypoint_prod.sh
ENTRYPOINT ["entrypoint_prod.sh"]

EXPOSE 3000
