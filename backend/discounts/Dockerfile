FROM golang:1.16-alpine as build

RUN mkdir -p /build
COPY go.mod go.sum /build/

COPY backend/discounts /build/backend/discounts

WORKDIR /build/backend/discounts
RUN go build

FROM alpine

RUN mkdir -p /deploy
COPY --from=build /build/backend/discounts/discounts /deploy/discounts

CMD ["/deploy/discounts"]
