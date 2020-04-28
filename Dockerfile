FROM alpine:latest

MAINTAINER 826466266@qq.com

WORKDIR /export/webapps/pinecms/

ADD pinecms /export/webapps/pinecms/pinecms
ADD resources /export/webapps/pinecms/resources
ADD data.db /export/webapps/pinecms/data.db

EXPOSE 2019

CMD ["./pinecms","serve", "start"]

