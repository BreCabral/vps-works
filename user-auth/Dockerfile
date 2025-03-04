FROM eclipse-temurin:21-jre-alpine

WORKDIR /app
COPY target/user-auth-0.0.1-SNAPSHOT.jar /app/user-auth.jar
ENTRYPOINT ["java", "-jar", "user-auth.jar"]