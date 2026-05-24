FROM busybox:1.38

COPY ./gowebarduino /home/
EXPOSE 8000

CMD ["/home/gowebarduino"]
