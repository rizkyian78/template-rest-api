
# Use an image that contains the Google Cloud SDK (gcloud)
FROM gcr.io/google.com/cloudsdktool/google-cloud-cli:alpine AS config

RUN echo "ls"

RUN gcloud --version

SLEEP 

# Stage 2: Create the final image with Jenkins and the configured gcloud
FROM jenkins/jenkins:lts-jdk17
# Switch to root to install software
USER root
# Install the Google Cloud SDK by copying it from the first stage
COPY --from=config /google-cloud-sdk /google-cloud-sdk
