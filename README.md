# Kansuler/http-redirect

![License](https://img.shields.io/github/license/Kansuler/http-redirect) ![Version](https://img.shields.io/github/go-mod/go-version/Kansuler/http-redirect) [![Codacy Badge](https://app.codacy.com/project/badge/Grade/32399e0c2b5e4e1aa32e2f4570baad0d)](https://www.codacy.com/gh/Kansuler/http-redirect/dashboard?utm_source=github.com&utm_medium=referral&utm_content=Kansuler/http-redirect&utm_campaign=Badge_Grade)

A small container package that are configurable through environment variables to redirect to another domain.

Docker hub image: [kansuler/http-redirect](https://hub.docker.com/r/kansuler/http-redirect)

API and detailed documentation can be found at <https://godoc.org/github.com/Kansuler/http-request>

| environment variable | description                                                                                                                 |
| -------------------- | --------------------------------------------------------------------------------------------------------------------------- |
| target_host          | the target domain that the visitor will be redirected to. Format should be `https://example.com`                            |
| append_uri           | if `true` the URI path will append to the Location header. resulting in a redirect with URI like `https://example.com/path` |
| log_format           | Set default settings for different cloud providers. Available formats are `stackdriver`                                     |
| http_status          | set either `307` temporary redirect or `308` permanent redirect. Default is `307`                                           |

## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.
