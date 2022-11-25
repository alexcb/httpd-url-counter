## httpd-url-counter

displays the number of times an endpoint has been HTTP GET'ed

## Example

Start the server:

    $ go run main.go
    2022-11-25T18:11:47Z 127.0.0.1:51814 /foo/bar 0
    2022-11-25T18:11:50Z 127.0.0.1:44272 /foo/bar 1
    2022-11-25T18:11:53Z 127.0.0.1:44306 /foo 0
    2022-11-25T18:11:54Z 127.0.0.1:44314 /foo/bar 2

Perform some HTTP GET's

    $ curl http://127.0.0.1:1234/foo/bar
    0
    
    $ curl http://127.0.0.1:1234/foo/bar
    1
    
    $ curl http://127.0.0.1:1234/foo
    0
    
    $ curl http://127.0.0.1:1234/foo/bar
    2
