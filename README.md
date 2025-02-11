# ![HAProxy](assets/images/haproxy-weblogo-210x49.png "HAProxy")

## HAProxy Swagger Models

This project contains structs and validation methods that are autogenerated using [go-swagger](https://github.com/go-swagger/go-swagger) from the swagger spec found [here](https://github.com/haproxytech/dataplaneapi-specification/blob/master/build/haproxy_spec.yaml).These structs and are used in the [DataPlaneAPI](http://github.com/haproxytech/dataplaneapi) and [client-native](http://github.com/haproxytech/client-native).

## Contributing

You can generate the models using the following command:

```
./swagger generate model -f haproxy_spec.yaml \
    -m models \
    -t $GOPATH/src/github.com/haproxytech/
```

For commit messages and general style please follow the haproxy project's [CONTRIBUTING guide](https://github.com/haproxy/haproxy/blob/master/CONTRIBUTING) and use that where applicable.