# Requirements
### Bento 4
Bento4 is a C++ class library and tools designed to read and write ISO-MP4 files. This format is defined in international specifications ISO/IEC 14496-12, 14496-14 and 14496-15. The format is a derivative of the Apple Quicktime file format, so Bento4 can be used to read and write most Quicktime files as well. In addition to supporting ISO-MP4, Bento4 includes support for parsing and multiplexing H.264 and H.265 elementary streams, converting ISO-MP4 to MPEG2-TS, packaging HLS and MPEG-DASH, CMAF, content encryption, decryption, and much more.

Upload to GCP and send message to RabbitMQ

### Dependencies
* `github.com/asaskevich/govalidator` for data validation
* `github.com/stretchr/testify/require` for testing
* `github.com/satori/go.uuid` for uuid management

# Domains
Video is an entity of this project.

Job are steps which competes to convert the videos into another codec. Job should have:
* id
* local to store