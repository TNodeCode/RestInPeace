FROM scratch
COPY dist/linux/arm64/rip ./
CMD [ "./rip" ]