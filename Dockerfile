FROM busybox:1.31

COPY ./gowebarduino /home/
EXPOSE 8000

CMD ["/home/gowebarduino"]
