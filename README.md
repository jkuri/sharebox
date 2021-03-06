# sharebox

Sharebox is a web file manager application. Its intent is to share files quickly from your host
machine to local network or over the internet with minimal or zero configuration. It comes bundled in a single executable file.

For sharing over the internet, it has included [vex](https://github.com/bleenco/vex) client, so you don't need to worry about firewall or NAT which is configured on your local network.

<p align="center">
  <img src="https://user-images.githubusercontent.com/1796022/43529504-bb6955a0-95ab-11e8-966e-cb445774abea.gif" alt="ShareBox">
</p>

## Usage

### Sharing specific folder

```sh
sharebox -dir /home/bleenco/Videos
```

### Using vex tunnels

You would like to use `-vex` parameter if your machine is running behind firewall or NAT.
This can also come handy if you don't want to expose your public IP.

```sh
sharebox -vex
```

Sample generated URL:

```sh
Serving files from ./ at 0.0.0.0:4505 ...
[vexd] Generated URL: https://6a197324.bleenco.space
```

### Help

For complete command line usage, please see

```sh
sharebox --help
```

## License

> The MIT License
>
> Copyright (c) 2018 Bleenco GmbH https://bleenco.com
>
> Permission is hereby granted, free of charge, to any person obtaining a copy
> of this software and associated documentation files (the "Software"), to deal
> in the Software without restriction, including without limitation the rights
> to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
> copies of the Software, and to permit persons to whom the Software is
> furnished to do so, subject to the following conditions:
>
> The above copyright notice and this permission notice shall be included in
> all copies or substantial portions of the Software.
>
> THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
> IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
> FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
> AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
> LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
> OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
> THE SOFTWARE.
