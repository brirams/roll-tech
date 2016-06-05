# roll-tech
Baby's first Go and boostrap project. I was hoping to crank this out in a day but that never happened.

## `server`
A really simple REST API that allows for creating and reading contact details.

### Endpoints

- `POST /alumni`: Create an alumn
- `GET /alumni/{id}`: Get alumn with id

### Running the thing
This relies on some environment variables being set to define the db connection details:

``` sh
(source $path_to_secrets && go run server/*.go)
```

## `app`
Can hardly pass for an app but it's a simple static html page that uses some mixture of bootstrap,
custom css, and jquery to hit the above api to save data. It's severely lacking in robustness but
think it's a good place to start from.

## Running the thing
Navigate to `app/index.html` and things should just load. You can mess around and input some data that'll
hit the aforementioned api.
