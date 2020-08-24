FROM frolvlad/alpine-glibc
WORKDIR /app
COPY lab1  /app/
EXPOSE 8080
ENTRYPOINT ["/app/lab1"]