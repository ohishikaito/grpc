FROM ruby:2.6.3

RUN apt-get update -qq && \
  apt-get install -y apt-utils \
  vim

WORKDIR /app

ADD Gemfile .
ADD Gemfile.lock .
# add ↓の2つを明示的にinstallしないと、bundle installとrails sでエラー吐く
RUN gem install bundler:2.2.17
RUN gem install rails -v 6.0.3.7

RUN bundle install -j4

ADD . /app

COPY entrypoint.sh /usr/bin/
RUN chmod +x /usr/bin/entrypoint.sh
ENTRYPOINT ["entrypoint.sh"]

EXPOSE 3000