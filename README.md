# stemcell-tracker

Pivotal hackday project used to track stemcell compatibility across products

## Stemcell Badge

Add a badge like this <img src="http://stemcell-tracker-hackday.cfapps.io/badge?product_name=cf-mysql" alt="Stemcell"> to your OSS projects to show the latest stemcell that has passed your CI.

#### Embed Badge
```
<img src="http://stemcell-tracker-hackday.cfapps.io/badge?product_name=PRODUCT_NAME">
```

#### Update latest stemcell for product
```
curl -X PUT -d STEMCELL_VERSION http://stemcell-tracker-hackday.cfapps.io/stemcell?product_name=PRODUCT_NAME
```

#### Get latest stemcell for product
```
curl http://stemcell-tracker-hackday.cfapps.io/stemcell?product_name=PRODUCT_NAME
```