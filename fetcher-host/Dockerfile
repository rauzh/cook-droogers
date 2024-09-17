FROM python:3.9-slim

WORKDIR /usr/src/app

RUN pip install fastapi && pip install uvicorn

COPY . .

CMD ["uvicorn", "main:app", "--host", "0.0.0.0", "--port", "1337"]
