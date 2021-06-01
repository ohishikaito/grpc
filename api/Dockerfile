FROM ruby:2.6.3 AS builder

RUN apt-get update -qq && \
  apt-get install -y apt-utils \
  vim

WORKDIR /app

ADD Gemfile .
ADD Gemfile.lock .
# add ↓の2つを明示的にinstallしないと、bundle installでエラー吐く
RUN gem install bundler:2.2.17
# 前はgem install railsしてたけど、docker-composeのcommandをbundle execにしたら解消

RUN bundle install -j4


FROM ruby:2.6.3

RUN apt-get update -qq && \
  apt-get install -y apt-utils \
  vim

WORKDIR /app

COPY --from=builder /usr/local/bundle /usr/local/bundle
COPY . /app

COPY entrypoint.sh /usr/bin/
RUN chmod +x /usr/bin/entrypoint.sh
ENTRYPOINT ["entrypoint.sh"]

EXPOSE 3000