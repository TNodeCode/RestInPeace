FROM scratch
COPY dist/arm/rip ./
CMD [ "./rip" ]