FROM rockylinux:9.2.20230513

WORKDIR /app
COPY . .
RUN chmod 0755 plusone_someone
RUN chown root:root plusone_someone
CMD ["./plusone_someone"]
EXPOSE 8000