FROM golang:1.24 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN find /app
RUN ls -la /app
RUN ls -la /app/internal
RUN ls -la /app/internal/bot

