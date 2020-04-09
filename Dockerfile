FROM alpine

# ca_root_nss

WORKDIR /app
COPY ./app /app
COPY ./zoneinfo.zip /usr/local/go/lib/time/zoneinfo.zip
RUN apk add ca-certificates

USER root
ENV PORT=80
ENV DBSERVER=localhost
ENV LineChannelSecret="53277b183aafd125a52274ce376232e0"
ENV LineChannelAccessToken="95cP2gvG2TGxQI/dwwwU0H9oxEtlRfD8CLcuRVuIBEhQomtnG6hU2Ffj0LkYmmQhk2+KDkVUm5Y58VwdbO+tg9wriQFqn6Ba+uoF20ni6HNts7w9UgYUcCXWhGWPQTv5I3tb8BNPNs/51GwHG2T0OQdB04t89/1O/w1cDnyilFU="

ENTRYPOINT ["/app/app"]
