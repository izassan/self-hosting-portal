FROM python:3
USER root

RUN mkdir -p /self-hosting-portal

WORKDIR /self-hosting-portal

RUN apt-get update
RUN apt-get -y install locales && \
    localedef -f UTF-8 -i ja_JP ja_JP.UTF-8
ENV LANG ja_JP.UTF-8
ENV LANGUAGE ja_JP:ja
ENV LC_ALL ja_JP.UTF-8
ENV TZ JST-9
ENV TERM xterm

RUN pip install --upgrade pip
RUN pip install --upgrade setuptools

ADD services.json /self-hosting-portal
ADD main.py /self-hosting-portal
ADD requirements.txt /self-hosting-portal
RUN mkdir -p templates
ADD templates /self-hosting-portal/templates

RUN pip install -r requirements.txt

ENTRYPOINT ["python", "main.py"]
