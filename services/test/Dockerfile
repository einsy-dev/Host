# --- STAGE 1: Build ---
FROM golang:alpine AS builder
WORKDIR /app
# Best practice: copy mods first to cache dependencies
COPY go.mod ./ 
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# --- STAGE 2: Run ---
FROM alpine:latest
WORKDIR /root/
# Copy ONLY the binary from the builder
COPY --from=builder /app/main .
EXPOSE 80
CMD ["./main"]