# Video Transcoding Service

Written in [Go](https://go.dev/)

* Uploads video to AWS S3 object storage
* Utitlizes Ffmpeg for segmentation
* AWS ECS for scaling transcoding containers (scale)

##### Check List

* [x] Repository Setup
* [x] DB Connection Setup
* [x] User Management APIs (Authentication)
* [x] S3 Presigned Upload
* [x] Lambda Trigger Queue Offload uploaded video metadata 
* [ ] Transcoding Service
* [ ] Dockerize Transcoder Service
* [ ] Deploy Transcoder Service AWS ECS
* [ ] Dockerize Upload Service( + User Management)
* [ ] Deploy to Upload Service AWS ECS