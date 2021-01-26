FROM scratch
COPY dist/amd64/linux/rip ./
CMD [ "./rip" ]