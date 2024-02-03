# FROM python:3.7-slim

# COPY . /app

# WORKDIR /app

# RUN pip install --trusted-host pypi.python.org -r requirements.txt

# CMD ["python", "main.py"]


FROM ubuntu:latest

RUN apt-get update
RUN apt-get install python3 -y
RUN apt-get install python3-pip -y
RUN pip install selenium



WORKDIR /usr/app/src

# COPY python.py /app/script.py
# COPY go.mod /app
COPY python.py /usr/app/src


# WORKDIR 

#WORKDIR /app

# COPY . .

CMD ["python3", "./python.py"]
