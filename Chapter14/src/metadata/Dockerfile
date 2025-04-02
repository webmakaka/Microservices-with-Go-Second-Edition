FROM alpine:latest

# Copy our binary and configs to the container.
COPY main .
COPY configs/. .
# Expose the port for accepting incoming requests.
EXPOSE 8081
# Execute our service.
CMD ["/main"]
