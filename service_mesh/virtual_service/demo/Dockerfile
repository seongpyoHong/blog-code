FROM adoptopenjdk:13-jdk-hotspot as builder
COPY gradlew .
COPY gradle gradle
COPY build.gradle .
COPY settings.gradle .
COPY src src
RUN chmod +x ./gradlew && ./gradlew bootJar

FROM adoptopenjdk:13-jdk-hotspot
COPY --from=builder build/libs/*.jar app.jar
ARG version=NA
ENV VERSION=$version

EXPOSE 8080
ENTRYPOINT ["java", "-jar", "/app.jar"]