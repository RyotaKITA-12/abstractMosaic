FROM python:3.8-buster as builder

WORKDIR /app

COPY . .
COPY ./app/requirements.txt .
RUN pip3 install -r requirements.txt

CMD ["uvicorn", "main:app", "--reload", "--host", "0.0.0.0", "--port", "8008"]

