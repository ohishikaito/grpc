# localでdocker-compose upすると
# Could not load the 'listen' gem. Add `gem 'listen'` to the development group of your Gemfile (LoadError)
# って言われるから一旦保留

FROM ruby:2.6.3

WORKDIR /app

ADD Gemfile .
ADD Gemfile.lock .
# add: ↓を明示的にinstallしないと、bundle installでエラー吐く
RUN gem install bundler:2.2.17
# 本番用はbundle installするgroupを絞る
# RUN bundle config set without "development test"
RUN bundle install -j4

COPY . /app

COPY entrypoint_prod.sh /usr/bin/
RUN chmod +x /usr/bin/entrypoint_prod.sh
ENTRYPOINT ["entrypoint_prod.sh"]

EXPOSE 3000
