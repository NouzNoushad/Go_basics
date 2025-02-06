#Use the official Golang image
FROM golang:1.23.3

#Set the working directory
WORKDIR /app

#Copy the Go module files
COPY go.mod .
COPY go.sum .

#Download dependencies
RUN go mod download

#Copty the source code
COPY . .

#Build the application
RUN go build -o main .

#Expose the application port
EXPOSE 8080

#Run the application
CMD [ "./main" ]