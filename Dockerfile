FROM jenkins/jenkins:lts-jdk17

USER root

RUN	curl "https://awscli.amazonaws.com/awscli-exe-linux-x86_64.zip" -o "awscliv2.zip"
RUN	unzip awscliv2.zip && ./aws/install



RUN aws ecr get-login-password --region us-east-1

