FROM golang:1.18-alpine

# set a directory for the app
WORKDIR /usr/src/app

# copy all the files to the container
COPY . .

# RUN go get "github.com/gin-gonic/gin"
# RUN go get "go.mongodb.org/mongo-driver/mongo"
# RUN go get "go.mongodb.org/mongo-driver/mongo/options"
# RUN go get "go.mongodb.org/mongo-driver/bson/primitive"
# RUN go get "go.mongodb.org/mongo-driver/bson"
# RUN go get "github.com/aws/aws-sdk-go/aws"
# RUN go get "github.com/aws/aws-sdk-go/aws/session"
# RUN go get "github.com/aws/aws-sdk-go/service/s3"
# RUN go get "github.com/aws/aws-sdk-go/service/s3/s3manager"
# RUN go get "github.com/dgrijalva/jwt-go"
# RUN go get "github.com/andru100/Social-Network"

RUN go mod download

RUN go build -o /Backend

EXPOSE 8080 8081

CMD ["/Backend"]