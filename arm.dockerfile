FROM scratch
COPY dist/linux/arm/rip ./
CMD [ "./rip" ]